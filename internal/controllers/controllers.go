package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/robin-andreasson/fox"
	"github.com/robin-andreasson/yt-aitool/internal/types"
	"github.com/robin-andreasson/yt-aitool/internal/utils"
	"github.com/robin-andreasson/yt-aitool/internal/yt/subtitles"
	"github.com/robin-andreasson/yt-aitool/internal/yt/video"
)

func Subtitles(c *fox.Context) error {
	payload := Get(time.Now(), c.Body)

	status := fox.Get[string](payload, "status")
	code := fox.Status.Ok

	if status == "error" {
		code = fox.Status.BadRequest
	}

	return c.JSON(code, payload)
}

func Summarize(c *fox.Context) error {
	return GetAIResponse("summarize", c)
}

func Explain(c *fox.Context) error {
	return GetAIResponse("explain", c)
}

func GetAIResponse(variant string, c *fox.Context) error {
	start := time.Now()

	payload := Get(start, c.Body)

	status := fox.Get[string](payload, "status")

	if status == "error" {
		return c.JSON(fox.Status.BadRequest, payload)
	}

	subtitles := fox.Get[string](payload, "result")

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", strings.NewReader(utils.GenerateRequestBody(variant, subtitles)))

	req.Header.Add("Authorization", "Bearer "+os.Getenv("CHAT_GPT_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return c.JSON(fox.Status.InternalServerError, utils.GenerateResponse("error", start, "internal error, try again later"))
	}

	body, _ := io.ReadAll(res.Body)

	var data types.ChatGptResponse

	json.Unmarshal(body, &data)

	if len(data.Choices) == 0 {
		return c.JSON(fox.Status.InternalServerError, utils.GenerateResponse("error", start, "internal error, try again later"))
	}

	return c.JSON(fox.Status.Ok, utils.GenerateResponse("success", start, utils.Escape(data.Choices[0].Text)))
}

func Get(start time.Time, body any) map[string]any {

	url := fox.Get[string](body, "url")
	language := fox.Get[string](body, "language")

	if language == "" {
		language = "en"
	}

	videoData, err := video.FetchVideo(url)

	if err != nil {
		return utils.GenerateResponse("error", start, "error getting video")
	}

	captionTracks := videoData.Captions.PlayerCaptionsTracklistRenderer.CaptionTracks

	if len(captionTracks) == 0 {
		return utils.GenerateResponse("error", start, "video does not have subtitles")
	}

	for _, caption := range captionTracks {
		if caption.LanguageCode != language {
			continue
		}

		subtitles, err := subtitles.FetchSubtitles(caption.BaseUrl)

		if err != nil {
			return utils.GenerateResponse("error", start, "could not get subtitles")
		}

		return utils.GenerateResponse("success", start, subtitles)
	}

	return utils.GenerateResponse("error", start, "video does not support the specified language ( "+language+" )")
}

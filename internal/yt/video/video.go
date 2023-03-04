package video

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/robin-andreasson/yt-aitool/internal/types"
)

func FetchVideo(url string) (types.Video, error) {
	res, err := http.Get(url)

	if err != nil {
		return types.Video{}, fmt.Errorf("could not get video")
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return types.Video{}, fmt.Errorf("could not get video")
	}

	strBody := string(body)

	rex := regexp.MustCompile("var ytInitialPlayerResponse = (.*})")

	result := rex.FindStringSubmatch(strBody)

	if len(result) < 2 {
		return types.Video{}, fmt.Errorf("invalid response")
	}

	var video types.Video

	err = json.Unmarshal([]byte(result[1]), &video)

	return video, err
}

package subtitles

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/robin-andreasson/yt-aitool/internal/types"
	"github.com/robin-andreasson/yt-aitool/internal/utils"
	tokenizer "github.com/sandwich-go/gpt3-encoder"
)

func FetchSubtitles(url string) (string, error) {
	res, err := http.Get(url + "&fmt=json3")

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	var sub types.Subtitles

	err = json.Unmarshal(body, &sub)

	if err != nil {
		return "", err
	}

	var subtitles string

	for _, e := range sub.Events {
		for _, s := range e.Segs {
			subtitles += s.Utf8
		}
	}

	subtitles = utils.Escape(Filter(subtitles))

	return subtitles, nil
}

func Filter(str string) string {
	encoder, err := tokenizer.NewEncoder()
	if err != nil {
		log.Fatal(err)
	}

	encoded, _ := encoder.Encode(str)

	if len(encoded) > 2000 {
		encoded = encoded[:2000]
	}

	return encoder.Decode(encoded)
}

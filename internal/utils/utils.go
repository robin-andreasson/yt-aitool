package utils

import (
	"fmt"
	"strings"
	"time"
)

func GenerateRequestBody(variant string, prompt string) string {
	return fmt.Sprint(`
	{
		"model":  "text-davinci-003",
		"prompt": "` + generatePrompt(variant, prompt) + `",
        "max_tokens": 2000
	}
	`)
}

func generatePrompt(variant string, subtitles string) string {
	what := `I want you to explain and teach me the topic discussed in this text:\n`
	if variant == "summarize" {
		what = `I want you to summarize this text while using present tense:\n`
	}

	return fmt.Sprint(what + `\"` + subtitles + `\"`)
}

func GenerateResponse(status string, start time.Time, result string) map[string]any {
	return map[string]any{
		"status": status,
		"time":   fmt.Sprint(time.Since(start).Microseconds(), "μ"),
		"result": result,
	}
}

func Escape(str string) string {
	str = strings.Replace(str, "\n", `\n`, -1)

	return strings.Replace(str, `"`, `\"`, -1)
}

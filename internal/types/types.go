package types

type Subtitles struct {
	Events []Event `json:"events"`
}

type Event struct {
	//TStartMs     float64 `json:"tStartMs"`
	//DDurationMs  float64 `json:"dDurationMs"`
	//Id           float64 `json:"id"`
	//WpWinPosId   float64 `json:"wpWinPosId"`
	//WsWinStyleId float64 `json:"wsWinStyleId"`
	//
	//WWinId float64 `json:"wWinId"`

	Segs []Seg `json:"segs"`
}

type Seg struct {
	Utf8 string `json:"utf8"`
	//TOffsetMs float64 `json:"tOffsetMs"`
	//AcAsrConf float64 `json:"acAsrConf"`
}

type Video struct {
	Captions     Captions     `json:"captions"`
	VideoDetails VideoDetails `json:"videoDetails"`
}

type VideoDetails struct {
	Keywords         []string `json:"keywords"`
	ShortDescription string   `json:"shortDescription"`
}

type Captions struct {
	PlayerCaptionsTracklistRenderer PlayerCaptionsTracklistRenderer `json:"playerCaptionsTracklistRenderer"`
}

type PlayerCaptionsTracklistRenderer struct {
	CaptionTracks        []CaptionTracks `json:"captionTracks"`
	TranslationLanguages []LanguageCode  `json:"translationLanguages"`
}

type CaptionTracks struct {
	BaseUrl        string `json:"baseUrl"`
	LanguageCode   string `json:"languageCode"`
	IsTranslatable bool   `json:"isTranslatable"`
}

type TranslationLanguages struct {
	LanguageCodes []LanguageCode `json:"translationLanguages"`
}

type LanguageCode struct {
	LanguageCode string `json:"languageCode"`
}

type ChatGptResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text string `json:"text"`
}

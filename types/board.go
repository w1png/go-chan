package types

import (
	"io/ioutil"
	"net/http"
	"strconv"

	j "encoding/json"
)

type Board struct {
	Board           string `json:"board"`
	Title           string `json:"title"`
	Worksafe        int    `json:"ws_board"`
	PerPage         int    `json:"per_page"`
	Pages           int    `json:"pages"`
	MaxFilesize     int    `json:"max_filesize"`
	MaxWebmFilesize int    `json:"max_webm_filesize"`
	MaxCommentChars int    `json:"max_comment_chars"`
	MaxWebmDuration int    `json:"max_webm_duration"`
	BumpLimit       int    `json:"bump_limit"`
	ImageLimit      int    `json:"image_limit"`
	Cooldowns       struct {
		Threads int `json:"threads"`
		Replies int `json:"replies"`
		Images  int `json:"images"`
	} `json:"cooldowns"`
	MetaDescription string      `json:"meta_description"`
	IsArchived      int         `json:"is_archived"`
	BoardFlags      interface{} `json:"board_flags"`
	CountryFlags    int         `json:"country_flags"`
	UserIds         int         `json:"user_ids"`
	Oekaki          int         `json:"oekaki"`
	SjisTags        int         `json:"sjis_tags"`
	CodeTags        int         `json:"code_tags"`
	MathTags        int         `json:"math_tags"`
	TextOnly        int         `json:"text_only"`
	ForceAnon       int         `json:"force_anon"`
	WebmAudio       int         `json:"webm_audio"`
	RequireSubject  int         `json:"require_subject"`
	MinImageWidth   int         `json:"min_image_width"`
	MinImageHeight  int         `json:"min_image_height"`
}

func (b *Board) GetPage(page int) ([]Thread, error) {
	resp, err := http.Get("https://a.4cdn.org/" + b.Board + "/" + strconv.Itoa(page) + ".json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type JSON struct {
		Threads []Thread `json:"threads"`
	}
	var json JSON
	err = j.Unmarshal(body, &json)
	if err != nil {
		return nil, err
	}

	return json.Threads, nil
}

func (b *Board) GetPages() ([]Page, error) {
	resp, err := http.Get("https://a.4cdn.org/" + b.Board + "/catalog.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pages []Page
	err = j.Unmarshal(body, &pages)
	if err != nil {
		return nil, err
	}

	return pages, nil
}

func (b *Board) GetAllThreads() ([]Thread, error) {
	var threads []Thread

	pages, err := b.GetPages()
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		threads = append(threads, page.Threads...)
	}

	return threads, nil
}

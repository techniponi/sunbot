package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// DerpiResults is a struct to contain Derpibooru's JSON search results
type DerpiResults struct {
	Search []struct {
		ID              int       `json:"id"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		FirstSeenAt     time.Time `json:"first_seen_at"`
		Score           int       `json:"score"`
		CommentCount    int       `json:"comment_count"`
		Width           int       `json:"width"`
		Height          int       `json:"height"`
		FileName        string    `json:"file_name"`
		Description     string    `json:"description"`
		Uploader        string    `json:"uploader"`
		UploaderID      int       `json:"uploader_id"`
		Image           string    `json:"image"`
		Upvotes         int       `json:"upvotes"`
		Downvotes       int       `json:"downvotes"`
		Faves           int       `json:"faves"`
		Tags            string    `json:"tags"`
		TagIds          []int     `json:"tag_ids"`
		AspectRatio     float64   `json:"aspect_ratio"`
		OriginalFormat  string    `json:"original_format"`
		MimeType        string    `json:"mime_type"`
		Sha512Hash      string    `json:"sha512_hash"`
		OrigSha512Hash  string    `json:"orig_sha512_hash"`
		SourceURL       string    `json:"source_url"`
		Representations struct {
			ThumbTiny  string `json:"thumb_tiny"`
			ThumbSmall string `json:"thumb_small"`
			Thumb      string `json:"thumb"`
			Small      string `json:"small"`
			Medium     string `json:"medium"`
			Large      string `json:"large"`
			Tall       string `json:"tall"`
			Full       string `json:"full"`
		} `json:"representations"`
		IsRendered  bool `json:"is_rendered"`
		IsOptimized bool `json:"is_optimized"`
	} `json:"search"`
	Total        int           `json:"total"`
	Interactions []interface{} `json:"interactions"`
}

// Perform a Derpibooru search query with a given string of tags and an API key
func DerpiSearchWithTags(tags string, key string) (DerpiResults, error) {

	// format for URL query
	derpiTags := strings.Replace(tags, " ", "+", -1)

	// make URL query
	urlQuery := "https://derpibooru.org/search.json?q=" + derpiTags
	if key != "" {
		urlQuery += "&key=" + key
	}
	resp, err := http.Get(urlQuery)
	if err != nil {
		fmt.Println(err)
		return DerpiResults{}, fmt.Errorf("Failed with HTTP error.")
	}

	// read response body
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return DerpiResults{}, fmt.Errorf("Failed with error reading response body.")
	}

	// parse json
	results := DerpiResults{}
	err = json.Unmarshal(respBody, &results)
	if err != nil {
		fmt.Println(err)
		return DerpiResults{}, fmt.Errorf("Failed with JSON parsing error.")
	}

	return results, nil

}

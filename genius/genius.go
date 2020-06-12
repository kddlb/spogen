// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    searchResult, err := UnmarshalSearchResult(bytes)
//    bytes, err = searchResult.Marshal()

package genius

import "encoding/json"

func (r *SearchResult) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SearchResult struct {
	Meta     Meta     `json:"meta"`
	Response Response `json:"response"`
}

type Meta struct {
	Status int64 `json:"status"`
}

type Response struct {
	Hits []Hit `json:"hits"`
}

type Hit struct {
	Highlights []interface{} `json:"highlights"`
	Index      string        `json:"index"`
	Type       string        `json:"type"`
	Result     Result        `json:"result"`
}

type Result struct {
	AnnotationCount          int64         `json:"annotation_count"`
	APIPath                  string        `json:"api_path"`
	FullTitle                string        `json:"full_title"`
	HeaderImageThumbnailURL  string        `json:"header_image_thumbnail_url"`
	HeaderImageURL           string        `json:"header_image_url"`
	ID                       int64         `json:"id"`
	LyricsOwnerID            int64         `json:"lyrics_owner_id"`
	LyricsState              string        `json:"lyrics_state"`
	Path                     string        `json:"path"`
	PyongsCount              *int64        `json:"pyongs_count"`
	SongArtImageThumbnailURL string        `json:"song_art_image_thumbnail_url"`
	SongArtImageURL          string        `json:"song_art_image_url"`
	Stats                    Stats         `json:"stats"`
	Title                    string        `json:"title"`
	TitleWithFeatured        string        `json:"title_with_featured"`
	URL                      string        `json:"url"`
	PrimaryArtist            PrimaryArtist `json:"primary_artist"`
	SpogenDistance           int           `json:"SPOGENDistance,omitempty"`
}

type PrimaryArtist struct {
	APIPath        string `json:"api_path"`
	HeaderImageURL string `json:"header_image_url"`
	ID             int64  `json:"id"`
	ImageURL       string `json:"image_url"`
	IsMemeVerified bool   `json:"is_meme_verified"`
	IsVerified     bool   `json:"is_verified"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	Iq             *int64 `json:"iq,omitempty"`
}

type Stats struct {
	UnreviewedAnnotations int64  `json:"unreviewed_annotations"`
	Hot                   bool   `json:"hot"`
	Pageviews             *int64 `json:"pageviews,omitempty"`
	Concurrents           *int64 `json:"concurrents,omitempty"`
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    song, err := UnmarshalSong(bytes)
//    bytes, err = song.Marshal()
//
//    searchResult, err := UnmarshalSearchResult(bytes)
//    bytes, err = searchResult.Marshal()

package genius

import (
	"bytes"
	"encoding/json"
	"errors"
)

func UnmarshalSong(data []byte) (Song, error) {
	var r Song
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Song) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSearchResult(data []byte) (SearchResult, error) {
	var r SearchResult
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SearchResult) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Song struct {
	Meta     Meta         `json:"meta"`
	Response SongResponse `json:"response"`
}

type Meta struct {
	Status int64 `json:"status"`
}

type SongResponse struct {
	Song SongClass `json:"song"`
}

type SongClass struct {
	AnnotationCount          int64                   `json:"annotation_count"`
	APIPath                  string                  `json:"api_path"`
	AppleMusicID             string                  `json:"apple_music_id"`
	AppleMusicPlayerURL      string                  `json:"apple_music_player_url"`
	Description              Description             `json:"description"`
	EmbedContent             string                  `json:"embed_content"`
	FeaturedVideo            bool                    `json:"featured_video"`
	FullTitle                string                  `json:"full_title"`
	HeaderImageThumbnailURL  string                  `json:"header_image_thumbnail_url"`
	HeaderImageURL           string                  `json:"header_image_url"`
	ID                       int64                   `json:"id"`
	LyricsOwnerID            int64                   `json:"lyrics_owner_id"`
	LyricsPlaceholderReason  interface{}             `json:"lyrics_placeholder_reason"`
	LyricsState              string                  `json:"lyrics_state"`
	Path                     string                  `json:"path"`
	PyongsCount              int64                   `json:"pyongs_count"`
	RecordingLocation        string                  `json:"recording_location"`
	ReleaseDate              string                  `json:"release_date"`
	ReleaseDateForDisplay    string                  `json:"release_date_for_display"`
	SongArtImageThumbnailURL string                  `json:"song_art_image_thumbnail_url"`
	SongArtImageURL          string                  `json:"song_art_image_url"`
	Stats                    SongStats               `json:"stats"`
	Title                    string                  `json:"title"`
	TitleWithFeatured        string                  `json:"title_with_featured"`
	URL                      string                  `json:"url"`
	CurrentUserMetadata      SongCurrentUserMetadata `json:"current_user_metadata"`
	Album                    Album                   `json:"album"`
	CustomPerformances       []CustomPerformance     `json:"custom_performances"`
	DescriptionAnnotation    DescriptionAnnotation   `json:"description_annotation"`
	FeaturedArtists          []interface{}           `json:"featured_artists"`
	LyricsMarkedCompleteBy   interface{}             `json:"lyrics_marked_complete_by"`
	Media                    []Media                 `json:"media"`
	PrimaryArtist            Artist                  `json:"primary_artist"`
	ProducerArtists          []Artist                `json:"producer_artists"`
	SongRelationships        []SongRelationship      `json:"song_relationships"`
	VerifiedAnnotationsBy    []interface{}           `json:"verified_annotations_by"`
	VerifiedContributors     []interface{}           `json:"verified_contributors"`
	VerifiedLyricsBy         []interface{}           `json:"verified_lyrics_by"`
	WriterArtists            []Artist                `json:"writer_artists"`
}

type Album struct {
	APIPath     string `json:"api_path"`
	CoverArtURL string `json:"cover_art_url"`
	FullTitle   string `json:"full_title"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Artist      Artist `json:"artist"`
}

type Artist struct {
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

type SongCurrentUserMetadata struct {
	Permissions         []string           `json:"permissions"`
	ExcludedPermissions []string           `json:"excluded_permissions"`
	Interactions        PurpleInteractions `json:"interactions"`
	Relationships       IqByAction         `json:"relationships"`
	IqByAction          IqByAction         `json:"iq_by_action"`
}

type PurpleInteractions struct {
	Pyong     bool `json:"pyong"`
	Following bool `json:"following"`
}

type IqByAction struct {
}

type CustomPerformance struct {
	Label   string   `json:"label"`
	Artists []Artist `json:"artists"`
}

type Description struct {
	DOM DOM `json:"dom"`
}

type DOM struct {
	Tag      string     `json:"tag"`
	Children []DOMChild `json:"children"`
}

type PurpleChild struct {
	Tag      string        `json:"tag"`
	Children []StickyChild `json:"children"`
}

type FluffyChild struct {
	Tag        string        `json:"tag"`
	Attributes *Attributes   `json:"attributes,omitempty"`
	Data       *Data         `json:"data,omitempty"`
	Children   []IndigoChild `json:"children"`
}

type Attributes struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type TentacledChild struct {
	Tag      string   `json:"tag"`
	Children []string `json:"children"`
}

type Data struct {
	APIPath string `json:"api_path"`
}

type DescriptionAnnotation struct {
	Type                 string        `json:"_type"`
	AnnotatorID          int64         `json:"annotator_id"`
	AnnotatorLogin       string        `json:"annotator_login"`
	APIPath              string        `json:"api_path"`
	Classification       string        `json:"classification"`
	Fragment             string        `json:"fragment"`
	ID                   int64         `json:"id"`
	IsDescription        bool          `json:"is_description"`
	Path                 string        `json:"path"`
	Range                Range         `json:"range"`
	SongID               int64         `json:"song_id"`
	URL                  string        `json:"url"`
	VerifiedAnnotatorIDS []interface{} `json:"verified_annotator_ids"`
	Annotatable          Annotatable   `json:"annotatable"`
	Annotations          []Annotation  `json:"annotations"`
}

type Annotatable struct {
	APIPath          string           `json:"api_path"`
	ClientTimestamps ClientTimestamps `json:"client_timestamps"`
	Context          string           `json:"context"`
	ID               int64            `json:"id"`
	ImageURL         string           `json:"image_url"`
	LinkTitle        string           `json:"link_title"`
	Title            string           `json:"title"`
	Type             string           `json:"type"`
	URL              string           `json:"url"`
}

type ClientTimestamps struct {
	UpdatedByHumanAt int64 `json:"updated_by_human_at"`
	LyricsUpdatedAt  int64 `json:"lyrics_updated_at"`
}

type Annotation struct {
	APIPath             string                        `json:"api_path"`
	Body                Description                   `json:"body"`
	CommentCount        int64                         `json:"comment_count"`
	Community           bool                          `json:"community"`
	CustomPreview       interface{}                   `json:"custom_preview"`
	HasVoters           bool                          `json:"has_voters"`
	ID                  int64                         `json:"id"`
	Pinned              bool                          `json:"pinned"`
	ShareURL            string                        `json:"share_url"`
	Source              interface{}                   `json:"source"`
	State               string                        `json:"state"`
	URL                 string                        `json:"url"`
	Verified            bool                          `json:"verified"`
	VotesTotal          int64                         `json:"votes_total"`
	CurrentUserMetadata AnnotationCurrentUserMetadata `json:"current_user_metadata"`
	Authors             []Author                      `json:"authors"`
	CosignedBy          []interface{}                 `json:"cosigned_by"`
	RejectionComment    interface{}                   `json:"rejection_comment"`
	VerifiedBy          interface{}                   `json:"verified_by"`
}

type Author struct {
	Attribution int64       `json:"attribution"`
	PinnedRole  interface{} `json:"pinned_role"`
	User        User        `json:"user"`
}

type User struct {
	APIPath                     string                  `json:"api_path"`
	Avatar                      Avatar                  `json:"avatar"`
	HeaderImageURL              string                  `json:"header_image_url"`
	HumanReadableRoleForDisplay string                  `json:"human_readable_role_for_display"`
	ID                          int64                   `json:"id"`
	Iq                          int64                   `json:"iq"`
	Login                       string                  `json:"login"`
	Name                        string                  `json:"name"`
	RoleForDisplay              string                  `json:"role_for_display"`
	URL                         string                  `json:"url"`
	CurrentUserMetadata         UserCurrentUserMetadata `json:"current_user_metadata"`
}

type Avatar struct {
	Tiny   Medium `json:"tiny"`
	Thumb  Medium `json:"thumb"`
	Small  Medium `json:"small"`
	Medium Medium `json:"medium"`
}

type Medium struct {
	URL         string      `json:"url"`
	BoundingBox BoundingBox `json:"bounding_box"`
}

type BoundingBox struct {
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

type UserCurrentUserMetadata struct {
	Permissions         []interface{}      `json:"permissions"`
	ExcludedPermissions []string           `json:"excluded_permissions"`
	Interactions        FluffyInteractions `json:"interactions"`
}

type FluffyInteractions struct {
	Following bool `json:"following"`
}

type AnnotationCurrentUserMetadata struct {
	Permissions         []string              `json:"permissions"`
	ExcludedPermissions []string              `json:"excluded_permissions"`
	Interactions        TentacledInteractions `json:"interactions"`
	IqByAction          IqByAction            `json:"iq_by_action"`
}

type TentacledInteractions struct {
	Cosign bool        `json:"cosign"`
	Pyong  bool        `json:"pyong"`
	Vote   interface{} `json:"vote"`
}

type Range struct {
	Content string `json:"content"`
}

type Media struct {
	NativeURI *string `json:"native_uri,omitempty"`
	Provider  string  `json:"provider"`
	Type      string  `json:"type"`
	URL       string  `json:"url"`
	Start     *int64  `json:"start,omitempty"`
}

type SongRelationship struct {
	RelationshipType string        `json:"relationship_type"`
	Type             string        `json:"type"`
	Songs            []interface{} `json:"songs"`
}

type SongStats struct {
	AcceptedAnnotations   int64 `json:"accepted_annotations"`
	Contributors          int64 `json:"contributors"`
	IqEarners             int64 `json:"iq_earners"`
	Transcribers          int64 `json:"transcribers"`
	UnreviewedAnnotations int64 `json:"unreviewed_annotations"`
	VerifiedAnnotations   int64 `json:"verified_annotations"`
	Hot                   bool  `json:"hot"`
	Pageviews             int64 `json:"pageviews"`
}

type SearchResult struct {
	Meta     Meta                 `json:"meta"`
	Response SearchResultResponse `json:"response"`
}

type SearchResultResponse struct {
	Hits []Hit `json:"hits"`
}

type Hit struct {
	Highlights []interface{} `json:"highlights"`
	Index      string        `json:"index"`
	Type       string        `json:"type"`
	Result     Result        `json:"result"`
}

type Result struct {
	AnnotationCount          int64       `json:"annotation_count"`
	APIPath                  string      `json:"api_path"`
	FullTitle                string      `json:"full_title"`
	HeaderImageThumbnailURL  string      `json:"header_image_thumbnail_url"`
	HeaderImageURL           string      `json:"header_image_url"`
	ID                       int64       `json:"id"`
	LyricsOwnerID            int64       `json:"lyrics_owner_id"`
	LyricsState              string      `json:"lyrics_state"`
	Path                     string      `json:"path"`
	PyongsCount              *int64      `json:"pyongs_count"`
	SongArtImageThumbnailURL string      `json:"song_art_image_thumbnail_url"`
	SongArtImageURL          string      `json:"song_art_image_url"`
	Stats                    ResultStats `json:"stats"`
	Title                    string      `json:"title"`
	TitleWithFeatured        string      `json:"title_with_featured"`
	URL                      string      `json:"url"`
	PrimaryArtist            Artist      `json:"primary_artist"`
	SpogenDistance           int         `json:"SPOGENDistance,omitempty"`
}

type ResultStats struct {
	UnreviewedAnnotations int64  `json:"unreviewed_annotations"`
	Hot                   bool   `json:"hot"`
	Pageviews             *int64 `json:"pageviews,omitempty"`
}

type DOMChild struct {
	PurpleChild *PurpleChild
	String      *string
}

func (x *DOMChild) UnmarshalJSON(data []byte) error {
	x.PurpleChild = nil
	var c PurpleChild
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.PurpleChild = &c
	}
	return nil
}

func (x *DOMChild) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.PurpleChild != nil, x.PurpleChild, false, nil, false, nil, false)
}

type StickyChild struct {
	FluffyChild *FluffyChild
	String      *string
}

func (x *StickyChild) UnmarshalJSON(data []byte) error {
	x.FluffyChild = nil
	var c FluffyChild
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.FluffyChild = &c
	}
	return nil
}

func (x *StickyChild) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.FluffyChild != nil, x.FluffyChild, false, nil, false, nil, false)
}

type IndigoChild struct {
	String         *string
	TentacledChild *TentacledChild
}

func (x *IndigoChild) UnmarshalJSON(data []byte) error {
	x.TentacledChild = nil
	var c TentacledChild
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.TentacledChild = &c
	}
	return nil
}

func (x *IndigoChild) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.TentacledChild != nil, x.TentacledChild, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}

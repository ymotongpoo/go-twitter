package twitter

// http://dev.twitter.com/docs/platform-objects/tweets
type Tweets struct {
	Annotations        string        `json:"annotations"`
	Contributors       *Contributors `json:"contributors"`
	Coordinates        *Coordinates  `json:"coordinates"`
	CreatedAt          string        `json:"created_at"`
	CurrentUserRetweet struct {
		Id    int64  `json:"id"`
		IdStr string `json:"id_str"`
	} `json:"current_user_retweet"`
	Entities             *Entities              `json:"entities"`
	FavoriteCount        int                    `json:"favorite_count"`
	Favorited            bool                   `json:"favorited"`
	FilterLevel          string                 `json:"filter_level"`
	Id                   int64                  `json:"id"`
	IdStr                string                 `json:"id_str"`
	InReplyToScreenName  string                 `json:"in_reply_to_screen_name"`
	InReplyToStatusId    int64                  `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr string                 `json:"in_reply_to_status_id_str"`
	InReplyToUserId      int64                  `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   string                 `json:"in_reply_to_user_id_str"`
	Lang                 string                 `json:"lang"`
	Place                *Places                `json:"place"`
	PossiblySensitive    bool                   `json:"possibly_sensitive"`
	Scopes               map[string]interface{} `json:"scopes"`
	RetweetCount         int                    `json:"retweet_count"`
	Retweeted            bool                   `json:"retweeted"`
	Source               string                 `json:"source"`
	Text                 string                 `json:"text"`
	Truncated            bool                   `json:"truncated"`
	User                 *Users                 `json:"user"`
	WithheldCopyright    bool                   `json:"withheld_copyright"`
	WithheldInCountries  []string               `json:"withheld_in_countries"`
}

type Contributors struct {
	Id         int64  `json:"id"`
	IdStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

type Coordinates struct {
	Coordinates []float `json:"coordinates"`
	Type        string  `json:"type"`
}

// http://dev.twitter.com/docs/platform-objects/entities
type Entities struct {
	Hashtags    []*Hashtags    `json:"hashtags"`
	Media       []*Media       `json:"media"`
	Urls        []*Url         `json:"urls"`
	UserMetions []*UserMention `json:"user_mentions"`
}

type Hashtags struct {
	Indices []int  `json:"indices"`
	Text    string `json:"text"`
}

type Media struct {
	DisplayUrl        string `json:"display_url"`
	ExpandedUrl       string `json:"expanded_url"`
	Id                int64  `json:"id"`
	IdStr             string `json:"string"`
	Indices           []int  `json:"indices"`
	MediaUrl          string `json:"media_url"`
	MediaUrlHttps     string `json:"media_url_https"`
	Sizes             *Sizes `json:"sizes"`
	SourceStatusId    int64  `json:"source_status_id"`
	SourceStatusIdStr string `json:"source_stats_id_str"`
	Type              string `json:"type"`
	Url               string `json:"url"`
}

type Size struct {
	H      int    `json:"h"`
	Resize string `json:"resize"`
	W      int    `json:"w"`
}

type Sizes struct {
	Thumb  *Size `json:"thumb"`
	Large  *Size `json:"large"`
	Medium *Size `json:"medium"`
	Small  *Size `json:"small"`
}

type Url struct {
	DisplayUrl  string `json:"display_url"`
	ExpandedUrl string `json:"expanded_url"`
	Indices     []int  `json:"indices"`
	Url         string `json:"url"`
}

type UserMention struct {
	Id         int64  `json:"id"`
	IdStr      string `json:"id_str"`
	Indices    []int  `json:"indices"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

// https://dev.twitter.com/docs/platform-objects/places
type Places struct {
	Attributes  map[string]string `json:"attributes"`
	BoundingBox *BoundingBox      `json:"bounding_box"`
	Country     string            `json:"country"`
	CountryCode string            `json:"country_code"`
	FullName    string            `json:"full_name"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	PlaceType   string            `json:"place_type"`
	Url         string            `json:"url"`
}

type BoundingBox struct {
	Coordinates [][][]float `json:"coordinates"`
	Type        string      `json:"type"`
}

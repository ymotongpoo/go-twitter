package twitter

const (
	RestAPIBaseUrl       = "https://api.twitter.com/1.1/"
	StreamAPIBaseUrl     = "https://stream.twitter.com/1.1/"
	UserStreamAPIBaseUrl = "https://userstream.twitter.com/1.1/"
	SiteStreamAPIBaseUrl = "https://sitestream.twitter.com/1.1/"
)

var ResourceInfoMap = map[string]*ResourceInfo{
	"statuses/mentions_timeline": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/mentions_timeline.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{},
		OptionalArgs: []string{"count", "since_id", "max_id",
			"contributor_details", "include_entities"},
	},
	"statuses/user_timeline": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/user_timeline.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{},
		OptionalArgs: []string{"user_id", "screen_name", "since_id",
			"count", "trim_user", "exclude_replies", "contributor_details",
			"include_rts"},
	},
	"statuses/home_timeline": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/home_timeline.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{},
		OptionalArgs: []string{"count", "since_id", "max_id", "trim_user",
			"exclude_replies", "contributor_details", "include_entities"},
	},
	"statuses/retweets_of_me": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/retweets_of_me.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{},
		OptionalArgs: []string{"count", "since_id", "max_id", "trim_user",
			"include_entities", "include_user_entities"},
	},
	"statuses/show/:id": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/show/%v.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{"id"},
		OptionalArgs: []string{"count", "trim_user", "include_my_retweet",
			"include_entities"},
	},
	"statuses/destroy/:id": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/destroy/%v.json",
		Authentication: true,
		HttpMethod:     "POST",
		RequiredArgs:   []string{"id"},
		OptionalArgs: []string{"count", "trim_user", "include_my_retweet",
			"include_entities"},
	},
	"statuses/update": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/update.json",
		Authentication: true,
		HttpMethod:     "POST",
		RequiredArgs:   []string{"status"},
		OptionalArgs: []string{"in_reply_to_status_id", "lat", "long",
			"place_id", "display_coordinates", "trim_user"},
	},
	"statuses/retweet/:id": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/retweet/%v.json",
		Authentication: true,
		HttpMethod:     "POST",
		RequiredArgs:   []string{"id"},
		OptionalArgs:   []string{"trim_user"},
	},
	"statuses/update_with_media": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/update_with_media.json",
		Authentication: true,
		HttpMethod:     "POST",
		RequiredArgs:   []string{"status", "media[]"},
		OptionalArgs: []string{"possibly_sensitive", "in_reply_to_status_id",
			"lat", "long", "place_id", "display_coordinates"},
	},
	"statuses/oembed": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "statuses/oembed.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{"id", "url"},
		OptionalArgs: []string{"maxwidth", "hide_media", "hide_thread",
			"omit_script", "align", "related", "lang"},
	},
	"search/tweets": &ResourceInfo{
		EndPoint:       RestAPIBaseUrl + "search/tweets.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{"q"},
		OptionalArgs: []string{"geocode", "lang", "locate", "result_type",
			"count", "until", "since_id", "max_id", "include_entites",
			"callback"},
	},
	"statuses/filter": &ResourceInfo{
		EndPoint:       StreamAPIBaseUrl + "statuses/filter.json",
		Authentication: true,
		HttpMethod:     "POST",
		RequiredArgs:   []string{},
		OptionalArgs: []string{"follow", "track", "locations", "delimited",
			"stall_warning"},
	},
	"statuses/sample": &ResourceInfo{
		EndPoint:       StreamAPIBaseUrl + "statuses/sample.json",
		Authentication: true,
		HttpMethod:     "POST",
		RequiredArgs:   []string{},
		OptionalArgs:   []string{"delimited", "stall_warning"},
	},
	"statuses/firehose": &ResourceInfo{
		EndPoint:       StreamAPIBaseUrl + "statuses/firehose.json",
		Authentication: true,
		HttpMethod:     "POST",
		RequiredArgs:   []string{},
		OptionalArgs:   []string{"count", "delimited", "stall_warning"},
	},
	"user": &ResourceInfo{
		EndPoint:       UserStreamAPIBaseUrl + "user.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{},
		OptionalArgs: []string{"delimited", "stall_warnings", "with",
			"replies", "track", "locations"},
	},
	"site": &ResourceInfo{
		EndPoint:       SiteStreamAPIBaseUrl + "site.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{"follow"},
		OptionalArgs: []string{"delimited", "stall_warnings", "with",
			"replies"},
	},
}

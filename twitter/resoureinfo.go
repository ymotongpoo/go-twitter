package twitter

const APIBaseUrl = "https://api.twitter.com/1.1/"

var ResourceInfoMap = map[string]*ResourceInfo{
	"status/mentions_timeline": &ResourceInfo{
		EndPoint:       APIBaseUrl + "status/mentions_timeline.json",
		Authentication: true,
		HttpMethod:     "GET",
		RequiredArgs:   []string{},
		OptionalArgs: []string{"count", "since_id", "max_id",
			"contributor_details", "include_entities"},
	},
}

package octogopher

type GistOwner struct {
	Login      string `json:"login"`
	Id         int64  `json:"id"`
	AvatarUrl  string `json:"avatar_url"`
	GravatarId string `json:"gravatar_id"`
	URL        string `json:"url"`
}

type GistFile struct{}

type GistFork struct{}

type HistoricalGist struct{}

type Gist struct {
	Files map[string]*GistFile
}

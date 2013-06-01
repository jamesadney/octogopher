package octogopher

import (
	"time"
)

const REPOS_URL = "/repositories"
const AUTHUSER_REPOS_URL = "/user/repos"

type Repo struct {

	// Fields for create and edit
	Name         string `json:"name"`
	Description  string `json:"description"`
	Homepage     string `json:"homepage"`
	Private      bool   `json:"private"`
	HasIssues    bool   `json:"has_issues"`
	HasWiki      bool   `json:"has_wiki"`
	HasDownloads bool   `json:"has_downloads"`

	// Fields for create only
	TeamId            int    `json:"team_id"`
	AutoInit          bool   `json:"auto_init"`
	GitignoreTemplate string `json:"gitignore_template"`

	// Fields for edit only
	DefaultBranch string `json:"default_branch"`

	// Viewable Fields
	Id    int `json:"id"`
	Owner struct {
		Login      string `json:"login"`
		Id         int    `json:"id"`
		AvatarUrl  string `json:"avatar_url"`
		GravatarId string `json:"gravatar_id"`
		Url        string `json:"url"`
	} `json:"owner"`
	FullName      string    `json:"full_name"`
	Fork          bool      `json:"fork"`
	Url           string    `json:"url"`
	HtmlUrl       string    `json:"html_url"`
	CloneUrl      string    `json:"clone_url"`
	GitUrl        string    `json:"git_url"`
	SshUrl        string    `json:"ssh_url"`
	SvnUrl        string    `json:"svn_url"`
	MirrorUrl     string    `json:"mirror_url"`
	Language      string    `json:"language"`
	Forks         int       `json:"forks"`
	ForksCount    int       `json:"forks_count"`
	Watchers      int       `json:"watchers"`
	WatchersCount int       `json:"watchers_count"`
	Size          int       `json:"size"`
	MasterBranch  string    `json:"master_branch"`
	OpenIssues    int       `json:"open_issues"`
	PushedAt      time.Time `json:"pushed_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Organization  struct {
		Login      string `json:"login"`
		Id         int    `json:"id"`
		AvatarUrl  string `json:"avatar_url"`
		GravatarId string `json:"gravatar_id"`
		Url        string `json:"url"`
		Type       string `json:"type"`
	} `json:"organization"`
	Parent *Repo `json:"parent"`
	Source *Repo `json:"source"`

	KeysUrl          string `json:"keys_url"`
	CollaboratorsUrl string `json:"collaborators_url"`
	TeamsUrl         string `json:"teams_url"`
	HooksUrl         string `json:"hooks_url"`
	IssueeventsUrl   string `json:"issue_events_url"`
	EventsUrl        string `json:"events_url"`
	AssigneesUrl     string `json:"assignees_url"`
	BranchesUrl      string `json:"branches_url"`
	TagsUrl          string `json:"tags_url"`
	BlobsUrl         string `json:"blobs_url"`
	GittagsUrl       string `json:"git_tags_url"`
	GittefsUrl       string `json:"git_refs_url"`
	TreesUrl         string `json:"trees_url"`
	StatusesUrl      string `json:"statuses_url"`
	LanguagesUrl     string `json:"languages_url"`
	StargazersUrl    string `json:"stargazers_url"`
	ContributorsUrl  string `json:"contributors_url"`
	SubscribersUrl   string `json:"subscribers_url"`
	SubscriptionUrl  string `json:"subscription_url"`
	CommitsUrl       string `json:"commits_url"`
	GitcommitsUrl    string `json:"git_commits_url"`
	CommentsUrl      string `json:"comments_url"`
	Issue_commentUrl string `json:"issue_comment_url"`
	ContentsUrl      string `json:"contents_url"`
	CompareUrl       string `json:"compare_url"`
	MergesUrl        string `json:"merges_url"`
	ArchiveUrl       string `json:"archive_url"`
	DownloadsUrl     string `json:"downloads_url"`
	IssuesUrl        string `json:"issues_url"`
	PullsUrl         string `json:"pulls_url"`
	MilestonesUrl    string `json:"milestones_url"`
	NotificationsUrl string `json:"notifications_url"`
	LabelsUrl        string `json:"labels_url"`
}

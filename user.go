package octogopher

import (
	"time"
)

const USERS_URL = "/users"
const AUTHUSER_URL = "/user"

type User struct {

	// Updateable Fields
	Name     string `json:"name"`
	Email    string `json:"email"`
	Blog     string `json:"blog"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Hireable bool   `json:"hireable"`
	Bio      string `json:"bio"`

	// Public, Read-Only Fields
	Login       string    `json:"login"`
	Id          int64     `json:"id"`
	AvatarUrl   string    `json:"avatar_url"`
	GravatarId  string    `json:"gravatar_id"`
	URL         string    `json:"url"`
	PublicRepos int       `json:"public_repos"`
	PublicGists int       `json:"public_gists"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	HtmlURL     string    `json:"html_url"`
	CreatedAt   time.Time `json:"created_at"`
	Type        string    `json:"type"`

	// Read-Only Fields for Authenticated User
	TotalPrivateRepos int `json:"total_private_repos"`
	OwnedPrivateRepos int `json:"owned_private_repos"`
	PrivateGists      int `json:"private_gists"`
	DiskUsage         int `json:"disk_usage"`
	Collaborators     int `json:"collaborators"`

	Plan struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}

// TODO:
func (u *User) Update() error {
	return nil
}

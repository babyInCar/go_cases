package main

import (
	"fmt"
	"net/http"
)

type GithubAccountInfo struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	Events_url        string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog:`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	TwitterName       string `json:"twitter_user"`
	PublicRepos       int    `json:"public_repos"`
	PublicGists       int    `json:"public_gists"`
	Folloers          int    `json:"followers"`
	Following         int    `json:"following"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func GithubUserStorager(name string) GithubAccountInfo {
	url := fmt.Sprintf("https://api.github.com/users/%s", name)
	//发起网络请求
	request, _ := http.NewRequest("GET", url, nil)
	// todo: header里面内容待完善
	request.Header.Add("User-Agent", "Mozilla/5.0")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return GithubAccountInfo{}
	}
	defer response.Body.Close()
	// 解析数据
}

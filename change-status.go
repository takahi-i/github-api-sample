package main

import (
	"fmt"

	"github.com/google/go-github/github"
	"code.google.com/p/goauth2/oauth"
)

func main() {

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: "foobar"},
	}
	client := github.NewClient(t.Client())

	repositories := &client.Repositories
	status, _, err := repositories.CreateStatus("takahi-i",
		"walter-github-sample",
		"foobar",
		&github.RepoStatus{
			State: github.String("failure"),
			TargetURL: github.String("https://github.com/walter-cd"),
			Description: github.String("うほっ")})
	if err != nil {
		fmt.Errorf("Repositories.CreateStatus returned error: %v", err)
	}
	fmt.Println(status)
}

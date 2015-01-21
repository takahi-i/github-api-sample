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

	commits, _, _ := client.Repositories.ListCommits("takahi-i", "walter-github-sample", &github.CommitsListOptions{})
	fmt.Println(commits)
}

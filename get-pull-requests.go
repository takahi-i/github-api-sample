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

	pullreqs, _, _ := client.PullRequests.List("takahi-i", "walter-github-sample", &github.PullRequestListOptions{})
	fmt.Println(pullreqs)
}

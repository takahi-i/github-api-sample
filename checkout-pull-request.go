package main

import (
	"os"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/google/go-github/github"
	"code.google.com/p/goauth2/oauth"
)

func main() {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: "foobar"},
	}

	client := github.NewClient(t.Client())

	pullreqs, _, _ := client.PullRequests.List("takahi-i", "walter-github-sample", &github.PullRequestListOptions{})
	pullreq := pullreqs[0]
	num := *pullreq.Number
	fmt.Printf("%d\n", num)

	out, err := exec.Command("git", "fetch", "origin", "refs/pull/" + strconv.Itoa(num) + "/head:pr_"+strconv.Itoa(num)).Output()

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
	
	fmt.Println(string(out))
	fmt.Println(num) 
	out2, err2 :=	exec.Command("git", "checkout", "pr_"+strconv.Itoa(num)).Output()
	fmt.Println(string(out2))	
        if err2 != nil {
                fmt.Println(err2)
                os.Exit(1)
        }

	

}


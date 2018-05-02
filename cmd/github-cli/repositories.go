package main

import (
	"fmt"
	"context"
	"github.com/google/go-github/github"
)

func reposList() {
	ctx := context.Background()
	client := github.NewClient(nil)

	// list public repositories for org "github"
	//opt := &github.RepositoryListByOrgOptions{Type: "public"}
	opt := &github.RepositoryListOptions{Type: "public"}
	//repos, _, err := client.Repositories.ListByOrg(ctx, "v2-labs", opt)
	repos, _, err := client.Repositories.List(ctx, "juvenal", opt)

	if err == nil {
		fmt.Println(repos)
	} else {
		fmt.Println(err)
	}
}

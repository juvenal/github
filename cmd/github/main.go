package main

import (
	//"os"
	"fmt"
	//"flag"
	"context"
	"github.com/google/go-github/github"
)

func main() {

	// Check presence of command and verb
	//if ! (len(os.Args) < 2) && ! ((len(os.Args) == 2 && os.Args[1] ==  "-h" || os.Args[1] == "--help"))  {
	//    fmt.Println("Process the parameters given...")
	//    return
	//} else {
	//    helpBaseUsage(os.Args)
	//    return
	//}

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

	// Check if you asked for the '[-]-help' flag
	//helpPtr := flag.Bool("help", false, "Presents this help instructions")

	//flag.Parse()

	//fmt.Println(*helpPtr)

	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]
	//command := os.Args[1]

	//fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)
	//fmt.Println(command)

	// Branch to process valid command and verbs
	//switch command {
	//	case "login": {

	//	}
	//	case "repo": {

	//	}
	//	default:
	//}
}

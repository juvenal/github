package main

import (
	"os"
	"fmt"
	"flag"
)

func main() {
	// Prepare verb processing
	orgVerb := flag.NewFlagSet("org", flag.ExitOnError)
	repoVerb := flag.NewFlagSet("repo", flag.ExitOnError)
	loginVerb := flag.NewFlagSet("login", flag.ExitOnError)

	// Options for login verb
	loginTokenPtr := loginVerb.String("token", "", "Github `token_ID` (Required).")
	loginResetPtr := loginVerb.Bool("reset", false, "Reset stored Github token ID.")

	// Options for repo verb
	repoListAllPtr := repoVerb.Bool("list", false, "Return list of available repositories.")
	repoListOrgPtr := repoVerb.String("org", "", "Inform the Organization name.")
	repoNamePtr := repoVerb.String("name", "", "Inform the name of the repository to access.")
	repoCloneURLPtr := repoVerb.String("cloneURL", "", "Return the informed repository clone URL.")
	repoCloneTypePtr := repoVerb.String("cloneType", "https", "Inform the type of repository clone URL returned.")

	// Options for org verb

	// Check for arguments given
	if (len(os.Args) < 2) {
		fmt.Printf("github-cli: ")
		fmt.Println("missing command. Available ones are:")
		fmt.Println("\n- 'login' with parameters:")
		loginVerb.PrintDefaults()
		fmt.Println("\n- 'repo' with paramters:")
		repoVerb.PrintDefaults()
		fmt.Println("\n- 'org' with parameters:")
		orgVerb.PrintDefaults()
		os.Exit(1)
	}
	//	helpBaseUsage(os.Args)
	//} else if (len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help")) {
	//	helpBaseUsage(os.Args)
	//} else {
	//	command := os.Args[1]

	switch os.Args[1] {
		case "login": {
			fmt.Println("login verb given.")
			loginVerb.Parse(os.Args[2:])
			fmt.Printf("Options received:\n- Token: %s\n- Reset: %t\n", *loginTokenPtr, *loginResetPtr)
			fmt.Println()
		}
		case "repo": {
			fmt.Println("repo verb given.")
			repoVerb.Parse(os.Args[2:])
			fmt.Printf("Options received:\n- List: %t\n- Org: %s\n- Name: %s\n- URL: %s\n- Type: %s\n", *repoListAllPtr, *repoListOrgPtr, *repoNamePtr, *repoCloneURLPtr, *repoCloneTypePtr)
			reposList()
		}
		case "org": {
			fmt.Println("orgs verb given.")
			orgVerb.Parse(os.Args[2:])
			orgsList()
		}
		default: {
			fmt.Printf("github-cli: ")
			fmt.Println("unrecognized command. Available ones are:")
			fmt.Println("\n- 'login' with parameters:")
			loginVerb.PrintDefaults()
			fmt.Println("\n- 'repo' with paramters:")
			repoVerb.PrintDefaults()
			fmt.Println("\n- 'org' with parameters:")
			orgVerb.PrintDefaults()
			os.Exit(1)
		}
	}

	// Check presence of command and verb
	//if ! (len(os.Args) < 2) && ! ((len(os.Args) == 2 && os.Args[1] ==  "-h" || os.Args[1] == "--help"))  {
	//    fmt.Println("Process the parameters given...")
	//    return
	//} else {
	//    helpBaseUsage(os.Args)
	//    return
	//}

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

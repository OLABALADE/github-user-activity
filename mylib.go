package mylib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

type Activity struct {
	Type string `json:"type"`

	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`

	Payload struct {
		Action string `json:"action"`

		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`

		RefType string `json:"ref_type"`
	} `json:"payload"`
}

func GetActivity(user *string) (*[]Activity, error) {

	url := fmt.Sprintf("https://api.github.com/users/%s/events", *user)
	req, err := http.Get(url)

	if req.StatusCode == http.StatusNotFound {
		log.Fatalf("User %s not found", *user)
	}

	if err != nil {
		return &[]Activity{}, err
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		return &[]Activity{}, err
	}

	var data []Activity

	if err = json.Unmarshal(body, &data); err != nil {
		return &[]Activity{}, err
	}

	return &data, nil
}

func DisplayActivity(a *[]Activity) {
	for _, act := range *a {

		switch act.Type {
		case "CreateEvent":
			fmt.Printf("Created %s -> %s\n", act.Payload.RefType, act.Repo.Name)

		case "DeleteEvent":
			fmt.Printf("Deleted %s in %s\n", act.Payload.RefType, act.Repo.Name)

		case "ForkEvent":
			fmt.Printf("Forked  %s\n", act.Repo.Name)

		case "IssueEvent":
			fmt.Printf("%s an issue in %s\n", act.Payload.Action, act.Repo.Name)

		case "PushEvent":
			fmt.Printf("%d commit to %s\n", len(act.Payload.Commits), act.Repo.Name)

		case "PullRequestEvent":
			fmt.Printf("%s a pull request in %s\n", act.Payload.Action, act.Repo.Name)

		case "WatchEvent":
			fmt.Printf("Starred  %s\n", act.Repo.Name)

		default:
			reg, err := regexp.Compile(`p*Event`)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%s in %s\n", reg.ReplaceAllString(act.Type, ""), act.Repo.Name)
		}
	}
}

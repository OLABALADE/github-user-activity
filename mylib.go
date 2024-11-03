package mylib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Activity struct {
	Type string `json:"type"`

	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`

	Payload struct {
		Action  string `json:"action"`
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
        RefType string `json:"ref_type"`
	} `json:"payload"`
}

func GetActivity(user *string) *[]Activity {

	url := fmt.Sprintf("https://api.github.com/users/%s/events", *user)
	req, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data []Activity

	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	return &data
}

func DisplayActivity(a *[]Activity) {
	for _, act := range *a {
		switch act.Type {
		case "PushEvent":
			fmt.Printf("%d commit to %s\n", len(act.Payload.Commits), act.Repo.Name)
		case "IssueEvent":
			fmt.Printf("%s an issue in %s\n", act.Payload.Action, act.Repo.Name)
		case "CreateEvent":
			fmt.Printf("Created %s in %s\n", act.Payload.RefType, act.Repo.Name)
		default:
			fmt.Printf("%s in %s\n", act.Type, act.Repo.Name)
		}
	}
}

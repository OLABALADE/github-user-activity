package main

import (
	"flag"
	"github-user-activity"
	"log"
)

func main() {
	user := flag.String("user", "", "User's github name")
	flag.Parse()

	activity, err := mylib.GetActivity(user)

	if err != nil {
		log.Fatal(err)
	}

	mylib.DisplayActivity(activity)
}

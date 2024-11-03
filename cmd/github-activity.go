package main

import (
	"flag"
    "github-user-activity"
)

func main() {
	user := flag.String("user", "", "User's github name")
    flag.Parse()

    activity := mylib.GetActivity(user)

    mylib.DisplayActivity(activity)



}

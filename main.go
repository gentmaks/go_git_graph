package main

import (
	"flag"
	"fmt"
)

func stats(email string) {
	fmt.Println("stats")
}

func add(folder string) {
	fmt.Println("scan")
}

func main(){
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "Add a new folder to scan for your git repositories")
	flag.StringVar(&email, "email", "your@email.com", "The email to scan")
	flag.Parse()
	if folder != "" {
		add(folder)
		return
	}
	stats(email)
}

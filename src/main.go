package main

import (
	"func/github"
	"os"
	"log"
	"fmt"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %s %s %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
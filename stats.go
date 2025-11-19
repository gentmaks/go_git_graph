package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const outOfRange = 99999
const daysInLastSixMonths = 183
const weeksInLastSixMonths = 26

func stats (email string) {
	commits := processRepositories(email)
	printCommitsStats(commits)
}

func processRepositories(email string) map[int]int {
	filePath := getDotFilePath()
	repos := parseFileLinesToSlice(filePath)	
	daysInMap := daysInLastSixMonths

	commits := make(map[int]int, daysInMap)
	for i := daysInMap; i > 0; i-- {
		commits[i] = 0	
	}
	
	for _, path := range(repos) {
		commits = fillCommits(email, path, commits)
	}
	return commits
}

func fillCommits (email string, path string, commits map[int]int) map[int]int{
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}

	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}

	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}

	offset := calcOffset()
    err = iterator.ForEach(func(c *object.Commit) error {
        daysAgo := countDaysSinceDate(c.Author.When) + offset

        if c.Author.Email != email {
            return nil
        }

        if daysAgo != outOfRange {
            commits[daysAgo]++
        }

        return nil
    })
    if err != nil {
        panic(err)
    }

    return commits
}

func printCommitsStats(commits map[int]int) {
}


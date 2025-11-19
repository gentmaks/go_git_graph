package main

import (
	"fmt"
	"time"
	"sort"
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
	keys := sortMapIntoSlice(commits)
	cols := buildCols(keys, commits)
	printCells(cols)
}

func calcOffset() int {
    var offset int
    weekday := time.Now().Weekday()

    switch weekday {
    case time.Sunday:
        offset = 7
    case time.Monday:
        offset = 6
    case time.Tuesday:
        offset = 5
    case time.Wednesday:
        offset = 4
    case time.Thursday:
        offset = 3
    case time.Friday:
        offset = 2
    case time.Saturday:
        offset = 1
    }

    return offset
}

func getBeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return startOfDay
}

func countDaysSinceDate(date time.Time) int {
	days := 0
	now := getBeginningOfDay(time.Now())
	for date.Before(now) {
		date = date.Add(time.Hour * 24)
		days++
		if days > daysInLastSixMonths {
			return outOfRange
		}
	}
	return days
}
func sortMapIntoSlice(m map[int]int) []int {
	var keys []int
	for key, _ := range(m) {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
} 
func buildCols(keys []int, commits map[int]int) []int {
	return nil 
}
func printCells(cols []int) {
}

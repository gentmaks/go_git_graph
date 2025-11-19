package main

import (
	"fmt"
	"strings"
	"log"
	"os"
	"os/user"
	"bufio"
	"github.com/go-git/go-git/v5"
)

func scan(path string) {
    fmt.Printf("Found folders:\n\n")
    repos := recursiveScanFolder(path)

    // Filter out non-Git folders just in case
    validRepos := make([]string, 0, len(repos))
    for _, repoPath := range repos {
        if _, err := git.PlainOpen(repoPath); err == nil {
            validRepos = append(validRepos, repoPath)
            fmt.Println(repoPath)
        }
    }

    filePath := getDotFilePath()
    addNewSliceElementToFile(filePath, validRepos)
    fmt.Printf("\n\nSuccessfully added %d repos\n\n", len(validRepos))
}

func recursiveScanFolder(path string) []string {
	return scanGitFolders(make([]string, 0), path)
}

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dotFile := usr.HomeDir + "/.gogitlocalstats"
	return dotFile
}

func addNewSliceElementToFile(filePath string, newRepos []string) {
    existingRepos := parseFileLinesToSlice(filePath)
    repos := joinSlices(newRepos, existingRepos)
    dumpStringsSliceToFile(repos, filePath)
}

func parseFileLinesToSlice(filePath string) []string {
    f, err := os.Open(filePath) // read-only
    if err != nil {
        if os.IsNotExist(err) {
            return []string{} // file doesn’t exist yet → empty slice
        }
        log.Fatal(err)
    }
    defer f.Close()

    var lines []string
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return lines
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range(new) {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}		
	return existing
}

func dumpStringsSliceToFile(repos []string, filePath string) {
    content := strings.Join(repos, "\n") + "\n" // add newline at end
    if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
        log.Fatal(err)
    }
}

func sliceContains(slice []string, val string) bool {
	for _, v := range(slice) {
		if v == val {
			return true
		}
	}	
	return false
}

func openFile(filePath string) *os.File {
    f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0755)
    if err != nil {
        if os.IsNotExist(err) {
            _, err = os.Create(filePath)
            if err != nil {
                panic(err)
            }
        } else {
            panic(err)
        }
    }

    return f
}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")
	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.ReadDir(-1)
	if err != nil {
		log.Fatal(err)
	}
	var path string;
	for _, file :=  range(files) {
		if file.IsDir(){ 
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/")
				// fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}
	return folders
}

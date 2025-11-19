package main

import (
	"flag"
	"fmt"
	"strings"
	"log"
	"os"
	"os/user"
	"bufio"
	"io"
)
func scan(path string) {
	fmt.Printf("Found folders: \n\n")
	repos := recursiveScanFolder(path)
	for _, file := range(repos) {
		fmt.Println(file)
	}
	filePath := getDotFilePath()
	fmt.Println(filePath)
	addNewSliceElementToFile(filePath, repos)
	fmt.Printf("\n\nSuccessfully added\n\n")
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
	f := openFile(filePath)
	defer f.Close()
	
	var lines []string
	scanner := bufio.NewScanner(f)
	for {
		if !scanner.Scan() {
			break
		}
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
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

func dumpStringsSliceToFile(repos []string , filePath string) {
	content := strings.Join(repos, "\n")
	if err := os.WriteFile(filePath, []byte(content), 0755); err != nil {
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

func stats(email string) {
	fmt.Println("stats")
}

func main(){
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "Add a new folder to scan for your git repositories")
	flag.StringVar(&email, "email", "your@email.com", "The email to scan")
	flag.Parse()
	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}

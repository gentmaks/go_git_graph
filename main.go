package main

import (
	"flag"
	"fmt"
	"strings"
	"log"
	"os"
)
func scan(path string) {
	fmt.Printf("Found folders: \n\n")
	folders := scanGitFolders(make([]string, 0), path)
	for _, file := range(folders) {
		fmt.Println(file)
	}
	// repos := resursiveScanFolder(path)
	// filePath := getDotFilePath()
	// addNewSliceElementToFile(filePath, repos)
	fmt.Printf("\n\nSuccessfully added\n\n")
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
				fmt.Println(path)
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

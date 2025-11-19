package main

import (
	"flag"
)
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
// 	m := map[int]int {
// 		1: 10,
// 		5: 20,
// 		4: 20,
// 		3: 90,
// 		9: 20,
// 		7: 70,
// 	}
// 	sorted_keys := sorting_map(m)
// 	for _, key := range(sorted_keys){
// 		fmt.Println(key)
// 	}
// }
//
// func sorting_map(m map[int]int) []int {
// 	var keys []int
// 	for key, _ := range(m) {
// 		keys = append(keys, key)
// 	}
// 	sort.Ints(keys)
// 	return keys
} 

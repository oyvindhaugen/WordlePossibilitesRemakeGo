package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var words []string

func main() {
	list, err := http.Get("https://gist.githubusercontent.com/cfreshman/a7b776506c73284511034e63af1017ee/raw/845966807347a7b857d53294525263408be967ce/wordle-nyt-answers-alphabetical.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer list.Body.Close()
	if list.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", list.StatusCode, list.Body)
	}
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(list.Body)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	for {
		fmt.Println("Which letter do you wish to remove?")
		letter := ""
		fmt.Scanln(&letter)
		if letter != "wq" {
			for i, wn := range words {
				if strings.Contains(wn, letter) {
					words = remove(words, i)
				}
			}
			fmt.Println(len(words))
		} else {
			for _, word := range words {
				fmt.Println(word)
			}
			fmt.Printf("You had %v possibilities left.", len(words))
			break
		}
	}
}
func remove(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = ""
	s = s[:len(s)-1]
	return s
}

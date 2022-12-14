package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var words []string
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
	fmt.Println("Any letter you input will remove all words containing that letter from the list.")
	fmt.Println("When you wish to print out the list, just write 'wq'.")
	letter := ""
	x := 0
	for {
		fmt.Scan(&letter)
		if letter == "wq" {
			break
		}
		for i := 0; i < len(words); i++ {
			if strings.ContainsAny(words[i], letter) {
				copy(words[i:], words[i+1:])
				words[len(words)-1] = "no"
				words = words[:len(words)-1]
				i = 0

			}
		}
		fmt.Printf("Removed all words with %s\n", letter)
	}
	for num, _ := range words {
		if words[num] == "no" {
			x++
		}
	}
	fmt.Printf("There are %v elements left\n", len(words[:len(words)-x]))
	fmt.Printf("Do you want to print it out? y/n\n")
	ans := ""
	fmt.Scan(&ans)
	if ans == "y" {
		for _, word := range words {
			if word != "no" {
				fmt.Println(word)
			}
		}
	} else {
		os.Exit(0)
	}
}

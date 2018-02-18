package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"
	"unicode/utf8"
)

func init() {
	// Seed RNG since it is not normally seeded
	rand.Seed(time.Now().Unix())

	// Basic flags
	flag.IntVar(&numWords, "n", 4, "number of words per password")
	flag.IntVar(&numPass, "p", 1, "number of passwords")
	flag.Parse()
}

var (
	numWords int
	numPass  int
)

func main() {
	// Not portable right now
	f, _ := os.Open("/usr/share/dict/words")
	s := bufio.NewScanner(f)

	var words []string
	for s.Scan() {
		word := s.Text()

		// Avoid proper words and numbers
		first, _ := utf8.DecodeRuneInString(word)
		if unicode.IsLower(first) {
			words = append(words, word)
		}
	}

	for i := 0; i < numPass; i++ {
		for j := 0; j < numWords; j++ {
			fmt.Printf("%s ", words[rand.Intn(len(words))])
		}
		fmt.Println()
	}
}

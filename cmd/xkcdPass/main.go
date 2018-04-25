package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	numPass  int
	numWords int
	verbose  bool
)

func init() {
	// Seed RNG since it is not normally seeded
	rand.Seed(time.Now().Unix())

	// Basic flags
	flag.IntVar(&numPass, "p", 1, "number of passwords to generate")
	flag.IntVar(&numWords, "n", 4, "number of words per password")
	flag.BoolVar(&verbose, "v", false, "verbose output to stderr")
	flag.Parse()
}

func main() {
	// Check common words file locations
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		// Try again
		f, err = os.Open("/usr/dict/words")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to find a words file.")
			os.Exit(1)
		}
	}
	s := bufio.NewScanner(f)

	// Load words into slice for random access
	var words []string
	for s.Scan() {
		word := s.Text()

		// Avoid proper words and numbers
		first, _ := utf8.DecodeRuneInString(word)
		if unicode.IsLower(first) {
			words = append(words, word)
		}
	}

	// Password generation
	for i := 0; i < numPass; i++ {
		for j := 0; j < numWords; j++ {
			fmt.Printf("%s ", words[rand.Intn(len(words))])
		}
		fmt.Println()
	}

	// Entropy calculation
	if verbose {
		entropy := math.Round(math.Log2(float64(len(words))) * float64(numWords))
		fmt.Fprintf(os.Stderr, "Generated %d password(s) with ~%.f bits of entropy each.\n",
			numPass, entropy)
	}
}

package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
	"unicode"
	"unicode/utf8"
)

var (
	passCount int
	wordCount int
	verbose   bool
)

func init() {
	flag.IntVar(&passCount, "p", 1, "number of passwords to generate")
	flag.IntVar(&wordCount, "w", 4, "number of words per password")
	flag.BoolVar(&verbose, "v", false, "verbose output to stderr")
	flag.Parse()
}

func main() {
	// Check common words file locations
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		// Try the other common location.
		f, err = os.Open("/usr/dict/words")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to find a words file.")
			os.Exit(1)
		}
	}
	s := bufio.NewScanner(f)

	// Load dictionary into slice for random access.
	var dict []string
	for s.Scan() {
		w := s.Text()

		// Avoid using proper nouns and numbers.
		first, _ := utf8.DecodeRuneInString(w)
		if unicode.IsLower(first) {
			dict = append(dict, w)
		}
	}

	// Store dictionary length in a big.Int for use in the CSPRNG.
	var dictLen big.Int
	dictLen.SetInt64(int64(len(dict)))

	// Generate passwords using the crypto/rand CSPRNG.
	for i := 0; i < passCount; i++ {
		for j := 0; j < wordCount; j++ {
			n, err := rand.Int(rand.Reader, &dictLen)
			if err != nil {
				fmt.Fprint(os.Stderr, "CSPRNG error: ", err)
			}

			fmt.Printf("%s ", dict[n.Int64()])
		}
		fmt.Println()
	}

	// Calculate the number of bits of entropy for a generated password.
	if verbose {
		entropy := math.Round(math.Log2(float64(len(dict))) * float64(wordCount))
		fmt.Fprintf(os.Stderr, "Generated %d password(s) with ~%.f bits of entropy each.\n",
			passCount, entropy)
	}
}

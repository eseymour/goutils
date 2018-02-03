package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// The frequencies of byte values in fortunes.
var byteFreqs [1 << 8]int

// The distribution of byte values in fortunes.
var byteDist [1 << 8]float64

// Check prints an error message and exits with 1 if err != nil. It does nothing
// otherwise.
func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func main() {
	path, err := exec.LookPath("fortune")
	check(err)

	// Match on all fortune lists with regex matching anything.
	cmd := exec.Command(path, "-m", "[\\s\\S]*")
	check(err)

	cmdOut, err := cmd.StdoutPipe()
	check(err)
	check(cmd.Start())

	// Account for additive smoothing.
	count := 256
	scanner := bufio.NewScanner(cmdOut)
	for scanner.Scan() {
		line := scanner.Bytes()

		if bytes.Equal(line, []byte("%")) {
			// Ignore strfile delimiter line.
			continue
		}

		// Increment count and frequencies.
		count += len(line)
		for _, b := range line {
			byteFreqs[b]++
		}
	}
	check(scanner.Err())

	// Fortune exits with one for some reason. Checking for that condition is too
	// cumbersome, so we ignore Wait's return value.
	cmd.Wait()

	// Calculate distribution.
	for i := range byteFreqs {
		// Apply additive smoothing.
		byteFreqs[i]++

		byteDist[i] = float64(byteFreqs[i]) / float64(count)
	}

	// Print results in a form that can be inserted into a Go program.
	fmt.Printf("var byteFreqs = %#v\n", byteFreqs)
	fmt.Printf("var byteDist = %#v\n", byteDist)
}

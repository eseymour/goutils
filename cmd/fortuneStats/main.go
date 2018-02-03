package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var byteFreqs [1 << 8]int
var byteDist [1 << 8]float64

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func main() {
	path, err := exec.LookPath("fortune")
	check(err)

	// Match on all fortune lists with regex matching anything
	cmd := exec.Command(path, "-am", ".*")
	check(err)

	cmdOut, err := cmd.StdoutPipe()
	check(err)

	check(cmd.Start())

	count := 256 // Account for additive smoothing
	scanner := bufio.NewScanner(cmdOut)
	for scanner.Scan() {
		line := scanner.Bytes()

		// Ignore strfile delimiter line
		if bytes.Equal(line, []byte("%")) {
			continue
		}

		count += len(line)
		for _, b := range line {
			byteFreqs[b]++
		}
	}
	check(scanner.Err())
	cmd.Wait()

	// Calculate distribution
	for i := range byteFreqs {
		// Additive smoothing
		byteFreqs[i]++

		byteDist[i] = float64(byteFreqs[i]) / float64(count)
	}

	fmt.Printf("var byteFreqs = %#v\n", byteFreqs)
	fmt.Printf("var byteDist = %#v\n", byteDist)
}

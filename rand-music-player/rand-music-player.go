/*
Play a random song in a specified directory.
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	// Ensure a single argument is sent in
	if len(os.Args) != 2 {
		fmt.Println("Expected 1 argument. Specify the directory.")
		os.Exit(1)
	}

	// Build a list of all files in the specified directory
	var files []string

	root := os.Args[1]
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path != root {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	// Set the random seed
	rand.Seed(time.Now().UnixNano())

	// Get a random song
	song := files[rand.Intn(len(files))]
	fmt.Printf("Now Playing: %s\n", song)

	// First, ensure no other songs are playing
	exec.Command("killall", "omxplayer.bin").Run()

	// Begin playing
	cmd := exec.Command("omxplayer", song)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
}

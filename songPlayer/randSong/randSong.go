/*
Play a random song in a specified directory.
*/

package randSong

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// Get a random song from the given directory
func GetRandomSong(root string) string {
	songs := getSongs(root)
	// Set the random seed
	rand.Seed(time.Now().UnixNano())

	// Get a random song
	song := songs[rand.Intn(len(songs))]
	return song
}

// Get available songs
func getSongs(root string) []string {
	// Build a list of all files in the specified directory
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path != root {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Failed to read the directory?")
		panic(err)
	}
	return files
}

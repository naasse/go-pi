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
	"strings"
)

// Check args.
// 0: Root directory to find a song in.
// 1: [optional] BLUETOOTH (case-insensitive) will play via omxplayer alsa output
// returns whether or not bluetooth is set.
func checkArgs() bool {
	// Ensure a directory argument is sent in
	if len(os.Args) < 2 {
		fmt.Println("Expected 1 or 2 arguments. Specify the directory. BLUETOOTH flag is optional.")
		os.Exit(1)
	}
	blue := false
	if len(os.Args) > 2 {
		blue = strings.EqualFold(os.Args[2], "BLUETOOTH")
	}
	return blue
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

// Kill any existing omxplayers
func killOmx() {
	err := exec.Command("killall", "omxplayer.bin").Run()
	if err != nil {
		fmt.Println("Failed to kill players. This may be expected, if none exist.")
	}
	// Sleep for 2 seconds in case there's any process cleanup not complete
	time.Sleep(2 * time.Second)
}

// Get a random song
func getRandomSong(songs []string) string {
	// Set the random seed
	rand.Seed(time.Now().UnixNano())

	// Get a random song
	song := songs[rand.Intn(len(songs))]
	fmt.Printf("Now Playing: %s\n", song)
	return song
}

// Play the given song
func playSong(song string, blue bool) {
	args := []string{"-b", "--vol", "-500", song}
	cmd := exec.Command("omxplayer", args...)
	if blue {
		args = append(args, "-o")
		args = append(args, "alsa")
		cmd = exec.Command("omxplayer", args...)
	}
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to start song.")
		killOmx()
		panic(err)
	}
}

// Main method
func main() {
	// Check args, determine audio output device
	blue := checkArgs()
	// Find list of available songs
	songs := getSongs(os.Args[1])

	// Ensure no other songs are playing
	killOmx()

	// Play a random song
	song := getRandomSong(songs)

	// Play the song
	playSong(song, blue)

	// Begin playing
	//cmd := exec.Command("omxplayer", "-b", "--vol", "-500", song)
	//if blue {
	//	cmd = exec.Command("omxplayer", "-b", "--vol", "-500", "-o", "alsa", song)
	//}
	//err = cmd.Start()
	//if err != nil {
	//	fmt.Printf("Failed to start song.")
	//	panic(err)
	//}
}

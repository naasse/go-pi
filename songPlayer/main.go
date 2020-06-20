/*
Play a specified song from disk
*/

package main

import (
	"fmt"
	"go-pi/songPlayer/randSong"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Check args.
// 0: File to play/directory to find random song.
func checkArgs() {
	// Ensure a file argument is sent in
	if len(os.Args) < 2 {
		fmt.Println("Expected 1  argument. Specify the song file, or a directory for random song.")
		os.Exit(1)
	}
}

// Kill any existing vlc processes
func killVlc() {
	exec.Command("killall", "vlc").Run()
}

// Play the given song
func playSong(song string) {
	fmt.Printf("Now Playing: %s\n", song)
	cmd := exec.Command("vlc", song)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to start song.")
		killVlc()
		panic(err)
	}
	time.Sleep(1 * time.Second)
}

// Main method
func main() {
	// Check args
	checkArgs()

	song := os.Args[1]
	// Check if a song or directory was passed in.
	if !strings.HasSuffix(song, ".mp3") {
		// It's not an mp3. Get a random song from the directory.
		song = randSong.GetRandomSong(song)
	}

	// Ensure no other songs are playing
	killVlc()

	// Play the song
	playSong(song)
}

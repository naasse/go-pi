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
// 1: [optional] BLUETOOTH (case-insensitive) will play via omxplayer alsa output
// returns whether or not bluetooth is set.
func checkArgs() bool {
	// Ensure a file argument is sent in
	if len(os.Args) < 2 {
		fmt.Println("Expected 1 or 2 arguments. Specify the song file, or a directory for random song.")
		fmt.Println("BLUETOOTH flag is optional.")
		os.Exit(1)
	}
	blue := false
	if len(os.Args) > 2 {
		blue = strings.EqualFold(os.Args[2], "BLUETOOTH")
		if blue {
			fmt.Println("Will play via Bluetooth.")
		} else {
			fmt.Println("Will play via analag audio.")
		}
	}
	return blue
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

// Play the given song
func playSong(song string, blue bool) {
	args := []string{"omxplayer", "-b", "--vol", "-500", song}
	if blue {
		args = append(args, "-o")
		args = append(args, "alsa")
	}
	fmt.Printf("Now Playing: %s\n", song)
	cmd := exec.Command("omxplayer", args...)
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

	song := os.Args[1]
	// Check if a song or directory was passed in.
	if !strings.HasSuffix(song, ".mp3") {
		// It's not an mp3. Get a random song from the directory.
		song = randSong.GetRandomSong(song)
	}

	// Ensure no other songs are playing
	killOmx()

	// Play the song
	playSong(song, blue)
}

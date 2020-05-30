/*
Toggle current on the given GPIO pin.
This program will ensure that translates to a physical pin.
Intended to be used in a blinking LED example.
*/

package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
	"strconv"
)

var (
	// Use GPIO pin 17, corresponds to physical pin 11 on the pi
	pin = rpio.Pin(17)

	// Supported pins. Map the GPIO pin to the physical pin.
	pins = make(map[int]int)
)

func initPins() {
	pins[17] = 11
}

func main() {
	// Ensure a single argument is sent in
	if len(os.Args) != 2 {
		fmt.Println("Expected 1 argument. Specify the GPIO pin to blink.")
		os.Exit(1)
	}

	// Ensure the argument is numeric.
	pinNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number.\n", os.Args[1])
		os.Exit(1)
	}

	// Initialize the pin mappings.
	initPins()

	// Ensure the mcu pin is in our pins mappings
	physicalPin, ok := pins[pinNumber]
	if !ok {
		fmt.Printf("%d is not a valid pin.\n", pinNumber)
		os.Exit(1)
	}

	fmt.Printf("%d is the physical pin.\n", physicalPin)

	// The pin to toggle
	pin = rpio.Pin(pinNumber)

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Output()

	fmt.Println("Blinking!")
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
	fmt.Println("Done.")
}

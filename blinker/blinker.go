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
	"math/rand"
)

var (
	// Use GPIO pin 17, corresponds to physical pin 11 on the pi
	pin = rpio.Pin(17)

	// Supported pins. Map the GPIO pin to the physical pin.
	pins = make(map[int]int)
)

func initPins() {
	pins[12] = 32
	pins[16] = 36
	pins[17] = 11
	pins[20] = 38
	pins[21] = 40
	pins[27] = 13
}

func main() {
	// Ensure a single argument is sent in
	if len(os.Args) == 1 {
		fmt.Println("Expected 1 or more arguments. Specify the GPIO pins to blink.")
		os.Exit(1)
	}

	// Initialize the pins to blink
	var togglePins = make([]rpio.Pin, len(os.Args) - 1 )

	// Initialize the pin mappings
	initPins()

	for i:= 1; i < len(os.Args); i++ {
		pinNumber, err:= strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Printf("%q is not a number.\n", os.Args[i])
			os.Exit(1)
		}

		// Ensure the GPIO pin is in our pin mappings
		_, ok := pins[pinNumber]
		if !ok {
			fmt.Printf("%d is not a valid pin.\n", pinNumber)
			os.Exit(1)
		}

		// Add it to the array of pins to blink
		togglePins[i - 1] = rpio.Pin(pinNumber)
	}

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pins to output mode
	for i:= 0; i < len(togglePins); i++ {
		togglePins[i].Output()
	}

	fmt.Println("Blinking!")
	for x := 0; x < 20; x++ {
		togglePin := togglePins[rand.Intn(len(togglePins))]
		state := togglePin.Read()
		fmt.Printf("Pin %d: %t -> %t\n", togglePin, state != 0, state == 0)
		togglePin.Toggle()
		time.Sleep(time.Second / 2)
	}
	fmt.Println("Done.")
}

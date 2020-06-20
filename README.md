# go-pi

Learning projects to work with Raspberry pi GPIO with Go

## Blinker

Toggle a given GPIO pin on (high/low) to make LEDs blink.

### Dependencies

* [go-rpio](github.com/stianeikeland/go-rpio)

### Usage

`go run blinker/blinker.go {pin}`

## rfid

Begin listening on RFID (RFC522). When a known RFID UID is detected, play the song that it maps to via omxplayer over a Bluetooth speaker.

### Usage

`python2 rfid/read.py`

TODO: Convert to Go.

## songPlayer

Play a given song, or a random song in a given directory. Uses omxplayer, which comes native on Raspberry Pi OS. Supports playing via Bluetooth or analog audio out.

### Usage

`go run songPlayer/main.go {song | directory} {BLUETOOTH?}`

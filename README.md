# go-pi

Learning projects to work with Raspberry pi GPIO with Go

## Blinker

Toggle a given GPIO pin on (high/low) to make LEDs blink.

### Dependencies

* [go-rpio](github.com/stianeikeland/go-rpio)

### Usage

`go run blinker/blinker.go {pin}`

## rfid

Begin listening on RFID (RFC522). When a known RFID UID is detected, play the song that it maps to via VLC.

Volume controls should be controlled by the device, defaults, and VLC should be marked to Play and exit.

### Dependencies

* [psutil](https://pypi.org/project/psutil/)

### Usage

`python2 rfid/read.py`

The primary intent for this script is to run on boot.

Modify or create `~/.config/lxsession/LXDE-pi/autostart` as follows:

```
@lxpanel --profile LXDE-pi
@pcmanfm --desktop --profile LXDE-pi
@xscreensave -no-splash
@lxterminal --comand="python2 /home/pi/go/src/go-pi/rfid/read.py"
```

Only the final line is necessary, but since this will override `/etc/xdg/lxsession/LXDE-pi/autostart`, you probably want to copy that file's contents over.

### TODO

Convert to Go.

## songPlayer

Play a given song, or a random song in a given directory. Uses VLC media player, which comes native on Raspberry Pi OS.

### Usage

`go run songPlayer/main.go {song | directory}`

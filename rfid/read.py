#!/usr/bin/env python
# -*- coding: utf8 -*-
#
#    Copyright 2014,2018 Mario Gomez <mario.gomez@teubi.co>
#
#    This file is part of MFRC522-Python
#    MFRC522-Python is a simple Python implementation for
#    the MFRC522 NFC Card Reader for the Raspberry Pi.
#
#    MFRC522-Python is free software: you can redistribute it and/or modify
#    it under the terms of the GNU Lesser General Public License as published by
#    the Free Software Foundation, either version 3 of the License, or
#    (at your option) any later version.
#
#    MFRC522-Python is distributed in the hope that it will be useful,
#    but WITHOUT ANY WARRANTY; without even the implied warranty of
#    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#    GNU Lesser General Public License for more details.
#
#    You should have received a copy of the GNU Lesser General Public License
#    along with MFRC522-Python.  If not, see <http://www.gnu.org/licenses/>.
#

import RPi.GPIO as GPIO
import MFRC522
import signal
import subprocess
import time
import psutil

continue_reading = True

current_song = None

# Capture SIGINT for cleanup when the script is aborted
def end_read(signal,frame):
    global continue_reading
    print "Ctrl+C captured, ending read."
    continue_reading = False
    GPIO.cleanup()

# Hook the SIGINT
signal.signal(signal.SIGINT, end_read)

# Create an object of the class MFRC522
MIFAREReader = MFRC522.MFRC522()

# Welcome message
print "Welcome to the MFRC522 data read example."
print "Press Ctrl-C to stop."

# Dictionary of known RFID UID
uids = {
    "136.4.52.177": "/home/pi/Music/Evie/The_Wheels_On_The_Bus.mp3", # Elmo
    "136.4.68.177": "/home/pi/Music/Evie/Twinkle_Twinkle_Little_Star.mp3", # Cookie Monster
    "136.4.158.176": "/home/pi/Music/Evie/Baby_Shark.mp3", # Oscar
    "136.4.148.176": "/home/pi/Music/Evie/Head_Shoulders_Knees_and_Toes.mp3", # Grover
    "136.4.140.176": "/home/pi/Music/Evie/Old_MacDonald_Had_A_Farm.mp3" # Abby
}

# The last time an RFID was read
last_scan = 0

# Minimum time between RFID reads (seconds)
cooldown = 2

# This loop keeps checking for chips. If one is near it will get the UID and authenticate
while continue_reading:
    
    # Scan for cards    
    (status, _) = MIFAREReader.MFRC522_Request(MIFAREReader.PICC_REQIDL)

    # If a card is found
    if status == MIFAREReader.MI_OK:
        # Ensure scan hasn't occurred too rapidly, or the same RFID wasn't scanned twice
        now = time.time()
        diff = now - last_scan
        if diff >= cooldown:
            last_scan = now 

            # Get the UID of the card
            (status, uid) = MIFAREReader.MFRC522_Anticoll()

            # If we have the UID, continue
            if status == MIFAREReader.MI_OK:

                # Get UID
                scanned = "{}.{}.{}.{}".format(uid[0], uid[1], uid[2], uid[3])

                if scanned in uids:
                    song = uids[scanned]

                    # Check if there's anything playing
                    vlc_running = False
                    for proc in psutil.process_iter():
                        try:
                            if "vlc" in proc.name().lower():
                                vlc_running = True
                        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
                            pass

                    if song != current_song or not vlc_running:
                        current_song = song
                        subprocess.call(["/home/pi/go/bin/songPlayer", song, "BLUETOOTH"])
                else:
                    print scanned, "is not mapped."
            else:
                print "Error reading RFID."


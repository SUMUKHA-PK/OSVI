#!/usr/bin/env python3


from grovepi import *
import RPi.GPIO as GPIO

"""
Input - Output relay
3 - 13
4 - 12
5 - 7
6 - 26
"""




# Makes the relay input pin high, checks for existance of relay at the pin
# First argument is the relay input
# Second argument is the output of the relay
def relayOn(inputR, outputR, t) : 
    
    print("Inside relayOn")

    # triggering the relay
    assert pinMode(inputR, "OUTPUT") == 1
    assert digitalWrite(inputR, 1) == 1
    print("Relay pin: %d made high"%(inputR))
    
    GPIO.setmode(GPIO.BOARD)
    GPIO.setup(outputR, GPIO.IN, pull_up_down = GPIO.PUD_UP)
    # return the state of the output pin of the relay
    return GPIO.input(outputR)

# switches the relay off
def relayOff(inputR,t):
    assert pinMode(inputR,"OUTPUT") == 1
    assert digitalWrite(inputR,0) == 1
    print("Relay switched off")

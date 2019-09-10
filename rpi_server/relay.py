#!/usr/bin/env python3

from grovepi import *
from time import sleep

def relay(i, t) : 
    assert pinMode(i, "OUTPUT") == 1
    assert digitalWrite(i, 1) == 1
    print("High")
    sleep(t)
    assert digitalWrite(i, 0) == 1
    print("Low")

#if __name__ == "__main__":
#    for i in range(2, 8):
#        relay(i, 1)

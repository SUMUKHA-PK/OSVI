from grovepi import *
from time import sleep


def testRelay(i, t):
    assert pinMode(i, "OUTPUT") == 1
    assert digitalWrite(i, 1) == 1
    print("High")
    sleep(t)
    assert digitalWrite(i, 0) == 1
    print("Low")


def check(pin):
    try:
        if pinMode(pin, "INPUT") != 1:
            raise
        if digitalRead(pin) == 1:
            raise ValueError
    except AssertionError:


if __name__ == "__main__":
   for i in range(2, 8):
       testRelay(i, 1)

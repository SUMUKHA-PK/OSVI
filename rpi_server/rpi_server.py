#!/usr/bin/env python3

import os
import sys
from http import server
from http import client
import relay
import json
import threading
import time

trig_to_relay = {'test' : (3, 13), "reset" :(4, 12) , "start" : (5, 7)}


class MyTimer(object):
    def __init__(self):
        self._lock = Lock()
        self.start()

    def start(self):

        with self._lock:
            self._timer = Timer(2, self.update)
            self._timer.start()
            self._running = True

    def stop(self):

        with self._lock:
            if self._timer.is_alive():
                self._timer.cancel()

            self._running = False

    def restart(self):

        with self._lock:
            if not self._running:
                return

        self.start()

    def timeout_handler(self): 

        conn = client.HTTPConnection("10.53.88.119:55555")
        headers = {'Content-type' : 'application/json'}
        foo = {'text' : 'This makes sure experiment is complete'}
        json_data = json.dumps(foo)
        conn.request("POST", "/experimentComplete", json_data, headers)

        response = conn.getresponse()
        print(response.read().decode())

def timeout_handler(): 

    conn = client.HTTPConnection("10.53.88.119:55555")
    headers = {'Content-type' : 'application/json'}
    foo = {'text' : 'This makes sure experiment is complete'}
    json_data = json.dumps(foo)
    conn.request("POST", "/experimentComplete", json_data, headers)

    response = conn.getresponse()
    print(response.read().decode())

class HTTPClientHandler(server.BaseHTTPRequestHandler) : 
    
    # Health check routine!
    def do_GET(self) : 
        self.send_response(200)

        # Send headers
        self.send_header('Content-type', 'text/html')
        self.end_headers()

        # Send message back to client
        message = "Health check okay!"
        # Write content as utf-8 data
        self.wfile.write(bytes(message, 'utf-8'))


    def do_POST(self) :     

        try: 

            """
            Basic Idea: 

            1. There are 5 buttons. So, PK sends a json message.
        
            2. Decode each message. Call the correponding manish's API - This is 
                blocking. Then send back the result to PK.

            3. Let us see how this blocking mechanism works. Later, we can use an               
                event loop and make it non-blocking.
            """
            trigger_type = None
            trigger_machine = None
            status_msg = None
            content_len = None
            msg = None
            
            # Timer object
            #timer = MyTimer()

            # Read the message body
            status_msg = self.headers.get('Status-Message')
            content_len = int(self.headers.get('Content-Length'))
            msg = self.rfile.read(content_len)
            print(msg)

            # Decode the json-request.
            decoded_req = json.loads(msg)
            print(decoded_req)
             
            trigger_type = decoded_req['TriggerType']
            trigger_machine = decoded_req['Machine']

            print(trigger_type, trigger_machine)

            # At this point, I know which relay to trigger.
            # Use the relay API - Blocking action.
            # Use trig_to_relay
            
            message = None
            if relay.relayOn(trig_to_relay["reset"][0], trig_to_relay["reset"][1], 1) == 0 :
                message = "Relay switched on"
                print("Relay switched on")
                relay.relayOff(trig_to_relay["reset"][0],1)
            else :
                message = "Relay isn't connected to the system"
                print("Relay isn't connected to the system")
            
            time.sleep(1)
            
            if relay.relayOn(trig_to_relay["test"][0], trig_to_relay["test"][1], 1) == 0 :
                message = "Relay switched on"
                print("Relay switched on")
                relay.relayOff(trig_to_relay["test"][0],1)
            else :
                message = "Relay isn't connected to the system"
                print("Relay isn't connected to the system")

            time.sleep(1)
            
            if relay.relayOn(trig_to_relay["start"][0], trig_to_relay["start"][1], 1) == 0 :
                message = "Relay switched on"
                print("Relay switched on")
                relay.relayOff(trig_to_relay["start"][0],1)
            else :
                message = "Relay isn't connected to the system"
                print("Relay isn't connected to the system")
            
            time.sleep(1)

            if relay.relayOn(trig_to_relay["start"][0], trig_to_relay["start"][1], 1) == 0 :
                message = "Relay switched on"
                print("Relay switched on")
                relay.relayOff(trig_to_relay["start"][0],1)
            else :
                message = "Relay isn't connected to the system"
                print("Relay isn't connected to the system")
            
            print("Before crafting message")

            # Craft the body
            self.send_response_only(200, "Relay status: " + message)
            self.send_header('Content-type','application/json')
            self.end_headers()

            print("After crafting")
            timeout_handler()

            #self.wfile.write(bytes("Working?", 'utf-8'))
            print("BEfore timer")
            #timer.start()
            print("After timer")
            
            # Once the timer goes off, I need to send PK a post request.
            """
            At this point, response is sent. 
            I need to start a timer for X minutes. 
            * Send back a message to PK once X minutes is over.
            """
            
        except:
            self.send_response_only(400, "Bad Request")
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            message = "Error in handling the POST Request"
            self.wfile.write(bytes(message, 'utf-8'))

        return


def main(server_ip_addr, server_port_no) : 
    
    server_address = (server_ip_addr, server_port_no)
    try: 
        httpd = server.HTTPServer(server_address, HTTPClientHandler)
        print("Running HTTPServer at ", server_address)
        httpd.serve_forever()
        
    except: 
        print("Unable to run server at ", server_address)
        sys.exit(-1)


if __name__ == "__main__": 
    if len(sys.argv) != 3: 
        print("Usage: $ ", sys.argv[0], "<IPAddress> <PortNo>")
        sys.exit(-1)

    main(sys.argv[1], int(sys.argv[2]))

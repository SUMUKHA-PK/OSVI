#!/usr/bin/env python3

import os
import sys
import threading
from http import server
import relay
import json


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
        self.send_response(200)
        
        # Send headers
        self.send_header('Content-type','text/html')
        self.end_headers()
 
        """
        Basic Idea: 

        1. There are 5 buttons. So, PK can send 5 different json-encoded 
            messages. 
        
        2. Decode each message. Call the correponding manish's API - This is 
            blocking. Then send back the result to PK.

        3. Let us see how this blocking mechanism works. Later, we can use an 
            event loop and make it non-blocking.
        """

        message = "Hello world!"
        self.wfile.write(bytes(message, "utf8"))
        for i in range(0, 5) : 
            relay.relay(5, 1)

        return


def main(server_ip_addr, server_port_no) : 
    
    server_address = (server_ip_addr, server_port_no)
    httpd = server.HTTPServer(server_address, HTTPClientHandler)
    print("Running HTTPServer at ", server_address)
    httpd.serve_forever()


if __name__ == "__main__": 
    if len(sys.argv) != 3: 
        print("Usage: $ ", sys.argv[0], "<IPAddress> <PortNo>")
        sys.exit(-1)

    main(sys.argv[1], int(sys.argv[2]))

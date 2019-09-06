#!/usr/bin/env python3

import os
import sys
import threading
from http import server

class HTTPClientHandler(server.BaseHTTPRequestHandler) : 
    
    def do_POST(self) : 
        self.send_response(200)
        
        # Send headers
        self.send_header('Content-type','text/html')
        self.end_headers()
 
        # Send message back to client
        message = "Hello world!"
        # Write content as utf-8 data
        self.wfile.write(bytes(message, "utf8"))
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

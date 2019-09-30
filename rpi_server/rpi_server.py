#!/usr/bin/env python3

import os
import sys
import threading
from http import server
import relay
import json
import types



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
        
            # Objects needed.
            trigger_req = types.trigger_request()


            # Read the message body
            status_msg = self.headers.get('Status-Message')
            content_len = int(self.headers.get('Content-Length'))
            msg = self.rfile.read(content_len)
            print(msg)


            # Decode the json-request.
            decoded_req = json.loads(msg)
            print(decoded_req)
            """ 
            trigger_req.trigger_type = decoded_req['TriggerType']
            trigger_req.machine = decoded_req['Machine']
            
            print(trigger_req.trigger_type, trigger_req.machine)
            """



            # Craft the body
            self.send_response_only(200, "Successfully received a trigger request")
            self.send_header('Content-type','application/json')
            self.end_headers()

            message = "Hello world!"
            self.wfile.write(bytes(message, "utf8"))
            """
            for i in range(0, 5) : 
                relay.relay(5, 1)
            """
        except: 
            self.send_error(400, "Bad Request: " + str(sys.exc_info()[0]))

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

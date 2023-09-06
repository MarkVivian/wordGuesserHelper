import http.server
# allows us to set the server to a specific port.
import socketserver

# this is the port of the server.
port: int = 3005

# this is the simple server. its incharge of incoming requests.
handler = http.server.SimpleHTTPRequestHandler

#  creates a TCP server on the specified port and associates it with the handler (the file server).
with socketserver.TCPServer(("", port), handler) as httpd:
    print(f"server at port {port} is running")
    # starts the server and makes it listen for incoming requests indefinitely.
    httpd.serve_forever()
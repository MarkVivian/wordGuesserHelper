import http.server
import socketserver

port: int = 3005
handler = http.server.SimpleHTTPRequestHandler

with socketserver.TCPServer(("", port), handler) as httpd:
    print(f"server at port {port} is running")
    httpd.serve_forever()
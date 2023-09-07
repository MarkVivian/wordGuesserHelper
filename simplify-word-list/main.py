import http.server
import Api_handler

# Create an HTTP server with the custom request handler
server = http.server.HTTPServer(('localhost', 3002), Api_handler.RequestHandling)

# Start the server
print('Server listening on http://localhost:3002')
server.serve_forever()
#http://localhost:3002/filterWords
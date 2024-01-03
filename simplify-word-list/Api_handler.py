import http.server
import json
import dataclasses
import dataTypes
import filterData


class DataClassEncoder(json.JSONEncoder):
    def default(self, obj):
        if dataclasses.is_dataclass(obj):
            # Convert the data class instance to a dictionary
            return dataclasses.asdict(obj)
        return super().default(obj)


class RequestHandling(http.server.BaseHTTPRequestHandler):
    def do_POST(self):
        if self.path == "/filterWords":
            # Read the request data
            # Read the request body
            content_length = int(self.headers['Content-Length'])
            request_data = self.rfile.read(content_length).decode('utf-8')

            # Deserialize the JSON data
            try:
                # this allows me to use the data class to access the values.
                '''
                    without it:
                        getting data -> data["firstLetter"]
                    with it:
                        getting data -> data.firstLetter
                '''
                data = dataTypes.DataValues(**json.loads(request_data))
                print(data.firstLetter)
            except json.JSONDecodeError as e:
                print("an error occured while decoding the data")
                self.send_response(400)
                self.end_headers()
                self.wfile.write(f"Invalid JSON: {e}".encode('utf-8'))
                return

            # Process the data
            processed_data: dataTypes.DataValues = filterData.FilterData(data).checkFirstAndLastLetter()
            # Serialize the response data back to JSON
            response_data = json.dumps(processed_data, cls=DataClassEncoder)

            # Set response headers
            self.send_response(200)
            self.send_header('Content-Type', 'application/json')
            self.end_headers()

            # Send the response
            self.wfile.write(response_data.encode("utf-8"))
        else:
            # Handle other paths if needed
            self.send_response(404)
            self.end_headers()
            self.wfile.write(b'Not Found')

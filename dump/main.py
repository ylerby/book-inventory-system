from http.server import BaseHTTPRequestHandler, HTTPServer
import json


class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):
    def send_dump(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()

        data = {
            "books": {
                1: {
                    "title": "Book 1",
                    "author": "Author 1",
                    "year": 2022
                },
                2: {
                    "title": "Book 2",
                    "author": "Author 2",
                    "year": 2021
                },
                3: {
                    "title": "Book 3",
                    "author": "Author 3",
                    "year": 2020
                }
            }
        }
        self.wfile.write(json.dumps(data).encode())


def run(server_class=HTTPServer, handler_class=SimpleHTTPRequestHandler, port=8000) -> None:
    server_address = ("", port)
    httpd = server_class(server_address, handler_class)
    httpd.serve_forever()


if __name__ == "__main__":
    run(port=8080)

from http.server import BaseHTTPRequestHandler, HTTPServer
import json
import pandas as pd


class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):
    def read_and_process_data(self):
        books = pd.read_pickle("books.pkl")
        instances = pd.read_pickle("instances.pkl")
        admins = pd.read_pickle("admins.pkl")
        users = pd.read_pickle("users.pkl")
        readers = pd.read_pickle("readers.pkl")
        language = pd.read_pickle("language.pkl")
        genre = pd.read_pickle("genre.pkl")
        format_df = pd.read_pickle("format.pkl")
        authors = pd.read_pickle("authors.pkl")
        productions = pd.read_pickle("productions.pkl")

        books_dict = books.to_dict(orient="index")
        instances_dict = instances.set_index("instanceID").T.to_dict("dict")
        admins_dict = admins.set_index("adminID").T.to_dict("dict")
        users_dict = users.set_index("userID").T.to_dict("dict")
        readers_dict = readers.set_index("readerID").T.to_dict("dict")
        language_dict = language.set_index("languageID").T.to_dict("dict")
        genre_dict = genre.set_index("genreID").T.to_dict("dict")
        format_dict = format_df.set_index("formatID").T.to_dict("dict")

        authors_dict = {
            row["authorId"]: {"name": row["name"], "productionID": row["productionID"]}
            for index, row in authors.iterrows()
        }

        productions_dict = productions.set_index("productionID").T.to_dict("dict")

        final_json = {
            "books": books_dict,
            "instances": instances_dict,
            "admins": admins_dict,
            "users": users_dict,
            "readers": readers_dict,
            "language": language_dict,
            "genre": genre_dict,
            "format": format_dict,
            "authors": authors_dict,
            "productions": productions_dict,
        }

        return final_json

    def send_dump(self):
        self.send_response(200)
        self.send_header("Content-type", "application/json")
        self.end_headers()

        data = self.read_and_process_data()

        self.wfile.write(json.dumps(data).encode())


def run(
    server_class=HTTPServer, handler_class=SimpleHTTPRequestHandler, port=8000
) -> None:
    server_address = ("", port)
    httpd = server_class(server_address, handler_class)
    httpd.serve_forever()


if __name__ == "__main__":
    run(port=8080)

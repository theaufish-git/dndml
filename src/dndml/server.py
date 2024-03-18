from pathlib import Path

from flask import Flask, send_from_directory

app = Flask(__name__)

root = Path(__file__).parent.parent.parent/"tests/data"


@app.route("/<path:filename>")
def serve_file(filename):
    print(root / filename)
    return send_from_directory(root, filename)


if __name__ == "__main__":
    app.run()

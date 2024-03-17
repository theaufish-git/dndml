from typing import Any
from pathlib import Path
from urllib.parse import urlparse
from urllib import request

class Registry:
    _include_directories: list[str]
    _registry: dict[str, Any]

    def __init__(self, directories: list[str]) -> None:
        self._registry = {}
        self._include_directories = [*directories]
        self._include_directories.append(".")
        self._include_directories.append(Path("~/.dndml/manifests").expanduser())

    def register(self, uri: str) -> None:
        parsed_url = urlparse(uri)
        
        if parsed_url.scheme == "":
            parsed_url.scheme = "file"

        response = request.urlopen(parsed_url.geturl())
        response.read()
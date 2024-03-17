from typing import Any, Optional, Protocol

from pydantic import BaseModel, ImportString
from yaml import safe_load
from .registry import Registry


class InvalidManifestError(Exception): ...


class SupportsRead(Protocol):
    def read(self, __length: int = ...) -> str | bytes: ...


class Manifest(BaseModel):
    kind: ImportString
    kwargs: dict[str, Any]
    source: Optional[str] = None
    registry: Registry

    def unpack(self) -> Any:
        for key, value in self.kwargs.items():
            if isinstance(value, Manifest):
                self.kwargs[key] = value.unpack()
            elif isinstance(value, dict):
                self.kwargs[key] = {
                    k: v.unpack() if isinstance(v, Manifest) else v
                    for k, v in value.items()
                }
            elif isinstance(value, list):
                self.kwargs[key] = [
                    item.unpack() if isinstance(item, Manifest) else item
                    for item in value
                ]

        

        return self.kind(**self.kwargs)


class ManifestParser:
    @staticmethod
    def parse(stream: str | bytes | SupportsRead) -> Manifest:
        root = safe_load(stream)

        if not isinstance(root, dict):
            raise InvalidManifestError("Manifest must be a dictionary")

        return ManifestParser._parse_object(root)

    @staticmethod
    def _parse_object(root: Any) -> Any:
        if isinstance(root, dict) and "kind" in root:
            kind = root["kind"]
            del root["kind"]
            kwargs = ManifestParser._parse_object(root)

            return Manifest(kind=kind, kwargs=kwargs)
        elif isinstance(root, dict):
            return {
                key: ManifestParser._parse_object(value)
                for key, value in root.items()
            }
        elif isinstance(root, list):
            return [ManifestParser._parse_object(item) for item in root]

        return root

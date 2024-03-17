from enum import StrEnum, auto

from caseconverter import (camelcase, flatcase, kebabcase, macrocase,
                           pascalcase, snakecase)


class StrEnum(StrEnum):
    def __str__(self):
        return self.name

    def camelcase(self):
        return camelcase(self.name)

    def flatcase(self):
        return flatcase(self.name)

    def kebabcase(self):
        return kebabcase(self.name)

    def macrocase(self):
        return macrocase(self.name)

    def pascalcase(self):
        return pascalcase(self.name)

    def snakecase(self):
        return snakecase(self.name)

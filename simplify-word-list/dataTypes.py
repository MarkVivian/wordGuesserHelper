from dataclasses import dataclass


@dataclass
class DataValues:
    firstLetter: str
    lastLetter: str
    length: int
    lettersPresent: list[str]
    wordList: list[str]

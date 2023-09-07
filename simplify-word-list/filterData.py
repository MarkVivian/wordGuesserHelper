import dataTypes


class FilterData:
    def __init__(self, data):
        self.data: dataTypes.DataValues = data

    def checkFirstAndLastLetter(self) -> dataTypes.DataValues:
        values_combined: list[str] = self.data.lettersPresent
        values_combined.append(self.data.firstLetter)
        values_combined.append(self.data.lastLetter)
        returning_values: list[str] = []
        for word in self.data.wordList:
            if all(letter in word for letter in self.data.lettersPresent):
                returning_values.append(word)

        self.data.wordList = returning_values
        print(returning_values)
        return self.data


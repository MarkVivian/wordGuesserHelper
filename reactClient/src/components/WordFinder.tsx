"use client"

import type React from "react"
import { useState } from "react"
import { useTheme } from "../context/ThemeContext"

const WordFinder: React.FC = () => {
  const { theme } = useTheme()
  const [wordLength, setWordLength] = useState<string>("")
  const [firstLetter, setFirstLetter] = useState<string>("")
  const [lastLetter, setLastLetter] = useState<string>("")
  const [results, setResults] = useState<string[]>([])
  const [showValidationMessage, setShowValidationMessage] = useState<boolean>(false)
  const [lettersPresentInput, setLettersPresentInput] = useState("")
  const [lettersNotPresentInput, setLettersNotPresentInput] = useState("")

  // Only allow A-Z for firstLetter and lastLetter
  const handleFirstLetterChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const val = e.target.value.toUpperCase()
    if (/^[A-Z]?$/.test(val)) setFirstLetter(val.toLowerCase())
  }
  const handleLastLetterChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const val = e.target.value.toUpperCase()
    if (/^[A-Z]?$/.test(val)) setLastLetter(val.toLowerCase())
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    // Validation - at least one field must be filled
    if (
      (!wordLength &&
      !firstLetter &&
      !lastLetter &&
      !lettersPresentInput &&
      !lettersNotPresentInput) || !wordLength
    ) {
      setShowValidationMessage(true)
      return
    } else {
      setShowValidationMessage(false)
    }

    // Split lettersPresentInput and lettersNotPresentInput into arrays of unique uppercase letters
    const presentArr = Array.from(
      new Set((lettersPresentInput.toUpperCase().replace(/[^A-Z]/g, "")).toLowerCase().split(""))
    )
    const notPresentArr = Array.from(
      new Set((lettersNotPresentInput.toUpperCase().replace(/[^A-Z]/g, "")).toLowerCase().split(""))
    )

    const body_data = JSON.stringify({
      firstLetter: firstLetter,
      lastLetter: lastLetter,
      length: parseInt(wordLength) || 0,
      lettersPresent: presentArr,
      lettersNotPresent: notPresentArr,
      wordList: [],
    })

    console.log(body_data)
    const response = await fetch("http://localhost:3001/searchWord", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: body_data,
    });
    const data = await response.json()
    if (data.error) {
      console.error(data.error)
      return
    }
    console.log(data)
    setResults(data["wordList"])
  }

  const buttonClass = theme === "purple" ? "purple-button" : "blue-button"

  return (
    <div className="w-full md:w-1/2 bg-white/10 backdrop-blur-md rounded-lg p-6 flex flex-col h-full">
      <h2 className="text-xl font-semibold text-white mb-4">Word Finder</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label className="block text-white text-sm mb-1">Word Length</label>
            <input
              type="number"
              min="3"
              max="15"
              value={wordLength}
              onChange={(e) => setWordLength(e.target.value)}
              className="w-full bg-white/20 border-none text-white rounded px-3 py-2 caret-white"
            />
          </div>
          <div>
            <label className="block text-white text-sm mb-1">First Letter</label>
            <input
              type="text"
              maxLength={1}
              value={firstLetter}
              onChange={handleFirstLetterChange}
              className="w-full bg-white/20 border-none text-white rounded px-3 py-2 caret-white"
            />
          </div>
          <div>
            <label className="block text-white text-sm mb-1">Letters Present</label>
            <input
              type="text"
              value={lettersPresentInput}
              onChange={(e) =>
                setLettersPresentInput(e.target.value.replace(/[^a-zA-Z]/g, ""))
              }
              className="w-full bg-white/20 border-none text-white rounded px-3 py-2 caret-white"
              placeholder="Type all present letters (A-Z)"
            />
          </div>
          <div>
            <label className="block text-white text-sm mb-1">Last Letter</label>
            <input
              type="text"
              maxLength={1}
              value={lastLetter}
              onChange={handleLastLetterChange}
              className="w-full bg-white/20 border-none text-white rounded px-3 py-2 caret-white"
            />
          </div>
          <div className="col-span-1 sm:col-span-2">
            <label className="block text-white text-sm mb-1">Letters Not Present</label>
            <input
              type="text"
              value={lettersNotPresentInput}
              onChange={(e) =>
                setLettersNotPresentInput(e.target.value.replace(/[^a-zA-Z]/g, ""))
              }
              className="w-full bg-white/20 border-none text-white rounded px-3 py-2 caret-white"
              placeholder="Type all not present letters (A-Z)"
            />
          </div>
        </div>
        {showValidationMessage && (
          <div className="text-red-300 text-sm">Please provide at least one search criteria and word length</div>
        )}
        <button type="submit" className={`w-full py-2 rounded-button whitespace-nowrap ${buttonClass}`}>
          Find Words
        </button>
      </form>
      <div className="mt-4 flex-grow overflow-hidden">
        <h3 className="text-white font-semibold mb-2">Results:</h3>
        <div className="bg-white/10 rounded-lg p-4 h-[calc(100%-2rem)] overflow-y-auto">
          {(results != null) && (results.length > 0) ? (
            <div className="results-container">
              {results.map((word, index) => (
                <div key={index} className="result-item text-white">
                  {word}
                </div>
              ))}
            </div>
          ) : (
            <p className="text-white/70 text-center">Use the form above to find words matching your criteria</p>
          )}
        </div>
      </div>
    </div>
  )
}

export default WordFinder

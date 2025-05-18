"use client"

import type React from "react"
import { useState, useEffect, useCallback } from "react"
import HowToPlayModal from "./HowToPlayModal"
import Keyboard from "./Keyboard"
import { useTheme } from "../context/ThemeContext"
import GameModal from "./GameModal"

const WordleGame: React.FC = () => {
  const { theme } = useTheme()
  const [gridSize, setGridSize] = useState<number>(5)
  const [gridSizeInput, setGridSizeInput] = useState<string>("5")
  const [gridSizeError, setGridSizeError] = useState<string>("")
  const [currentRow, setCurrentRow] = useState<number>(0)
  const [currentCol, setCurrentCol] = useState<number>(0)
  const [secretWord, setSecretWord] = useState<string>("")
  const [grid, setGrid] = useState<Array<Array<{ letter: string; state: string; locked: boolean }>>>([])
  const [keyboardState, setKeyboardState] = useState<Record<string, string>>({})
  const [showModal, setShowModal] = useState<boolean>(false)
  const [showGiveUpModal, setShowGiveUpModal] = useState<boolean>(false)
  const [showWinModal, setShowWinModal] = useState<boolean>(false)
  const [gameOver, setGameOver] = useState<boolean>(false)
  const [maxGridSize, setMaxGridSize] = useState<number>(8)
  const maxAttempts = 5

  // Fetch a new word and reset game
  const initializeGame = useCallback(async (newGridSize?: number) => {
    const size = newGridSize ?? gridSize
    const newGrid = Array(maxAttempts)
      .fill(null)
      .map(() =>
        Array(size)
          .fill(null)
          .map(() => ({ letter: "", state: "", locked: false })),
      )
    setGrid(newGrid)
    setCurrentRow(0)
    setCurrentCol(0)
    setKeyboardState({})
    setGameOver(false)

    const body_data = JSON.stringify({
      length: size,
      randomWord: ""
    })

    const response = await fetch("http://localhost:3001/randomWord", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: body_data,
    });
    const data = await response.json()
    setSecretWord(data["randomWord"].toUpperCase())
  }, [gridSize, maxAttempts])

  // Initialize game on page load and when grid size changes
  useEffect(() => {
    initializeGame(gridSize)
  }, [gridSize, initializeGame])

  // Handle grid size input
  const handleGridSizeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const val = e.target.value.replace(/[^0-9]/g, "")
    setGridSizeInput(val)
    setGridSizeError("")
    initializeGame(Number(val))
  }

  // Update grid size and start new game
  const updateGrid = () => {
    const num = Number(gridSizeInput)
    if (num < 4 || num > maxGridSize) {
      setGridSizeError(`Grid size must be between 4 and ${maxGridSize}`)
      return
    }
    setGridSize(num)
    setGridSizeError("")
    initializeGame(num)
  }

  // Handle give up
  const giveUp = () => {
    setGameOver(true)
    setShowGiveUpModal(true)
  }

  const handleGiveUpModalClose = () => {
    setShowGiveUpModal(false)
    initializeGame(gridSize)
  }

  const handleWinModalClose = () => {
    setShowWinModal(false)
    initializeGame(gridSize)
  }

  // Handle key input
  const handleKeyInput = useCallback(
    (key: string) => {
      if (gameOver) return

      if (key === "backspace") {
        if (currentCol > 0 && !grid[currentRow][currentCol - 1].locked) {
          const newGrid = grid.map(row => row.map(cell => ({ ...cell })))
          newGrid[currentRow][currentCol - 1].letter = ""
          setGrid(newGrid)
          setCurrentCol(currentCol - 1)
        }
      } else if (key === "enter") {
        if (currentCol === gridSize) {
          checkWord()
        }
      } else if (/^[a-z]$/.test(key) && currentCol < gridSize && !grid[currentRow][currentCol].locked) {
        const newGrid = grid.map(row => row.map(cell => ({ ...cell })))
        newGrid[currentRow][currentCol].letter = key.toUpperCase()
        setGrid(newGrid)
        setCurrentCol(currentCol + 1)
      }
    },
    [currentCol, currentRow, gameOver, grid, gridSize],
  )

  // Keyboard event listener
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      const active = document.activeElement
      if (active && (active.tagName === "INPUT" || active.tagName === "TEXTAREA" || (active as HTMLElement).isContentEditable)) {
        return
      }
      if (e.key === "Backspace") {
        handleKeyInput("backspace")
      } else if (e.key === "Enter") {
        handleKeyInput("enter")
      } else if (/^[a-zA-Z]$/.test(e.key)) {
        handleKeyInput(e.key.toLowerCase())
      }
    }

    window.addEventListener("keydown", handleKeyDown)
    return () => window.removeEventListener("keydown", handleKeyDown)
  }, [handleKeyInput])

  // Check the guessed word
  const checkWord = () => {
    const currentWord = grid[currentRow].map((cell) => cell.letter).join("")
    const newGrid = grid.map(row => row.map(cell => ({ ...cell })))
    const newKeyboardState = { ...keyboardState }

    // Prepare arrays for marking
    const secretArr = secretWord.split("")
    const guessArr = currentWord.split("")
    const letterCount: Record<string, number> = {}

    // Count letters in secret word for yellow marking
    secretArr.forEach((char) => {
      letterCount[char] = (letterCount[char] || 0) + 1
    })

    // First pass: mark correct (green)
    for (let i = 0; i < gridSize; i++) {
      const letter = guessArr[i]
      if (letter === secretArr[i]) {
        newGrid[currentRow][i].state = "correct"
        newKeyboardState[letter.toLowerCase()] = "correct"
        letterCount[letter] -= 1
      }
    }

    // Second pass: mark present (yellow) or absent (grey)
    for (let i = 0; i < gridSize; i++) {
      const letter = guessArr[i]
      if (newGrid[currentRow][i].state === "correct") continue
      if (secretArr.includes(letter) && letterCount[letter] > 0) {
        newGrid[currentRow][i].state = "present"
        if (newKeyboardState[letter.toLowerCase()] !== "correct") {
          newKeyboardState[letter.toLowerCase()] = "present"
        }
        letterCount[letter] -= 1
      } else {
        newGrid[currentRow][i].state = "absent"
        if (!newKeyboardState[letter.toLowerCase()]) {
          newKeyboardState[letter.toLowerCase()] = "absent"
        }
      }
      newGrid[currentRow][i].locked = true
    }

    setGrid(newGrid)
    setKeyboardState(newKeyboardState)

    if (currentWord === secretWord) {
      setGameOver(true)
      setTimeout(() => {
        setShowWinModal(true)
      }, 300)
    } else if (currentRow + 1 >= maxAttempts) {
      setGameOver(true)
      setTimeout(() => {
        setShowGiveUpModal(true)
      }, 300)
    } else {
      setCurrentRow(currentRow + 1)
      setCurrentCol(0)
    }
  }

  // Set max grid size based on screen width
  useEffect(() => {
    const handleResize = () => {
      const width = window.innerWidth
      if (width < 640) {
        setMaxGridSize(6) // Mobile
      } else if (width < 1024) {
        setMaxGridSize(8) // Tablet
      } else {
        setMaxGridSize(11) // Desktop
      }
    }
    handleResize()
    window.addEventListener("resize", handleResize)
    return () => window.removeEventListener("resize", handleResize)
  }, [])

  // Dynamic button class based on theme
  const buttonClass = theme === "purple" ? "bg-purple-700 hover:bg-purple-800" : "bg-blue-700 hover:bg-blue-800"

  return (
    <div className="flex flex-col items-center justify-center w-full h-full min-h-[500px] p-4">
      <div className="w-full max-w-2xl bg-white/10 backdrop-blur-md rounded-lg p-4 sm:p-8 flex flex-col mb-6 md:mb-0 shadow-lg">
        <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-4 gap-4">
          <div className="flex items-center gap-4">
            <h2 className="text-xl font-semibold text-white">Daily Challenge</h2>
            <button
              onClick={() => setShowModal(true)}
              className={`px-3 py-1 rounded-button text-sm ${buttonClass} text-white`}
            >
              <i className="ri-question-line mr-1"></i>How to Play
            </button>
          </div>
          <div className="flex items-center gap-2">
            <label className="text-white text-sm">Grid Size:</label>
            <input
              type="number"
              min={4}
              max={maxGridSize}
              value={gridSizeInput}
              onChange={handleGridSizeInput}
              className="w-16 bg-white/20 border-none text-white rounded px-3 py-1 text-center"
              inputMode="numeric"
            />
            <button onClick={updateGrid} className={`px-3 py-1 rounded-button ${buttonClass} text-white`}>
              Update
            </button>
          </div>
        </div>
        {gridSizeError && (
          <div className="text-red-300 text-sm mb-2">{gridSizeError}</div>
        )}

        <div
          className="game-grid flex-grow grid gap-2 mb-4"
          style={{
            gridTemplateColumns: `repeat(${gridSize}, minmax(0, 1fr))`,
            minHeight: `${maxAttempts * 3}rem`,
          }}
        >
          {grid.map((row, rowIndex) =>
            row.map((cell, colIndex) => (
              <div
                key={`${rowIndex}-${colIndex}`}
                className={`letter-box ${cell.state} ${cell.locked ? "locked" : ""}`}
                style={{
                  width: "3rem",
                  height: "3rem",
                  fontSize: "2rem",
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "center",
                  border: "2px solid #ccc",
                  background:
                    cell.state === "correct"
                      ? "#22c55e" // green
                      : cell.state === "present"
                      ? "#eab308" // yellow
                      : cell.state === "absent"
                      ? "#6b7280" // grey
                      : cell.locked
                      ? "rgba(255,255,255,0.15)"
                      : "rgba(255,255,255,0.05)",
                  color: "#fff",
                  borderRadius: "0.5rem",
                  margin: "2px",
                  transition: "background 0.2s",
                }}
              >
                {cell.letter}
              </div>
            )),
          )}
        </div>

        <div className="flex justify-center mt-4 gap-2">
          <button onClick={giveUp} className={`px-6 py-2 rounded-button ${buttonClass} text-white`}>
            Give Up
          </button>
        </div>

        <div className="w-full mt-4 overflow-x-auto">
          <Keyboard onKeyPress={handleKeyInput} keyState={keyboardState} theme={theme} />
        </div>
      </div>

      {showModal && <HowToPlayModal onClose={() => setShowModal(false)} />}

      {showGiveUpModal && (
        <GameModal onClose={handleGiveUpModalClose}>
          <div className="flex flex-col items-center">
            <h3 className="text-xl font-semibold text-white mb-4">The word was:</h3>
            <div className="text-3xl font-mono mb-6 text-purple-300">{secretWord}</div>
            <button
              className={`px-6 py-2 rounded-button ${theme === "purple" ? "bg-purple-700 hover:bg-purple-800" : "bg-blue-700 hover:bg-blue-800"} text-white font-semibold`}
              onClick={handleGiveUpModalClose}
            >
              Play Again
            </button>
          </div>
        </GameModal>
      )}

      {showWinModal && (
        <GameModal onClose={handleWinModalClose}>
          <div className="flex flex-col items-center">
            <h3 className="text-xl font-semibold text-white mb-4">Congratulations!</h3>
            <div className="text-lg mb-6 text-purple-300">You found the word!</div>
            <button
              className={`px-6 py-2 rounded-button ${theme === "purple" ? "bg-purple-700 hover:bg-purple-800" : "bg-blue-700 hover:bg-blue-800"} text-white font-semibold`}
              onClick={handleWinModalClose}
            >
              Play Again
            </button>
          </div>
        </GameModal>
      )}
    </div>
  )
}

export default WordleGame
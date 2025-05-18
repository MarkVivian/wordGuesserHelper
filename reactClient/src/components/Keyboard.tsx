"use client"

import type React from "react"

interface KeyboardProps {
  onKeyPress: (key: string) => void
  keyState: Record<string, string>
  theme: string
}

const Keyboard: React.FC<KeyboardProps> = ({ onKeyPress, keyState, theme }) => {
  const rows = [
    ["q", "w", "e", "r", "t", "y", "u", "i", "o", "p"],
    ["a", "s", "d", "f", "g", "h", "j", "k", "l"],
    ["z", "x", "c", "v", "b", "n", "m"],
  ]

  const getKeyClass = (key: string) => {
    if (keyState[key]) {
      return keyState[key] // correct, present, or absent
    }
    return theme === "purple" ? "unused-purple" : "unused-blue"
  }

  return (
    <div className="mt-4 pb-2">
      {rows.map((row, rowIndex) => (
        <div key={rowIndex} className="flex justify-center gap-1 mb-1">
          {rowIndex === 2 && (
            <div
              className={`keyboard-key ${theme === "purple" ? "unused-purple" : "unused-blue"}`}
              onClick={() => onKeyPress("enter")}
              style={{ padding: "0 8px" }}
            >
              ENTER
            </div>
          )}

          {row.map((key) => (
            <div key={key} className={`keyboard-key ${getKeyClass(key)}`} onClick={() => onKeyPress(key)}>
              {key.toUpperCase()}
            </div>
          ))}

          {rowIndex === 2 && (
            <div
              className={`keyboard-key ${theme === "purple" ? "unused-purple" : "unused-blue"}`}
              onClick={() => onKeyPress("backspace")}
              style={{ padding: "0 8px" }}
            >
              ⌫
            </div>
          )}
        </div>
      ))}
    </div>
  )
}

export default Keyboard

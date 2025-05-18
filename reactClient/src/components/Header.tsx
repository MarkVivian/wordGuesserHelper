"use client"

import type React from "react"
import { useTheme } from "../context/ThemeContext"

const Header: React.FC = () => {
  const { theme, changeTheme } = useTheme()

  const handleThemeToggle = (e: React.ChangeEvent<HTMLInputElement>) => {
    changeTheme(e.target.checked ? "blue" : "purple")
  }

  return (
    <header className="flex justify-between items-center mb-6">
      <h1 className="text-3xl font-['Pacifico'] text-white">Wordle Game</h1>
      <div className="flex items-center space-x-3">
        <span className="text-white text-sm">Purple</span>
        <label className="custom-switch">
          <input type="checkbox" checked={theme === "blue"} onChange={handleThemeToggle} />
          <span className="switch-slider"></span>
        </label>
        <span className="text-white text-sm">Blue</span>
      </div>
    </header>
  )
}

export default Header

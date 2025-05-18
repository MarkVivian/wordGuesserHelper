"use client"

import type React from "react"
import { createContext, useState, useContext, type ReactNode } from "react"

type Theme = "purple" | "blue"

interface ThemeContextType {
  theme: Theme
  changeTheme: (newTheme: Theme) => void
}

const ThemeContext = createContext<ThemeContextType | undefined>(undefined)

interface ThemeProviderProps {
  children: ReactNode
}

export const ThemeProvider: React.FC<ThemeProviderProps> = ({ children }) => {
  const [theme, setTheme] = useState<Theme>("purple")

  const changeTheme = (newTheme: Theme) => {
    setTheme(newTheme)
  }

  return (
    <ThemeContext.Provider value={{ theme, changeTheme }}>
      <div className={`${theme === "purple" ? "purple-theme" : "blue-theme"}`}>{children}</div>
    </ThemeContext.Provider>
  )
}

export const useTheme = (): ThemeContextType => {
  const context = useContext(ThemeContext)
  if (context === undefined) {
    throw new Error("useTheme must be used within a ThemeProvider")
  }
  return context
}

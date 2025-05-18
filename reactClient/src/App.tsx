import type React from "react"
import Header from "./components/Header"
import WordleGame from "./components/WordleGame"
import WordFinder from "./components/WordFinder"
import { ThemeProvider } from "./context/ThemeContext"
import "./App.css"

const App: React.FC = () => {
  return (
    <ThemeProvider>
      <div className="app-container min-h-screen bg-gradient-to-br from-purple-900 to-blue-900">
        <div className="container mx-auto px-4 py-6 max-w-7xl flex flex-col min-h-screen">
          <Header />
          <div className="flex flex-col md:flex-row gap-6 flex-1 overflow-y-auto">
            <WordleGame />
            <WordFinder />
          </div>
        </div>
      </div>
    </ThemeProvider>
  )
}

export default App

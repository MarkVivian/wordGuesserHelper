"use client"

import type React from "react"
import { useRef, useEffect } from "react"

interface HowToPlayModalProps {
  onClose: () => void
}

const HowToPlayModal: React.FC<HowToPlayModalProps> = ({ onClose }) => {
  const modalRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (modalRef.current && !modalRef.current.contains(event.target as Node)) {
        onClose()
      }
    }

    document.addEventListener("mousedown", handleClickOutside)
    return () => {
      document.removeEventListener("mousedown", handleClickOutside)
    }
  }, [onClose])

  return (
    <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50">
      <div ref={modalRef} className="bg-white/10 backdrop-blur-md p-6 rounded-lg max-w-md w-full mx-4">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-xl font-semibold text-white">How to Play</h3>
          <button onClick={onClose} className="text-white hover:text-white/80">
            <i className="ri-close-line text-2xl"></i>
          </button>
        </div>
        <ul className="text-white/90 text-sm space-y-3">
          <li className="flex items-center gap-2">
            <span className="w-4 h-4 bg-green-600 rounded"></span>
            <span>Green: Correct letter in correct position</span>
          </li>
          <li className="flex items-center gap-2">
            <span className="w-4 h-4 bg-yellow-600 rounded"></span>
            <span>Yellow: Correct letter in wrong position</span>
          </li>
          <li className="flex items-center gap-2">
            <span className="w-4 h-4 bg-gray-600 rounded"></span>
            <span>Gray: Letter not in the word</span>
          </li>
        </ul>
      </div>
    </div>
  )
}

export default HowToPlayModal

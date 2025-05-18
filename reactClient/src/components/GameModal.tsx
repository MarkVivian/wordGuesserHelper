import React, { useRef, useEffect } from "react"

interface GameModalProps {
  onClose: () => void
  children: React.ReactNode
}

const GameModal: React.FC<GameModalProps> = ({ onClose, children }) => {
  const modalRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (modalRef.current && !modalRef.current.contains(event.target as Node)) {
        onClose()
      }
    }
    document.addEventListener("mousedown", handleClickOutside)
    return () => document.removeEventListener("mousedown", handleClickOutside)
  }, [onClose])

  return (
    <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50">
      <div ref={modalRef} className="bg-white/10 backdrop-blur-md p-6 rounded-lg max-w-md w-full mx-4">
        {children}
      </div>
    </div>
  )
}

export default GameModal
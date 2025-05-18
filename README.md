# Word Guesser Helper

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue)](https://golang.org)
[![React](https://img.shields.io/badge/react-%3E%3D17-blue)](https://reactjs.org)

**Word Guesser Helper** is a **two-part** application that aids in word-guessing games (e.g., Wordle) by providing:

* 🎲 **Random Word Generation** (by length)
* 🔍 **Word Search** (by constraints)

---

## 🔖 Table of Contents

1. [Features](#features)
2. [Architecture](#architecture)
3. [Prerequisites](#prerequisites)
4. [Installation & Setup](#installation--setup)

   * [Server](#server)
   * [Client](#client)
5. [Configuration](#configuration)
6. [API Endpoints](#api-endpoints)
7. [Running Locally](#running-locally)
8. [Project Structure](#project-structure)
9. [Contributing](#contributing)
10. [License](#license)
11. [Contact](#contact)

---

## ✨ Features

* **Random Word**: Returns a random word of specified length from `read_data/english_words.txt`.
* **Word Search**: Filters words by:

  * `length` (optional)
  * `first_letter` / `last_letter` (optional)
  * `include_letters` (optional)
  * `exclude_letters` (optional)

---

## 🏗 Architecture

| Component                  | Language / Framework | Responsibility   |
| -------------------------- | -------------------- | ---------------- |
| **Server**                 | Go                   | • HTTP REST API  |
| • Load & query word list   |                      |                  |
| **Client**                 | React SPA            | • Wordle-like UI |
| • Sends requests to Server |                      |                  |

---

## 🔧 Prerequisites

Before you begin, ensure you have:

* **Go** `>=1.18`
* **Node.js** `>=14` & **npm** or **Yarn**
* **Git** (optional, for cloning)

---

## ⚙️ Installation & Setup

### 🔹 Server

```bash
# 1. Clone the server repo
git clone https://github.com/MarkVivian/wordGuesserHelper
cd word-guesser-helper/goServer

# 2. Install dependencies & tidy modules
go mod tidy

# 3. Add your word list
#    you can use the pre-existing 'english_words.txt' file or replace with your own word file in read_data/. 
#    ensure to add the path of the word file to '.env'.
```

### 🔹 Client

```bash
# 1. Clone the client repo
git clone https://github.com/MarkVivian/wordGuesserHelper
cd word-guesser-helper/client

# 2. Install dependencies
npm install    # or yarn install
```

---

## 🔒 Configuration

### ▶️ Server

| ENV Variable     | Default                 | Description            |
| ---------------- | ----------------------- | ---------------------- |
| `PORT`           | `3001`                  | Server listening port  |
| `WORD_LIST_PATH` | `./read_data/words.txt` | Path to word list file |

```bash
# Example
export PORT=3001
export WORD_LIST_PATH=./read_data/words.txt
```

### ▶️ Client

Create a `.env.local` in `client/`:

```ini
REACT_APP_API_BASE_URL=http://localhost:3001
```

---

## 🚀 API Endpoints

### 1. **Get Random Word**

* **URL**: `/randomWord`
* **Method**: `POST`
* **Payload**:

```json
{  
  "length": 5
}
```

**Response**:

```json
{  
  "length": 5,
  "randomWord": "apple"  
}
```

### 2. **Search Words**

* **URL**: `/api/search`
* **Method**: `POST`
* **Payload**:

```json
{  
  "length": 5,               
  "first_letter": "a",     
  "last_letter": "e",      
  "lettersPresent": ["p"],
  "lettersNotPresent": ["x"] 
}
```

**Response**:

```json
{  
  "length": 5,               
  "first_letter": "a",     
  "last_letter": "e",      
  "lettersPresent": ["p"],
  "lettersNotPresent": ["x"],
  "wordList": ["apple", "angle", "amble"]  
}
```

---

## 🖥 Running Locally

1. **Start Server**

   ```bash
   cd word-guesser-helper/goServer
   go run main.go
   ```

   *Server available at* `http://localhost:3001`

2. **Start Client**

   ```bash
   cd word-guesser-helper/client
   npm start    # or yarn start
   ```

   *App opens at* `http://localhost:3000`

3. **Play!**

   * 🎮 Use the Wordle-like UI to fetch random words or search by constraints.

---

## 📁 Project Structure

```
word-guesser-helper/
├── server/                 # Go server
│   ├── main.go             # Entry point
|   |──read_data/           # contains word list and its handler.
|   |   |── english_words.txt # word file
|   |   └── readFile.go     # handles the word file.
│   |── wordChoice/          # handles logic for search and random mechanism
|   │    |── randomWordChoice.go # handles the random word generation.
│   │    └── findRelatedWords.go # handles the logic for searching for words.
|   |── handlesApis/          # contains the api handler.
|   |   └──  apiHandler.go   # assigns the number generation and search functionality to the respective API.
│   |── components/          
│   |    └── structs.go     # contains the struct of the variables in the project.
│   └── go.mod


└── client/                 # React SPA
    ├── public/             # Static assets
    ├── src/                # Source code
    │   ├── components/     # Reusable UI
    │   ├── App.js          # Root component
    │   └── ...
    └── package.json        # Dependencies & scripts
```

---

## 🤝 Contributing

Contributions are **welcome**! Please:

1. **Fork** the repo
2. **Create** a feature branch: `git checkout -b feature/YourFeature`
3. **Commit** changes: `git commit -m "Add YourFeature"`
4. **Push** branch: `git push origin feature/YourFeature`
5. **Open** a Pull Request

---

## 📄 License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.

---

## 📬 Contact

Maintainer: [markviviantech@gmail.com](mailto:markviviantech@gmail.com)

---

*Happy guessing!* 🎉

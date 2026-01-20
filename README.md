# Golang Mini Projects

A growing collection of **Golang mini projects** created to strengthen core Go fundamentals through hands-on practice and practical implementations.

This repository will be continuously updated with new projects covering different aspects of the Go programming language.

---

## 🎯 Purpose of This Repository

The goal of this repository is to:
- Practice and solidify Go fundamentals
- Build small but complete projects
- Learn Go by writing real, runnable programs
- Gradually move from basics to more advanced concepts

Each project is self-contained and focuses on specific concepts.

---

## 🚀 How to Run a Project

Make sure Go is installed:
```bash
go version
```
Run any project using:
```bash
go run <filename>.go
```
Example:
```bash
go run cli-calculator.go
```

---

## 📂 Projects

### 1. CLI Calculator
**Description:**  
A command-line calculator built in Go that supports basic arithmetic operations and runs interactively until the user exits.

**Concepts Covered:**
- Functions
- `switch` statements
- Loops (`for`)
- Input handling using `bufio.Scanner`
- String processing and type conversion
- Basic error handling

**Features:**
- Addition, subtraction, multiplication, division, modulo
- Input validation
- Division and modulo by zero handling
- User-friendly CLI flow

---

### 2. Number Guessing Game
**Description:**  
A beginner-friendly command-line number guessing game written in Go where the player tries to guess a randomly generated number within a limited number of attempts.

**Concepts Covered:**
- Random number generation (`math/rand`)
- Seeding randomness using `time`
- Loops and conditional logic
- User input handling with `bufio.Scanner`
- String-to-integer conversion using `strconv`
- Basic game logic and flow control

**Features:**
- Randomly generated secret number (1–100)
- User-defined maximum number of attempts
- Feedback for each guess (too low / too high)
- Graceful handling of invalid input
- Game ends on correct guess or when attempts are exhausted

---

### 3. To-Do CLI Application
**Description:**  
A command-line to-do list application written in Go that allows users to manage daily tasks interactively through a menu-driven interface.

**Concepts Covered:**
- Structs for data modeling
- Slices for dynamic data storage
- Menu-driven program design
- User input handling using `bufio.Scanner`
- String manipulation and validation
- Basic state management in CLI applications

**Features:**
- Add new tasks
- View all tasks with completion status
- Mark tasks as completed
- Delete tasks by index
- Continuous menu loop until user exits
- Input validation for safe operations



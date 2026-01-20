package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title string
	Done  bool
}

var tasks []Task

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())

	if title == "" {
		fmt.Println("Task title can't be empty.")
		return
	}

	tasks = append(tasks, Task{
		Title: title,
		Done:  false,
	})
	fmt.Println("Task added successfully")

}

func listTask() {

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("Your Tasks: ")
	for i, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
	}
}

func markTaskDone(scanner *bufio.Scanner) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to mark.")
		return
	}
	listTask()

	fmt.Print("Enter the task number to mark as done: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	i, err := strconv.Atoi(input)

	if err != nil || i < 1 || i > len(tasks) {
		fmt.Println("Invalid task number")
		return
	}
	tasks[i-1].Done = true
	fmt.Printf("Task %d marked as done\n", i)

}

func deleteTask(scanner *bufio.Scanner) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to mark.")
		return
	}
	listTask()

	fmt.Print("Enter the task number to delete: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	i, err := strconv.Atoi(input)

	if err != nil || i < 1 || i > len(tasks) {
		fmt.Println("Invalid task number")
		return
	}

	tasks = append(tasks[:i-1], tasks[i:]...)
	fmt.Println("Task Deleted Successfully")

}

func main() {
	fmt.Println("-------- Welcome to Kartike's To-Do List app --------")
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Let's start, %s\n", name)

	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Mark Task as Done")
	fmt.Println("5. Exit")

	for {
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choiceStr := scanner.Text()
		choice, _ := strconv.Atoi(strings.TrimSpace(choiceStr))

		switch choice {

		case 1:

			// Add task

			addTask(scanner)

		case 2:
			// View tasks
			listTask()

		case 3:
			// Delete task
			deleteTask(scanner)

		case 4:
			// Mark task as done
			markTaskDone(scanner)

		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")

		}
	}

}

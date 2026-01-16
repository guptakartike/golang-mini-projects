package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

type Task struct{
	Title string
	Done bool
}


func main(){
	fmt.Println("-------- Welcome to Kartike's To-Do List app --------")
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name:= scanner.Text()
	fmt.Printf("Let's start, %s\n",name)

	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Mark Task as Done")
	
	
	for{
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choiceStr := scanner.Text()
		choice,_ := strconv.Atoi(strings.TrimSpace(choiceStr))

		switch choice{
		case 1:
			// Add task
		
			fmt.Println("1")
			continue
		case 2:
			// View tasks
			fmt.Println("2")
			continue
		case 3:
			// Delete task
			fmt.Println("3")
			continue
		case 4:
			// Mark task as done
			fmt.Println("4")
			continue
		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}
	}


	

}
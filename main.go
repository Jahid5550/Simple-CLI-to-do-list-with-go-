package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents a single to-do item
type Task struct {
	Description string
	Done        bool
}

func main() {
	//slice
	var tasks []Task
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Simple To-Do List ---")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Remove Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		// Read input as string, then convert to int
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			viewTasks(tasks)
		case 2:
			tasks = addTask(tasks, reader)
		case 3:
			tasks = markTaskDone(tasks, reader)
		case 4:
			tasks = removeTask(tasks, reader)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func viewTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("\nYour Tasks:")
	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "✔"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Description)
	}
}

func addTask(tasks []Task, reader *bufio.Reader) []Task {
	fmt.Print("Enter task description: ")
	desc, _ := reader.ReadString('\n')
	desc = strings.TrimSpace(desc)
	if desc == "" {
		fmt.Println("Task cannot be empty.")
		return tasks
	}
	tasks = append(tasks, Task{Description: desc, Done: false})
	fmt.Println("Task added.")
	return tasks
}

func markTaskDone(tasks []Task, reader *bufio.Reader) []Task {
	if len(tasks) == 0 {
		fmt.Println("No tasks to mark as done.")
		return tasks
	}
	viewTasks(tasks)
	fmt.Print("Enter task number to mark as done: ")
	input, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return tasks
	}
	tasks[index-1].Done = true
	fmt.Println("Task marked as done.")
	return tasks
}

func removeTask(tasks []Task, reader *bufio.Reader) []Task {
	if len(tasks) == 0 {
		fmt.Println("No tasks to remove.")
		return tasks
	}
	viewTasks(tasks)
	fmt.Print("Enter task number to remove: ")
	input, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return tasks
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("Task removed.")
	return tasks
}

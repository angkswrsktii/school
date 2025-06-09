package main

import (
	"Goland/app"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	app.InitializeActivities()
	fmt.Println("Welcome to School Activity App!")

	var choice string
	var exit = false

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("")
		fmt.Println("Select an action:")
		fmt.Println("1. List all activities")
		fmt.Println("2. Add a new activity")
		fmt.Println("3. Mark an activity as completed")
		fmt.Println("4. Remove an activity")
		fmt.Println("5. Exit")

		fmt.Print("Choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		switch choice {
		case "1":
			activities := app.GetAllActivities()
			if len(activities) == 0 {
				fmt.Println("You don't have any activities.")
				break
			}

			for id, activity := range activities {
				mark := "[]"
				if activity.IsCompleted() {
					mark = "[v]"
				}
				fmt.Println(id+1, ".", mark, activity.Name())
			}
			break

		case "2":
			fmt.Println("Enter your activity:")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			input = strings.TrimSpace(input)

			err = app.InsertActivity(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Success!")
			break

		case "3":
			fmt.Println("Enter activity id:")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			input = strings.TrimSpace(input)

			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}

			activity, err := app.GetActivity(id - 1)
			if err != nil {
				fmt.Println(err)
				continue
			}

			activity.ToggleCompletion()

			fmt.Println("Success!")
			break

		case "4":
			fmt.Println("Enter activity id:")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			input = strings.TrimSpace(input)

			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}

			err = app.RemoveActivity(id - 1)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Success!")
			break

		case "5":
			exit = true
			break

		default:
			fmt.Println("Invalid choice")
		}

		if exit {
			fmt.Println("Bye.")
			break
		}
	}
}

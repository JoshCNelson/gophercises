package cmd

import (
	"fmt"
	"gophercises/task_manager/db"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("you have not tasks to complete. Why not take a vacation")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

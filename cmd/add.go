package cmd

import (
	"fmt"

	"github.com/gustavop-fausto/todo-app/internal/storage"
	"github.com/gustavop-fausto/todo-app/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding new task to the list.",

	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.Load()
		if err != nil {
			return fmt.Errorf("erro ao carregar tasks: %w", err)
		}

		tasks = task.Add(tasks, args[0])

		err = storage.Save(tasks)
		if err != nil {
			return fmt.Errorf("erro ao adicionar task: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

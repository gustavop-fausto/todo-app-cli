package cmd

import (
	"fmt"

	"github.com/gustavop-fausto/todo-app/internal/storage"
	"github.com/gustavop-fausto/todo-app/internal/task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks registered",

	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.Load()
		if err != nil {
			return fmt.Errorf("erro ao carregar tasks: %w", err)
		}

		task.List(tasks)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

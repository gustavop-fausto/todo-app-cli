package cmd

import (
	"fmt"
	"strconv"

	"github.com/gustavop-fausto/todo-app/internal/storage"
	"github.com/gustavop-fausto/todo-app/internal/task"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark the task as done",

	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.Load()
		if err != nil {
			return err
		}

		if tasks, err = markAsDoneByID(tasks, args); err != nil {
			return err
		}

		if err = storage.Save(tasks); err != nil {
			return err
		}

		return nil
	},
}

func markAsDoneByID(tasks []task.Task, args []string) ([]task.Task, error) {
	for _, value := range args {
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("erro id inválido: %w", err)
		}

		tasks, err = task.MarkAsDone(tasks, id)
		if err != nil {
			return nil, fmt.Errorf("erro id inválido: %w", err)
		}
	}

	return tasks, nil
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

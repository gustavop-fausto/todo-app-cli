/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/gustavop-fausto/todo-app/internal/storage"
	"github.com/gustavop-fausto/todo-app/internal/task"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark the task as done",

	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.Load()
		if err != nil {
			return fmt.Errorf("erro ao carregar tasks: %w", err)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("erro id inválido: %w", err)
		}

		tasks, err = task.MarkAsDone(tasks, id)
		if err != nil {
			return fmt.Errorf("erro id inválido: %w", err)
		}

		err = storage.Save(tasks)
		if err != nil {
			return fmt.Errorf("erro ao adicionar task: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

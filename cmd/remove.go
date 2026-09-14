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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the task with the id given",

	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.Load()
		if err != nil {
			return fmt.Errorf("erro ao carregar tasks: %w", err)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("erro argumento dado não é um número: %w", err)
		}

		tasks, err = task.Remove(tasks, id)
		if err != nil {
			return fmt.Errorf("erro id fornecido inválido: %w", err)
		}

		err = storage.Save(tasks)
		if err != nil {
			return fmt.Errorf("erro ao adicionar task: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

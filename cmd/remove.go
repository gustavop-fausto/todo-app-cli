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
	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.Load()
		if err != nil {
			return fmt.Errorf("erro ao carregar tasks: %w", err)
		}

		var removedTasks []string
		var taskRemoved string
		for _, value := range args {
			id, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("erro argumento dado não é um número: %w", err)
			}

			tasks, taskRemoved, err = task.Remove(tasks, id)
			if err != nil {
				return fmt.Errorf("erro id fornecido inválido: %w", err)
			}

			removedTasks = append(removedTasks, taskRemoved)
		}

		err = storage.Save(tasks)
		if err != nil {
			return fmt.Errorf("erro ao adicionar task: %w", err)
		}

		for _, value := range removedTasks {
			fmt.Printf("🗑️  Tarefa removida: \"%s\"\n", value)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

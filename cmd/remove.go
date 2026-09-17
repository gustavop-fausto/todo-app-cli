package cmd

import (
	"fmt"
	"strconv"

	"github.com/gustavop-fausto/todo-app/internal/storage"
	"github.com/gustavop-fausto/todo-app/internal/task"
	"github.com/spf13/cobra"
)

var removeAll bool

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the task with the id given",

	Args: func(cmd *cobra.Command, args []string) error {
		if removeAll {
			return nil
		}
		return cobra.MinimumNArgs(1)(cmd, args)
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.Load()
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			return fmt.Errorf("🗑️ Não há nenhuma task para remover 🗑️")
		}

		tasks, removedTasks, err := runRemove(tasks, args)
		if err != nil {
			return err
		}

		if err = storage.Save(tasks); err != nil {
			return err
		}

		printingRemovedTasks(removedTasks)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolVarP(&removeAll, "all", "a", false, "Remove all tasks")
}

func runRemove(tasks []task.Task, args []string) ([]task.Task, []string, error) {
	var removedTasks []string

	if removeAll {
		tasks = []task.Task{}
	} else {
		var err error
		if tasks, removedTasks, err = removeTaskByID(tasks, args); err != nil {
			return nil, nil, err
		}
	}

	return tasks, removedTasks, nil
}

func removeTaskByID(tasks []task.Task, args []string) ([]task.Task, []string, error) {
	var removedTasks []string
	var taskRemoved string

	for _, value := range args {
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, nil, fmt.Errorf("erro argumento dado não é um número: %w", err)
		}

		tasks, taskRemoved, err = task.Remove(tasks, id)
		if err != nil {
			return nil, nil, fmt.Errorf("erro id fornecido inválido: %w", err)
		}

		removedTasks = append(removedTasks, taskRemoved)
	}

	return tasks, removedTasks, nil
}

func printingRemovedTasks(removedTasks []string) {
	if len(removedTasks) == 0 {
		fmt.Printf("🗑️  Todas as tarefas foram removidas com sucesso! 🗑️")
		return
	}

	for _, value := range removedTasks {
		fmt.Printf("🗑️  Tarefa removida: \"%s\"\n", value)
	}
}

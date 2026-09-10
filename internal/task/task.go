package task

import (
	"fmt"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Done        bool   `json:"done"`
	Description string `json:"description"`
}

func Add(tasks []Task, name, description string) []Task {
	tasks = append(tasks, Task{ID: nextID(tasks), Name: name, Done: false, Description: description})
	return tasks
}

func Remove(tasks []Task, id int) ([]Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, nil
		}
	}

	return nil, fmt.Errorf("erro id não encontrado")
}

func MarkDone(tasks []Task, id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return nil
		}
	}

	return fmt.Errorf("erro id não encontrado")
}

func nextID(tasks []Task) int {
	greater := 0
	for _, value := range tasks {
		if value.ID > greater {
			greater = value.ID
		}
	}

	return greater + 1
}

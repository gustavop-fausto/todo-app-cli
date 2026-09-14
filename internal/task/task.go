package task

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
)

type Task struct {
	ID   int    `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func Add(tasks []Task, todo string) []Task {
	tasks = append(tasks, Task{ID: nextID(tasks), Todo: todo, Done: false})
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

func MarkAsDone(tasks []Task, id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return nil
		}
	}

	return fmt.Errorf("erro id não encontrado")
}

func List(tasks []Task) {
	t := table.New(os.Stdout)

	t.SetHeaders("ID", "Task", "Done")
	t.SetHeaderStyle(table.StyleBold)
	t.SetLineStyle(table.StyleBrightBlue)
	t.SetBorders(true)
	t.SetDividers(table.UnicodeRoundedDividers)

	for _, value := range tasks {
		coloredText := coloringTasks(value.Todo)
		t.AddRow(strconv.Itoa(value.ID), tml.Sprintf(coloredText), strconv.FormatBool(value.Done))
	}

	t.Render()
}

func coloringTasks(todo string) string {
	colors := []string{
		"red", "green", "yellow", 
	"magenta", "white",
	}

	sortedColor := colors[rand.Intn(len(colors))]
	return fmt.Sprintf("<%s>%s<%s>", sortedColor, todo, sortedColor)
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

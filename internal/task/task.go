package task

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
	"github.com/mergestat/timediff"
)

type Task struct {
	ID        int       `json:"id"`
	Todo      string    `json:"todo"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
}

func Add(tasks []Task, todo string) []Task {
	tasks = append(tasks, Task{
		ID:        nextID(tasks),
		Todo:      todo,
		Done:      false,
		CreatedAt: time.Now().Add(-10 * time.Second),
	})

	return tasks
}

func Remove(tasks []Task, id int) ([]Task, string, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			taskRemoved := tasks[i].Todo
			tasks = append(tasks[:i], tasks[i+1:]...)

			return tasks, taskRemoved, nil
		}
	}

	return nil, "", fmt.Errorf("erro id não encontrado")
}

func MarkAsDone(tasks []Task, id int) ([]Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return tasks, nil
		}
	}

	return nil, fmt.Errorf("erro id não encontrado")
}

func List(tasks []Task) {
	t := table.New(os.Stdout)

	t.SetHeaders("ID", "Task", "Done", "Created At")
	t.SetHeaderStyle(table.StyleBold)

	t.SetLineStyle(table.StyleBrightBlue)
	t.SetBorders(true)
	t.SetDividers(table.UnicodeRoundedDividers)

	t.SetAlignment(table.AlignLeft, table.AlignLeft, table.AlignCenter)

	for i, value := range tasks {
		coloredText := coloringTasks(value.Todo, i)
		t.AddRow(strconv.Itoa(value.ID), tml.Sprintf(coloredText), isDone(value.Done), timediff.TimeDiff(value.CreatedAt))
	}

	t.Render()
}

func coloringTasks(todo string, i int) string {
	colors := []string{
		"red", "green", "yellow",
		"magenta", "white",
	}

	sortedColor := colors[i%len(colors)]
	return fmt.Sprintf("<%s>%s<%s>", sortedColor, todo, sortedColor)
}

func isDone(isDone bool) string {
	if isDone {
		return "✅"
	}
	return "❌"
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

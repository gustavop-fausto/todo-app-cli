package task

import (
	"slices"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name           string // Nome do teste que vai aparecer quando eu rodar
		tasks          []Task
		todoToAdd      string
		expectedID     int
		expectedLength int
	}{
		{
			name:           "Adicionar task em lista vazia",
			tasks:          []Task{},
			todoToAdd:      "Primeira tarefa",
			expectedID:     1,
			expectedLength: 1,
		},
		{
			name: "Adicionar task em lista já com valores definidos",
			tasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
			},
			todoToAdd:      "Tarefa 3",
			expectedID:     3,
			expectedLength: 3,
		},
		{
			name: "Adicionar task sem o reaproveitamento do ID removido",
			tasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
				{ID: 3, Todo: "Tarefa 3", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
			},
			todoToAdd:      "Tarefa 4",
			expectedID:     4,
			expectedLength: 3,
		},
	}

	for _, tt := range tests {
		// cria subtestes independetes que nao quebram caso algum teste falhe
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.tasks, tt.todoToAdd)

			// testa o tamanho final do slice
			if len(result) != tt.expectedLength {
				t.Fatalf("expected length of %d, got %d", tt.expectedLength, len(result))
			}

			newTask := result[len(result)-1]

			// testa se o nome da task foi salvo corretamente
			if tt.todoToAdd != newTask.Todo {
				t.Errorf("expected task %q, got %q", tt.todoToAdd, newTask.Todo)
			}

			// testa se o ID que ele gerou foi o correto para aquela posição
			if tt.expectedID != newTask.ID {
				t.Errorf("expected id %d, got %d", tt.expectedID, newTask.ID)
			}

			// testa se cada task inicializa como nao concluida
			if newTask.Done != false {
				t.Errorf("expect initialize the task as not done")
			}
		})
	}
}

func TestRemove(t *testing.T) {
	createdAt := time.Now().Add(-10 * time.Second)

	tests := []struct {
		name           string
		tasks          []Task
		remainingTasks []Task
		removedTask    string
		idToRemove     int
		wantErr        bool
	}{
		{
			name:           "Remover uma task de uma slice vazio",
			tasks:          []Task{},
			remainingTasks: []Task{},
			removedTask:    "",
			idToRemove:     1,
			wantErr:        true,
		},
		{
			name: "Remover uma task com um id inválido",
			tasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: createdAt},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: createdAt},
			},
			remainingTasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: createdAt},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: createdAt},
			},
			removedTask: "",
			idToRemove:  100,
			wantErr:     true,
		},
		{
			name: "Remover uma task com um id válido",
			tasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: createdAt},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: createdAt},
			},
			remainingTasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: createdAt},
			},
			removedTask: "Tarefa 2",
			idToRemove:  2,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			remainingTasks, removedTask, err := Remove(tt.tasks, tt.idToRemove)

			if (err != nil) != tt.wantErr {
				t.Errorf("expected error - %v, got %v: %v", tt.wantErr, err != nil, err)
			}

			if tt.wantErr {
				return
			}

			if !slices.Equal(tt.remainingTasks, remainingTasks) {
				t.Errorf("expected slice: %#v, got %#v", tt.remainingTasks, remainingTasks)
			}

			if tt.removedTask != removedTask {
				t.Errorf("expected removed task: %s, got %s", tt.removedTask, removedTask)
			}
		})
	}
}

func TestMarkAsDone(t *testing.T) {
	createdAt := time.Now().Add(-10 * time.Second)

	tests := []struct {
		name     string
		tasks    []Task
		idToMark int
		wantErr  bool
	}{
		{
			name:     "Marcar como concluído uma task em um slice vazio",
			tasks:    []Task{},
			idToMark: 1,
			wantErr:  true,
		},
		{
			name: "Marcar uma task como conclúida com um id inválido",
			tasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: createdAt},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: createdAt},
			},
			idToMark: 100,
			wantErr:  true,
		},
		{
			name: "Marcar uma task como conclúida com um id válido",
			tasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: createdAt},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: createdAt},
			},
			idToMark: 1,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tasks, err := MarkAsDone(tt.tasks, tt.idToMark)

			if (err != nil) != tt.wantErr {
				t.Errorf("expected erro: %v, got %v - %v", tt.wantErr, err != nil, err)
			}

			if tt.wantErr {
				return
			}

			for i := range tasks {
				if tasks[i].ID == tt.idToMark && tasks[i].Done == false {
					t.Errorf("expected the task marked as done it, got not done")
					return
				}
			}
		})
	}
}

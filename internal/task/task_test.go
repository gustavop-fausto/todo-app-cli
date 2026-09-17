package task

import (
	"testing"
	"time"
	"slices"
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
				t.Logf("expected task %q, got %q", tt.todoToAdd, newTask.Todo)
			}

			// testa se o ID que ele gerou foi o correto para aquela posição
			if tt.expectedID != newTask.ID {
				t.Logf("expected id %d, got %d", tt.expectedID, newTask.ID)
			}

			// testa se cada task inicializa como nao concluida
			if newTask.Done != false {
				t.Logf("expect initialize the task as not done")
			}
		})
	}
}

func TestRemove(t *testing.T) {
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
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
			},
			remainingTasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
			},
			removedTask: "",
			idToRemove:  100,
			wantErr:     true,
		},
		{
			name: "Remover uma task com um id válido",
			tasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
			},
			remainingTasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
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
				t.Logf("expected error - %v, got %v: %v", tt.wantErr, err != nil, err)
			}
			
			if tt.wantErr {
				return
			}

			if slices.Equal(tt.remainingTasks, remainingTasks) {
				t.Logf("expected slice: %v, got %v", tt.wantErr, err != nil)
			}

			if tt.removedTask != removedTask {
				t.Logf("expected removed task: %s, got %s", tt.removedTask, removedTask)
			}
		})
	}
}

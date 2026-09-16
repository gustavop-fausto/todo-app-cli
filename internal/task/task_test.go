package task

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name           string // Nome do teste que vai aparecer quando eu rodar
		initialTasks   []Task
		todoToAdd      string
		expectedID     int
		expectedLength int
	}{
		{
			name:           "Adicionar task em lista vazia",
			initialTasks:   []Task{},
			todoToAdd:      "Primeira tarefa",
			expectedID:     1,
			expectedLength: 1,
		},
		{
			name: "Adicionar task em lista já com valores definidos",
			initialTasks: []Task{
				{ID: 1, Todo: "Tarefa 1", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
				{ID: 2, Todo: "Tarefa 2", Done: false, CreatedAt: time.Now().Add(-10 * time.Second)},
			},
			todoToAdd:      "Tarefa 3",
			expectedID:     3,
			expectedLength: 3,
		},
		{
			name: "Adicionar task sem o reaproveitamento do ID removido",
			initialTasks: []Task{
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
			result := Add(tt.initialTasks, tt.todoToAdd)

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

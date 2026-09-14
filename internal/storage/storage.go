package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gustavop-fausto/todo-app/internal/task"
)

func GetFilePath() (string, error) {
	// Pega o path da home do usuário
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("erro ao tentar buscar path do home do usuário: %w", err)
	}

	path := filepath.Join(home, ".todo-cli")

	// Cria se não existir o diretório, não faz nada se já existir
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return "", fmt.Errorf("erro ao tentar criar diretório: %w", err)
	}

	// Retorna o fullpath do arquivo json
	return filepath.Join(path, "data.json"), nil
}

func Load() ([]task.Task, error) {
	// Pega o path do arquivo json
	path, err := GetFilePath()
	if err != nil {
		return nil, fmt.Errorf("erro ao tentar buscar path do arquivo json: %w", err)
	}

	// Tenta abrir o arquivo
	file, err := os.Open(path)
	if err != nil {
		// Caso o arquivo não exista, eu chamo o Save() para poder criar ele
		if os.IsNotExist(err) {
			err := Save([]task.Task{})
			if err != nil {
				return nil, fmt.Errorf("erro ao tentar criar arquivo json: %w", err)
			}

			return []task.Task{}, nil
		}
		return nil, fmt.Errorf("erro ao tentar ler arquivo json: %w", err)
	}

	defer func() { _ = file.Close() }()

	// Pegando informações para poder saber se o arquivo é vazio
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("erro ao tentar pegar informações do arquivo: %w", err)
	}

	if fileInfo.Size() == 0 {
		return []task.Task{}, nil
	}

	// Transforma o slice de bytes do json para uma struct em go
	var tasks []task.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("erro na decodificação do arquivo json: %w", err)
	}

	return tasks, nil
}

func Save(tasks []task.Task) error {
	// Pega o path do arquivo json
	path, err := GetFilePath()
	if err != nil {
		return fmt.Errorf("erro ao pegar o path do arquivo json: %w", err)
	}

	// Transforma o slice de tasks para um slice de bytes (só aceita bytes)
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("erro ao transformar tasks em um slice de bytes: %w", err)
	}

	// Sobrescreve tudo que tem no arquivo, criando caso não exista
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever em arquivo json: %w", err)
	}

	return nil
}

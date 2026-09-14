# 📝 todo-app

Uma CLI simples e direta para gerenciar suas tarefas do dia a dia, construída em **Go** com [Cobra](https://github.com/spf13/cobra).

```
$ todo add "Estudar Go" "Revisar Grafos"
✅ Tarefa adicionada: "Estudar Go"
✅ Tarefa adicionada: "Revisar PR"

$ todo list
╭────┬────────────────┬──────╮
│ ID │      Task      │ Done │
├────┼────────────────┼──────┤
│ 1  │ Estudar Go     │  ❌  │
│ 2  │ Revisar Grafos │  ❌  │
╰────┴────────────────┴──────╯
```

## ✨ Funcionalidades

- ➕ Adicionar uma ou várias tarefas de uma vez
- 📋 Listar todas as tarefas em uma tabela colorida no terminal
- ✅ Marcar tarefas como concluídas
- 🗑️ Remover tarefas específicas por ID, ou todas de uma vez
- 💾 Persistência automática em disco — os dados continuam lá mesmo depois de fechar o terminal

## 📦 Instalação

### Pré-requisitos

- [Go](https://go.dev/dl/) 1.27 ou superior instalado

### Compilando a partir do código-fonte

```bash
git clone https://github.com/gustavop-fausto/todo-app.git
cd todo-app
go build -o todo
```

Isso gera um binário `todo` na pasta atual. Para rodá-lo:

```bash
./todo --help
```

## 🚀 Uso

### Adicionar tarefas

```bash
todo add "Comprar leite"
todo add "Comprar leite" "Pagar a conta" "Ligar para o dentista"
```

Você pode adicionar uma ou várias tarefas de uma vez, cada uma como um argumento separado.

### Listar tarefas

```bash
todo list
```

Exibe todas as tarefas cadastradas em uma tabela, com o status de conclusão de cada uma.

### Marcar como concluída

```bash
todo done 1
```

Marca a tarefa de ID `1` como concluída.

### Remover tarefas

```bash
todo remove 2
todo remove 2 3 5        # remove várias de uma vez
todo remove --all         # remove todas as tarefas
todo remove -a             # forma curta de --all
```

## 🗂️ Onde os dados são salvos

As tarefas ficam armazenadas em um arquivo JSON no seu diretório home, independente de onde você executa o comando:

```
~/.todo-cli/data.json
```

O arquivo é criado automaticamente na primeira vez que você roda qualquer comando — não é necessário nenhum passo manual de configuração.

## 🏗️ Estrutura do projeto

```
todo-app/
├── main.go                    # ponto de entrada
├── cmd/                        # comandos da CLI (Cobra)
│   ├── root.go                 # comando raiz
│   ├── add.go
│   ├── list.go
│   ├── done.go
│   └── remove.go
└── internal/
    ├── task/
    │   └── task.go              # struct Task e regras de negócio
    └── storage/
        └── storage.go           # leitura/escrita do arquivo JSON
```

O projeto segue uma separação simples de responsabilidades:

- **`cmd/`** — só lida com entrada/saída do terminal (parsing de argumentos, flags, mensagens). Não contém lógica de negócio.
- **`internal/task/`** — regras de negócio puras (adicionar, remover, marcar como concluída), sem saber nada sobre arquivos ou terminal.
- **`internal/storage/`** — persistência em disco (ler e escrever o JSON), sem saber nada sobre regras de negócio.

## 🛠️ Tecnologias

- [Go](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra) — construção da CLI
- [aquasecurity/table](https://github.com/aquasecurity/table) — renderização da tabela no terminal
- [liamg/tml](https://github.com/liamg/tml) — cores no terminal

## 📄 Licença

Este projeto está sob a licença especificada no arquivo [LICENSE](./LICENSE).

---

Feito em Go, por [gustavop-fausto](https://github.com/gustavop-fausto).

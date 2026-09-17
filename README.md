# 📝 todo-app

Uma CLI simples para gerenciar tarefas do dia a dia, construída em **Go** com [Cobra](https://github.com/spf13/cobra).

![Preview do todo-app](./assets/preview.gif)

## ✨ Funcionalidades

- ➕ Adicionar uma ou várias tarefas de uma vez
- 📋 Listar todas as tarefas em uma tabela colorida no terminal
- ✅ Marcar tarefas como concluídas
- 🗑️ Remover tarefas específicas por ID, ou todas de uma vez
- 💾 Persistência automática em disco — os dados continuam lá mesmo depois de fechar o terminal

## 📦 Instalação

### Pré-requisitos

- Go 1.27 ou superior instalado

### Compilando a partir do código-fonte

Faça o clone do repositório e faça o build do projeto. Caso queira, use a flag `-o` para nomeáo-lo do jeito que preferir.

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
todo done 1 2 3
```

Você também pode marcar mais de uma tarefa como concluída de uma vez.

### Remover tarefas

```bash
todo remove 2
todo remove 2 3 5
todo remove --all
todo remove -a
```

## 🗂️ Onde os dados são salvos

As tarefas ficam armazenadas em um arquivo JSON no seu diretório home, independente de onde você executa o comando:

```
~/.todo-cli/data.json
```

## 🛠️ Tecnologias

- [Go](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra) — construção da CLI
- [aquasecurity/table](https://github.com/aquasecurity/table) — renderização da tabela no terminal
- [liamg/tml](https://github.com/liamg/tml) — cores no terminal

## 📄 Licença

Este projeto está sob a licença especificada no arquivo [LICENSE](./LICENSE).

---

Feito em Go, por [gustavop-fausto](https://github.com/gustavop-fausto).

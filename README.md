
# Bookingo - é o meu primeiro projeto simples feito na linguagem Golang. 


## Descrição do Projeto

É uma apis REST que tem como objetivo gerenciar os livors de uma livraria. É um projeto super simples que utilizei para aplicar os meus conhecimentos a medida que fui adquirindo no curso de Formação aprendendo GO.

## Funcionalidades - endpoints

- **Cadastrar novo livro**: POST /books.
- **Consulta de livro por ID:** GET /books/:id;
- **Lista todos os livros:** GET /books;
- **Atualiza informações do livro**: PUT /books/{id};
- **Publicar uma avaliação do livro por id**: PPST /:id/reviews; 

## Estrutura do Projeto

```bash
bookingo
├── cmd
│   ├── main.go
├── dados
│   ├── books.json
├── internal
│   ├── data
│   │   ├── book.go
│   ├── handler
│   │   ├── bookHandler.go
│   │   ├── reviewHandler.go
│   ├── models
│   │   ├── bookModel.go
│   │   ├── reviewModel.go
│   ├── service
│   │   ├── bookService.go
│   │   ├── bookService.go
└── go.mod
└── go.sum
 
```

## Passos para Executar o Projeto

### Pré-requisitos 

- **Go instalado na máquina**
- **Git instalado na maquina para efetuar o clone do projeto**

### Executar a Aplicação

1. Clone o repositório:

    ```bash
       git clone https://github.com/andresavasconcelos/bookingo-primeiro-projeto-em-go.git
       cd bookingo
    ```

2. Compile:
    ```bash
      go run .\cmd\main.go
    ```

3. Acesse o serviço:

  - Serviço: `http://localhost:8080`


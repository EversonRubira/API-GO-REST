# API GO REST - Personalidades

Este projeto é uma API REST desenvolvida em Go (Golang) que permite gerenciar informações sobre personalidades históricas. Ele utiliza o framework Gorilla Mux para roteamento e GORM para interagir com o banco de dados PostgreSQL. A API implementa operações CRUD (Create, Read, Update, Delete) para a manipulação de dados de personalidades.

## Funcionalidades

- **Listar todas as personalidades**: Retorna uma lista de todas as personalidades cadastradas no banco de dados.
- **Buscar uma personalidade por ID**: Retorna os detalhes de uma personalidade específica com base no ID fornecido.
- **Criar uma nova personalidade**: Permite a criação de uma nova personalidade através de um payload JSON.
- **Editar uma personalidade existente**: Atualiza os dados de uma personalidade com base no ID fornecido.
- **Deletar uma personalidade**: Remove uma personalidade do banco de dados com base no ID fornecido.

## Endpoints

### 1. Listar todas as personalidades
**GET** `/api/personalidades`

Retorna uma lista de todas as personalidades cadastradas no banco de dados.

**Exemplo de Resposta:**
```json
[
    {
        "id": 1,
        "nome": "Deodato Petit Wertheimer",
        "historia": "Deodato Petit Wertheimer foi um médico e político brasileiro..."
    },
    {
        "id": 2,
        "nome": "Carmela Dutra",
        "historia": "Carmela Teles Leite Dutra foi a primeira-dama do Brasil..."
    }
]


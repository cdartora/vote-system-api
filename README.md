# Vote System API
Uma API em Go + mongodb que gerencia um sistema de upvotes com usuários, autenticação de autorização.

## Rotas:
### **Users**
As rotas de usuários consistem em:
#### Registrar um novo usuário
- `POST /user`
  - Body: 
  `{
    name: "John Doe", "password": "123456"
  }`
  - Retorno: 201 Created

#### Login
- `POST /user/login`
  - Body: 
  `{
    name: "John Doe", "password": "123456"
  }`
  - Retorno: 200 OK
    - Body: `{"token": "token"}`

#### Retornar todos usuários
- `GET /user`
  - Necessário token válido na requisição
  - Retorno: 200
    - Body: `[{"id": "uuid", "name": "John Doe", "password": "123456"}, ..., ...]`

#### Retornar um usuário pelo id
- `GET /user/:id`
  - Necessário token válido na requisição
  - Retorno: 200
    - Body: `{"id": "uuid", "name": "John Doe", "password": "123456"}`

#### Atualizar um usuário pelo id
- `PATCH /user/:id`
  - Necessário token válido na requisição
  - Retorno: 200
    - Body: `{"name": "John Example", "password": "12345678"}`

#### Deletar um usuário pelo id
- `DELETE /user/:id`
  - Necessário token válido na requisição
  - Retorno: 200
---
### Coins
#### Registrar uma nova moeda
- `POST /coin`
  - Body: 
  `{ name: "Bitcoin", "code": "BTC" }`
  - Necessário token válido na requisição
  - Retorno: 200
  
#### Votar em uma moeda
- `POST /coin/:coin-id`
  - Body: 
  `{ vote: 1 }` (Upvote),
  `{ vote: -1}` (Downvote)
  - Necessário token válido na requisição
  - Retorno: 200
    - Body: 
  `{ votes: númeroDeVotosDaMoeda }`

#### Votos de todas as moedas
- `GET /coin`
  - Necessário token válido na requisição
  - Retorno: 200
    - Body: 
  `[{ "name": "Bitcoin", "code": "BTC", votes: númeroDeVotosDaMoeda }, ..., ...]`

#### Votos de uma moeda
- `GET /coin/:coin-id`
  - Necessário token válido na requisição
  - Retorno: 200
    - Body: 
  `{ "name": "Bitcoin", "code": "BTC", votes: númeroDeVotosDaMoeda }`

## Banco de dados
A estrutura do mongo, conectado através do mongo-driver, foi a seguinte:
Usuários estão presentes na coleção `users` com uma estrutura de id, nome e senha: 
```
{
  _id: "7e138192-4198-4ac5-a485-d40c96338079",
  name: "John Doe",
  password: "123456",
}
```
Já as moedas estão na coleção `coins` com uma estrutura de id, nome, código da moeda e votos:
```
{
  _id: "9b2038ed-ee8a-4d2e-a0d1-df38fc9919f5",
  name: "Bitcoin",
  code: "BTC",
  votes: [
    {
      userId: 7e138192-4198-4ac5-a485-d40c96338079,
      vote: 1,
    },
    {
      userId: 7e138192-4198-4ac5-a485-d40c96338079,
      vote: -1,
    },
    {
      userId: 7e138192-4198-4ac5-a485-d40c96338079,
      vote: 0,
    }
  ]
}
```
Sempre que uma nova moeda é criada um documento com estrutura de id do usuário e voto é gerado para cada usuário dentro da coleção `users`. Sendo assim, a contagem de votos de uma moeda se dá pela soma do valor `vote` dentro do campo `votes` de qualquer moeda. E, sempre que um usuário altera seu voto o campo vote do respectivo documento é alterado na coleção `coins`.

Deploy da API feito através da railway app neste [link](vote-system-api-production.up.railway.app).

**Aviso:** API não está com todas as rotas implementas, essa é apenas a documentação da estrutura planejada para o funcionamento da API. 

*Feito com ♡ por @cdartora.*

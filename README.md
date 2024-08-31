
# Desafio Clean Architecture

Especificações:

Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.




## Portas


| Serviço | Porta     |
| :-------- | :------- |
| `Webserver` | `:8000` |  
| `gRPC` | `:50051` | 
| `GraphQL` | `:8080` | 



## Deployment

Suba o banco de dados MySQL e o RabbitMQ

```bash
  docker-compose up -d
```
Execute o migration para criar a tabela no banco de dados

```bash
  migrate -path=internal/infra/database/migrations -database="mysql://root:root@tcp(localhost:3306)/orders_db" -verbose up
```
Inicie a aplicação

```bash
  migrate -path=internal/infra/database/migrations -database="mysql://root:root@tcp(localhost:3306)/orders_db" -verbose up
```
Inicie a aplicação no diretório cmd/ordersystem/

```bash
  go run main.go wire_gen.go
```
## Testando

### webserver

O arquivo create_order.http possui 3 requisões para testar o webserver. Execute os dois POSTs e o GET para listar as ordens.

### gRPC

Inicie o cliente evans pelo CLI e chame os serviços

```bash
  evans -r repl

  package pb
  service OrderService
  call CreateOrder
  call ListOrders
```
### GraphQL

Acesse o cliete na porta 8080 e execute as chamadas:

```graphql
mutation createOrder {
  createOrder(input: { id: "uu", Price: 855, Tax: 0.5 }) {
    id
    Tax
    Price
    FinalPrice
  }
}

query queryOrder {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```


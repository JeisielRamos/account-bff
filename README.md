
# Account-bff

API de transferencia entre contas Internas de um banco digital.


### Pré requisitos

- [Go (lang)](https://go.dev/)
- [Mysql](https://www.mysql.com/)


## Ajuda
- Você deve ter na raiz do projeto o arquivo `.env`, os nomes das variaveis estão no arquivo `.env.example`.
- A variavel `SECRET_JWT` pode ser qualquer senha aleatoria, pois ela é utilizada para gerar o Token.

Para rodar a aplicação, abra o terminal na raiz do projeto e execute o comando
```bash
 $ go run cmd/main.go
```

#### Comando para gerar as tabelas do banco

Tabela Contas
```bash
CREATE TABLE `accounts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `cpf` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `secret` varchar(100) NOT NULL,
  `balance` double DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `accounts_cpf_UN` (`cpf`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

```
    
Tabela Transferência
```bash
CREATE TABLE `transfers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account_origin_id` int NOT NULL,
  `account_destination_id` int NOT NULL,
  `amount` double DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `transfers_origin_FK` (`account_origin_id`),
  KEY `transfers_destination_FK` (`account_destination_id`),
  CONSTRAINT `transfers_destination_FK` FOREIGN KEY (`account_destination_id`) REFERENCES `accounts` (`id`),
  CONSTRAINT `transfers_origin_FK` FOREIGN KEY (`account_origin_id`) REFERENCES `accounts` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
## API Reference

### HealthCheck
```http
   GET /healthCheck
```
 - Verifica se a aplicação esta levantada

Retorno esperado 
```json
  { "status": "success" }
```
##
### Login
```http
  POST /login
```
 - Valida o usuário, se estiver correto irá retornar um Token para ser utilizado nos endpoits de Account e Transfer

Body 
```json
{   
  "cpf": "39341828015",
  "secret": "1234"
}
```
Retorno esperado 
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcGYiOiIzOTM0MTgyODAxNSIsImV4cCI6MTY2MDE4MjYzNn0.gjxR-cY9tKnSCYjC-hZEWCvLCD321ciUtQ6-MdCZRfU"
}
```
##
### Account

```http
  POST /api/accounts
```
- Cria uma nova conta
- O `balance` pode começar com 0;
- O `cpf` é unico;
Body 
```json   
{
  "name": "teste",
  "cpf": "39341828015",
  "secret": "1234",
  "balance": 0.0
}
```
Headers

| key | values     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`      | `string` | **Required**. o token que é retornado no login |

Retorno esperado 
```json
{
  "id": "1"
  "name": "teste",
  "cpf": "39341828015",
  "secret": "1234",
  "balance": 0.0
}
```
##

```http
  Get /api/accounts
```
- Retorna todas as contas

Headers

| key | values     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`      | `string` | **Required**. o token que é retornado no login |

Retorno esperado 
```json
[{
  "id": "1"
  "name": "teste",
  "cpf": "39341828015",
  "secret": "1234",
  "balance": 0.0
},{
  "id": "2"
  "name": "teste2",
  "cpf": "25991674000",
  "secret": "1234",
  "balance": 0.0
}]
```
##
```http
  Get /api/accounts/{account_id}/balance
```
- Retorna o saldo de uma conta

Params - Path Variables

| key | values     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `account_id`      | `int` | **Required**. Deve ser passado o id de uma conta |

Headers

| key | values     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`      | `string` | **Required**. o token que é retornado no login |

Retorno esperado 
```json
{
    "account_id": 1,
    "balance": 0
}
```
##

### Transfer

```http
  POST /api/transfers
```
- Realiza a transferencia de uma conta para outra;
- O valor da transferencia não pode ser maior que o saldo da conta de origem;
- Deve ser passado o id da conta de destino `(account_destination_id)` e o valor da transferência `(amount)`
Body 
```json   
{
    "account_destination_id": "2",
    "amount": 2.0
}
```
Headers

| key | values     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`      | `string` | **Required**. o token que é retornado no login |

Retorno esperado 
```json
{
    "id": 1,
    "account_origin_id": "1",
    "account_destination_id": "2",
    "amount": 2,
    "created_at": "2022-08-11T11:38:17"
}
```
##
```http
  Get /api/transfers
```
- Retorna as transferências realizadas pela conta `(que vem através do token)`

Headers

| key | values     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `token`      | `string` | **Required**. o token que é retornado no login |

Retorno esperado 
```json
[
    {
        "id": 1,
        "account_origin_id": "1",
        "account_destination_id": "2",
        "amount": 2,
        "created_at": "2022-08-10 19:38:21"
    },
    {
        "id": 2,
        "account_origin_id": "1",
        "account_destination_id": "2",
        "amount": 2,
        "created_at": "2022-08-10 21:51:17"
    }
]
```
##

Todos os Endpoits utiliza esse retorno de erro
```json
{
    "StatusCode": 400,
    "Message": "Error "
}
```
## Ajustes a fazer

- Ajustar o retorno do campo `created_at` no endpoint `POST /api/transfers`.
- Ajustar o Docker, está dando erro na conecção com o mysql
- Incluir `swagger` para a documentação, assim deixando o readme mais limpo.
- Incluir o Rollback nas alterações feitas no banco caso de algum erro. `(endpoint POST /api/transfers)`

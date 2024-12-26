# Wallet UUID.

## Технологии и инструменты:
- Go
- PostgreSQL
- Docker
- Docker-compose
- ENV
- Echo
- Тесты 
- Крипто

## Запросы:

- POST: 
Запрос на создание кошелька `/api/v1/create`

{
  "valletId": "123e4567-e89b-12d3-a456-426614174001"
}





- POST: 
Запрос на обновнение кошелька `/api/v1/wallet`

{
  "valletId": "123e4567-e89b-12d3-a456-4266141740001",
  "operationType": "`DEPOSIT`" or "`WITHDRAW`",
  "amount": 1000
}



- GET: 
Запрос на баланс кошелька `/api/v1/wallets/{WALLET_UUID}`

Вместо `{WALLET_UUID}` - номер кошелька, например `/api/v1/wallets/123e4567-e89b-12d3-a456-426614174001`


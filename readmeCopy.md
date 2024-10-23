### API Documentation

1. `CREATE` Register User
    - `POST` -> `/api/v1/auth/register`
    - Request:
      ```json
      {
         "username": "hunayn",
         "password": "rahasia"
      }
      ``` 
    - Response:
      ```json
      "status": {
         "code": 201,
         "message": "Created"
      },
      "data": {
        "id_user": "95c4507a-7dca-4f6b-8c5f-200d894c7cc2",
        "name": "hunayn",
        "password": "$2a$10$Nrvz/c0orjmOpFn96YyCOexqBR64z/Brm.SW0/thia6.R6stPUBYW",
        "role": "employee"

      }
2. `POST` Create New Product
     - `POST` -> `/api/v1/product`
    - Request:
      ```json
      {
         "nameProvider": "Indosat",
         "nominal": 5000,
         "price": 6000,
         "idSupliyer": "ce18dad5-c003-4b98-b811-7d372ca75439"
      }
      ``` 
    - Response:
      ```json
      "status": {
         "code": 201,
         "message": "Created"
      },
      "data": {
        "idProduct": "blabla",
        "nameProvider": "Indosat",
         "nominal": 5000,
         "price": 6000,
         "idSupliyer": "ce18dad5-c003-4b98-b811-7d372ca75439"
      }

3. `POST` Create New Merchant
     - `POST` -> `/api/v1/merchant`
    - Request:
      ```json
      {
        "idUser": "35706923-3177-4731-9247-cd33abcb9944",
         "nameMerchant": "Konter Pak Eko",
         "address": "jakarta",
         "idProduct": "72aebbb7-a955-4614-afc7-9275fc58dc22",
         "balance": 0
      }
      ``` 
    - Response:
      ```json
      "status": {
         "code": 201,
         "message": "Created"
      },
      "data": {
       {
        "idMerchant": "uuid_merchant",
        "idUser": "35706923-3177-4731-9247-cd33abcb9944",
         "nameMerchant": "Konter Pak Eko",
         "address": "jakarta",
         "idProduct": "72aebbb7-a955-4614-afc7-9275fc58dc22",
         "balance": 0
      }
      }
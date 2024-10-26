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
      - Status Code: 201 Created
      - Body :

```json
{
  "Message": "string",
  "data": {
        "id_user": "95c4507a-7dca-4f6b-8c5f-200d894c7cc2",
        "name": "hunayn",
        "password": "$2a$10$Nrvz/c0orjmOpFn96YyCOexqBR64z/Brm.SW0/thia6.R6stPUBYW",
        "role": "employee"

      }
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

4. `POST` Create New Transaction
     - `POST` -> `/api/v1/transaction`
    - Request:
      ```json
      {
        "merchantId": "",
         "userId": "",
         "customerName":"",
         "destinationNumber":"",
         "transactionDate":"",
         "transactionDetail": [
          {
            "productId":""
          }
         ]
      }
      ``` 
    Response :

- Status Code: 201 Created
- Body :

```json
{
    "message": "Transaction list",
    "data": 
        {
            "transactionId": "83881046-778b-4ad5-b1a0-d5d96f44de03",
            "customerName": "rahman",
            "destinationNumber": "087654326453",
            "user": {
                "id_user": "53a6e29c-fe7e-4f41-9c82-76e188513cff",
                "name": "hunayn",
                "role": "employee"
            },
            "merchant": {
                "idMerchant": "ba61ec93-a8cf-4798-9c1c-2d617fcf3d58",
                "nameMerchant": "Konter Pak Eko",
                "address": "jakarta"
            },
            "transactionDate": "2024-10-25T00:00:00Z",
            "transactionDetail": [
                {
                    "transactionDetailId": "120940bc-87c5-4575-ae2d-0a6859531b5c",
                    "transactionId": "83881046-778b-4ad5-b1a0-d5d96f44de03",
                    "product": {
                        "idProduct": "6fb3f317-aa06-4cf0-9cd2-0f6c5d5915ec",
                        "nameProvider": "Indosat",
                        "nominal": 5000,
                        "price": 6000
                    }
                }
            ]
        }
}
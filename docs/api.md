## **Add Account**

Creates new account.

- **URL**

  /accounts

- **Method:** `POST`

- **Data Params**

  ```
  {"id":"john789", "balance":10, "currency":"AUD"}
  ```

- **Success Response:**

  - **Code:** 200 <br />
    **Content:** `{}`

* **Error Response:**

  Happens when the account is already registered on the database

  - **Code:** 200 <br />
    **Content:** `{"err": "account already exist"}`

- **Sample Call:**
  ```
    curl --location --request POST 'http://localhost:8080/v1/accounts' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "id": "bob123",
          "balance": 100,
          "currency": "USD"
      }'
  ```

## **Show Accounts**

Returns all available accounts.

- **URL**

  /accounts

- **Method:** `GET`

- **Success Response:**

  - **Code:** 200 <br />
    **Content:** `{"res":[{"id":"bob123","balance":100,"currency":"USD"}]}`

- **Sample Call:**
  ```
    curl --request GET 'http://localhost:8080/v1/accounts'
  ```

## **Show Payments**

Returns all recorded payments.

- **URL**

  /payments

- **Method:** `GET`

- **Success Response:**

  - **Code:** 200 <br />
    **Content:** `{"res":[{"account":"bob123","amount":100,"to_account":"alice456","direction":"outgoing"}]}`

- **Sample Call:**
  ```
    curl --request GET 'http://localhost:8080/v1/payments'
  ```

## **Transfer**

Sending a payment from one account to another.

- **URL**

  /transfer

- **Method:** `POST`

- **Data Params**

  ```
  {"senderId":"bob123", "receiverId":"alice456", "amount":100}
  ```

- **Success Response:**

  - **Code:** 200 <br />
    **Content:** `{}`

* **Error Response:**

  - Happens when the amount inserted is not more than zero
    
    **Code:** 200 <br />
    **Content:** `{"err": "amount should be greater than 0"}`

  - Happens when the sender or receiver not found in the database
    
    **Code:** 200 <br />
    **Content:** `{"err": "record not found"}`

  - Happens when sender doesn't have enough balance in their wallet
    
    **Code:** 200 <br />
    **Content:** `{"err": "insufficient balance"}`

  - Happens when sender and receiver have different currency registered in their account
    
    **Code:** 200 <br />
    **Content:** `{"err": "currency mismatch"}`

- **Sample Call:**
  ```
    curl --location --request POST 'http://localhost:8080/v1/transfer' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "senderId": "bob123",
          "receiverId": "alice456",
          "amount": 100
      }'
  ```
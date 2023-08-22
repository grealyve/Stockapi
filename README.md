# StocksAPI Documentation

Welcome to the documentation for the StocksAPI, a simple RESTful API for managing stock information. This API is built using the Go programming language and the Gorilla Mux router.

## Table of Contents

- [Endpoints](#endpoints)
  - [Get Stock](#get-stock)
  - [Get All Stocks](#get-all-stocks)
  - [Create Stock](#create-stock)
  - [Update Stock](#update-stock)
  - [Delete Stock](#delete-stock)
- [Setup](#setup)
- [Usage](#usage)

## Endpoints

### Get Stock

Get information about a specific stock by providing its ID.

- **Endpoint:** `/api/stock/{id}`
- **Method:** GET
- **Response:**
  ```json
  {
    "name": "Apple Inc."
    "price": 150,
    "company":"Apple"
  }
  ```

### Get All Stocks

Get a list of all available stocks.

- **Endpoint:** `/api/stock`
- **Method:** GET
- **Response:**
  ```json
  {
    "stockid":1,
    "name": "Apple Inc.",
    "price": 150,
    "company":"Apple"
  },
  // Other stock objects
  ```
![getall](https://github.com/grealyve/Stockapi/assets/41903311/47ed9dd5-7e2c-429b-bd14-9fde15b41b11)


### Create Stock

Create a new stock entry.

- **Endpoint:** `/api/newstock`
- **Method:** POST
- **Response:**
  ```json
  {
    "id": 1,
    "message": "stock created successfully",
  }
  ```
- Sample Request:
![newstock](https://github.com/grealyve/Stockapi/assets/41903311/6801eb35-afd1-433e-977d-728850d5738f)

### Update Stock

Update information about a specific stock by providing its ID.

- **Endpoint:** `/api/stock/{id}`
- **Method:** PUT
- **Response:**
  ```json
  {
    "id": 1,
    "message": "stock created successfully",
  }
  ```
- Sample Request:
![update](https://github.com/grealyve/Stockapi/assets/41903311/ff526d4e-59b1-4b55-bbe2-45868d3545ae)

### Delete Stock

Delete a stock entry by providing its ID.

- **Endpoint:** `/api/delete/{id}`
- **Method:** DELETE
- Sample request and response:
![delete](https://github.com/grealyve/Stockapi/assets/41903311/3175d343-1b56-4f48-a595-883e4034f343)

## Setup
To run the StocksAPI on your local machine, follow these steps:

1. Clone this repository.
2. Make sure you have Go (Golang) installed on your system.
3. Run the following command to start the server:
```go
go run main.go
```
4. The API will be accessible at http://localhost:8000
5. To be able to create the necessary SQL table, you need to execute these commands in the postgresql docker:
```bash
psql -U grealyve -d stockapi
CREATE DATABASE stockapi;
CREATE TABLE stocks(stockid SERIAL PRIMARY KEY, name TEXT, price INT, company TEXT);
```

## Usage
Feel free to use and modify this API for your own projects. You can integrate it with databases, add authentication, and more. If you encounter any issues or have suggestions, please open an issue on this repository.

### Referrence:
https://youtu.be/1nLH4J-DRLg?list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9

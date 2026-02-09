# **Placebo Palace**

![Placebo Palace Logo](assets/drugs.webp)

Welcome to **Placebo Palace**, an online shop for drugs. This repository contains the **backend API** for the online shop.

## **Features**

* **Authentication** – simple login/logout with JWT tokens
* **Product Listing** – view all drug products
* **Product Details** – fetch details for individual drug
* **Checkout** – place orders with token-based authentication

## **API Endpoints**

| Method | Endpoint               | Description                      |
| ------ | ---------------------- | -------------------------------- |
| POST   | `/auth/login`          | Log in and receive a JWT token   |
| POST   | `/auth/logout`         | Log out (dummy endpoint)         |
| GET    | `/products`            | List all products                |
| GET    | `/products/{id}`       | Get details of a single product  |
| POST   | `/checkout/placeorder` | Place an order (JWT required)    |

## **Getting Started**

1. **Clone the repository**

```bash
git clone https://github.com/inf23056/sose-26-devops.git
cd sose-26-devops
```

2. **Install dependencies**

```bash
go mod tidy
```

3. **Debugging**

   Please note that a valid and compatible Delve installation is necessary for this.

   In order to debug, use Visual Studio Code as IDE. A `launch.json` is set up in the `.vscode` folder which starts the Go server and passes through environment variables (e.g. the path to dlv for Delve).

   

4. **Run the server**

```bash
go run main.go
```

5. **Access the API**

   Server will run on `http://localhost:8080`. Use Postman, curl or any other API client to test the endpoints.

## **Authentication**

* **Login credentials (for testing):**

  * Username: `user`
  * Password: `pass`
* Include the returned JWT token in `Authorization: Bearer <token>` for protected endpoints

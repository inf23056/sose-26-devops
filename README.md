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

## Git Conventions

### 1. Commit Messages

**Recommended Format:**

&lt;type&gt;(&lt;scope&gt;): &lt;short description in present tense&gt;

&lt;optional: longer explanation&gt;

&lt;optional: footer (Issues, Breaking Changes)&gt;

**Tense & Style**

Present tense, imperative:

- add login validation
- fix null pointer in payment service
- remove deprecated API endpoint

Not recommended:

- added login validation
- fixes bug
- login validation added

**Types**

| Type     | Meaning                                  |
|----------|------------------------------------------|
| feat     | New feature                              |
| fix      | Bug fix                                  |
| docs     | Documentation                            |
| style    | Formatting (no logic change)             |
| refactor | Code restructuring                       |
| test     | Tests                                    |
| chore    | Build, CI, Dependencies                  |
| perf     | Performance improvement                  |

**Examples:**

Simple commits:
- `feat(auth): add JWT refresh token functionality`
- `fix(api): handle empty response in user endpoint`
- `docs(readme): update installation instructions`
- `chore(deps): delve to v1.22.0`

With scope and detailed description:
- `feat(user-profile): add avatar upload with image validation`
- `fix(checkout): prevent duplicate payment submissions`
- `refactor(api): extract validation logic into separate service`
- `perf(database): add index on user_id for faster queries`

Multi-line commit with body:
```
feat(auth): implement two-factor authentication

Add SMS and authenticator app support for 2FA.
Users can enable this in account settings.
```

Breaking changes:
```
feat(api)!: change user endpoint response structure

BREAKING CHANGE: The user endpoint now returns `userData` 
instead of `user` to align with other endpoints.
```

---

### 2. Branch Names

**General Pattern:**

&lt;type&gt;/&lt;short-description&gt;

**Common Types:**

- feature/
- bugfix/
- hotfix/
- release/
- chore/
- experiment/

**Examples:**

Feature branches:
- `feature/user-authentication`
- `feature/123-add-payment-integration`
- `feature/dashboard-redesign`

Bugfix branches:
- `bugfix/456-fix-login-redirect`
- `bugfix/null-pointer-in-cart`
- `bugfix/safari-css-rendering`

Hotfix branches:
- `hotfix/payment-timeout`
- `hotfix/critical-security-patch`
- `hotfix/database-connection-leak`

Release and chore branches:
- `release/v2.1.0`
- `chore/update-dependencies`
- `chore/migrate-to-typescript`

**Rules:**

- all lowercase  
- use hyphens, no spaces  
- keep it short but clear
- optionally prefix with ticket/issue number

## Docker

Docker Hub:

 [inf23056/sose-26-devops](https://hub.docker.com/repository/docker/inf23056/sose-26-devops/general)


### Build

docker build -t sose-26-devops:`major.minor.patch` 

### Run

docker run -p 8080:8080 sose-26-devops:`major.minor.patch`


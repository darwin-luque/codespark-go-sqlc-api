# CodeSpark - Realworld API with Golang and SQLC

> **Notes**:
>
> - This guide will be chaning as we move forward with the project. Please check for changes as we move forward through new branches
> - The guide is in English but the session itself will be in Spanish

This is a simple API that uses Golang and SQLC to interact with a PostgreSQL database. This is an API that allows to authenticate users, manage articles, comment on articles, favorite articles, follow users, and manage user profiles. The objectve of this project is to demonstrate how to build scalable, maintainable, and testable APIs using Golang and SQLC. We'll dive into the following topics:

- Use SQLC to generate type-safe Go code from SQL queries
- Start an HTTP server using the standard library
- Authenticate and manage users using JWT and bcrypt
- Proper naming conventions and best practices for endpoints
- Design database schemas for scalability and maintainability
- Using a domain-driven design approach to structure the codebase

## How we will tackle this CodeSpark

This project will be divided into multiple branches. Each branch will cover a specific topic. The branches will be as follows:

- **start-here**: This branch will contain the initial setup of the project. It will have the basic structure of the project and the initial setup of the database (using docker-compose file).
- **sqlc-setup**: This branch will cover the setup of SQLC. We will write some SQL queries and generate Go code using SQLC.
- **user-authentication**: This branch will cover the user authentication. We will use JWT for authentication and bcrypt for password hashing. We will implement the endpoints to manage users (CRUD operations).
- **articles**: This branch will cover the articles. We will implement the endpoints to manage articles (CRUD operations).
- **comments**: This branch will cover the comments. We will implement the endpoints to manage comments (CRUD operations).
- **favorites**: This branch will cover the favorites. We will implement the endpoints to manage favorites (CRUD operations).
- **follows**: This branch will cover the follows. We will implement the endpoints to manage follows (CRUD operations).
- **pre-launch**: This branch will cover the final touches before launching the API. We will fine-tune the API, add some tests, have a well-made Dockerfile, and make sure everything is working as expected.

## Prerequisites

- Go 1.21 or higher. You can download it [here](https://golang.org/dl/)
- Docker and Docker Compose. You can download it [here](https://www.docker.com/products/docker-desktop)
- SQLC. You can download it [here](https://docs.sqlc.dev/en/latest/overview/install.html)
- Air for live reload. You can download it [here](https://github.com/air-verse/air)

## Getting Started

1. Clone the repository

   ```bash
   git clone github.com/darwin-luque/codespark-go-sqlc-api
   ```

2. Change directory

   ```bash
   cd codespark-go-sqlc-api
   ```

3. Start the PostgreSQL database

   ```bash
   docker compose up -d
   ```

4. Run the main.go file with air

   ```bash
   air
   ```

## API Endpoints

### Authentication

> **Note**: You can check how implement these endpoints in the `e2e/auth.rest` file.

- POST /api/users/sign-up
- POST /api/users/sign-in
- GET /api/users/me

### Articles

> **Note**: You can check how implement these endpoints in the `e2e/articles.rest` file.

- POST /api/articles
- GET /api/articles
- GET /api/articles/:slug
- PUT /api/articles/:slug
- DELETE /api/articles/:slug
- POST /api/articles/:slug/publish
- GET /api/articles/favorites

### Comments

> **Note**: You can check how implement these endpoints in the `e2e/comments.rest` file.

- POST /api/articles/:slug/comments
- GET /api/articles/:slug/comments
- DELETE /api/comments/:id

### Favorites

> **Note**: You can check how implement these endpoints in the `e2e/favorites.rest` file.

- POST /api/articles/:slug/favorite
- DELETE /api/articles/:slug/favorite

## Tools Used

- [VSCode](https://code.visualstudio.com/)
- [TablePlus](https://tableplus.com/)
- [Warp](https://warp.dev/)

In the `.vscode` folder, you will find the recommended extensions

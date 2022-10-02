# Build Golang RESTful API with Gorm, Gin and Postgres

## 1. How to Setup Golang GORM RESTful API Project with Postgres

![How to Setup Golang GORM RESTful API Project with Postgres](https://codevoweb.com/wp-content/uploads/2022/08/How-to-Setup-Golang-GORM-RESTful-API-Project-with-Postgres.webp)

This article will guide you on how you can set up a Golang project with the GORM library and PostgreSQL to build CRUD RESTful API to perform the basic Create/Get/Update/Delete operations

### Topics Covered

- What is an ORM?
- Setup the Golang Project
    - Initialize the Golang Project
    - Create a PostgreSQL Docker Container
    - Load and Validate the Environment Variables
    - Create a Utility Function to Connect to PostgreSQL
- Data Modeling and Migration with GORM
    - Creating the Database Model with GORM
    - Install the UUID OSSP Module for PostgreSQL
    - Migrating the Schema with GORM
- Create the Golang Server with Gin Gonic
- Testing the Golang API

Read the entire article here: [https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres](https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres)

## 2. API with Golang + GORM + PostgreSQL: Access & Refresh Tokens

![API with Golang + GORM + PostgreSQL: Access & Refresh Tokens](https://codevoweb.com/wp-content/uploads/2022/08/API-with-Golang-GORM-PostgreSQL-Access-Refresh-Tokens.webp)

This article will teach you how to secure a Golang application with access and refresh tokens using GORM, Postgres, Docker, and Docker-compose. Also, you will learn how to use the RS256 algorithm with private and public keys to sign the tokens.

### Topics Covered

- Golang, Gin & GORM JWT Authentication Overview
- JWT Authentication Example with Golang and GORM
- Generate the Private and Public Keys
- Load and Validate the Environment Variables
- Create the Database Models with GORM
- Run the Database Migration with GORM
- Hash and Verify the Passwords with Bcrypt
- Sign and Verify the RS256 JSON Web Tokens
    - Function to Generate the Tokens
    - Function to Verify the Tokens
- Create the Authentication Route Controllers
    - Register User Controller
    - Login User Controller
    - Refresh Access Token Controller
    - Logout Controller
- Create a Middleware Guard
- Create a User Controller
    - Authentication Routes
    - User Routes
- Add the Routes to the Gin Middleware Stack

Read the entire article here: [https://codevoweb.com/golang-gorm-postgresql-user-registration-with-refresh-tokens](https://codevoweb.com/golang-gorm-postgresql-user-registration-with-refresh-tokens)


## 3. Golang and GORM - User Registration and Email Verification

![Golang and GORM - User Registration and Email Verification](https://codevoweb.com/wp-content/uploads/2022/08/Golang-and-GORM-User-Registration-and-Email-Verification.webp)

In this comprehensive guide, you will learn how to secure a Golang RESTful API with JSON Web Tokens and Email verification. We will start by registering the user, verifying the user's email address, logging in the registered user, and logging out the authenticated user.

### Topics Covered

- Golang and GORM JWT Authentication Overview
- Create the Database Models with GORM
- Database Migration with GORM
- Generate and Verify the Password with Bcrypt
- Sign and Verify the JWT (JSON Web Tokens)
    - Update the Environment Variables File
    - Validate the Variables with Viper
    - Generate the JSON Web Tokens
    - Verify the JSON Web Tokens
- Create the SMTP Credentials
- Setup the HTML Templates
    - Add the HTML Email Base Template
    - Add the HTML Email CSS Styles
    - Add the Email Verification Template
    - Create the Email Utility Function
- Create the Controller Functions
    - Function to Generate the Verification Code
    - User Registration Controller
    - Verify Email Controller
    - Login User Controller
    - Logout User Controller
    - Get User Profile Controller
- Create the Authentication Guard
- Create Routes for the Controllers
    - Auth Routes
    - User Routes
- Register the Routes and Start the Golang Server

Read the entire article here: [https://codevoweb.com/golang-and-gorm-user-registration-email-verification](https://codevoweb.com/golang-and-gorm-user-registration-email-verification)

## 4. Forgot/Reset Passwords in Golang with SMTP HTML Email

![Forgot/Reset Passwords in Golang with SMTP HTML Email](https://codevoweb.com/wp-content/uploads/2022/08/Forgot-Reset-Passwords-in-Golang-with-SMTP-HTML-Email.webp)

This article will teach you how to add a secure forgot/reset password feature to a Golang RESTful API application. We will generate the HTML Email templates with the standard Golang template package and send them via SMTP with the Gomail package.

### Topics Covered

- Forgot Password and Password Reset Flow
- Create the Database Models with GORM
- Create an SMTP Account
- Setup the HTML Templates
    - Add the Email Template CSS
    - Add the Password Reset HTML Template
- Encoding/Decoding the Password Reset Code
- Create a Utility Function to Send the Emails
- Add the Forgot Password Route Handler
- Add the Reset Password Route Handler
- Add the API Routes to the Gin Middleware Stack

Read the entire article here: [https://codevoweb.com/forgot-reset-passwords-in-golang-with-html-email](https://codevoweb.com/forgot-reset-passwords-in-golang-with-html-email)

## 5. Build a RESTful CRUD API with Golang

![Build a RESTful CRUD API with Golang](https://codevoweb.com/wp-content/uploads/2022/08/Build-a-RESTful-CRUD-API-with-Golang.webp)

This article will teach you how to create a CRUD RESTful API in a Golang environment that runs on a Gin Gonic server and uses a PostgreSQL database. We’ll also discuss how you can build models, connect to the running SQL database server, and run database migrations with the GORM library.

### Topics Covered

- What is a REST API
- What is a REST API
- What is GORM
- What is Gin Gonic
- Create the Database Models with GORM
- Generating the SQL Table with GORM
- Creating CRUD Functions in a RESTful API
    - Create a Constructor for the CRUD Operations
    - Create Operation Route Handler
    - Update Operation Route Handler
    - Retrieve a Single Record Route Handler
    - Retrieve All Records Route Handler
    - Delete Operation Route Handler
- Creating Routes for the CRUD Operations
- Update/Configure the Golang API Server
- Testing the Golang CRUD API with Postman
    - Log into the API
    - Creating a New Record
    - Updating the Record
    - Request a Single Record
    - Retrieve all Records with Paginated Results
    - Delete a Record

Read the entire article here: [https://codevoweb.com/build-restful-crud-api-with-golang](https://codevoweb.com/build-restful-crud-api-with-golang)

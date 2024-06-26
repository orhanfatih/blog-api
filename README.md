# Blog API

This is a simple RESTful API for a blogging platform built with Go.

## Features

- User registration and authentication with JWT token-based authentication
- User management: create, read, update, delete user profiles
- Post management: create, read, update, delete blog posts

### Tech Stack
- Go, Echo, Gorm, PostgreSQL, JWT

## API Endpoints

### Auth Endpoints

- `POST v1/auth/register`: Register a new user
- `POST v1/auth/login`: Authenticate and obtain a JWT token
- `POST v1/auth/logout`: Logout and invalidate the JWT token.

### User Endpoints

- `GET v1/user/me`: Get user profile
- `PATCH v1/user/`: Update user profile
- `DELETE v1/user/`: Delete user profile

### Blog Post Endpoints

- `POST v1/posts/`: Create a new blog post
- `GET v1/posts/:id`: Get a blog post by ID
- `PUT v1/posts/:id`: Update a blog post
- `DELETE v1/posts/:id`: Delete a blog post
- `GET v1/posts/`: Get blog posts

## Requirements:

* Docker

# How to run

## Running the Application

1. **Set up environment variables:**
   - Create a file named `.env` in the root directory of your project.
   - Inside the `.env` file, define your environment variables (copy fields from .env.template) and fill values

2. **Build and Run the application:**
   ```bash
   docker-compose up --build
   ```
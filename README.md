# Blog API

This is a simple RESTful API for a blogging platform built with Go.

## Features

- User registration and authentication with JWT token-based authentication
- User management: create, read, update, delete user profiles
- Post management: create, read, update, delete blog posts

## API Endpoints

### Auth Endpoints

- `POST v1/auth/register`: Register a new user
- `POST v1/auth/login`: Authenticate and obtain a JWT token
- `POST v1/auth/logout`: Logout and invalidate the JWT token.

### User Endpoints

- `GET v1/users/me`: Get user profile
- `PUT v1/users/`: Update user profile
- `DELETE v1/users/`: Delete user profile

### Blog Post Endpoints

- `POST v1/posts`: Create a new blog post
- `GET v1/posts/:id`: Get a blog post by ID
- `PUT v1/posts/:id`: Update a blog post
- `DELETE v1/posts/:id`: Delete a blog post


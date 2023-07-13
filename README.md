# Blog API

This is a simple RESTful API for a blogging platform built with Go.

## Features

- User registration and authentication with JWT token-based authentication
- User management: create, update, delete user profiles
- Post management: create, update, delete blog posts

## API Endpoints

### User Endpoints

- `POST /api/v1/register`: Register a new user.
- `POST /api/v1/login`: Authenticate and obtain a JWT token.
- `POST /api/v1/logout`: Logout and invalidate the JWT token.

### Blog Post Endpoints

- `POST /api/v1/posts`: Create a new blog post.
- `GET /api/v1/posts/:id`: Get a blog post by ID.
- `PUT /api/v1/posts/:id`: Update a blog post.
- `DELETE /api/v1/posts/:id`: Delete a blog post.


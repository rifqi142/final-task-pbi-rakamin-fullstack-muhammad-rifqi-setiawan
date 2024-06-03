# ⚡️ Final Project - Project Based Learning Rakamin ⚡️

This repository contains models and request/response structs for a photo-sharing application built in Go.


### 📃 Introduction


This project provides the necessary data structures and validation mechanisms for managing users and photos in a photo-sharing application. It includes structs for users, photos, and various request and response formats required for user authentication, photo creation, and updating.




## 🧑🏻‍💻 Author

- [@Muhammad Rifqi Setiawan](https://github.com/rifqi142)


## 📚 Tech Backend

- Golang
- Gin Gonic Framework
- JWT-GO
- Gorm
- Go validator


##  📲End Point

### User
- POST /users/register: Register a new user.
- POST /users/login: Log in an existing user.
- PUT /users/: Update user data. Requires authentication.
- DELETE /users/:UserId Delete user account. Requires authentication.


### Photos
- POST /photos/: Upload a new photo.
- GET /photos/:PhotoId: Get a photo by ID.
- PUT /photos/:PhotoId: Update photo data.
- DELETE /photos/:PhotoId: Delete a photo.

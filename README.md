# CRUD_GO

A simple **CRUD (Create, Read, Update, Delete) REST API** built using **Golang** and **Gorilla Mux**.  
This project demonstrates how to build a basic backend API in Go without using a database (in-memory storage).

---

## 🚀 Features

- Create a movie
- Get all movies
- Get a movie by ID
- Update a movie
- Delete a movie
- RESTful API structure
- JSON request & response handling

---

## 🛠️ Tech Stack

- **Go (Golang)**
- **net/http**
- **github.com/gorilla/mux**
- JSON for data exchange

---
## PROJECT STRUCTURE
CRUD_GO/
│── main.go
│── README.md

---

## ▶️ How to Run the Project

### 1️⃣ Clone the repository
```bash
git clone https://github.com/shahid-923/CRUD_GO.git

2️⃣ Go to project directory
cd CRUD_GO

3️⃣ Install dependencies
go mod init crud_go
go get github.com/gorilla/mux

4️⃣ Run the server
go run main.go
Server will start at:
http://localhost:8080

📌 API Endpoints
➤ Get all movies
GET /movies

➤ Get movie by ID
GET /movie/{id}

➤ Create a movie
POST /movie


Body (JSON):

{
  "isbn": "12345",
  "title": "New Movie",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}

➤ Update a movie
PUT /movie/{id}

➤ Delete a movie
DELETE /movie/{id}


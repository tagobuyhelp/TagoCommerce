# 🛒 TagoCommerce – Scalable eCommerce Backend

**TagoCommerce** is a modular, high-performance eCommerce backend built with **Golang (Gin)** and **MongoDB**. Designed to power modern frontend clients (React, Vue, Next.js, mobile apps), it provides core APIs for authentication, product management, cart, and order processing with a clean, multi-store-ready architecture.


## 🚀 Features

- ✅ **User Authentication** (JWT, bcrypt)
- 🛍️ **Product & Category Management**
- 🛒 **Cart & Order APIs**
- 🧑‍💼 **Role-Based Access** (Admin, Vendor, Customer)
- 🌐 **MongoDB** as NoSQL document store
- 📦 **Modular Structure** for easy scaling
- 🛠️ Built with **Gin** for blazing-fast HTTP routing



## 🗂️ Project Structure


tagocommerce-backend/
├── cmd/                  # Entrypoint (main.go)
├── internal/
│   ├── config/           # Environment loader
│   ├── database/         # MongoDB connection
│   ├── models/           # Data models (User, Product, etc.)
│   ├── repository/       # MongoDB queries
│   ├── service/          # Business logic
│   ├── handler/          # Gin route handlers
│   ├── middleware/       # JWT, auth, logging
│   └── routes/           # Route groups
├── pkg/                  # Utilities (jwt, password hashing, etc.)
├── .env                  # Environment variables
├── go.mod                # Go module definition
└── README.md             # You’re here!



## 🔧 Installation & Run

### 1. Clone the Repository

```bash
git clone https://github.com/tagobuyhelp/TagoCommerce.git
cd TagoCommerce


### 2. Setup `.env`

Create a `.env` file in the root:
.env
PORT=8080
MONGO_URI=mongodb://localhost:27017
JWT_SECRET=your_secret_key_here


### 3. Run the Server

bash
go mod tidy
go run cmd/tagocommerce/main.go


> Server runs at `http://localhost:8080`


## 📌 Example Endpoints

| Method | Endpoint             | Description             |
| ------ | -------------------- | ----------------------- |
| POST   | `/api/auth/register` | Register a new user     |
| POST   | `/api/auth/login`    | Login and get token     |
| GET    | `/api/products`      | List products           |
| POST   | `/api/products`      | Add new product (admin) |



## 🧠 Tech Stack

* **Language**: Golang 1.21+
* **Framework**: [Gin](https://github.com/gin-gonic/gin)
* **Database**: MongoDB
* **Auth**: JWT + Bcrypt
* **Env**: `github.com/joho/godotenv`
* **Validation**: `go-playground/validator` (recommended)



## 📈 Roadmap

* [x] Auth System (JWT)
* [ ] Products & Categories CRUD
* [ ] Cart & Orders
* [ ] Admin Panel APIs
* [ ] Multi-Vendor Store Logic
* [ ] Payment Gateway Integration
* [ ] Public API Documentation (Swagger)



## 🧑‍💻 Contributing

Pull requests are welcome. For major changes, open an issue first to discuss what you'd like to change.



## 📄 License

[MIT](LICENSE)



## 🔗 Connect

* **Author**: [Tarik Aziz](https://github.com/tagobuyhelp)
* **Company**: [Tagobuy IT Services](https://tagobuy.net)



![Go Reference](https://pkg.go.dev/badge/github.com/go-telegram-bot-api/telegram-bot-api/v5.svg)
![Test](https://github.com/go-telegram-bot-api/telegram-bot-api/actions/workflows/test.yml/badge.svg)
<div style="width:100%; display: flex; align-items: center;">
  <h1>Api Go REST - GIN + MongoDB
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" height="60" width="70" style="margin-bottom: -15px; z-index: -10; margin-left: 1.25rem"/>
  </h1> 
</div>


### 🛠  Description   

</br>

This project aims to create a REST API that connects directly to MongoDB, performing CRUD operations. I developed an API using the Go programming language and the GIN framework, and all endpoints are documented with Swagger.

The project design is modular, following recognized development standards. I carefully selected the logic and structure, creating a specific folder for settings. This folder includes log files, connection to MongoDB and database initialization.

I developed a custom logger to handle various responses and information from endpoints, from success responses to errors.

I opted for a personalized approach to data validation, avoiding external libraries or libraries that allowed me to validate JSON structures accurately.

To speed up development, I also created a Makefile with customized commands that facilitate routine tasks.


## Swagger

</br>

<p align="center">
  <kbd>
 <img width="100%" style="border-radius: 10px" height="600" src="https://github.com/JuanCampbsi/Preview_README/blob/86f37a264c34d108e5e1f52e9acc8c144fa81a12/assets/go-oportunities-swaggers.png" alt="Intro"> 
  </kbd>
  </br>
</p>

</br>


### ⌨ Installation
To use it, you need to clone the repository, install the dependencies and run the project.

```bash
# Open terminal/cmd and then Clone this repository
$ git clone https://github.com/JuanCampbsi/go-opportunities.git

# Access project folder in terminal/cmd
$ cd go_api_mongo_db

# Install the dependencies
$ go mod tidy

# Run the application in development mode
$ go run main.go                           

```

</br>	

### ⌨ Stack of technologies and libraries

-   [Golang](https://go.dev/doc/) - version 1.20
-   [GIN](https://github.com/gin-gonic/gin) - version 1.9.0
-   [GORM](https://gorm.io/gorm ) - version 1.25.1
-   [MongoDB](go.mongodb.org/mongo-driver) - version 1.12.1
-   [Swaggo](https://github.com/swaggo/swag) - version 1.16.1
 
</br>

👨‍💻 **Author** 💻

Developed by [_Juan Campos_](https://www.linkedin.com/in/juancampos-ferreira/)
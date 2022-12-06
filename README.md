# aircompany_backend
A simple application that allows you to register new users, get all avaliable flights from the database and tickets purchased by the user.


# How to install
## Golang

Download and install [golang](https://golang.org)

## PostgreSQL

Download and install https://www.postgresql.org/download/

Open pg4admin and execute the sql file in the folder sql. Create the database at your pc.

Create .env file in the root of the project 
```bash
db_name = <your_name>
db_pass = <your_password>
db_user = <your_user>
db_host = <your_host>
db_port = <your_port>
db_sslmode = disable
secret_key = <your_key> // a string with random letters, digits and symbols. It's used for jwt
```


# How to run
```bash
> go run main.go
```

# Tools used
Golang, PostgreSQL, Go-chi, go-jwt

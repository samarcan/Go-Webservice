# Go-Webservice

The main objective of this project is learn how to create webservices with Go.

The database engine used is PostgreSQL

In order to run the project will be neccesary a `settings.env` file in the root directory in which will be defined all the database configuration parameters. This file look like this:

```
db_type = postgres
db_host = Database_Host
db_port = Database_Port
db_name = Database_Name
db_user = Database_User
db_pass = Database_Password
```

**Dependencides**

* github.com/gorilla/mux
* github.com/jinzhu/gorm
* github.com/lib/pq
* github.com/joho/godotenv

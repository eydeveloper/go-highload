# Project Setup Instructions

To set up this project, follow the steps below:

1. Create a .env file in the root directory of your project and add the following line: 
   DB_PASSWORD=your_database_password

2. Run migrations to create the necessary database tables. Use the following command:
   migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up


3. Finally, to start the server, execute the following command:
   go run cmd/main.go

That's it! Your project should now be set up and running.
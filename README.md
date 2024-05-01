# Highload Social

## Description

This project serves as a practical learning tool directly related to my studies in software architecture. As a student
enrolled in a software architecture course, I've developed this project to deepen my understanding of architectural
principles and refine my practical skills. Synchronous replication is configured in this application to ensure data
integrity of slave replicas.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- You have installed Docker
- You have installed Go

## Setup

To set up this project, follow these steps:

1. Clone the repository:
   ```shell
   git clone https://github.com/eydeveloper/highload-social.git
   ```

2. Create a .env file in the root directory of the project and add the following configuration:
   ```dotenv
   DB_PASSWORD=<database_password>
   ```

3. Start the Docker services and prepare slave replicas.
   ```shell
   make up
   make postgres-slave-1-bash
   pg_basebackup -h postgres-master -D /var/lib/postgresql/data -U replicator -R --wal-method=stream
   exit
   make postgres-slave-2-bash
   pg_basebackup -h postgres-master -D /var/lib/postgresql/data -U replicator -R --wal-method=stream
   exit
   ```

4. Run the migrations to set up the database schema:
   ```shell
   migrate -path ./database/schema -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up
   ```

5. Create GIN indexes:
   ```shell
   make postgres-master-psql
   create index <index_name> on users using gin (to_tsvector('english', first_name), to_tsvector('english', last_name));
   exit
   ```

6. Start the server by running the following command:
   ```shell
   go run cmd/main.go
   ```

## Available Actions

### This project supports the following actions:

- Register: Allows users to register for a new account.
- Login: Enables users to authenticate and log in to their accounts.
- Get User by ID: Retrieves user information based on the provided user ID.
- Search users by name: Finds users by first name and last name.
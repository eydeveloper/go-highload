# Highload Social

## Description

This project serves as a practical learning tool directly related to my studies in software architecture. As a student
enrolled in a software architecture course, I've developed this project to deepen my understanding of architectural
principles and refine my practical skills. It incorporates RabbitMQ for message queuing and distribution, along with
WebSocket services for real-time communication, reflecting modern architectural patterns. Various replication types,
including physical sync, async, and logical replication, are configured within the application to ensure data integrity
across slave replicas, enhancing fault tolerance and reliability.

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

## Scaling

### WebSocket services

WebSocket services achieve linear scalability through load balancing, horizontal scaling, and optimized resource
utilization. Load balancers evenly distribute incoming connections across multiple servers, while horizontal scaling
involves adding more servers dynamically to handle increasing loads. Optimized resource utilization includes techniques
like asynchronous I/O and event-driven architectures to maximize server throughput and responsiveness. Monitoring and
auto-scaling mechanisms ensure optimal resource allocation and responsiveness during peak loads. Fault tolerance is
ensured through redundant components and failover mechanisms, allowing WebSocket services to seamlessly handle growing
numbers of concurrent connections while maintaining high performance and reliability.

### RabbitMQ

Scaling RabbitMQ involves both horizontal and vertical strategies. Horizontal scaling encompasses cluster formation,
load balancing, and sharding queues across multiple nodes, while vertical scaling entails increasing node resources and
optimizing configurations. Monitoring metrics and employing auto-scaling mechanisms are crucial for dynamically
adjusting RabbitMQ clusters. High availability is ensured through replication, mirroring, and geographic distribution of
RabbitMQ clusters, enhancing fault tolerance and disaster recovery capabilities. By implementing these strategies and
best practices, RabbitMQ can efficiently handle larger workloads while maintaining reliability and performance in
distributed systems.
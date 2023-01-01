# Go SM

## Installation

### Docker

The application requires a database. For local development purposes, you can use docker compose to create your own database.

#### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

#### Setup

1. Clone this repository:

```bash
git clone https://github.com/mikededo/go-sm.git
```

2. Create an .env file in the root of the repository and set the MYSQL_ROOT_PASSWORD variable to the password you want to use for the root user: 

```bash
MYSQL_ROOT_PASSWORD=my-secret-password
```

3. Start the MySQL container:

```bash
docker-compose up -d
```

#### Connecting to the database

To connect to the database using a database management system, use the following connection details:

- Host: localhost
- Port: 3306
- User: root
- Password: my-secret-password (or the value you specified in the MYSQL_ROOT_PASSWORD variable in the .env file)

#### Stopping the Container

To stop the MySQL container, run the following command:

```bash
docker-compose down
```
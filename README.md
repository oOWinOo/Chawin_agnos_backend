
# Chawin_agnos_backend

 Backend Strong Password Recommendation steps


# Deployment

Follow these steps to deploy the project locally:

### 1. Clone Repository

```bash
git clone https://github.com/oOWinOo/Chawin_agnos_backend.git
```

### 2. Open Docker
###### Make sure Docker is installed on your machine.
####    Warning Port Usage
##### This project uses the following localhost ports:

- Port 8000: API endpoint
- Port 8080: Web application
- Port 5432: PostgreSQL database
- Port 80: Web server
- Port 5050: pgAdmin for database administration
Ensure that no other processes are using these ports on your machine to avoid conflicts. If any of these ports are already in use, you may need to stop the conflicting processes or modify the Docker configuration.

### 3. Docker Compose Up

```bash
docker compose up -d
```

### 4. Docker Compose Up Go service

```bash
docker compose up go_service --build -d
```
## View Logs in PostgreSQL

#### Access localhost:5050 for pgAdmin.

- Username: admin@email.com
- Password: admin

#### Set up the connection with the following details:

- Database Name: Anything you want
- Hostname: postgres
- Maintenance: postgres
- Username: postgres
- Password: postgres

#### Now, you can query the logs with:

```postgres
select * from log_in_outs
```
# Testing

You can choose any option below for testing, but make sure you have already run `docker-compose up -d`

### 1. Unit test
#### Run Unit Tests
```bash
go test
```
This command will execute unit tests defined in main_test.go.

### 2. Test API Endpoints
##### 1. Use a tool such as Insomnia or Postman.
##### 2. Send a POST request to localhost:8000/api/strong_password_steps with the following body:
```json
{
  "init_password": "Insert password here"
}
```
The Response will be 
```json
{
"num_of_steps": "number_of_steps (int)"
}
```

### 3. Open Index Page
#### Visit localhost:8080 to open index.html. 


# API Reference

#### Submit Password

```bash
  POST /api/strong_password_steps
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `init_password` | `string` | **Required**. Your Password |


Takes init_password and returns "num_of_steps" the minimum number of actions to make the password strong.


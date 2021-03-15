
# Installation and usage

## Environment Properties


You can either use an .env file or export env variables to pass below values.

**.env** file:

```dotenv
#Port for manager service
PORT=8080
#database driver
DRIVER_NAME=postgres
#database connection url
DATASOURCE_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
#database migration properties. version should be logical with sql files prefixes; 1_,2_ etc
MIGRATE_VERSION=2
MIGRATE_SCRIPT_URL=file://scripts/postgresql
MIGRATE_DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
```

## Application Setup

### Database Migration

We can use either an in memory database or a relational database to store our data. Here, primarily Postgresql is used and for fast testing, sqlite is preferred.
Once you make ready your database, then migration process should be triggered.
- If you want to complete it on your local; run `/cmd/migrate/migrate.go`. Make sure you have correct values for env variables; **MIGRATE_VERSION, MIGRATE_SCRIPT_URL, MIGRATE_DATABASE_URL**.
  Once this is successfull, you need to see message `migration has been completed`, then check the database and the tables.


- If you want to run in your docker environment, docker-compose handles it directly. Please check `docker-compose.yml` for more information.


- You can also check the github action flow to understand how it is handled during deployment.

### Docker in CI/CD pipeline

After the CI/CD pipeline builds manager in the **Build step**, the binary is picked up and put into the Docker image in the next step.

Once your image is ready, the container can be created with;

`docker run -p 8080:8080 <repo-domain-name>/task-manager ./manager`

### Dockerfile dev

We can use **Dockerfile.dev** file, if we want to build the image in our local environment. Once the process is completed successfully, we can use the same commands above to run the containers.

### Docker Compose

Docker compose provides us a compact environment, which includes all dependencies with Postgresql. Once it runs successfully, we can see **task-manager** service up

**Example**

###### **Request**

>curl -X POST -v http://localhost:8080/api/task
> 
> name=new_task

###### **Response**

> < HTTP/1.1 201 Created >

###### **Request**

>curl -X GET -v http://localhost:8080/api/task

###### **Response**

> < HTTP/1.1 200 OK >
>
> [
{
"task_id": "4bfe34e1-eaa4-4e10-97e7-0b21fda9d8a2",
"task_name": "task_65564",
"task_done": false
},
{
"task_id": "b934552d-ea0f-4b64-b43e-246b91e37285",
"task_name": "task_65564e",
"task_done": false
}
]

###### **Request**

>curl -X GET -v http://localhost:8080/api/task/<task_id>

###### **Response**

> < HTTP/1.1 200 Created >
>
> {
"task_id": "4bfe34e1-eaa4-4e10-97e7-0b21fda9d8a2",
"task_name": "task_65564",
"task_done": false
}

###### **Request**

>curl -X DELETE -v http://localhost:8080/api/task/<task_id>

###### **Response**

> < HTTP/1.1 204 No Content >

###### **Request**

>curl -X PUT -v http://localhost:8080/api/task/<task_id>

###### **Response**

> < HTTP/1.1 204 No Content >


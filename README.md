# Ambush Journey Program - Medical Clinic

This is a repository used by the AJP Backend team to practice building a web server in Go, developing a software for a medical clinic.

The software is used for scheduling, planning and managing medical appointments, as well as medical exams and services.

This project uses Docker, Goose for database migrations, and SQLBoiler for generating models in Golang.

## Workflows and Rules

Project tasks and user stories are listed in [Trello](https://trello.com/b/UtCP3tht/ajp-medical-clinic)

Branch guidelines should follow the syntax `<type>/<subject>` where \<type\> should be one of `feat`, `bugfix`, `test` or `chore`, and \<subject\> should be a very short description on the changes made.

Commit messages can follow the same guidelines from branches, but it's not necessary as long as it describes the changes made clearly.

Make sure to have `go fmt` running on your editor's save process. For VS Code you can follow the instructions [here](https://stackoverflow.com/a/53513867).

Once you have finished your changes on your feature branch, please open a pull request to the main branch and add the mentors as reviewers.

## Building the Server

To build the server, you need to have Docker and Docker Compose installed on your machine. Once installed, you can build and start the server by running the following command in the terminal:

```bash
$ docker-compose up
```

This command reads the docker-compose.yml file and starts all the services defined in it.

## Database Migrations with Goose

Goose is a database migration tool. You manage your database schema by creating incremental SQL or Go scripts.

To create a new migration script, use the goose create command followed by the name of your migration:
```bash
goose create AddSomeTable sql
```
This will create a new SQL file in the migrations directory. You can then edit this file to define your migration.

To apply all available migrations, use the goose up command:

```bash
$ goose up
```

To roll back the most recently applied migration, use the goose down command:

```bash
$ goose down
```
## Generating Models with SQLBoiler

SQLBoiler is a tool to generate a Go ORM tailored to your database schema.

First, you need to install SQLBoiler and its driver. For PostgreSQL, the driver is psql:


```golang
go get -u -t github.com/volatiletech/sqlboiler
go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
```

Then, you can generate the models by running the following command:

```bash
$ sqlboiler psql
```

This will generate a models directory containing Go files for each table in your database.

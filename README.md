# Requiriments
1. Golang
2. Docker
<br>

##### Observations: This app was tested on windows operational system, maybe will not work on linux because the way that I implemented the migrations on application start-up. If its the case, it will be fixed to be compatible with Linux and Windows, or Linux only.

## How to create a new migration sql file?

The easiest way is to download the golang migrate CLI [clicking here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate),
and after that, run the following command on terminal at project root: *migrate create -ext sql -dir /migrations -seq your_migration_name*, edit the file to do your changes on database

## How to run docker database image?

Run the following command on terminal at project root: *docker-compose up*

## How to run the app?

After start the docker image, run the following command on terminal at project root: go run ./cinema-app

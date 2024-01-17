
# FullCycle Clean Architecture Code Challenge

Ordering System with Golang

## Pre-requisites to run the app locally

- Migrate tool installed. Have a look at: [migrate tool](https://github.com/golang-migrate/migrate)
- Grpc client installed. Suggestion: Evans: [evans](https://github.com/ktr0731/evans)
- Docker and docker compose installed

## Local environment setup

1. Set up all the containers defined in docker compose file. Run: </br>```docker compose up -d```</br></br>
2. Execute database migrations. Database migrations are included in the `./sql/migrations` folder. Using migrate tool run: </br> </br>```migrate -path=sql/migrations -database "mysql://root:root@(localhost:3306)/orders" -verbose up```</br></br>
Note that the command above is ready to execute database migrations with the MySql database configuration defined.</br></br>
3. Create a RabbitMQ Queue with the name of your preference lets say "orders". After create the queue, add a binding from exchange: `amq.direct`.</br></br>
4. Change directory to `cmd/orderingsystem` </br>
Run command ```go run main.go wire_gen.go```

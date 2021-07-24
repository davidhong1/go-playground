initdb:
	mysql -uroot -p < ./sql/init.sql

dep:
	go mod download

run:
	go run main.go

build:
	go build -o bin/week02 main.go
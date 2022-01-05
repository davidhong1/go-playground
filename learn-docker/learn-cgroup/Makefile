all:
	GOOS=linux go build -o bin/$$(go mod why | tail -n 1) main.go
	docker run --privileged=true -it -v $$PWD/bin:/app ubuntu /bin/bash
	# apt update && apt install stress
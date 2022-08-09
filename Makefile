# docker build
d-build:
	docker build -f build/package/Dockerfile -t aut-node-app

# docker run 
d-run:
	docker run -d --name aut-node-app-container -p 127.0.0.1:80:8080 aut-node-app

# go build
build:
	go build -o main

# starting application
start:
	chmod +x ./internal/cmd/darolfonoon/main
	./main

# go run
run:
	go run ./internal/cmd/darolfonoon/main.go
	

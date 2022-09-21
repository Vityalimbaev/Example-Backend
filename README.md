
run database:

        docker-compose up -d example-db

run app: 

        cd /app
        go mod vendor
        go run cmd/main/main.go

run all services in docker: 

        docker-compose up -d

swagger: 127.0.0.1:8000/swagger/index.html
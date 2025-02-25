DB_DSN = "postgres://postgres:yourpassword@localhost:5432/main?sslmode=disable"

run:
		go run ./cmd/app/main.go

migrate-new:
		migrate create -ext sql -dir ./migrations -seq ${NAME}

migrate-up:
		migrate -path ./migrations -database $(DB_DSN) up

migrate-down:
		migrate -path ./migrations -database $(DB_DSN) down

restart:
		 stop run



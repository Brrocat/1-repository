DB_DSN = "postgres://postgres:yourpassword@localhost:5432/main?sslmode=disable"

run:
		go run ./cmd/app/main.go

migrate-new:
		migrate create -ext sql -dir ./migrations $(NAME)

migrate-up:
		migrate -pach ./migrations -database $(DB_DSN) up

migrate-down:
		migrate -pach ./migrations -database $(DB_DSN) down


# Usage: make create-migrate name=create_users_table
create-migrate:
	goose -dir ./internal/db/migrations create ${name} sql

migrate:
	source .env && goose -dir ./internal/db/migrations postgres "$$DATABASE_URL" up

seed-customers:
	go run main.go seed CustomerSeeder
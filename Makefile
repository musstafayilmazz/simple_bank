start-docker:
	docker-compose up -d
stop-docker:
	docker-compose down
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
run-code:
	@go run main.go
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: start-docker createdb dropdb stop-docker migratedown migrateup

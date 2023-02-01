postgres:
	docker run --name postgres-12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres-12 createdb --username=root --owner=root nobank

dropdb:
	docker exec -it postgres-12 dropdb nobank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nobank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nobank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
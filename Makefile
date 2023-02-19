postgres:
	docker run --name postgres-12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres-12 createdb --username=root --owner=root nobank

dropdb:
	docker exec -it postgres-12 dropdb nobank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nobank?sslmode=disable" -verbose up

prevmigrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nobank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nobank?sslmode=disable" -verbose down

prevmigratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/nobank?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown prevmigrateup prevmigratedown sqlc test server
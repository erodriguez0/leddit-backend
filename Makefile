postgres:
	docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16.2

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root leddit

dropdb:
	docker exec -it postgres16 dropdb leddit

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/leddit?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/leddit?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/leddit?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/leddit?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 server
postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:12-alpine

postgresstart:
	sudo docker start postgres12

postgresstop:
	sudo docker stop postgres12

createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgres://root:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://root:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgres://root:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./db/sqlc/

main:
	go run main.go

buildbank:
	 sudo docker buildx build -t simplebank:latest .

runbank:
	sudo docker run --name simplebank -p 2207:2207 -d simplebank

startbank:
	sudo docker start simplebank

stopbank:
	sudo docker stop simplebank

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/simple_bank/db/sqlc Store

.PHONY: postgres postgresstart postgresstop createdb dropdb migrateup migratedown sqlc test mock

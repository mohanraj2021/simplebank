postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:12-alpine

postgresstart:
	sudo docker start postgres12

postgresstop:
	sudo docker stop postgres12

createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root simplebank

dropdb:
	sudo docker exec -it postgres12 dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgres://root:admin@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:admin@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres postgresstart postgresstop createdb dropdb migrateup migratedown sqlc 

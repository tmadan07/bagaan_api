network:
	docker network create bank-network

postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlci:
	docker pull kjconroy/sqlc

sqlcg:
	docker run --rm -v %cd%:/src -w /src kjconroy/sqlc generate
	
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate 

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go api/db/sqlc Store

dbreset:
	liquibase update --changelog-file changelog.reset.yaml --URL jdbc:postgresql://localhost/ --log-level ERROR

dbupdate:
	liquibase update --changelog-file changelog.master.yaml --log-level ERROR

dbtestseed:
	liquibase update --changelog-file db/migration/changelog.part.seedtest.yaml --log-level ERROR

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlci sqlcg test server mock dbreset dbupdate dbtestseed

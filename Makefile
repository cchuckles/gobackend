# project commands

buildpostgres:
	@go build -tags=postgres -o bin/backend.exe cmd/backend/main.go

runpostgres: buildpostgres
	@./bin/backend

test:
	@go test -v ./...


# postgres commands

postgresinit:
	docker run --name pglatest -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:latest

postgres: 
	docker exec -it pglatest psql

postgrescreatedb:
	docker exec -it pglatest createdb --username=root --owner=root gobackend

postgresdropdb:
	docker exec -it pglatest dropdb gobackend
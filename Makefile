createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root ngekost_db

dropdb:
	sudo docker exec -it postgres dropdb ngekost_db

migrateup:
	migrate -database "postgres://root:root@localhost:5431/ngekost_db?sslmode=disable" -path database/migrations up

migratedown:
	migrate -database "postgres://root:root@localhost:5431/ngekost_db?sslmode=disable" -path database/migrations down

run:
	go run main.go

build:
	go build -o ./dist .

image:
	sudo docker build -t ngekost-api:latest .

.PHONY: createdb dropdb migrateup migratedown run build image
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
	go build -o ./dist/main .

image:
	sudo docker build -t arvinpaundra/ngekost-api:latest .

container: image
	sudo docker run --name ngekost-app -v ./.env:/.env -p 9012:9012 --rm arvinpaundra/ngekost-api:latest

test:
	go test -v -cover ./internal/app/...

coverage:
	go test -v -coverprofile=cover.out ./... && go tool cover -html=cover.out

.PHONY: createdb dropdb migrateup migratedown run build image test testcoverage
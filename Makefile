DB_URL=postgresql://uksf4elpbyzvlgivn9lr:RGRWrsbEmoRzqYKUGdB6K20qMCuKt2@brmmhqx43v3g38ns2ugl-postgresql.services.clever-cloud.com:50013/brmmhqx43v3g38ns2ugl

postgres:
	docker run --name todo -e DB_PORT=5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e DB_HOST=docker.for.mac.host.internal -d postgres:14-alpine

createdb:
	docker exec -it todo createdb --username=root --owner=root todo

dropdb:
	docker exec -it todo dropdb todo

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

test: 
	go test -v -cover ./...

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/demola234/todogrpc/db/sqlc Store

air:
	air

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb \
	--grpc-gateway_opt paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge,merge_file_name=payzone \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

db_docs:
	dbdocs build doc/doc.dbml

db_schema:
	dbml2sql --postgress -o doc/schema.sql doc/doc.dbml


.PHONY: createdb postgres dropdb migrateup migratedown sqlc server air proto evans db_docs db_schema mock test
mysql:
	docker run --name mysql -e MYSQL_ROOT_PASSWORD=passwd -p 3307:3306 -d mysql:latest

migrateup:
	 migrate -path db/migration -database "mysql://root:passwd@tcp(localhost:3307)/bank?parseTime=true&net_write_timeout=6000" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:passwd@tcp(localhost:3307)/bank" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...  # 执行所有包下面的测试

server:
	go run main.go

.PHONY: mysql migrateup migratedown sqlc test server
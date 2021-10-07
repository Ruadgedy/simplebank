mysql:
	docker run --name mysql -e MYSQL_ROOT_PASSWORD=passwd -p 3307:3306 -d mysql:latest

migrateup:
	 migrate -path db/migration -database "mysql://root:passwd@tcp(localhost:3307)/bank?parseTime=true" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:passwd@tcp(localhost:3307)/bank" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...  # 执行所有包下面的测试

.PHONY: mysql migrateup migratedown sqlc test
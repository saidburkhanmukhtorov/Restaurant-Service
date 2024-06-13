CURRENT_DIR=$(shell pwd)
DBURL='postgres://postgres:20005@localhost:5432/voting2?sslmode=disable'

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path migrations -database $(DBURL) -verbose up

mig-down:
	migrate -path migrations -database $(DBURL) -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table

swag-init:s
	swag init -g api/api.go -o api/doc

proto:
	protoc --twirp_out=paths=source_relative:. --go_out=paths=source_relative:. rpc/twirpAPI/twirp.proto

server:
	go run cmd/server/main.go


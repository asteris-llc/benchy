benchy: rpc/pb/root.pb.go rpc/pb/root.pb.gw.go
	go build -ldflags="-s -w" -o benchy main.go

rpc/pb/root.pb.go: rpc/pb/root.proto
	protoc -I rpc/pb \
	 -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:rpc/pb \
	 rpc/pb/root.proto

rpc/pb/root.pb.gw.go: rpc/pb/root.proto
	protoc -I rpc/pb \
	 -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --grpc-gateway_out=logtostderr=true:rpc/pb \
	 rpc/pb/root.proto

.PHONY: benchy

protoc --proto_path=protos/user_service --proto_path=/usr/local/include --gofast_out=plugins=grpc:. user.proto
protoc --proto_path=protos/post_service --proto_path=/usr/local/include --gofast_out=plugins=grpc:. post.proto
protoc --proto_path=protos/notification_service --proto_path=protos/third_party --gofast_out=plugins=grpc:. notification.proto
protoc --proto_path=protos/media_service --proto_path=protos/third_party --gofast_out=plugins=grpc:. services.proto
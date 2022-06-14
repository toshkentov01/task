protoc --proto_path=protos/user_service --gofast_out=plugins=grpc:. user.proto
protoc --proto_path=protos/user_service --proto_path=protos/user_service --gofast_out=plugins=grpc:. user.proto
protoc --proto_path=protos/media_service --gofast_out=plugins=grpc:. media.proto
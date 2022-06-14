protoc --proto_path=task-protos/crud_service --gofast_out=plugins=grpc:. crud.proto
protoc --proto_path=task-protos/data_service --gofast_out=plugins=grpc:. data.proto


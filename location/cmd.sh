protoc -I pb/location \
pb/location/location.proto \
--go_out=plugins=grpc:.
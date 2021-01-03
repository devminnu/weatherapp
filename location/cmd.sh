protoc -I pb/location \
pb/location/location.proto \
--go_out=plugins=grpc:.


//from weather app dir

docker image build -t location -f location/Dockerfile .
docker container run -d -p 33001:33001 --name location location  
docker container ls -a 
docker container rm -f location
docker container logs location
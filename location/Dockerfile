FROM golang

WORKDIR /app

# COPY location/go.mod .
# COPY location/go.sum .

COPY location .

COPY pb .

EXPOSE 33001

WORKDIR /Users/minhaj/go/src/grpc-apps/weatherapp/location

RUN [ "go", "build", "-o","app", "main.go"]

RUN [ "chmod", "777", "app" ]

CMD [ "./app" ]

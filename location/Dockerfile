FROM golang

ENV GO111MODULE=on

WORKDIR  /Users/minhaj/go/src/apps/weatherapp

COPY . .

WORKDIR  /Users/minhaj/go/src/apps/weatherapp/location

EXPOSE 30001

EXPOSE 50051

RUN [ "go", "build", "-o","locationapp", "main.go"]

RUN [ "chmod", "777", "locationapp" ]

CMD [ "./locationapp" ]

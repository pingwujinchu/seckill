FROM golang

ENV GO111MODULE=on \
    GIN_MODE=release \
    PORT=8080
    
WORKDIR /app

COPY . .

RUN export CGO_ENABLED=0 && export GOSUMDB=off && go build -mod=vendor -a -v -o main

EXPOSE 8080

ENTRYPOINT ["./main"]
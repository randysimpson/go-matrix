FROM golang:1.15 as builder
WORKDIR /go/src
RUN mkdir go-matrix
WORKDIR /go/src/go-matrix
ADD . .
#WORKDIR /go/src/go-matrix/matrix
#RUN go mod init github.com/randysimpson/go-matrix/matrix
#RUN ls
#RUN cat go.mod
#RUN go test

#clean up mod file
RUN go mod tidy

#run tests
RUN go test ./...

#output test coverage as html
RUN go test ./... -cover -coverprofile=c.out
RUN go tool cover -html=c.out -o coverage.html

RUN ls
RUN cat coverage.html

#list will create go.sum file with imports
RUN go list -m all

#RUN ls
#RUN cat go.sum

#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
#FROM scratch
#COPY --from=builder /go/src/go-matrix/main /app/
#WORKDIR /app
#CMD ["./main"]
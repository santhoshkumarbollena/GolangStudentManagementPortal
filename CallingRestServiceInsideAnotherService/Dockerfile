FROM golang:1.8
RUN go get github.com/gorilla/mux
ADD RestSerFromAnotherService.go RestSerFromAnotherService.go
Expose 10000
ENTRYPOINT ["go","run","RestSerFromAnotherService.go"]

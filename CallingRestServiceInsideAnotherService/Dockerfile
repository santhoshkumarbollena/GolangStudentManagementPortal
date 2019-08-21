FROM golang:1.11.13
RUN go get github.com/gorilla/mux && \
    go get gopkg.in/yaml.v2
ADD RestSerFromAnotherService.go RestSerFromAnotherService.go
Add Properties.yaml Properties.yaml

Expose 10000
ENTRYPOINT ["go","run","RestSerFromAnotherService.go"]

FROM golang:1.8
RUN go get github.com/gorilla/mux && \
    go get gopkg.in/oauth2.v3/errors && \
   go get gopkg.in/oauth2.v3/manage && \
   go get gopkg.in/oauth2.v3/server && \
   go get gopkg.in/oauth2.v3/store && \
   go get gopkg.in/oauth2.v3/models 
ADD SecuredServiceInAnotherService.go SecuredServiceInAnotherService.go
Expose 10000
ENTRYPOINT ["go","run","SecuredServiceInAnotherService.go"]

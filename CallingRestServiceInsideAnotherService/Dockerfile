FROM golang:1.11.13
RUN go get github.com/gorilla/mux && \
    go get gopkg.in/yaml.v2
ADD RestSerFromAnotherService.go RestSerFromAnotherService.go
Add Properties.yaml Properties.yaml
ENV SSH_PASSWD "root:Docker!"
RUN apt-get update \
        && apt-get install -y --no-install-recommends dialog \
        && apt-get update \
	&& apt-get install -y --no-install-recommends openssh-server \
	&& echo "$SSH_PASSWD" | chpasswd 

COPY sshd_config /etc/ssh/
COPY init.sh /usr/local/bin/
	
RUN chmod u+x /usr/local/bin/init.sh
Expose 10000 2222
ENTRYPOINT ["go","run","RestSerFromAnotherService.go"]

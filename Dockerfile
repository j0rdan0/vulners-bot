FROM golang
COPY main.go go.mod model.go /root
WORKDIR /root
RUN go mod tidy
ENTRYPOINT ["go","run","."]

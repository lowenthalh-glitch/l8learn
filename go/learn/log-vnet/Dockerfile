FROM saichler/builder:latest AS build

ENV CGO_ENABLED 1
COPY main.go /home/src/github.com/saichler/build/main.go
COPY .netrc /root/.netrc
RUN go mod init
RUN GOPROXY=direct GOPRIVATE=github.com go mod tidy
RUN rm -f /root/.netrc
RUN go mod vendor

WORKDIR /
RUN /patch.sh

WORKDIR /home/src/github.com/saichler/build
RUN go build -o logvnet

FROM saichler/l8learn-security:latest AS final
COPY --from=build /home/src/github.com/saichler/build/logvnet /home/run/logvnet

ENTRYPOINT ["/home/run/logvnet"]

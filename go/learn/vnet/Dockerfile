FROM saichler/builder:latest AS build
COPY main.go /home/src/github.com/saichler/build/
RUN go mod init
RUN GOPROXY=direct GOPRIVATE=github.com go mod tidy
RUN go build -o learn-vnet

FROM saichler/learn-security:latest AS final
COPY --from=build /home/src/github.com/saichler/build/learn-vnet /home/run/learn-vnet
ENTRYPOINT ["/home/run/learn-vnet"]

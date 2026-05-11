FROM saichler/builder:latest AS build
COPY main/main.go /home/src/github.com/saichler/build/
RUN go mod init
RUN GOPROXY=direct GOPRIVATE=github.com go mod tidy
RUN go build -o learn-web

FROM saichler/learn-security:latest AS final
COPY --from=build /home/src/github.com/saichler/build/learn-web /home/run/learn-web
COPY web /home/run/web
ENTRYPOINT ["/home/run/learn-web"]

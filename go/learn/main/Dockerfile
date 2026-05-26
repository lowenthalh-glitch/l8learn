FROM saichler/builder:latest AS build
COPY main.go /home/src/github.com/saichler/build/
COPY .netrc /root/.netrc
RUN go mod init
RUN GOPROXY=direct GOPRIVATE=github.com go mod tidy
RUN rm -f /root/.netrc
RUN go build -o learn

FROM saichler/l8learn-postgres:latest AS final
COPY --from=build /home/src/github.com/saichler/build/learn /home/run/learn
ENTRYPOINT ["/home/run/learn"]

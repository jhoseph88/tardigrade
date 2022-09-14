FROM us.gcr.io/rsg-base-prod/golang:1.17.13 as builder
 
WORKDIR /src

COPY . /src
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 go build -o tardigrade
 
FROM gcr.io/distroless/static
 
COPY --from=builder /src/tardigrade /
CMD ["/tardigrade"]

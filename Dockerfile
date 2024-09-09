FROM golang:1.22.1 as build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /shifter-auth

#FROM build-stage AS run-test-stage
#RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /bpjs /bpjs

USER nonroot:nonroot

ENTRYPOINT ["/bpjs"]
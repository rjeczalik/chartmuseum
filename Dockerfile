# This will be our builder image

FROM golang:alpine

ARG version=0.14.1

ARG revision=main

COPY . /go/src/github.com/helm/chartmuseum

WORKDIR /go/src/github.com/helm/chartmuseum

RUN CGO_ENABLED=0 GO111MODULE=on go build \
   -tags netgo,osusergo \
   -v --ldflags='-w -X main.Version=${version} -X main.Revision=${revision} -extldflags "-static"' \
   -o /chartmuseum \
   cmd/chartmuseum/main.go


# This will be the final image

FROM alpine:latest

RUN apk add --no-cache cifs-utils ca-certificates libc6-compat

COPY --from=0 /chartmuseum /chartmuseum
COPY chartmuseum.yaml /chartmuseum.yaml

EXPOSE 8080

USER 1000:1000

ENTRYPOINT ["/chartmuseum"]
CMD ["--config", "chartmuseum.yaml"]

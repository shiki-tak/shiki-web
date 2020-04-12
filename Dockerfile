FROM golang:1.13-alpine AS builder
ENV PACKAGES make bash git
RUN apk add --no-cache $PACKAGES
# Set working directory for the build
WORKDIR /work_dir

# Add source files
COPY . .

ENV PATH /go/bin:$PATH

RUN make build

FROM alpine

ENV PACKAGES curl make bash git
RUN apk add --no-cache $PACKAGES

COPY --from=builder /work_dir/build/shiki-web /shiki-web/shiki-web

WORKDIR /shiki-web
CMD ./shiki-web server start
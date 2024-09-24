### BUILD ###

# NOTE:
# Used to build Controller image
# In this file, we delete the *.ts intentionally
# Any other changes to Dockerfile should be reflected in Publish

FROM golang:1.22 AS go-builder

WORKDIR /app

COPY ./interpreter/go.mod ./interpreter/go.sum ./
RUN go mod download

COPY ./interpreter/ ./

RUN go build -o interpreter . && chmod +x interpreter


# crane digest cgr.dev/chainguard/node-lts:latest-dev
FROM cgr.dev/chainguard/node:latest-dev@sha256:7b649277768dd9c8c63097877fe9e32004b577c590ecf11422eaeb630cbcc395 AS build

WORKDIR /app

# Copy the node config files
COPY --chown=node:node app/package*.json ./

COPY ./app/controller.js ./controller.js 
COPY ./app/package.json ./package.json

##### DELIVER #####

# crane digest cgr.dev/chainguard/node-lts:latest
FROM cgr.dev/chainguard/node:latest@sha256:7d2170d090ad459647aff186ae85f79520832a35310d71ab2882719623921619

WORKDIR /app

COPY --from=build --chown=node:node /app/controller.js ./controller.js
COPY --from=build --chown=node:node /app/package.json ./package.json
COPY --from=go-builder /app/interpreter ./interpreter

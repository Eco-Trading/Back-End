FROM golang:latest AS build
WORKDIR /app
COPY . .
COPY ./main.go ./cmd/account_manager

ARG GRPC_PORT
ARG DATABASE_MONGODB_URL

ENV GRPC_PORT=${GRPC_PORT}
ENV DATABASE_MONGODB_URL=${DATABASE_MONGODB_URL}

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -x -a -ldflags '-extldflags "-static" -s -w' -o /account_manager

FROM gcr.io/distroless/static-debian11
WORKDIR /app
COPY --from=build /account_manager /account_manager
EXPOSE ${GRPC_PORT}
USER nonroot:nonroot
ENTRYPOINT ["/account_manager"]

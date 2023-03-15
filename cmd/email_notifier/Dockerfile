FROM golang:latest AS build
WORKDIR /app
COPY . .
COPY ./main.go ./cmd/email_manager

ARG GRPC_PORT
ARG DATABASE_MONGODB_URL

ENV GRPC_PORT=${GRPC_PORT}
ENV DATABASE_MONGODB_URL=${DATABASE_MONGODB_URL}

RUN go mod download
RUN CGO_ENABLED=0 go build -v -x -a -ldflags '-extldflags "-static" -s -w' -o /email_manager

FROM gcr.io/distroless/static-debian11
WORKDIR /app
COPY --from=build /email_manager /email_manager
EXPOSE ${GRPC_PORT}
USER nonroot:nonroot
ENTRYPOINT ["/email_manager"]

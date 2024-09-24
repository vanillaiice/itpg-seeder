FROM golang:1.23.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags="-s -w" -o /seeder .

FROM scratch

LABEL org.opencontainers.image.source=https://github.com/vanillaiice/itpg-seeder

WORKDIR /

COPY --from=build /seeder /seeder

ENTRYPOINT [ "/seeder" ]

CMD [ "--help" ]

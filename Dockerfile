FROM golang:1.23

WORKDIR /abramed

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

ENV POSTGRES_USER=victor
ENV POSTGRES_PASSWORD=victor
ENV POSTGRES_DB=pydata2
ENV POSTGRES_HOST=db

RUN migrate -path=./cmd/api/migrations -database=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}/${POSTGRES_DB} up



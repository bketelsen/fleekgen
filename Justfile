set dotenv-load

default: build


lint:
  golangci-lint run

snapshot:
  goreleaser release --clean --snapshot

build:
  @source ./.env
  @go build -a -tags netgo -ldflags '-w -extldflags "-static"' github.com/bketelsen/fleekgen

run: build
  ./fleekgen
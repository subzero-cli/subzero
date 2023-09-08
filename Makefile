build:
	go build -ldflags "-X main.version=$(git describe --tags) -X main.buildTime=$(date -u '+%Y-%m-%dT%H:%M:%SZ') -X main.commit=$(git rev-parse HEAD)"

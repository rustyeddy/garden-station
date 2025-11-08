target	= garden-station

all: test garden-station

garden-station:
	go build -v

run:
	go run -v .

run-mock:
	go run -v . -mock

strip:
	go build -ldflags="-s -w" -v -o "${target}_strip" .

pi:
	env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -v -o "${target}_pi" .

zero:
	env GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-s -w" -v -o "${target}-zero"

serve:
	go run -v . -mock

test:
	go test ./...

test-v:
	go test -v ./...

$(SUBDIRS):
	$(MAKE) -C $@

.PHONY: all test build $(SUBDIRS)


proj := devlog2stderr

default: runtest

build: $(proj)

runtest: $(proj) entry-test.sh
	docker run --rm \
		-v $(PWD)/$(proj):/usr/local/sbin/$(proj):ro \
		-v $(PWD)/entry-test.sh:/entry-test.sh:ro \
		\
		debian:stable /entry-test.sh

clean:
	rm -f *~ */*~ .*~ $(proj)
	go clean

$(proj): Makefile *.go
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $@
	go vet

.PHONY: build runtest clean

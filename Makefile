go-clean:
	rm -fr ./pkg

go-build:
	go build  -o ./pkg/page_checker ./cli/page_checker

build:
	$(MAKE) go-clean
	$(MAKE) go-build

check: 
	$(MAKE) build
	./pkg/page_checker img --json test/data/iana.json


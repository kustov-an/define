GOCMD=go
OUT=define

build:
	$(GOCMD) build -o $(OUT) cmd/*.go
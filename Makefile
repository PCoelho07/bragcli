BIN=brag

BD=dist
CDR=.

INSTALL_DIR=/usr/local/bin

test:
	go test ./...

fmt:
	gofmt -w .

build:
	go build -o $(BD)/$(BIN) $(CDR)/main.go

install: build
	sudo mv $(BD)/$(BIN) $(INSTALL_DIR)/$(BIN)

clean:
	rm -rf $(BD)

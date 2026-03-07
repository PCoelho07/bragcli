BIN=brag

BD=dist
CDR=.

INSTALL_DIR=/usr/local/bin

build:
	go build -o $(BD)/$(BIN) $(CDR)/main.go

install: build
	sudo mv $(BD)/$(BIN) $(INSTALL_DIR)/$(BIN)

clean:
	rm -rf $(BD)

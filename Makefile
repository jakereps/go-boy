GO = go
GOM = GO111MODULE=on $(GO)

BINDIR = ./bin
BINARY = go-boy

build: $(BINDIR)/$(BINARY)

$(BINDIR)/$(BINARY): $(BINDIR)
	$(GOM) build -o $(BINDIR)/$(BINARY) ./cmd/go-boy/...

$(BINDIR):
	mkdir -p $(BINDIR)

clean:
	$(GO) clean
	rm $(BINDIR)/$(BINARY)
	rmdir $(BINDIR)

test:
	$(GOM) test -cover ./...
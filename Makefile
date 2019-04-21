GO = go
GOM = GO111MODULE=on $(GO)

BINDIR = ./bin
BINARY = go-boy

build: $(BINDIR)/$(BINARY)

$(BINDIR)/$(BINARY): $(BINDIR)
	$(GO) build -o $(BINDIR)/$(BINARY) ./cmd/go-boy/...

$(BINDIR):
	mkdir -p $(BINDIR)

clean:
	$(GOM) clean
	rm $(BINDIR)/$(BINARY)
	rmdir $(BINDIR)

test:
	$(GOM) test -cover ./...
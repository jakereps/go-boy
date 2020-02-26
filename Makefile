GO := go
BINDIR := ./bin
BINARY := go-boy
GOFLAGS=-mod=vendor

bin: $(BINDIR)/$(BINARY)

$(BINDIR)/$(BINARY):
	$(GO) build -o $(BINDIR)/$(BINARY) ./cmd/go-boy/...

clean:
	$(GO) clean
	rm $(BINDIR)/$(BINARY)

test:
	$(GO) test -cover ./...

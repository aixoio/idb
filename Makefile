.DEFAULT_GOAL=prod

GO=go
PRODARGS=-ldflags "-s -w"
BIN=target
DB=idb.sqlite3
ARGS=-o $(BIN)
RM=rm

prod:
	$(GO) build $(PRODARGS) $(ARGS)

build:
	$(GO) build $(ARGS)

clean:
	$(GO) clean
	-$(RM) $(BIN)
	-$(RM) $(DB)

.PHONY: prod build clean

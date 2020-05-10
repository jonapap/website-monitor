BIN=bin
CMD=cmd

PROG=websitedeleter websitechecker websitesaver
LIST=$(addprefix $(BIN)/, $(PROG))

all: $(LIST)

$(BIN)/%: FORCE | $(BIN)
	go build -o $(BIN)/ ./$(CMD)/$*

FORCE:

$(BIN):
	mkdir $(BIN)

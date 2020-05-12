BIN=bin
CMD=cmd

PROG=websitedeleter websitechecker websitesaver
LIST=$(addprefix $(BIN)/, $(PROG))

.PHONEY: all
all: $(LIST)

$(BIN)/%: FORCE | $(BIN)
	go build -o $(BIN)/ ./$(CMD)/$*

FORCE:

$(BIN):
	mkdir $(BIN)

#Enable the compilation of indiviual program (ex: make websitedeleter)
.SECONDEXPANSION:
$(PROG): $(BIN)/$$@

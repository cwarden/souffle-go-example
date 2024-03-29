SHELL=/bin/bash

INTERFACE := SwigInterface/SwigInterface.h SwigInterface/SwigInterface.swigcxx
APP := SwigInterface/example.cpp example.go
DATA := edge.db

run: $(APP) $(DATA)
	go run -x .

build: souffle-go-example

souffle-go-example: $(APP) $(DATA)
	go build -x .

SwigInterface/%.cpp: %.dl $(INTERFACE) cmd-exists-souffle
	souffle --jobs auto -g $@ $<

%.db: %.csv cmd-exists-sqlite3
	-rm -f $@
	sqlite3 -echo $@ -cmd ".mode csv" ".import $< $*"

cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)

clean:
	-rm -f *.db souffle-go-example path.csv SwigInterface/*.cpp

.DELETE_ON_ERROR:

.PHONY: clean run build

.SECONDARY:

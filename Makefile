SHELL=/bin/bash

run: SwigInterface/example.cpp edge.db
	go run -x .

souffle-go-example: SwigInterface/example.cpp edge.db
	go build -x .

SwigInterface/%.cpp: %.dl cmd-exists-souffle
	souffle  -g $@ $<

%.db: %.csv cmd-exists-sqlite3
	-rm -f $@
	sqlite3 -echo $@ -cmd ".mode csv" ".import $< $*"

cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)

clean:
	-rm -f *.db souffle-go-example path.csv

.DELETE_ON_ERROR:

.PHONY: clean run

.SECONDARY:

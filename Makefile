.PHONY: all test push run

all:
	go install -ldflags=-s

test:
	# go test -bench=. -benchtime=100x
	# go test -bench=. -benchtime=1s
	go test -bench=.

push:
	git push

run:
	go run .

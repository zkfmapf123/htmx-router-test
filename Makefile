run:
	go run main.go

.PHONY: build
build: clean
	cd static && templ generate

.PHONY: run
run: 
	air

.PHONY: clean
clean:
	clear && rm -rf static/*.go


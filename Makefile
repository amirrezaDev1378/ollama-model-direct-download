.PHONY: build clear


build:
	clear
	./scripts/build-for-all.sh

clear:
	rm -rf ./bin

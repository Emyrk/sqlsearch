.PHONY: shell gen

shell:
	nix develop -c $SHELL

gen:
	go generate ./...
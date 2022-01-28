INSTALL_DIR ?= ~/bin
VERSION=0.0.1

SRC=$(shell find . -name '*.go') go.mod
awsp: $(SRC)
	go build -o ./bin/_awsp_prompt .

install: awsp
	mkdir -p $(INSTALL_DIR)
	rm -f $(INSTALL_DIR)/_awsp_prompt
	rm -f $(INSTALL_DIR)/_awsp
	cp -a ./bin/_awsp $(INSTALL_DIR)/_awsp
	cp -a ./bin/_awsp_prompt $(INSTALL_DIR)/_awsp_prompt

uninstall:
	rm -f $(INSTALL_DIR)/_awsp_prompt
	rm -f $(INSTALL_DIR)/_awsp

release: SHA256SUMS
	@echo "\nTo create a new release run:\n\n    gh release create --title v$(VERSION) v$(VERSION) \
	bin/_awsp \
	bin/_awsp_prompt \
	SHA256SUMS\n"

SHA256SUMS:
	shasum -a 256 \
	  bin/_awsp \
	  bin/_awsp_prompt \
		> $@
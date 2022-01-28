INSTALL_DIR ?= ~/bin

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
.PHONY: clean build install go-format go-lint go-doc tf-doc



clean:
	$(CLEAN_CMD) $(CLEAN_OPTS) $(BUILD_OPTS_OUTDIR)

build: clean
	$(BUILD_MKDIR) ./$(BUILD_OPTS_OUTDIR)
	$(BUILD_COMPILER) $(BUILD_COMMAND) $(BUILD_OPTS) -o $(BUILD_OUT)

install: build
	$(INSTALL_MK_CMD) $(INSTALL_MK_OPTS) $(INSTALL_LOCATION)
	$(INSTALL_MV_CMD) $(INSTALL_MV_OPTS) $(BUILD_OUT) $(INSTALL_LOCATION)


go-format:
	go fmt ./...

go-lint:

go-doc:

tf-doc:

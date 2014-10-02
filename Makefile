SRC = *.go
build: $(SRC)
	@go build
	@echo 'done.'
	@echo 'golock binary size:'
	@ls -lh cronlock | awk '{print $$5}'



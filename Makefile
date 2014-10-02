SRC = *.go
build: $(SRC)
	@go build
	@echo 'done.'
	@echo 'golock binary size:'
	@ls -lh cronlock | awk '{print $$5}'
package: 
	@rake package
	@echo 'done.'
	@ls -lh pkg

.PHONY: package


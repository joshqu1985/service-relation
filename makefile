all:
	@go build -o relation cmd/*.go
	@echo done.

clean:
	cd ../../proto/pb/ && ./clear.sh
	@go clean cmd/*
	@rm -f ./relation

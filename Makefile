build:
	@echo "  >  Building binary..."
	go build -o podctl cmd/podctl/main.go
	chmod 755 podctl
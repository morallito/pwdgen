# Build the pwdgen executable
pwdgen:
	go build -o pwdgen .

# Clean build artifacts
clean:
	rm -f pwdgen

# Install dependencies and build
install: 
	go mod download
	go mod tidy
	make pwdgen

# Run the application (after building)
run: pwdgen
	./pwdgen

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o pwdgen-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o pwdgen-darwin-amd64 .
	GOOS=windows GOARCH=amd64 go build -o pwdgen-windows-amd64.exe .

.PHONY: pwdgen clean install run build-all

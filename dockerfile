# Use a base image suitable for your application
FROM golang:1.21.1-bookworm

# workdir for repositories
WORKDIR /app

CMD ["sleep", "infinity"]

# GOOS=windows GOARCH=amd64 go build -o myprogram-windows-amd64.exe
# golang:1.21.1-bookworm
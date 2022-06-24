all: jmserver

init:
	go mod init jyotish

tidy:
	go mod tidy

jmserver:
	go build -o $@

clean:
	rm -rf jmserver

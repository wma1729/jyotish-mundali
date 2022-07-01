all: jmserver

init:
	go mod init jyotish

tidy:
	go mod tidy
	go get github.com/lib/pq
	go get gopkg.in/yaml.v3

jmserver:
	go build -o $@

clean:
	rm -rf jmserver

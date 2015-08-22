.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	goconvey -port 7042 -cover -workDir=$(realpath .)/ -depth=0

#

.PHONY : build
build : fmt test
	@echo "可以提交"

.PHONY : fmt
fmt :
	@echo "格式化代码"
	@gofmt -l -w ./



.PHONY : test
test :
	@echo "测试代码"
	@go test -v ./...
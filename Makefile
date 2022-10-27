#

.PHONY : build
build : fmt test
	@echo "可以提交"

.PHONY : fmt
fmt :
	@echo "格式化代码"
	@gofmt -l -w ./


SUBDIRS := $(wildcard */)

.PHONY : test
test :
	@echo "测试代码"
	for dir in $(SUBDIRS); do \
        cd $$dir; \
        go test -v; \
        cd ..;\
	done

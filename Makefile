CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/goid.a build/goid.so

.PHONY: all
all: $(LIBS)

build/goid.a: goid.go
	$(CGO) $(STATIC) -o build/goid.a $<

build/goid.so: goid.go
	$(CGO) $(SHARED) -o build/goid.so $<

.PHONY: clean
clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete

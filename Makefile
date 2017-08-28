ver?=unknown
name=espectro
os=linux_amd64 linux_arm darwin_amd64
os_local=$(shell uname | tr A-Z a-z)_amd64
sources=$(shell find . -name \*.go | grep -v vendor)
test_files=$(shell find . -name \*_test.go | grep -v vendor)

build=GOOS=$(2) GOARCH=$(3) go build -ldflags "-s -X main.version=${ver}" -o bin/$(2)_$(3)/$(1) cmd/$(1)/*.go

default: cover

test: $(test_files)
	go test -coverprofile=cover.out

cover: test
	go tool cover -func=cover.out

local: bin/$(os_local)/$(name)

bin/%/$(name): os=$(firstword $(subst _, ,$*))
bin/%/$(name): arch=$(lastword $(subst _, ,$*))
bin/%/$(name): $(sources)
	$(call build,$(name),$(os),$(arch))

dist/$(name)_$(ver)_%.tgz: bin/%/$(name)
	mkdir -p $(@D) && tar -C $(<D) -czf $@ $(^F)

dist: $(os:%=dist/$(name)_$(ver)_%.tgz)

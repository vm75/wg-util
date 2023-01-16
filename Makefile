TAG=`git describe --tags`
VERSION ?= `[ -d ".git" ] && git describe --tags || echo "0.0.0"`
LDFLAGS=-ldflags "-s -w -X main.appVersion=${VERSION}"
BINARY="wg-util"

build = echo "\n\nBuilding $(1)-$(2)" && GO386=softfloat GOOS=$(1) GOARCH=$(2) go build ${LDFLAGS} -o dist/${BINARY}_${VERSION}_$(1)_$(2) \
	&& gzip dist/${BINARY}_${VERSION}_$(1)_$(2) \
	&& if [ $(1) = "windows" ]; then mv dist/${BINARY}_${VERSION}_$(1)_$(2).gz dist/${BINARY}_${VERSION}_$(1)_$(2).exe.gz; fi

build: *.go go.*
	go build ${LDFLAGS} -o ${BINARY}
	rm -rf /tmp/go-*

clean:
	rm -f ${BINARY}

release:
	mkdir -p dist
	rm -f dist/${BINARY}_${VERSION}_*
	$(call build,linux,arm)
	$(call build,linux,amd64)

# $(call build,linux,386)
# $(call build,linux,arm64)
# $(call build,darwin,arm64)
# $(call build,darwin,amd64)
# $(call build,windows,386)
# $(call build,windows,amd64)

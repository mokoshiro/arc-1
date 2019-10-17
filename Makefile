local-db-up:
	make -C infra/local up
local-db-down:
	make -C infra/local down

BZFLAGS =
ifdef EXTENDED_BAZELRC
	BZFLAGS += --bazelrc=${EXTENDED_BAZELRC}
endif

.PHONY: build
build:
	bazel ${BZFLAGS} build -k -- //pkg/...

.PHONY: test
test:
	bazel ${BZFLAGS} test -- //pkg/...

.PHONY: dep
dep:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
	bazel run //:gazelle -- update-repos -from_file=go.mod

.PHONY: gazelle
gazelle:
	bazel run //:gazelle

.PHONY: clean
clean:
	bazel clean

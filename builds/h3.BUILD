load("@io_bazel_rules_go//go:def.bzl", "go_library")


go_library(
    name = "go_default_library",
    srcs = glob(["*.go", "*.c",  "include/*.h"], exclude=["*_test.go"]),
    importpath = "github.com/uber/h3-go",
    cgo = True,
    copts = [
        "-Iexternal/com_github_uber_h3_go",
    ],
    cdeps = [
        // "@com_github_hayabusa_havana_clibs//libvips:libvips_library",
    ],
    visibility = ["//visibility:public"],
)
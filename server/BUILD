load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "server",
    srcs = ["server.go"],
    importpath = "test",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gin-gonic/gin:go_default_library",
    ],
)
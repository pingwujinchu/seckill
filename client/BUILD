load("@rules_cc//cc:defs.bzl", "cc_binary")

cc_binary(
    name = "client",
    srcs = ["main.cc","HttpGetPostMethod.h", "httpget.cpp"],
    deps = [
        "//tools:cmdline",
    ],
)


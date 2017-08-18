package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_prefix")

go_prefix("github.com/douglas-reid/mixer-noop-reporter")

go_library(
    name = "go_default_library",
    srcs = ["reporter.go"],
    data = ["@com_github_google_protobuf//:well_known_protos"],
    deps = [
        "//config:go_default_library",
        "@com_github_gogo_protobuf//proto:go_default_library",
        "@com_github_istio_api//:mixer/v1",  # keep
        "@com_github_istio_api//:mixer/v1/config/descriptor",  # keep        
        "@com_github_istio_mixer//pkg/adapter:go_default_library",
        "@com_github_istio_mixer//template/logentry:go_default_library",
    ],
)

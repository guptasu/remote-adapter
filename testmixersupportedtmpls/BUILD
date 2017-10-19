package(default_visibility = ["//visibility:public"])

load("@com_github_istio_mixer//tools/codegen:generate.bzl", "mixer_supported_template_library")

mixer_supported_template_library(
    name = "go_default_library",
    packages = {
        "@com_github_guptasu_remoteTemplate//:go_default_library_proto.descriptor_set": "github.com/guptasu/remoteRpt",
    },
    deps = [
        "@com_github_guptasu_remoteTemplate//:go_default_library",
    ],
)

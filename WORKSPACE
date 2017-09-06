workspace(name = "com_github_douglasreid_mixer_noop_reporter")




git_repository(
    name = "com_github_google_protobuf",
    commit = "52ab3b07ac9a6889ed0ac9bf21afd8dab8ef0014",  # Oct 4, 2016 (match pubref dep)
    remote = "https://github.com/google/protobuf.git",
)



git_repository(
    name = "io_bazel_rules_go",
    commit = "7991b6353e468ba5e8403af382241d9ce031e571",  # Aug 1, 2017 (gazelle fixes)
    remote = "https://github.com/bazelbuild/rules_go.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "go_repository")

go_repositories()

go_repository(
    name = "com_github_istio_mixer",
    commit = "bb397adeb8560516121493174136263c600ef6d2",
    importpath = "istio.io/mixer",
)

#load("@com_github_istio_mixer//:adapter_author_deps.bzl", "mixer_adapter_repositories")

#mixer_adapter_repositories()

git_repository(
    name = "org_pubref_rules_protobuf",
    commit = "9ede1dbc38f0b89ae6cd8e206a22dd93cc1d5637",  # Mar 31, 2017 (gogo* support)
    remote = "https://github.com/pubref/rules_protobuf",
)

bind(
    name = "protoc",
    actual = "@com_github_google_protobuf//:protoc",
)

go_repository(
    name = "org_golang_x_net",
    commit = "f5079bd7f6f74e23c4d65efa0f4ce14cbd6a3c0f",  # Jul 26, 2017 (no releases)
    importpath = "golang.org/x/net",
)

go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b4da8ac2ed8007337bc5eb5007998",  # Jan 26, 2016 (no releases)
    importpath = "github.com/golang/glog",
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "8ee79997227bf9b34611aee7946ae64735e6fd93",  # Nov 16, 2016 (match pubref dep)
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_gogo_protobuf",
    commit = "100ba4e885062801d56799d78530b73b178a78f3",  # Mar 7, 2017 (match pubref dep)
    importpath = "github.com/gogo/protobuf",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "8050b9cbc271307e5a716a9d782803d09b0d6f2d",  # Apr 7, 2017 (v1.2.1)
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "org_golang_x_text",
    build_file_name = "BUILD.bazel",
    commit = "f4b4367115ec2de254587813edaa901bc1c723a8",  # Mar 31, 2017 (no releases)
    importpath = "golang.org/x/text",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    commit = "ed905158d87462226a13fe39ddf685ea65f1c11f",  # Dec 16, 2016 (no releases)
    importpath = "github.com/hashicorp/go-multierror",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    commit = "7554cd9344cec97297fa6649b055a8c98c2a1e55",  # Oct 27, 2014 (no releases)
    importpath = "github.com/hashicorp/errwrap",
)

load("@com_github_istio_mixer//:istio_api.bzl", "go_istio_api_repositories")

go_istio_api_repositories()

load("@com_github_istio_mixer//:googleapis.bzl", "go_googleapis_repositories")

go_googleapis_repositories()

load("@com_github_istio_mixer//:x_tools_imports.bzl", "go_x_tools_imports_repositories")

go_x_tools_imports_repositories()

go_repository(
    name = "com_github_spf13_cobra",
    commit = "35136c09d8da66b901337c6e86fd8e88a1a255bd",  # Jan 30, 2017 (no releases)
    importpath = "github.com/spf13/cobra",
)

go_repository(
    name = "com_github_spf13_pflag",
    commit = "9ff6c6923cfffbcd502984b8e0c80539a94968b7",  # Jan 30, 2017 (no releases)
    importpath = "github.com/spf13/pflag",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    commit = "76626ae9c91c4f2a10f34cad8ce83ea42c93bb75",
    importpath = "github.com/inconshreveable/mousetrap",
)




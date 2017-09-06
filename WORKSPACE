workspace(name = "com_github_douglasreid_mixer_noop_reporter")

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

load("@com_github_istio_mixer//:adapter_author_deps.bzl", "mixer_adapter_repositories")

load("@com_github_istio_mixer//:istio_api.bzl", "go_istio_api_repositories")

load("@com_github_istio_mixer//:googleapis.bzl", "go_googleapis_repositories")

load("@com_github_istio_mixer//:x_tools_imports.bzl", "go_x_tools_imports_repositories")

mixer_adapter_repositories()

go_istio_api_repositories()

go_googleapis_repositories()

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




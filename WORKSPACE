# Bazel WORKSPACE file.
# https://docs.bazel.build/versions/master/build-ref.html

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "eb49eb1db1da156a597efed8707aec9d9f9ebc57",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains(go_version="1.9.4")

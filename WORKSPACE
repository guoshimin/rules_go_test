# Bazel WORKSPACE file.
# https://docs.bazel.build/versions/master/build-ref.html

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "f44e447a6fbbf26bfd31305375bf60547aad8a82",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

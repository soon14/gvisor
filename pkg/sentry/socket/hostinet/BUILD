load("//tools:defs.bzl", "go_library")

package(licenses = ["notice"])

go_library(
    name = "hostinet",
    srcs = [
        "device.go",
        "hostinet.go",
        "save_restore.go",
        "socket.go",
        "socket_unsafe.go",
        "sockopt_impl.go",
        "stack.go",
    ],
    visibility = ["//pkg/sentry:internal"],
    deps = [
        "//pkg/abi/linux",
        "//pkg/binary",
        "//pkg/context",
        "//pkg/fdnotifier",
        "//pkg/log",
        "//pkg/safemem",
        "//pkg/sentry/arch",
        "//pkg/sentry/device",
        "//pkg/sentry/fs",
        "//pkg/sentry/fs/fsutil",
        "//pkg/sentry/inet",
        "//pkg/sentry/kernel",
        "//pkg/sentry/kernel/time",
        "//pkg/sentry/socket",
        "//pkg/sentry/socket/control",
        "//pkg/syserr",
        "//pkg/syserror",
        "//pkg/tcpip/stack",
        "//pkg/usermem",
        "//pkg/waiter",
        "@org_golang_x_sys//unix:go_default_library",
    ],
)

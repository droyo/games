load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_embed_data")

go_embed_data(
	name = "assets",
	package = "main",
	srcs = [
		"roguelikeChar_transparent.png",
		"roguelikeSheet_transparent.png",
	],
	flatten = True,
	var = "BuiltinAssets",
	visibility = ["//visibility:private"],
)

go_binary(
    name = "mazewalker",
    srcs = [
        "mazewalker.go",
    	":assets",
    ],
    importpath = "aqwari.net/game/mazewalker",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_faiface_pixel//:pixel",
        "@com_github_faiface_pixel//:pixelgl",
        "@org_golang_x_image//:colornames",
    ],
)

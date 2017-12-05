
http_archive(
	name = "io_bazel_rules_go",
	urls = ["https://github.com/bazelbuild/rules_go/archive/0.7.1.tar.gz"],
	strip_prefix = "rules_go-0.7.1",
	sha256 = "99fb948c4fae06761b22a0b7a17fce86d517567c4ef9b90a4451c985b21c5e1c",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")

go_rules_dependencies()
go_register_toolchains()

go_repository(
	name = "org_golang_x_image",
	urls = ["https://github.com/golang/image/archive/f7e31b4ea2e3413ab91b4e7d2dc83e5f8d19a44c.tar.gz"],
	strip_prefix = "image-f7e31b4ea2e3413ab91b4e7d2dc83e5f8d19a44c",
	sha256 = "089f7f7a5f272e069dbf2edf31c23a66d4b9a64d2c470f60f9c038dc85907860",
	importpath = "golang.org/x/image",
)

go_repository(
	name = "com_github_pkg_errors",
	urls = ["https://github.com/pkg/errors/archive/v0.8.0.tar.gz"],
	strip_prefix = "errors-0.8.0",
	sha256 = "bacf6c58e490911398cee61742ddc6a90c560733e4c9dcb3d867b17a894c9dd5",
	importpath = "github.com/pkg/errors",
)

go_repository(
	name = "com_github_go_gl_mathgl",
	urls = ["https://github.com/go-gl/mathgl/archive/4c3fc6b4bf30179013667e5d0b926e024d9a13c1.tar.gz"],
	strip_prefix = "mathgl-4c3fc6b4bf30179013667e5d0b926e024d9a13c1",
	sha256 = "692ee6d871744d46a6bec5ede064a78fb663085fe82fdc100f5aecd30b3f3584",
	importpath = "github.com/go-gl/mathgl",
)

new_http_archive(
	name = "com_github_go_gl_glfw",
	urls = ["https://github.com/go-gl/glfw/archive/513e4f2bf85c31fba0fc4907abd7895242ccbe50.tar.gz"],
	strip_prefix = "glfw-513e4f2bf85c31fba0fc4907abd7895242ccbe50/v3.2/glfw",
	sha256 = "e2e90c772b333ee89aa377ef42bbc0bdfca6040ab21d3595b782bf8e95f42504",
	build_file = "third_party/BUILD.com_github_go_gl_glfw",
)

new_http_archive(
	name = "com_github_faiface_pixel",
	urls = ["https://github.com/faiface/pixel/archive/v0.6.tar.gz"],
	strip_prefix = "pixel-0.6",
	sha256 = "fc6188987306d426810bfa9ce184a746fae62db02b3148ee603108e0d0cc4d99",
	build_file = "third_party/BUILD.com_github_faiface_pixel",
)

go_repository(
	name = "com_github_faiface_mainthread",
	urls = ["https://github.com/faiface/mainthread/archive/8b78f0a41ae388189090ac4506612659fa53082b.tar.gz"],
	strip_prefix = "mainthread-8b78f0a41ae388189090ac4506612659fa53082b",
	sha256 = "8432cb8b9817a271d4b2cfe8cb082ceb95fb83b3f6cfc3c295ada290b0b4c971",
	importpath = "github.com/faiface/mainthread",
)

go_repository(
	name = "com_github_faiface_glhf",
	urls = ["https://github.com/faiface/glhf/archive/98c0391c0fd3f0b365cfe5d467ac162b79dfb002.tar.gz"],
	strip_prefix = "glhf-98c0391c0fd3f0b365cfe5d467ac162b79dfb002",
	sha256 = "e1577d40f1b265a8c826cdd7d306d3af2331190b45561c71404959aa25fc9442",
	importpath = "github.com/faiface/glhf",
)

go_repository(
	name = "com_github_go_gl_gl",
	urls = ["https://github.com/go-gl/gl/archive/ac0d3d2af0fe995e2bb09fa6619b716b0c0fc0fa.tar.gz"],
	strip_prefix = "gl-ac0d3d2af0fe995e2bb09fa6619b716b0c0fc0fa",
	sha256 = "12b5301f8c7739d4ee5fc88d2ecfae13abe36bbffef6abfaf843334010b35c54",
	importpath = "github.com/go-gl/gl",
)

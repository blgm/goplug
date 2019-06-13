## Go Plugins example

Works on Linux and Mac with Go 1.10 and higher.

The `plug-alpha` directory has a `go-flags` command which can be built as a shared object using `build.sh`

The `cmd` directory has a Go command which has a built-in `go-flags` command, and also dynamically (at runtime) loads
the `plugin-alpha` plugin and exposes that command too.

To run it, run: `go run main.go builtin` and `go run main.go alpha`.

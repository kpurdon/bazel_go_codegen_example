def _hellogen_impl(ctx):
    out = ctx.actions.declare_file("hello.go")

    args = ctx.actions.args()
    args.add("-config-file", ctx.file.config.path)
    args.add("-output-file", out.path)

    ctx.actions.run(
        inputs = [ctx.file.config],
        outputs = [out],
        executable = ctx.executable._generator,
        tools = [ctx.executable._generator],
        arguments = [args],

    )

    return [
        DefaultInfo(files = depset([out])),
    ]

hellogen = rule(
    implementation = _hellogen_impl,
    attrs = {
        "config": attr.label(
            allow_single_file = True,
        ),
        "_generator": attr.label(
            executable = True,
            cfg = "host",
            default = "//hellogen",
        ),
    },
)

# Library Versions

You'll notice the pattern "somelibrary/v2" for major updates.

One rule: "Import compatibility rule"

* https://research.swtch.com/vgo-import

> How do you deploy an incompatible change to an existing package? This is the
fundamental challenge, the fundamental decision, in any package management
system. The answer decides the complexity of the resulting system It decides how
easy or difficult package management will be to use. (It also decides how easy
or difficult package management will be to implement, but the user experience is
more important.)

> To answer this question, this post first presents the import compatibility rule for Go:

**If an old package and a new package have the same import path, the new package
must be backwards compatible with the old package.**

Another angle:

> If you break compatibility, you should change the import path of your package.
> With Go 1.11 modules, that advice is formalized into the import compatibility
> rule.

## Basics

For the module name, go from `my/thing` to `my/thing/v2` - so people can
explicitly name which major version they want.
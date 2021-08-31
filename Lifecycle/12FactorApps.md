# 12 Factor Apps

* https://12factor.net/

> In the modern era, software is commonly delivered as a service: called web
> apps, or software-as-a-service. The twelve-factor app is a methodology for
> building software-as-a-service apps that:

* Use declarative formats for setup automation, to minimize time and cost for new developers joining the project;
* Have a clean contract with the underlying operating system, offering maximum portability between execution environments;
* Are suitable for deployment on modern cloud platforms, obviating the need for servers and systems administration;
* Minimize divergence between development and production, enabling continuous deployment for maximum agility;
* And can scale up without significant changes to tooling, architecture, or development practices.

----

## Codebase

> One codebase tracked in revision control, many deploys

Go is loved for it's deployment story. You can build a static binary and ship
it. You can wrap it in minimalistic containers.

## Dependencies

> Explicitly declare and isolate dependencies

Dependency management converged to go modules.


## Config

> Store config in the environment

Examples:

* [https://github.com/kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) (env)
* [https://github.com/rakyll/globalconf](https://github.com/rakyll/globalconf) (flags, files, env)
* [https://github.com/spf13/viper](https://github.com/spf13/viper) (everything)

## Backing services

> Treat backing services as attached resources

See config.

## Build, release, run

> Strictly separate build and run stages

Go projects build fast. Also: it's easy to crosscompile.

## Processes

> Execute the app as one or more stateless processes

Push state into the backing store.

> A twelve-factor app prefers to do this compiling during the build stage.

Go allows to embed files into the binary (with tools, or since Go 1.16 with package [https://pkg.go.dev/embed](https://pkg.go.dev/embed))

## Port binding

> Export services via port binding

Starting a web server directly.

> The twelve-factor app is completely self-contained and does not rely on
> runtime injection of a webserver into the execution environment to create a
> web-facing service. The web app exports HTTP as a service by binding to a
> port, and listening to requests coming in on that port.

That's what `net/http` and `ListenAndServe` is for.

## Concurrency

> Scale out via the process model

Let the surrounding tool (e.g. systemd, docker, orchestrator) manage and scale
your app.

## Disposability

> Maximize robustness with fast startup and graceful shutdown

Go services start fast (no separate project for building native binaries necessary). 

## Dev/prod parity

> Keep development, staging, and production as similar as possible

With easier setup, multiple environments may become easier.

## Logs

> Treat logs as event streams

Go has a couple of logging libraries.

> A twelve-factor app never concerns itself with routing or storage of its
> output stream. It should not attempt to write to or manage logfiles. Instead,
> each running process writes its event stream, unbuffered, to stdout.

Structured logging can help.

* [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* [https://github.com/rs/zerolog](https://github.com/rs/zerolog)

## Admin processes

> Run admin/management tasks as one-off processes

Go is well suited for command line tools.
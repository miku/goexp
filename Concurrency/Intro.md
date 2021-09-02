# Intro

## CSP

* Communicating Sequential Processes (Hoare, 1978)
* Share memory by communicating, don't communicate by sharing memory

This mantra contains "communication" - more message, less queue.

Core notions: goroutines, channels, select.

## Misc

The channel axioms:

* A send to a nil channel blocks forever
* A receive from a nil channel blocks forever
* A send to a closed channel panics
* A receive from a closed channel returns the zero value immediately


## Considerations

* a goroutine is just a function, running in another thread, everything that is
  wrong in sequential code will be wrong in a goroutine as well
* e.g. an endless loop, the goroutine will never exit
* communicating errors is harder, we cannot just return an "error" since
  program execution continues immediately after the call
* there is no way to stop a goroutine from the outside

> **Cooperative multitasking**, also known as non-preemptive multitasking, is a
> style of computer multitasking in which the operating system never initiates
> a context switch from a running process to another process. Instead, in order
> to run multiple applications concurrently, processes voluntarily yield
> control periodically or when idle or logically blocked.

Package context was added only in Go 1.7, as a way to unify cancellation a bit.

* go concurrency gives you the parts, but no no level constructs
* you library should only expose channels when it must
* keep concurrency under the hood but you API sequential
* callbacks are possible, but it is not the pattern of choice in go
* scheduling is done by the go runtime, indeterminate
* when multiple select cases are possible one is chosen at random

A study found just as many bugs in Go code using CSP vs classic structures like mutexes


## Typical patterns

* fan out, fan out-fan it (workers and queue)
* bounded concurrency
* timeouts, [Timeout](timeout/main.go)
* signal handling, [Signal](signal/main.go)

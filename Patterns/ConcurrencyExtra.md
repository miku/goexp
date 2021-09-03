# Concurrency Extra

## Context

Context helps to implement cancellation.

* [Context](Context)

A couple of notes:

* it is a value that you pass into a function

For example, you can see it it in the `db/sql` package, some functions are named `...Context`, e.g. https://pkg.go.dev/database/sql#DB.QueryContext

* it signals that this might be a longer running operation and that the user has the chance to singal cancellation to this function after it has started
* there is an initial `context.Background` that serves as a starting point

There are 3+1 ways to created a context value:

```go
// WithCancel returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed or cancel is called.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// WithDeadline returns a copy of the parent context with the deadline adjusted
// to be no later than d. If the parent's deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
// context's Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is
// closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) 

// WithTimeout returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed, cancel is called, or timeout elapses. The new
// Context's Deadline is the sooner of now+timeout and the parent's deadline, if
// any. If the timer is still running, the cancel function releases its
// resources.
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

// ...
// WithValue returns a copy of parent in which the value associated with key is
// val.
//
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
//
// The provided key must be comparable and should not be of type
// string or any other built-in type to avoid collisions between
// packages using context. Users of WithValue should define their own
// types for keys. To avoid allocating when assigning to an
// interface{}, context keys often have concrete type
// struct{}. Alternatively, exported context key variables' static
// type should be a pointer or interface.
func WithValue(parent Context, key, val interface{}) Context
```

You get a context back and a function to cancel it. You want to call `cancel` excplicitly for `WithCancel`. For `WithDeadline` and `WithTimeout` you typically `defer` the cancel to release resources.


## Errgroup



## Aggregate Errors

## Hedged Requests

## Nested Goroutines

## Other Patterns
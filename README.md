# ptr
Simple pointer helpers for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/candiduslynx/ptr.svg)](https://pkg.go.dev/github.com/candiduslynx/ptr)

## Obtaining pointer to the value

Simply pass the value to the `To` function, the type will be inferred by the Go compiler.

```go
p := ptr.To(float32(0.5)) // p is *float32 pointing to the value 0.5
```

## Retrieving value designated by pointer

To retrieve the value designated by pointer simply pass the pointer to `From` function.
If the pointer might be `nil` you can supply optional default value as well.
if the pointer is nil and default value isn't provided, `From` will panic.

```go
// panics
_ = ptr.From((*float32)(nil))

var vv = floaat32(0.5)
value = ptr.From(&vv) // value = float32(0.5)
value = ptr.From((*float32)(nil), 0.7) // value = float32(0.7)
```
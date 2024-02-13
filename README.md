## copyloopvar
copyloopvar is a linter detects places where loop variables are copied.

cf. [Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview)

## Example
```go
for i, v := range []int{1, 2, 3} {
    i := i // The loop variable "i" doesn't need to be copied
    v := v // The loop variable "v" doesn't need to be copied
    _, _ = i, v
}

for i := 1; i <= 3; i++ {
    i := i // The loop variable "i" doesn't need to be copied``
    _ = i
}
```

## Install
```bash
go install github.com/karamaru-alpha/copyloopvar/cmd/copyloopvar
go vet -vettool=`which copyloopvar` ./...
```

# copyloopvar

copyloopvar is a linter detects places where loop variables are copied.

cf. [Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview)

## Example

```go
for i, v := range []int{1, 2, 3} {
    i := i // It's unnecessary to copy the loop variable "i"
    v := v // It's unnecessary to copy the loop variable "v"
    _, _ = i, v
}

for i := 1; i <= 3; i++ {
    i := i // It's unnecessary to copy the loop variable "i"
    _ = i
}
```

## Install

```bash
go install github.com/karamaru-alpha/copyloopvar/cmd/copyloopvar
go vet -vettool=`which copyloopvar` ./...
```

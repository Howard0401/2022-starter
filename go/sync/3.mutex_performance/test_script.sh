go test -bench . -benchmem main_test.go -cpu 2

go test -bench . -benchmem main_test.go -cpu 8

go test -bench . -benchmem main_test.go -cpu 16

go test -bench . -benchmem main_test.go -cpu 32

go test -bench . -benchmem main_test.go -cpu 64

go test -bench . -benchmem main_test.go -cpu 128

# https://stackoverflow.com/questions/35588474/what-does-allocs-op-and-b-op-mean-in-go-benchmark
# allocs/op means how many distinct memory allocations occurred per op (single iteration).
# B/op is how many bytes were allocated per op.
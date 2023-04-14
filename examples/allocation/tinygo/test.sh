cd testdata
tinygo build -o mytest.wasm -scheduler=none -target=wasi mytest.go
wasm2wat mytest.wasm > x
cd ..
go run greet.go

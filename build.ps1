$Env:GOOS='android'
$Env:GOARCH='arm64'
$Env:CGO_ENABLED='1'
$Env:CC='aarch64-linux-android31-clang.cmd'

go env
"请确认上述信息，按任意键开始编译"

[Console]::ReadKey() | Out-Null
"开始编译..."

go build -buildmode='c-shared' -ldflags='-s -w' -o='builds/libsm4.so' main.go

"完成！"
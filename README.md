# go-net
## Installing Go
1. Install GoLand
2. Install Go
   1. https://go.dev/doc/install
3. Install Go SDK
   1. File -> Settings Go|GOROOT
   2. Click '+' and select **Download**
   3. Select **OK**
   4. Wait a few minutes while it installs (Check the progress on the bottom of the screen)
4. Set the Global GOPATH as the install path
   1. File -> Settings Go|GOPATH
   2. Add the SDK to the Global Go PATH
## References
https://www.programming-books.io/essential/go/
https://www.rapid7.com/blog/post/2016/08/04/build-a-simple-cli-tool-with-golang/
https://github.com/blackhat-go/bhg/

## Go Basic Types
bool
string
int (int8, int16, int32, int64)
uint (uint8, ..., uint 64, uintptr)
byte (aka uint8)
rune (aka int32)
float 32, float64
compelx64, complex128

## Ideas
### Channels
Run a dictionary over passwords, or something and send those values to a channel to save time.
Compare run time without using channels and see if it saves any.
https://go.dev/tour/concurrency/2
https://github.com/kubernetes/kubernetes/blob/cd6ffff85d257ff9067d59339f2ffdbcc66dc164/staging/src/k8s.io/client-go/tools/portforward/portforward_test.go


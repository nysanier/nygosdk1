// You can edit this code!
// Click here and start typing.
package main

// import "github.com/nysanier/nygosdk1/lib"
import (
	"nylib" // gopath 指向当前repo, 否则就需要用下面的相对路径了
	// "./nylib"
	// "./nylib2/sub2"
	"nylib2/sub2" // go get 的时候如何依赖本地包?
)

func main() {
	nylib.Fc()
	sub2.Hh()
}

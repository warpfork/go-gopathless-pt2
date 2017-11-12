package gopathless2

import (
	"io"

	"github.com/warpfork/go-gopathless"
)

func FuncWithDependencies(w io.Writer) {
	gopathless.HelloLibrary(w)
}

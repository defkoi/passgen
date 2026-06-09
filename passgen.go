package passgen

import (
	"crypto/rand"
	_ "embed"
	"encoding/binary"

	"github.com/dop251/goja"
)

const Name = "passgen"

//go:embed passgen.js
var source string

var rt *goja.Runtime

func init() {
	rt = goja.New()
	rt.SetRandSource(randFloat)
}

func GeneratePassword(length int) (string, error) {
	rt.Set("LENGTH", length)

	v, err := rt.RunScript(Name, source)
	return v.String(), err
}

func SetScript(script string) {
	source = script
}

func randFloat() float64 {
	var buf [8]byte
	if _, err := rand.Read(buf[:]); err != nil {
		panic(err)
	}

	ui := binary.BigEndian.Uint64(buf[:]) >> 11
	return float64(ui) / (1 << 53)
}

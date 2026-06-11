package passgen

import (
	"crypto/rand"
	_ "embed"
	"encoding/binary"

	"github.com/dop251/goja"
)

const Name = "passgen"

type CharSet struct {
	Lowers string
	Uppers string
	Digits string
	Others string
}

//go:embed passgen.js
var source string

var rt *goja.Runtime

var defaultCharSet = CharSet{
	Lowers: "abcdefghijklmnopqrstuvwxyz",
	Uppers: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	Digits: "0123456789",
	Others: "`~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?",
}

func init() {
	rt = goja.New()
	rt.SetRandSource(randFloat)
	SetCharSet(defaultCharSet)
}

func GeneratePassword(length int) (string, error) {
	rt.Set("LENGTH", length)

	if v, err := rt.RunScript(Name, source); err != nil {
		return "", err
	} else {
		return v.String(), nil
	}
}

func SetScript(script string) {
	source = script
}

func SetCharSet(cs CharSet) {
	rt.Set("LOWERS", cs.Lowers)
	rt.Set("UPPERS", cs.Uppers)
	rt.Set("DIGITS", cs.Digits)
	rt.Set("OTHERS", cs.Others)
}

func randFloat() float64 {
	var buf [8]byte
	if _, err := rand.Read(buf[:]); err != nil {
		panic(err)
	}

	ui := binary.BigEndian.Uint64(buf[:]) >> 11
	return float64(ui) / (1 << 53)
}

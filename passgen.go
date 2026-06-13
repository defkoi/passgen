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

func init() {
	rt = goja.New()
	rt.SetRandSource(randFloat)
	SetCharSet(DefaultCharSet())
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
	obj := rt.NewObject()
	rt.Set("charset", obj)

	obj.Set("lowers", cs.Lowers)
	obj.Set("uppers", cs.Uppers)
	obj.Set("digits", cs.Digits)
	obj.Set("others", cs.Others)
}

func DefaultCharSet() CharSet {
	return CharSet{
		Lowers: "abcdefghijklmnopqrstuvwxyz",
		Uppers: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Digits: "0123456789",
		Others: "`~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?",
	}
}

func randFloat() float64 {
	var buf [8]byte
	if _, err := rand.Read(buf[:]); err != nil {
		panic(err)
	}

	ui := binary.BigEndian.Uint64(buf[:]) >> 11
	return float64(ui) / (1 << 53)
}

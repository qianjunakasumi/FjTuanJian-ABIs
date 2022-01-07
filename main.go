package main

import "C"
import (
	"encoding/hex"
	"strings"

	"github.com/tjfoc/gmsm/sm4"
)

// KEY from: https://m.fjcyl.com/admin/user/uindex#L107
/*
var keyRec = "A7E74D2B6282AEB1C5EA3C28D25660A7";
*/
var KEY = []byte{167, 231, 77, 43, 98, 130, 174, 177, 197, 234, 60, 40, 210, 86, 96, 167}

func main() {}

//export encode
func encode(s *C.char) *C.char {
	msg, _ := sm4.Sm4Ecb(KEY, []byte(C.GoString(s)), true)
	return C.CString(strings.ToUpper(hex.EncodeToString(msg)))
}

//export decode
func decode(s *C.char) *C.char {
	msg, _ := hex.DecodeString(C.GoString(s))
	o, _ := sm4.Sm4Ecb(KEY, msg, false)
	return C.CString(string(o))
}

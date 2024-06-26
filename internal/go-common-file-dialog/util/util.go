package util

import (
	"github.com/go-ole/go-ole"
	"github.com/google/uuid"
)


// ff:
// str:
func StringToUUID(str string) *ole.GUID {
	return ole.NewGUID(uuid.NewSHA1(uuid.Nil, []byte(str)).String())
}

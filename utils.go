package e

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

func init() {
	var seed int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed); err != nil {
		panic(err)
	}
	rand.Seed(seed)
}

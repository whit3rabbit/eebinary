package generate

import (
	"bytes"
	"compress/zlib"
)

func Compress(byteSlice []byte) []byte {

	// Compress byte array to zlib
	var b bytes.Buffer
	en := zlib.NewWriter(&b)
	if _, err := en.Write(byteSlice); err != nil {
		panic(err)
	}
	if err := en.Flush(); err != nil {
		panic(err)
	}
	if err := en.Close(); err != nil {
		panic(err)
	}

	//compressedString := base64.StdEncoding.EncodeToString(b.Bytes())

	return b.Bytes()

}

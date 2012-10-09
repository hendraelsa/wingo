package bindata

import (
	"bytes"
	"compress/gzip"
	"io"
)

// ClosePng returns the raw, uncompressed file data data.
func ClosePng() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xea, 0x0c,
		0xf0, 0x73, 0xe7, 0xe5, 0x92, 0xe2, 0x62, 0x60, 0x60, 0xe0, 0xf5, 0xf4,
		0x70, 0x09, 0x02, 0xd2, 0x0d, 0x20, 0xcc, 0xc1, 0x06, 0x24, 0x0f, 0xdb,
		0x25, 0x9e, 0x06, 0x52, 0x8c, 0xc5, 0x41, 0xee, 0x4e, 0x0c, 0xeb, 0xce,
		0xc9, 0xbc, 0x04, 0x72, 0xd8, 0x92, 0xbc, 0xdd, 0x5d, 0x18, 0xfe, 0x83,
		0xe0, 0x82, 0xbd, 0xcb, 0x27, 0x03, 0x45, 0x38, 0x0b, 0x3c, 0x22, 0x8b,
		0x19, 0x18, 0xb8, 0x85, 0x41, 0x98, 0x91, 0x61, 0xd6, 0x1c, 0x09, 0xa0,
		0x20, 0x7b, 0x89, 0xa7, 0xaf, 0x2b, 0xfb, 0x1d, 0x66, 0x69, 0x66, 0x3d,
		0xee, 0xef, 0xcc, 0x9b, 0xdc, 0x81, 0x42, 0xb2, 0x99, 0x21, 0x11, 0x25,
		0xce, 0xf9, 0xb9, 0xb9, 0xa9, 0x79, 0x25, 0x0c, 0x20, 0xe0, 0x5c, 0x94,
		0x9a, 0x58, 0x92, 0x9a, 0xa2, 0x50, 0x9e, 0x59, 0x92, 0xa1, 0xe0, 0xee,
		0xe9, 0x1b, 0x90, 0xa2, 0x97, 0xca, 0x0e, 0xb4, 0x4e, 0xda, 0xd3, 0xc5,
		0x31, 0xa4, 0xe2, 0xd6, 0xdb, 0xdb, 0xde, 0x7c, 0x87, 0x0c, 0x78, 0x58,
		0x0f, 0xb0, 0x3c, 0xfd, 0x17, 0x5e, 0xa6, 0x7f, 0x80, 0x6d, 0xa5, 0x49,
		0xef, 0xfd, 0x33, 0x62, 0x9c, 0x81, 0xb7, 0x3f, 0x6a, 0xef, 0x37, 0x64,
		0xc0, 0x02, 0x1a, 0xfe, 0x3e, 0x73, 0x5c, 0xbf, 0x35, 0xae, 0x20, 0xf6,
		0xea, 0xda, 0x6d, 0xf2, 0xe9, 0x0a, 0xb7, 0xb3, 0xaa, 0x62, 0x4b, 0x3f,
		0x6b, 0xae, 0x59, 0xc7, 0x7c, 0xa1, 0x96, 0x79, 0xca, 0xb3, 0x6e, 0x06,
		0x26, 0xe5, 0x07, 0x7d, 0x16, 0x2b, 0x5a, 0x18, 0x3d, 0x6e, 0x31, 0x2c,
		0xfd, 0xa9, 0xd4, 0x70, 0x80, 0xe3, 0x87, 0x2a, 0x9f, 0x56, 0x2e, 0x43,
		0x90, 0xb9, 0x43, 0xe6, 0x3e, 0x8e, 0x84, 0x39, 0x5f, 0x03, 0x3f, 0x68,
		0xaa, 0x97, 0xde, 0x62, 0xb8, 0xc1, 0xb0, 0x30, 0x3a, 0x41, 0xaa, 0xe1,
		0x8c, 0xa3, 0xc5, 0xb6, 0x3e, 0x25, 0x07, 0x96, 0x2d, 0x35, 0xde, 0xcb,
		0xdb, 0x1c, 0x99, 0x37, 0x64, 0x0b, 0x44, 0x5d, 0xf1, 0x63, 0xe9, 0x68,
		0x08, 0xf3, 0xba, 0xbb, 0x89, 0xfd, 0xc1, 0xfd, 0xe6, 0x0b, 0xeb, 0xe2,
		0x94, 0x14, 0xe6, 0x30, 0x1a, 0x30, 0xec, 0xe0, 0x9c, 0xf2, 0x37, 0x8f,
		0xc9, 0xf4, 0xfe, 0x73, 0x86, 0x0e, 0x06, 0x75, 0xfb, 0xb5, 0x6d, 0xfb,
		0x78, 0x79, 0xd6, 0x2f, 0xe6, 0x10, 0x60, 0xf4, 0xab, 0x2b, 0xb2, 0xfd,
		0xab, 0x66, 0xb0, 0x67, 0x19, 0xd0, 0x90, 0x23, 0xf3, 0xfe, 0xfc, 0xfc,
		0xf0, 0xf9, 0x36, 0x83, 0x44, 0x83, 0x79, 0xfe, 0xbe, 0x99, 0x11, 0x1d,
		0x4b, 0x3f, 0xc6, 0xcb, 0x7f, 0x7f, 0x9f, 0xc0, 0xe3, 0xf0, 0x89, 0x37,
		0xdf, 0x47, 0x46, 0xa7, 0xf7, 0x7c, 0x1f, 0xb3, 0xa0, 0x02, 0xd3, 0x9c,
		0x7f, 0x4f, 0x2c, 0x22, 0xa4, 0xe6, 0xcb, 0xb2, 0x2b, 0x61, 0xf3, 0x1e,
		0x43, 0xc2, 0x6e, 0x6e, 0x1d, 0x2b, 0x85, 0xb0, 0xd7, 0x86, 0x8d, 0x4b,
		0x41, 0x5c, 0x4f, 0x57, 0x3f, 0x97, 0x75, 0x4e, 0x09, 0x4d, 0x80, 0x00,
		0x00, 0x00, 0xff, 0xff, 0x67, 0x7f, 0x78, 0xda, 0xc4, 0x01, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}

package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func assets_templates_help_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x94, 0x54,
		0xdd, 0x4e, 0xdc, 0x3c, 0x10, 0xbd, 0xde, 0x3c, 0xc5, 0x88, 0xef, 0x96,
		0x24, 0xf0, 0xa1, 0xde, 0x6c, 0x43, 0x24, 0x40, 0x95, 0x5a, 0x89, 0x3f,
		0x05, 0xb8, 0x46, 0x26, 0x99, 0x6c, 0xac, 0x3a, 0xb1, 0x65, 0x4f, 0x80,
		0xa8, 0xea, 0xbb, 0x77, 0xec, 0xfc, 0x6c, 0x68, 0x45, 0x45, 0x25, 0x84,
		0x3c, 0x9e, 0xf1, 0x39, 0x67, 0xce, 0x4c, 0x36, 0x73, 0x34, 0x28, 0xcc,
		0xa3, 0x4d, 0x73, 0x0c, 0x3f, 0xa2, 0xcd, 0xa6, 0x15, 0x76, 0x27, 0xbb,
		0xf8, 0x49, 0x13, 0xe9, 0x76, 0x0b, 0x47, 0xc9, 0x27, 0x6c, 0x3f, 0x47,
		0x9b, 0x9f, 0x5c, 0x70, 0xb2, 0x2e, 0x20, 0x6d, 0xb6, 0x70, 0x1c, 0x72,
		0xef, 0xbf, 0xc9, 0xd2, 0x09, 0x3d, 0xca, 0x9a, 0xe3, 0xfc, 0x8e, 0xeb,
		0xe8, 0xf6, 0xec, 0x1a, 0xbe, 0xa2, 0x32, 0x59, 0xca, 0x37, 0x7c, 0xdf,
		0x2b, 0xe6, 0xce, 0x94, 0xcc, 0x33, 0x01, 0x8d, 0xc5, 0xfa, 0xf4, 0xe0,
		0xbf, 0xde, 0xc9, 0x6e, 0x17, 0x3b, 0x5f, 0x6d, 0x44, 0x17, 0xbf, 0x48,
		0x6a, 0xe2, 0x1d, 0xfa, 0xf3, 0x41, 0xfe, 0xe0, 0x73, 0xb0, 0x20, 0xf9,
		0x1c, 0x8c, 0xb9, 0x2c, 0x15, 0x79, 0x96, 0x32, 0xd0, 0x47, 0xe0, 0x4a,
		0x3e, 0xb5, 0xb1, 0xb6, 0x71, 0xc9, 0xb7, 0xfa, 0x1d, 0xdc, 0x50, 0x04,
		0xda, 0xc2, 0x45, 0x28, 0xda, 0x13, 0x64, 0xa9, 0x57, 0xed, 0x9b, 0x3a,
		0xf9, 0xab, 0x20, 0x4e, 0x73, 0x91, 0xc9, 0xc7, 0x18, 0x3a, 0x41, 0xf2,
		0x19, 0xd5, 0x00, 0xae, 0x37, 0x46, 0x5b, 0x72, 0xcb, 0xb3, 0x24, 0x4b,
		0xcd, 0x58, 0x7a, 0x2b, 0x9c, 0x03, 0x6a, 0x70, 0x8f, 0xd8, 0x4a, 0x6b,
		0x59, 0x43, 0x00, 0xf6, 0x89, 0xac, 0xd4, 0x15, 0xe6, 0x63, 0x3f, 0x59,
		0x1a, 0x02, 0x28, 0x75, 0xdb, 0x8a, 0xae, 0x02, 0x25, 0x3b, 0x04, 0x23,
		0xac, 0x68, 0x91, 0xd0, 0x6e, 0x67, 0x54, 0x8b, 0xb3, 0x84, 0xf1, 0x19,
		0x34, 0x44, 0x66, 0x9b, 0xa6, 0x46, 0x50, 0x93, 0x90, 0x4e, 0x67, 0x6f,
		0xb8, 0xde, 0xe2, 0xa8, 0xe3, 0xbe, 0x91, 0x0e, 0x5c, 0xa3, 0x7b, 0x55,
		0x81, 0x50, 0x4a, 0xbf, 0x4c, 0x4d, 0x01, 0x69, 0x50, 0xba, 0x14, 0x84,
		0x20, 0xba, 0x01, 0x9e, 0xd1, 0x3a, 0xa9, 0x3b, 0xd0, 0x75, 0x08, 0x5b,
		0x5d, 0xf5, 0x0a, 0xa1, 0xb6, 0xba, 0x8d, 0x7c, 0xcc, 0xd7, 0x5e, 0xb2,
		0x78, 0x16, 0x52, 0x89, 0x27, 0xce, 0xc8, 0xae, 0xc2, 0x57, 0x74, 0x53,
		0xbf, 0xef, 0x1a, 0xf8, 0x87, 0xf3, 0xb3, 0x93, 0x63, 0xc2, 0xb7, 0x3a,
		0x66, 0xb8, 0x5c, 0x29, 0xe8, 0x1d, 0x06, 0x9e, 0xa3, 0xff, 0x8d, 0x28,
		0xbf, 0x8b, 0x1d, 0x13, 0x54, 0x48, 0xcc, 0xe9, 0x12, 0x7a, 0x25, 0xa8,
		0x25, 0x33, 0xef, 0x75, 0x8f, 0x22, 0xdd, 0xde, 0xf3, 0xf3, 0x01, 0x2a,
		0xac, 0x45, 0xaf, 0xe8, 0x70, 0xad, 0x83, 0x81, 0x95, 0x74, 0x14, 0x90,
		0x5b, 0xcd, 0x07, 0x8b, 0x25, 0x76, 0xb4, 0xee, 0x19, 0x45, 0xd9, 0xcc,
		0x4d, 0xcb, 0x2e, 0x7a, 0xab, 0x21, 0xf0, 0xee, 0x59, 0x82, 0xa3, 0xfc,
		0x67, 0xd0, 0xc6, 0xc1, 0x86, 0x43, 0x70, 0x1a, 0x44, 0x59, 0xa2, 0x7b,
		0xeb, 0x40, 0xd8, 0xd7, 0xc0, 0x7a, 0xe1, 0x43, 0x76, 0xe1, 0x9c, 0x11,
		0xfd, 0xf1, 0xa1, 0xb8, 0x74, 0x41, 0x58, 0x44, 0x83, 0x91, 0x25, 0xcf,
		0x65, 0x60, 0x12, 0x36, 0xa3, 0x92, 0x75, 0x8d, 0xd6, 0x8b, 0x9b, 0xc4,
		0x4c, 0x1a, 0xdd, 0x9a, 0x7e, 0xb5, 0x53, 0x81, 0xff, 0x63, 0x2d, 0x8a,
		0x19, 0x72, 0x19, 0x62, 0x24, 0x3b, 0x58, 0x0d, 0x57, 0xf3, 0x3f, 0xbb,
		0x0c, 0x16, 0x1e, 0x16, 0xf5, 0x0b, 0x1b, 0xcb, 0x9e, 0x57, 0xa9, 0x64,
		0x4d, 0x4c, 0xc8, 0x1c, 0x93, 0xf4, 0xe0, 0x99, 0xe2, 0xb1, 0xb8, 0xdf,
		0xb5, 0x27, 0xab, 0xed, 0x1d, 0xc7, 0x1e, 0xc7, 0xd3, 0xb7, 0x30, 0xad,
		0xef, 0xa0, 0x7b, 0x9b, 0xcc, 0xbb, 0x9b, 0x5a, 0x34, 0x3a, 0x5d, 0x38,
		0xef, 0x74, 0x8b, 0xdb, 0xed, 0x55, 0x40, 0x8c, 0xfe, 0xfd, 0x79, 0xcc,
		0x23, 0x72, 0xc4, 0x06, 0xc7, 0x15, 0x1a, 0x07, 0x49, 0x74, 0xfb, 0xa5,
		0xb8, 0x7c, 0xbc, 0x38, 0x2b, 0xee, 0x6f, 0xae, 0x1f, 0xaf, 0xbe, 0x15,
		0xc5, 0x4d, 0x71, 0xfa, 0x01, 0x98, 0xf1, 0xd7, 0x05, 0x26, 0xac, 0xe9,
		0xf3, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x03, 0xa3, 0x2c, 0xc7, 0x72,
		0x05, 0x00, 0x00,
		},
		"assets/templates/help.html",
	)
}

func assets_templates_import2_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x94, 0x92,
		0xc1, 0x6a, 0xc3, 0x30, 0x10, 0x44, 0xcf, 0xf1, 0x57, 0x2c, 0xba, 0xc7,
		0x4a, 0x0f, 0xbd, 0x04, 0x45, 0x10, 0x1a, 0x0a, 0x0e, 0xb4, 0xf4, 0x17,
		0x9c, 0x58, 0x8e, 0x54, 0x6c, 0xcb, 0xac, 0xd7, 0x34, 0xc5, 0xe8, 0xdf,
		0xbb, 0xb2, 0x70, 0x29, 0xed, 0x21, 0xcd, 0x69, 0xd0, 0x58, 0xf3, 0x56,
		0x78, 0x47, 0xd9, 0x07, 0x5d, 0xb4, 0xbd, 0x47, 0x82, 0x77, 0x7f, 0x82,
		0x69, 0x82, 0x9c, 0xd5, 0x55, 0x10, 0x82, 0x92, 0xfc, 0x2d, 0x53, 0x16,
		0x75, 0x96, 0xb1, 0xef, 0x6a, 0xc8, 0x8f, 0x7c, 0x25, 0x84, 0x6c, 0xa5,
		0xa8, 0x3c, 0x35, 0x06, 0xce, 0x4d, 0x39, 0x0c, 0x3b, 0x31, 0x1f, 0x84,
		0xce, 0x56, 0xec, 0x63, 0x14, 0x56, 0xbb, 0x50, 0x5d, 0x47, 0x5e, 0x49,
		0x3e, 0x27, 0xbf, 0x82, 0x81, 0x3e, 0x1b, 0xb3, 0x13, 0x1f, 0xae, 0x22,
		0xbb, 0x85, 0x72, 0x24, 0x2f, 0x74, 0x1c, 0xcb, 0xec, 0xfc, 0xd9, 0x63,
		0x9b, 0xa7, 0x60, 0xc1, 0xb9, 0xf9, 0x11, 0x54, 0xcd, 0x64, 0x99, 0xd0,
		0x3f, 0x27, 0x3c, 0xf9, 0xb6, 0x6f, 0x0c, 0x99, 0x7f, 0xe3, 0x97, 0xc0,
		0x2d, 0xee, 0xdb, 0xfe, 0x15, 0x5a, 0x87, 0xe8, 0xf1, 0xbe, 0x97, 0xc7,
		0xe0, 0xcb, 0x9c, 0xbb, 0x31, 0x61, 0x3f, 0x92, 0xe5, 0x4b, 0xc5, 0xe1,
		0x3e, 0x7e, 0x8a, 0x15, 0x87, 0xbf, 0x74, 0x96, 0xb8, 0x04, 0xde, 0xd4,
		0x4a, 0xb9, 0x1a, 0xcb, 0xd6, 0x2c, 0x34, 0x6b, 0xdc, 0xc5, 0xd2, 0x16,
		0x1e, 0x37, 0x9b, 0xfe, 0x2a, 0x96, 0x95, 0xd5, 0x8c, 0x5b, 0x9f, 0x7d,
		0x47, 0xe8, 0x1b, 0x01, 0x03, 0x9e, 0x77, 0x42, 0xba, 0xf9, 0xbf, 0xcb,
		0x69, 0x4a, 0x0d, 0x08, 0x41, 0x0e, 0x84, 0xa6, 0x6c, 0x85, 0x56, 0x32,
		0x31, 0x75, 0xec, 0x81, 0x69, 0x06, 0x93, 0x4a, 0xd0, 0x7f, 0x17, 0xc0,
		0x5c, 0x69, 0x5d, 0x95, 0xdd, 0xc5, 0xa0, 0xd0, 0xc7, 0x5f, 0x35, 0x82,
		0xce, 0x13, 0xd4, 0x7e, 0xec, 0x2a, 0x25, 0xfb, 0x44, 0xe8, 0xa2, 0xff,
		0x15, 0x00, 0x00, 0xff, 0xff, 0x04, 0xe0, 0x1e, 0xbe, 0x7a, 0x02, 0x00,
		0x00,
		},
		"assets/templates/import2.html",
	)
}

func assets_templates_notfound_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x56,
		0x4d, 0x6f, 0xdc, 0x36, 0x10, 0x3d, 0x27, 0xbf, 0x62, 0xc2, 0x5e, 0x43,
		0x09, 0xdb, 0x5e, 0x7a, 0x90, 0x04, 0x38, 0x41, 0x81, 0xe6, 0xd0, 0xc0,
		0x80, 0x53, 0xa0, 0x57, 0x4a, 0x9c, 0x95, 0xb8, 0xa1, 0x48, 0x86, 0xa4,
		0xbc, 0x16, 0x02, 0xff, 0xf7, 0x0e, 0xa9, 0xd5, 0xee, 0xfa, 0xab, 0x6b,
		0x1f, 0xba, 0x80, 0xa1, 0xa1, 0x34, 0xf3, 0x1e, 0xf9, 0x1e, 0xc9, 0x71,
		0xf5, 0x41, 0xda, 0x2e, 0xce, 0x0e, 0x61, 0x88, 0xa3, 0x6e, 0xde, 0x57,
		0xe9, 0x01, 0x5a, 0x98, 0xbe, 0x66, 0x68, 0x58, 0xf3, 0xfe, 0x5d, 0x35,
		0xa0, 0x90, 0xf4, 0x7c, 0x57, 0x8d, 0x18, 0x05, 0x74, 0x83, 0xf0, 0x01,
		0x63, 0xcd, 0xa6, 0xb8, 0xe5, 0xbf, 0x53, 0x02, 0xd0, 0xef, 0xf0, 0x6d,
		0x88, 0xd1, 0x71, 0xfc, 0x31, 0xa9, 0xdb, 0x9a, 0xfd, 0xc3, 0xff, 0xbe,
		0xe2, 0x9f, 0xed, 0xe8, 0x44, 0x54, 0xad, 0x46, 0x06, 0x9d, 0x35, 0x11,
		0x0d, 0x15, 0x7e, 0xf9, 0xa3, 0x46, 0xd9, 0xe3, 0xc3, 0x52, 0x23, 0x46,
		0xac, 0xd9, 0xad, 0xc2, 0xbd, 0xb3, 0x3e, 0x9e, 0x65, 0xef, 0x95, 0x8c,
		0x43, 0x2d, 0xf1, 0x56, 0x75, 0xc8, 0xf3, 0xe0, 0x23, 0x28, 0xa3, 0xa2,
		0x12, 0x9a, 0x87, 0x4e, 0x68, 0xac, 0x37, 0x84, 0x94, 0xa6, 0x17, 0x55,
		0xd4, 0xd8, 0xdc, 0x8c, 0xc2, 0xc7, 0xeb, 0xab, 0xaf, 0xc0, 0xe1, 0x5a,
		0xf4, 0x08, 0xc6, 0x46, 0xd8, 0xda, 0xc9, 0xc8, 0xaa, 0x5c, 0x12, 0x72,
		0xae, 0x56, 0xe6, 0x3b, 0x78, 0xd4, 0x35, 0x0b, 0x71, 0xd6, 0x18, 0x06,
		0x44, 0x22, 0x1d, 0x3c, 0x6e, 0x6b, 0x56, 0x76, 0x21, 0x94, 0x12, 0xb7,
		0x62, 0xd2, 0xb1, 0xa0, 0x98, 0x35, 0x17, 0x2b, 0xca, 0x4e, 0x9a, 0x5d,
		0x28, 0x3a, 0x6d, 0x27, 0xb9, 0xd5, 0xc2, 0x63, 0xd1, 0xd9, 0xb1, 0x14,
		0x3b, 0x71, 0x57, 0x6a, 0xd5, 0x86, 0x32, 0xee, 0x55, 0x8c, 0xe8, 0x79,
		0x6b, 0x6d, 0x0c, 0xd1, 0x0b, 0x57, 0xfe, 0x56, 0x6c, 0x8a, 0x4d, 0x66,
		0x3a, 0xbe, 0x5b, 0xb9, 0xaa, 0xf2, 0xa0, 0x78, 0xd5, 0x5a, 0x39, 0x67,
		0x72, 0x23, 0x6e, 0xa1, 0xd3, 0x22, 0x84, 0x9a, 0x51, 0xd8, 0x0a, 0x0f,
		0xcb, 0x83, 0x1f, 0xa6, 0xb9, 0x0e, 0xb7, 0xea, 0x0e, 0x25, 0x8f, 0xd6,
		0x31, 0xf0, 0x96, 0xa4, 0x49, 0xd9, 0xaa, 0x27, 0x07, 0x6c, 0xb6, 0x92,
		0x90, 0xa4, 0x3a, 0x22, 0x25, 0x89, 0x85, 0x32, 0x34, 0xad, 0xad, 0x9e,
		0x94, 0x5c, 0x12, 0x00, 0xaa, 0x0f, 0x9c, 0xc3, 0x27, 0x2f, 0x8c, 0x84,
		0xf4, 0x17, 0x6d, 0xdf, 0x6b, 0x84, 0x1e, 0x23, 0xf4, 0xde, 0x4e, 0x0e,
		0x25, 0xc9, 0xe9, 0xa1, 0xc5, 0xb4, 0x20, 0x18, 0x6d, 0xab, 0xe8, 0xab,
		0x54, 0xc1, 0x69, 0x31, 0x03, 0xe7, 0x2b, 0xc8, 0x19, 0xcf, 0x61, 0x6e,
		0x69, 0x51, 0xe8, 0x57, 0x16, 0x4a, 0x69, 0xa7, 0x18, 0xad, 0x81, 0xb4,
		0xf9, 0x6a, 0xb6, 0x0c, 0xd8, 0xa3, 0x9a, 0x85, 0x9c, 0x81, 0x14, 0x51,
		0x1c, 0x06, 0x69, 0xde, 0x5a, 0x0b, 0x17, 0x8e, 0xaf, 0x85, 0xef, 0xd3,
		0x7e, 0xfc, 0xa5, 0x0d, 0x1c, 0xef, 0xc4, 0xe8, 0x34, 0xf2, 0x43, 0xf9,
		0x9a, 0xc9, 0x37, 0x27, 0x5a, 0x22, 0x0e, 0x4e, 0x98, 0x95, 0x28, 0x78,
		0x6e, 0x8d, 0x9e, 0x59, 0xf3, 0x6d, 0x59, 0xe7, 0x49, 0xb0, 0xaa, 0x4c,
		0x79, 0x2f, 0xd5, 0x29, 0x52, 0x8f, 0x13, 0x07, 0x6b, 0xfe, 0xe7, 0xbc,
		0xaa, 0x5c, 0xa4, 0x39, 0xbd, 0x10, 0x8f, 0x44, 0x6a, 0x93, 0x57, 0xc7,
		0xad, 0xc8, 0x8e, 0x27, 0xa0, 0x2a, 0xc5, 0xea, 0x46, 0x49, 0x76, 0xe4,
		0x8d, 0xbf, 0xfa, 0xfb, 0x99, 0x94, 0xc1, 0x2e, 0x42, 0x1c, 0xf2, 0x92,
		0x21, 0xed, 0xee, 0xf0, 0x31, 0x39, 0x3b, 0xd2, 0x23, 0xf9, 0x6e, 0xe9,
		0x93, 0x5f, 0x8f, 0x61, 0xb6, 0x3c, 0xeb, 0xaf, 0x4c, 0xff, 0xbc, 0xcb,
		0xab, 0xd6, 0xf0, 0x48, 0x7b, 0x06, 0x4a, 0x92, 0xbd, 0xaf, 0xf2, 0xa6,
		0x9a, 0xf4, 0xd9, 0xda, 0x56, 0x24, 0x7a, 0x3c, 0xb0, 0x4f, 0xab, 0x86,
		0x34, 0x38, 0x2e, 0xf7, 0x81, 0x8c, 0xbd, 0x9e, 0xdd, 0x90, 0xb4, 0x84,
		0x63, 0xc4, 0x03, 0x0a, 0xdf, 0x0d, 0x47, 0x61, 0xe1, 0x26, 0x8f, 0x93,
		0x3a, 0x15, 0x9d, 0xce, 0x17, 0x91, 0x5b, 0x6f, 0xf7, 0x34, 0xfd, 0xcb,
		0xf8, 0x5a, 0x85, 0x78, 0x42, 0xff, 0x94, 0xcb, 0x2e, 0xa2, 0xab, 0x31,
		0xdf, 0x73, 0x97, 0xd1, 0xa5, 0xdd, 0x1b, 0x6d, 0x85, 0x3c, 0x31, 0x7c,
		0xc9, 0xa5, 0x4f, 0x18, 0xaa, 0x72, 0xd2, 0xab, 0xc7, 0xff, 0x21, 0xe5,
		0x1a, 0x7a, 0xd5, 0x0f, 0xf1, 0xa9, 0xae, 0xcd, 0xcf, 0x9f, 0x50, 0x84,
		0x69, 0xa4, 0x3d, 0x34, 0x17, 0x37, 0x76, 0xf2, 0x1d, 0x06, 0xb8, 0xbf,
		0x87, 0xb0, 0x84, 0x27, 0x56, 0x78, 0xf2, 0x7b, 0x06, 0xe0, 0x6a, 0x8a,
		0x83, 0xf5, 0x19, 0x40, 0x2c, 0xe1, 0x1b, 0x01, 0xfe, 0xb2, 0x72, 0xd2,
		0xcb, 0x0c, 0xc6, 0x25, 0x7c, 0x23, 0xc0, 0xb5, 0xe8, 0xbe, 0x53, 0x1b,
		0xc8, 0x08, 0xee, 0x10, 0x5f, 0x86, 0x58, 0x5d, 0x1a, 0x50, 0xbb, 0x57,
		0x78, 0xf4, 0x63, 0xc2, 0x90, 0xae, 0x0c, 0x1e, 0x54, 0x6f, 0x4e, 0x46,
		0xfd, 0x49, 0xd5, 0x2f, 0x73, 0x2d, 0x7e, 0x9d, 0x9f, 0xcf, 0x74, 0x3f,
		0xaf, 0x51, 0x55, 0x92, 0x4f, 0x4b, 0xaf, 0x7a, 0xee, 0xca, 0x3e, 0xdc,
		0xe6, 0xc3, 0xa6, 0xf9, 0x7a, 0xea, 0x6f, 0x34, 0xca, 0x6f, 0x5d, 0xf3,
		0x8d, 0x8e, 0xb5, 0x4b, 0xdd, 0x6f, 0xb6, 0x13, 0x35, 0xad, 0x3c, 0x3f,
		0xba, 0xb8, 0x3b, 0x3b, 0x69, 0x99, 0x3b, 0x62, 0x8b, 0x6b, 0x91, 0x5b,
		0xd8, 0xd6, 0x0b, 0xa2, 0x0a, 0x9d, 0x57, 0x2e, 0x42, 0xf0, 0xdd, 0x2b,
		0xfa, 0xda, 0x8e, 0x90, 0xfd, 0x5c, 0xfe, 0x9a, 0x9b, 0xd9, 0x32, 0x28,
		0x46, 0x65, 0x8a, 0x5d, 0xc8, 0x2a, 0x64, 0xa8, 0xe6, 0xcd, 0xa8, 0x2f,
		0x75, 0xcb, 0xdd, 0x79, 0xb3, 0x7c, 0x44, 0x41, 0xd7, 0x64, 0x6e, 0x96,
		0xa4, 0x42, 0xfe, 0x3f, 0xe6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e,
		0x63, 0x48, 0x23, 0xd8, 0x08, 0x00, 0x00,
		},
		"assets/templates/notfound.html",
	)
}

func assets_templates_browse_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xa4, 0x53,
		0xc1, 0x8e, 0xd3, 0x30, 0x10, 0x3d, 0xb7, 0x5f, 0x61, 0x85, 0x0b, 0x1c,
		0x92, 0xec, 0x0a, 0xed, 0x81, 0x90, 0xcd, 0x01, 0xa1, 0x3d, 0x21, 0x84,
		0x90, 0xe0, 0xee, 0xc6, 0x93, 0xc4, 0x62, 0x62, 0x47, 0xb6, 0x53, 0x28,
		0x55, 0xff, 0x9d, 0xb1, 0xe3, 0xb4, 0xc9, 0x6e, 0x0b, 0x07, 0x4e, 0x76,
		0xc6, 0xf3, 0xde, 0x3c, 0xbf, 0xe7, 0x94, 0xd6, 0x1d, 0x10, 0xaa, 0xed,
		0x26, 0x1b, 0xb8, 0xeb, 0x06, 0x6e, 0x5c, 0xa1, 0xb4, 0x7b, 0x5d, 0x20,
		0xb7, 0x2e, 0xad, 0x3b, 0x89, 0xe2, 0x4d, 0xc1, 0x1b, 0x07, 0x86, 0x1d,
		0xb7, 0x9b, 0x4d, 0xad, 0x95, 0x03, 0xe5, 0x0a, 0x96, 0xe4, 0xc9, 0x7b,
		0xfa, 0xee, 0xb9, 0x69, 0xa5, 0x4a, 0x8d, 0x6c, 0x3b, 0x2a, 0xde, 0xdf,
		0x0d, 0xbf, 0x16, 0x55, 0x84, 0xe6, 0x52, 0x3c, 0xc5, 0x09, 0x91, 0x06,
		0xb5, 0x29, 0xd8, 0x2b, 0xce, 0xb9, 0xef, 0x6f, 0x88, 0x35, 0xb5, 0xf2,
		0x37, 0x14, 0xec, 0x2e, 0x7b, 0x07, 0xfd, 0xb9, 0xbd, 0xfe, 0xc1, 0x5b,
		0x08, 0x88, 0x48, 0xb9, 0xd3, 0xce, 0xe9, 0xde, 0xb7, 0x3d, 0x84, 0xb6,
		0xcd, 0xc0, 0x85, 0x90, 0xaa, 0xa5, 0x31, 0x0f, 0x8b, 0x31, 0x13, 0x2e,
		0x0e, 0xdb, 0x83, 0x71, 0xb2, 0xe6, 0x98, 0x72, 0x94, 0xad, 0x2a, 0x98,
		0xd3, 0xc3, 0x02, 0xf9, 0x4c, 0xbb, 0xc7, 0x0b, 0x69, 0x50, 0x5a, 0xc7,
		0x46, 0x0c, 0x78, 0xbf, 0x4f, 0x83, 0x4b, 0xa9, 0x3b, 0x0c, 0xa4, 0x51,
		0x69, 0x05, 0x4b, 0x86, 0x78, 0xcf, 0x87, 0x6b, 0x04, 0x28, 0xaf, 0xc9,
		0x7f, 0x1b, 0x5b, 0xcb, 0x3c, 0xba, 0x5f, 0x76, 0xf7, 0xd5, 0xf1, 0xc8,
		0x0c, 0x57, 0xa4, 0x3b, 0xfb, 0x42, 0x3e, 0x7d, 0x90, 0xce, 0xb2, 0xd3,
		0xa9, 0xb4, 0x03, 0x57, 0xac, 0xa6, 0x34, 0xec, 0x63, 0x32, 0x27, 0x94,
		0x54, 0x25, 0x67, 0x9d, 0x81, 0xe6, 0x31, 0xc9, 0x77, 0x46, 0xff, 0xb4,
		0x40, 0xd8, 0x80, 0x22, 0x44, 0xe2, 0x89, 0xb2, 0xcf, 0xbc, 0x07, 0x0f,
		0xcf, 0x79, 0x45, 0x43, 0x88, 0xc3, 0x57, 0x41, 0x89, 0x50, 0xa3, 0x61,
		0xdb, 0x2d, 0x7d, 0xcb, 0xc6, 0xa3, 0x26, 0xb3, 0x4e, 0xa4, 0x46, 0xc8,
		0xfd, 0x65, 0x54, 0x28, 0x27, 0xf4, 0x30, 0x4a, 0xc7, 0x77, 0xe1, 0x85,
		0xd0, 0xce, 0xf8, 0x85, 0xd6, 0xae, 0x8a, 0xb8, 0x32, 0xa7, 0xfd, 0x54,
		0x13, 0xd5, 0xa4, 0x22, 0xd4, 0xb3, 0xa7, 0x11, 0xf1, 0x2c, 0x82, 0xce,
		0x3c, 0x3c, 0x9f, 0xf0, 0x4b, 0x9a, 0xef, 0x60, 0xac, 0xd4, 0xea, 0x26,
		0x4d, 0x3c, 0xbf, 0xc2, 0xb2, 0xd6, 0x3f, 0xaf, 0xde, 0xb3, 0xe7, 0x4a,
		0x8d, 0xde, 0x4b, 0x01, 0x82, 0xed, 0x0e, 0xab, 0x31, 0x7e, 0xdd, 0x2c,
		0x4d, 0xbf, 0x46, 0xe4, 0x7b, 0x17, 0xb6, 0xf4, 0x5a, 0x8c, 0x08, 0x2f,
		0xfd, 0xcf, 0x57, 0x9a, 0xa5, 0x71, 0x23, 0xc7, 0x6f, 0x5f, 0x3f, 0x85,
		0x38, 0x56, 0x11, 0xb6, 0x78, 0x18, 0x3a, 0x49, 0xbf, 0x11, 0x3b, 0xef,
		0xd2, 0x5a, 0xf7, 0x83, 0x01, 0x6b, 0x41, 0x24, 0x73, 0x58, 0xec, 0x26,
		0xdf, 0x94, 0x28, 0x49, 0x3a, 0xeb, 0x9f, 0x52, 0x0d, 0xb7, 0x7a, 0x69,
		0xd1, 0x7c, 0x48, 0xa5, 0x29, 0xc5, 0x88, 0xbd, 0x1c, 0xad, 0x62, 0x8f,
		0x2f, 0x37, 0xc4, 0x3e, 0x62, 0xe4, 0x88, 0x06, 0x7d, 0x94, 0x66, 0x76,
		0x17, 0xe5, 0xc5, 0x81, 0xd5, 0xd3, 0xfb, 0xe7, 0x5d, 0x1b, 0x8d, 0x02,
		0x4c, 0x5a, 0xa3, 0xb6, 0xb0, 0xbe, 0xed, 0xea, 0xc5, 0xd2, 0x80, 0xb5,
		0xfe, 0x85, 0x8e, 0x27, 0x89, 0xf0, 0xff, 0x42, 0x6e, 0x99, 0xfe, 0x77,
		0x19, 0x65, 0xee, 0x5d, 0x99, 0x3c, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff,
		0x8d, 0xa4, 0x7c, 0x62, 0x3a, 0x05, 0x00, 0x00,
		},
		"assets/templates/browse.html",
	)
}

func assets_templates_error_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x56,
		0x41, 0x6f, 0xe3, 0x36, 0x13, 0x3d, 0x67, 0x7f, 0xc5, 0xac, 0xbe, 0xeb,
		0x52, 0x82, 0xf3, 0x5d, 0x7a, 0x90, 0x04, 0x64, 0x17, 0x41, 0x9b, 0x43,
		0x8b, 0x00, 0x69, 0x81, 0x5e, 0x29, 0x71, 0x2c, 0x31, 0x4b, 0x91, 0x5a,
		0x72, 0x14, 0xc7, 0x58, 0xe4, 0xbf, 0x77, 0x48, 0x59, 0xb6, 0xe3, 0x24,
		0x75, 0x72, 0xa8, 0x81, 0x80, 0x43, 0x71, 0xe6, 0xcd, 0xf0, 0x3d, 0x92,
		0x93, 0xf2, 0xb3, 0x72, 0x2d, 0x6d, 0x47, 0x84, 0x9e, 0x06, 0x53, 0x7f,
		0x2a, 0xe3, 0x00, 0x46, 0xda, 0xae, 0xca, 0xd0, 0x66, 0xf5, 0xa7, 0x8b,
		0xb2, 0x47, 0xa9, 0x78, 0xbc, 0x28, 0x07, 0x24, 0x09, 0x6d, 0x2f, 0x7d,
		0x40, 0xaa, 0xb2, 0x89, 0xd6, 0xe2, 0x17, 0x76, 0x00, 0xfe, 0xed, 0xd6,
		0x7a, 0xa2, 0x51, 0xe0, 0x8f, 0x49, 0x3f, 0x54, 0xd9, 0xdf, 0xe2, 0xaf,
		0x2b, 0xf1, 0xcd, 0x0d, 0xa3, 0x24, 0xdd, 0x18, 0xcc, 0xa0, 0x75, 0x96,
		0xd0, 0x72, 0xe0, 0xcd, 0x75, 0x85, 0xaa, 0xc3, 0xe7, 0xa1, 0x56, 0x0e,
		0x58, 0x65, 0x0f, 0x1a, 0x37, 0xa3, 0xf3, 0x74, 0xe4, 0xbd, 0xd1, 0x8a,
		0xfa, 0x4a, 0xe1, 0x83, 0x6e, 0x51, 0xa4, 0xc9, 0x17, 0xd0, 0x56, 0x93,
		0x96, 0x46, 0x84, 0x56, 0x1a, 0xac, 0x56, 0x8c, 0x14, 0xcb, 0x23, 0x4d,
		0x06, 0xeb, 0xbb, 0x41, 0x7a, 0xba, 0xbd, 0xfa, 0x03, 0x04, 0xdc, 0xca,
		0x0e, 0xc1, 0x3a, 0x82, 0xb5, 0x9b, 0xac, 0x2a, 0x8b, 0xd9, 0x21, 0xf9,
		0x1a, 0x6d, 0xbf, 0x83, 0x47, 0x53, 0x65, 0x81, 0xb6, 0x06, 0x43, 0x8f,
		0xc8, 0x49, 0x7b, 0x8f, 0xeb, 0x2a, 0x2b, 0xda, 0x10, 0x0a, 0x85, 0x6b,
		0x39, 0x19, 0xca, 0xd9, 0xce, 0xea, 0xb3, 0x11, 0x45, 0xab, 0xec, 0x7d,
		0xc8, 0x5b, 0xe3, 0x26, 0xb5, 0x36, 0xd2, 0x63, 0xde, 0xba, 0xa1, 0x90,
		0xf7, 0xf2, 0xb1, 0x30, 0xba, 0x09, 0x05, 0x6d, 0x34, 0x11, 0x7a, 0xd1,
		0x38, 0x47, 0x81, 0xbc, 0x1c, 0x8b, 0xff, 0xe7, 0xab, 0x7c, 0x95, 0x32,
		0xed, 0xbf, 0x2d, 0xb9, 0xca, 0x62, 0xc7, 0x78, 0xd9, 0x38, 0xb5, 0x4d,
		0xc9, 0xad, 0x7c, 0x80, 0xd6, 0xc8, 0x10, 0xaa, 0x8c, 0xcd, 0x46, 0x7a,
		0x98, 0x07, 0xb1, 0x2b, 0x73, 0x99, 0xae, 0xf5, 0x23, 0x2a, 0x41, 0x6e,
		0xcc, 0xc0, 0x3b, 0xa6, 0x26, 0x7a, 0xeb, 0x8e, 0x15, 0x70, 0x49, 0x4a,
		0x46, 0x52, 0x7a, 0x8f, 0x14, 0x29, 0x96, 0xda, 0x72, 0x59, 0x6b, 0x33,
		0x69, 0x35, 0x3b, 0x00, 0x94, 0x9f, 0x85, 0x80, 0xaf, 0x5e, 0x5a, 0x05,
		0xf1, 0x8f, 0x5c, 0xd7, 0x19, 0x84, 0x0e, 0x09, 0x3a, 0xef, 0xa6, 0x11,
		0x15, 0xd3, 0xe9, 0xa1, 0xc1, 0xb8, 0x21, 0x18, 0x5c, 0xa3, 0x79, 0x55,
		0xe9, 0x30, 0x1a, 0xb9, 0x05, 0x21, 0x16, 0x90, 0xa3, 0x3c, 0xbb, 0xda,
		0xe2, 0xa6, 0xd0, 0x2f, 0x59, 0xd8, 0xa5, 0x99, 0x88, 0x9c, 0x85, 0x78,
		0xf8, 0xaa, 0x6c, 0x9e, 0x64, 0x27, 0x31, 0x73, 0xf2, 0x0c, 0x94, 0x24,
		0xb9, 0x9b, 0xc4, 0xba, 0x8d, 0x91, 0x63, 0xd8, 0x7f, 0x96, 0xbe, 0x8b,
		0xe7, 0xf1, 0x7f, 0x4d, 0x10, 0xf8, 0x28, 0x87, 0xd1, 0xa0, 0xd8, 0x85,
		0x2f, 0x9e, 0x62, 0x75, 0x48, 0xcb, 0x89, 0xc3, 0x28, 0xed, 0x92, 0x28,
		0x78, 0xe1, 0xac, 0xd9, 0x66, 0xf5, 0x9f, 0xf3, 0x3e, 0x0f, 0x84, 0x95,
		0x45, 0xf4, 0x7b, 0x2b, 0x4e, 0x33, 0x7b, 0x82, 0x73, 0x64, 0xf5, 0x7f,
		0xec, 0x57, 0x16, 0x33, 0x35, 0x87, 0x0f, 0xf2, 0x84, 0xa4, 0x26, 0x6a,
		0xb5, 0x3f, 0x8a, 0xd9, 0xfe, 0x06, 0x94, 0x85, 0x5c, 0xd4, 0x28, 0x58,
		0x8e, 0x74, 0xf0, 0x17, 0x7d, 0xbf, 0x31, 0x33, 0xd8, 0x12, 0x50, 0x9f,
		0xb6, 0x0c, 0xf1, 0x74, 0x87, 0x2f, 0x51, 0xd9, 0x81, 0x87, 0xa8, 0xbb,
		0xe3, 0x25, 0xbf, 0x5c, 0xc3, 0x24, 0x79, 0xe2, 0x5f, 0xdb, 0xee, 0x75,
		0x95, 0x17, 0xae, 0xe1, 0x84, 0xfb, 0x0c, 0xb4, 0x62, 0x79, 0xdf, 0xa5,
		0x4d, 0x39, 0x99, 0xa3, 0xbd, 0x2d, 0x48, 0x3c, 0x3c, 0x93, 0xcf, 0xe8,
		0x9a, 0x39, 0xd8, 0x6f, 0xf7, 0x19, 0x8d, 0x9d, 0xd9, 0x8e, 0x7d, 0xe4,
		0x12, 0xf6, 0x96, 0x08, 0x28, 0x7d, 0xdb, 0xef, 0x89, 0x85, 0xbb, 0x34,
		0x8f, 0xec, 0x94, 0x7c, 0x3b, 0xdf, 0x44, 0x6e, 0xbc, 0xdb, 0x70, 0xf9,
		0xe7, 0xf1, 0x8d, 0x0e, 0x74, 0x40, 0xff, 0x9a, 0xc2, 0xce, 0xa2, 0xeb,
		0x21, 0xbd, 0x73, 0xe7, 0xd1, 0x95, 0xdb, 0x58, 0xe3, 0xa4, 0x3a, 0x64,
		0xb8, 0x49, 0xa1, 0x2f, 0x32, 0x94, 0xc5, 0x64, 0x16, 0x8d, 0xff, 0x85,
		0xca, 0xc5, 0xf4, 0xba, 0xeb, 0xe9, 0x25, 0xaf, 0xf5, 0xcf, 0x9f, 0x90,
		0x87, 0x69, 0xe0, 0x33, 0xb4, 0xcd, 0xef, 0xdc, 0xe4, 0x5b, 0x0c, 0xf0,
		0xf4, 0x04, 0x61, 0x36, 0x0f, 0x59, 0xe1, 0xc5, 0xef, 0x15, 0x80, 0xab,
		0x89, 0x7a, 0xe7, 0x13, 0x80, 0x9c, 0xcd, 0x0f, 0x02, 0xfc, 0xee, 0xd4,
		0x64, 0xe6, 0x0a, 0x86, 0xd9, 0xfc, 0x20, 0xc0, 0xad, 0x6c, 0xbf, 0x73,
		0x1b, 0x48, 0x08, 0xe3, 0xce, 0x3e, 0x0f, 0xb1, 0xa8, 0xd4, 0xa3, 0x19,
		0xdf, 0xa1, 0xd1, 0x8f, 0x09, 0x43, 0x7c, 0x32, 0x44, 0xd0, 0x9d, 0x3d,
		0x08, 0xf5, 0x1b, 0x47, 0xbf, 0x9d, 0x6b, 0xd6, 0xeb, 0xf8, 0x7e, 0xc6,
		0xf7, 0x79, 0xb1, 0xca, 0x82, 0x75, 0x9a, 0x7b, 0xd5, 0x6b, 0x4f, 0xf6,
		0xee, 0x35, 0xef, 0x57, 0xbc, 0xd9, 0xfc, 0x57, 0x47, 0xdc, 0x95, 0xf3,
		0x6b, 0xef, 0x9d, 0x7f, 0x7a, 0xe2, 0xf6, 0xb1, 0x9a, 0x57, 0xf9, 0x19,
		0xd0, 0x6d, 0xec, 0x78, 0x17, 0xc9, 0xf7, 0xb2, 0xbe, 0x7e, 0x6c, 0x71,
		0x8c, 0x85, 0xc2, 0x49, 0xd4, 0x8d, 0x4a, 0x71, 0x97, 0x3b, 0xd7, 0xd1,
		0xe3, 0x11, 0xee, 0x1d, 0x31, 0x6f, 0x71, 0x3d, 0x7e, 0x9e, 0x8b, 0x3c,
		0x42, 0x3e, 0xbc, 0x2d, 0x65, 0x68, 0xbd, 0x1e, 0x09, 0x82, 0x6f, 0xdf,
		0xd1, 0x12, 0xef, 0x99, 0x34, 0xbf, 0x2d, 0x2e, 0x53, 0x1f, 0x9c, 0x27,
		0xf9, 0xa0, 0x6d, 0x7e, 0x1f, 0x12, 0x81, 0x09, 0xaa, 0xfe, 0x30, 0xea,
		0x5b, 0x8d, 0xf6, 0xfe, 0xb8, 0xcf, 0x9e, 0xa4, 0xe0, 0x17, 0x36, 0xf5,
		0x59, 0x26, 0x20, 0xfd, 0x0b, 0xf4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
		0x55, 0xa4, 0xd0, 0xd9, 0x13, 0x09, 0x00, 0x00,
		},
		"assets/templates/error.html",
	)
}

func assets_templates_layout_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x56,
		0xc1, 0x6e, 0xe3, 0x36, 0x10, 0x3d, 0x6f, 0xbe, 0x62, 0x96, 0xbd, 0x2e,
		0x25, 0xa4, 0xbd, 0xf4, 0x20, 0x09, 0xc8, 0x06, 0x05, 0xba, 0x87, 0x16,
		0x01, 0xb2, 0x05, 0x7a, 0xa5, 0xc4, 0xb1, 0x44, 0x2f, 0x45, 0x2a, 0xe4,
		0xc8, 0x8e, 0x61, 0xf8, 0xdf, 0x4b, 0x8a, 0x96, 0xed, 0x78, 0x93, 0xd8,
		0x39, 0xd4, 0x40, 0xc0, 0xa1, 0x34, 0x9c, 0xf7, 0xf8, 0x1e, 0xc5, 0x49,
		0xf1, 0x59, 0xda, 0x86, 0x36, 0x03, 0x42, 0x47, 0xbd, 0xae, 0x6e, 0x8a,
		0x38, 0x80, 0x16, 0xa6, 0x2d, 0x19, 0x1a, 0x56, 0xdd, 0x7c, 0x2a, 0x3a,
		0x14, 0x32, 0x8c, 0x9f, 0x8a, 0x1e, 0x49, 0x40, 0xd3, 0x09, 0xe7, 0x91,
		0x4a, 0x36, 0xd2, 0x82, 0xff, 0x1e, 0x12, 0x20, 0xfc, 0xf6, 0xef, 0x3a,
		0xa2, 0x81, 0xe3, 0xd3, 0xa8, 0x56, 0x25, 0xfb, 0x97, 0xff, 0x73, 0xc7,
		0xef, 0x6d, 0x3f, 0x08, 0x52, 0xb5, 0x46, 0x06, 0x8d, 0x35, 0x84, 0x26,
		0x2c, 0xfc, 0xf6, 0x47, 0x89, 0xb2, 0xc5, 0x97, 0x4b, 0x8d, 0xe8, 0xb1,
		0x64, 0x2b, 0x85, 0xeb, 0xc1, 0x3a, 0x3a, 0xc9, 0x5e, 0x2b, 0x49, 0x5d,
		0x29, 0x71, 0xa5, 0x1a, 0xe4, 0xd3, 0xe4, 0x0b, 0x28, 0xa3, 0x48, 0x09,
		0xcd, 0x7d, 0x23, 0x34, 0x96, 0xb7, 0xa1, 0x52, 0xa4, 0x47, 0x8a, 0x34,
		0x56, 0xdb, 0x6d, 0xf6, 0x3d, 0x06, 0xbb, 0x5d, 0x91, 0xa7, 0x27, 0xd3,
		0x4b, 0xad, 0xcc, 0x0f, 0x70, 0xa8, 0x4b, 0xe6, 0x69, 0xa3, 0xd1, 0x77,
		0x88, 0x01, 0xa5, 0x73, 0xb8, 0x28, 0x59, 0xde, 0x78, 0x9f, 0x4b, 0x5c,
		0x88, 0x51, 0x53, 0x16, 0x62, 0x56, 0x5d, 0x5c, 0x91, 0x37, 0xd2, 0x2c,
		0x7d, 0xd6, 0x68, 0x3b, 0xca, 0x85, 0x16, 0x0e, 0xb3, 0xc6, 0xf6, 0xb9,
		0x58, 0x8a, 0xe7, 0x5c, 0xab, 0xda, 0xe7, 0xb4, 0x56, 0x44, 0xe8, 0x78,
		0x6d, 0x2d, 0x79, 0x72, 0x62, 0xc8, 0x7f, 0xcb, 0x6e, 0xb3, 0xdb, 0x09,
		0xe9, 0xf0, 0x6c, 0xc6, 0x2a, 0xf2, 0xbd, 0xc4, 0x45, 0x6d, 0xe5, 0x66,
		0x02, 0x37, 0x62, 0x05, 0x8d, 0x16, 0xde, 0x97, 0x2c, 0x84, 0xb5, 0x70,
		0x90, 0x06, 0xbe, 0xa7, 0x39, 0x4f, 0x17, 0xea, 0x19, 0x25, 0x27, 0x3b,
		0x30, 0x70, 0x36, 0x68, 0x11, 0xb3, 0x55, 0x1b, 0x24, 0xb7, 0x93, 0x77,
		0xa1, 0x92, 0x54, 0x87, 0x4a, 0x51, 0x53, 0xa1, 0x4c, 0xa0, 0xb5, 0xd0,
		0xa3, 0x92, 0x29, 0x01, 0xa0, 0xf8, 0xcc, 0x39, 0x7c, 0x75, 0xc2, 0x48,
		0x88, 0x7f, 0x64, 0xdb, 0x56, 0x23, 0xb4, 0x48, 0xd0, 0x3a, 0x3b, 0x0e,
		0x28, 0x61, 0x61, 0x1d, 0xd4, 0x18, 0x37, 0x04, 0xbd, 0xad, 0x55, 0x78,
		0x2b, 0x95, 0x1f, 0xb4, 0xd8, 0x00, 0xe7, 0x73, 0x91, 0x13, 0x9c, 0x3d,
		0xb7, 0xb8, 0x29, 0x74, 0x33, 0x4a, 0x48, 0xa9, 0x47, 0x22, 0x6b, 0x20,
		0x9e, 0xb6, 0x92, 0xa5, 0x09, 0x3b, 0x5b, 0x93, 0xc0, 0x19, 0x48, 0x41,
		0x62, 0x3f, 0x89, 0xbc, 0xb5, 0x16, 0x83, 0x3f, 0x3c, 0x16, 0xae, 0x8d,
		0x07, 0xf0, 0x97, 0xda, 0x73, 0x7c, 0x16, 0xfd, 0xa0, 0x91, 0xef, 0x97,
		0xcf, 0x99, 0xfc, 0xf6, 0x08, 0x1b, 0x80, 0xfd, 0x20, 0xcc, 0x0c, 0xe4,
		0x1d, 0xb7, 0x46, 0x6f, 0x58, 0xf5, 0x3d, 0xed, 0xf3, 0x28, 0x58, 0x91,
		0xc7, 0xbc, 0xb7, 0xd6, 0xa9, 0xa0, 0x1e, 0x0f, 0x18, 0xac, 0xfa, 0x9f,
		0xf3, 0x8a, 0x3c, 0x49, 0x73, 0x7c, 0x20, 0xce, 0x44, 0xaa, 0xa3, 0x57,
		0x87, 0xa3, 0xc8, 0xaa, 0xc7, 0x5e, 0x38, 0x7a, 0xb8, 0xfb, 0xbb, 0xc8,
		0xc5, 0xec, 0x46, 0x1e, 0xec, 0x98, 0x0e, 0xfe, 0xec, 0xef, 0x7d, 0x50,
		0x06, 0x1b, 0x02, 0xea, 0xa6, 0x2d, 0x43, 0x3c, 0xdd, 0xfe, 0x4b, 0x74,
		0xb6, 0x0f, 0x43, 0xf4, 0xdd, 0x86, 0x57, 0x6e, 0xfe, 0xee, 0x26, 0xcb,
		0x27, 0xfd, 0x95, 0x69, 0x5f, 0x77, 0x79, 0xd6, 0x1a, 0xce, 0xb4, 0x67,
		0xa0, 0x64, 0xb0, 0xf7, 0x2a, 0x6f, 0x8a, 0x51, 0x9f, 0xec, 0x6d, 0xae,
		0x14, 0x86, 0x17, 0xf6, 0x69, 0xb5, 0xdd, 0x82, 0x5a, 0x00, 0x3e, 0x41,
		0xf6, 0x20, 0x5a, 0x04, 0xf6, 0x88, 0xc2, 0x35, 0x1d, 0x83, 0xdd, 0x6e,
		0x5e, 0x2e, 0x1a, 0x52, 0x2b, 0x64, 0xdb, 0x2d, 0x1a, 0xb9, 0xdb, 0x55,
		0x41, 0xb3, 0x83, 0x3c, 0x2f, 0x64, 0x6f, 0xf5, 0x66, 0xe8, 0xa2, 0xf6,
		0x70, 0x88, 0xb8, 0x4f, 0xd5, 0x66, 0x23, 0x20, 0x55, 0x8f, 0x6a, 0x16,
		0xe1, 0x6b, 0x7e, 0x9f, 0xc9, 0x57, 0x67, 0xd7, 0x71, 0xcf, 0x57, 0x30,
		0xa9, 0x53, 0xea, 0x65, 0x3e, 0x5a, 0x79, 0x3a, 0xb2, 0x49, 0x08, 0xd7,
		0xb1, 0xf9, 0xd6, 0xa7, 0xdb, 0xf3, 0x0a, 0x36, 0x2a, 0xa5, 0x5e, 0x66,
		0x23, 0xed, 0xda, 0x68, 0x2b, 0xe4, 0x91, 0x51, 0x42, 0xf9, 0x89, 0x51,
		0x91, 0x8f, 0x7a, 0x3e, 0x73, 0xef, 0x58, 0x3b, 0x87, 0x4e, 0xb5, 0x1d,
		0x9d, 0xf9, 0x1c, 0xf8, 0x85, 0x0b, 0x1c, 0x32, 0x3f, 0xf6, 0xe1, 0x4c,
		0x6f, 0xb2, 0x47, 0x3b, 0xba, 0x06, 0x7d, 0xdc, 0x8f, 0x4f, 0xe1, 0x11,
		0x15, 0x7e, 0xfa, 0xbd, 0x52, 0xe0, 0x6e, 0xa4, 0xce, 0xba, 0xa9, 0x80,
		0x48, 0xe1, 0x07, 0x0b, 0xfc, 0x65, 0xe5, 0xa8, 0x13, 0x83, 0x3e, 0x85,
		0x1f, 0x2c, 0xf0, 0x20, 0x9a, 0x1f, 0xc1, 0x9b, 0xa9, 0xc2, 0xb0, 0x8f,
		0x2f, 0x94, 0x38, 0x77, 0xf5, 0x4f, 0xd4, 0xc3, 0x55, 0x9e, 0x76, 0x31,
		0xf1, 0xb2, 0xa3, 0x4f, 0x23, 0xfa, 0x78, 0xe1, 0x71, 0xaf, 0x5a, 0x73,
		0xb4, 0x35, 0xc2, 0xbc, 0xcd, 0x2c, 0xb9, 0x7b, 0x7a, 0xbb, 0xc4, 0xee,
		0x32, 0x47, 0x45, 0x1e, 0x5c, 0x4d, 0x9d, 0xf6, 0xb5, 0x86, 0x93, 0x6c,
		0x0e, 0xad, 0xf9, 0x3e, 0x5d, 0x2f, 0xbb, 0xdd, 0xcd, 0x61, 0x71, 0x8c,
		0x7c, 0xe3, 0xd4, 0x40, 0xe0, 0x5d, 0x73, 0x45, 0x6f, 0x5d, 0x06, 0xfe,
		0x6e, 0x93, 0xff, 0x3a, 0x35, 0xd4, 0x34, 0xc9, 0x7a, 0x65, 0xb2, 0xa5,
		0x9f, 0xf6, 0x32, 0x95, 0xaa, 0x3e, 0x5c, 0xf5, 0xad, 0x8e, 0xbd, 0x3c,
		0x6d, 0xd8, 0x67, 0x10, 0xe1, 0xaa, 0x9e, 0x1a, 0x76, 0x68, 0xe0, 0xd3,
		0x3f, 0x4f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x33, 0x00, 0xa4, 0xb9,
		0x4d, 0x09, 0x00, 0x00,
		},
		"assets/templates/layout.html",
	)
}

func assets_templates_import_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x56,
		0xc1, 0x6e, 0xdb, 0x38, 0x10, 0x3d, 0xdb, 0x5f, 0x31, 0xd0, 0x9e, 0x25,
		0xd9, 0xc8, 0xee, 0x1e, 0xbc, 0xb2, 0x80, 0x6c, 0xbc, 0xc1, 0x1a, 0x68,
		0x82, 0xa0, 0x40, 0x0e, 0x3d, 0x15, 0xb4, 0x44, 0x4b, 0x44, 0x29, 0x92,
		0x20, 0xa9, 0xc4, 0x6e, 0x90, 0x7f, 0xef, 0x90, 0x94, 0x1a, 0x59, 0xb5,
		0x93, 0x34, 0xc9, 0x21, 0x88, 0x44, 0x3e, 0xce, 0xbc, 0xf7, 0x38, 0x33,
		0x56, 0x56, 0xcf, 0xf3, 0x75, 0xa3, 0xa4, 0xb6, 0x70, 0x71, 0x73, 0x7e,
		0x0d, 0x8d, 0x2c, 0x5b, 0x4e, 0xb3, 0x14, 0x97, 0xa7, 0x59, 0xad, 0xf3,
		0xe9, 0x34, 0xdb, 0x4a, 0xdd, 0x80, 0x96, 0x9c, 0x2e, 0x23, 0xf7, 0x18,
		0x41, 0x43, 0x6d, 0x2d, 0xcb, 0x65, 0xa4, 0xa4, 0xb1, 0x11, 0x90, 0xc2,
		0x32, 0x29, 0x96, 0x51, 0x04, 0x54, 0x14, 0x76, 0xaf, 0x10, 0xd6, 0xb4,
		0xdc, 0x32, 0x45, 0xb4, 0x4d, 0xdd, 0x81, 0xb8, 0x24, 0x96, 0x44, 0xf9,
		0x74, 0xf2, 0xf0, 0x00, 0x6c, 0x0b, 0xc9, 0xb6, 0x4e, 0xfe, 0x27, 0xe6,
		0x3f, 0xad, 0xa5, 0x36, 0xf0, 0xf8, 0x38, 0x9d, 0x64, 0x0a, 0x8c, 0xdd,
		0xbb, 0xf8, 0x1b, 0xa9, 0x4b, 0xaa, 0x17, 0x30, 0x57, 0x3b, 0x30, 0x92,
		0xb3, 0x12, 0xfe, 0xd8, 0xce, 0x66, 0xff, 0x40, 0x21, 0xb9, 0xc4, 0x65,
		0xf7, 0xe2, 0x02, 0x4d, 0x32, 0x63, 0xb5, 0x14, 0x55, 0xee, 0x83, 0x64,
		0x69, 0xf7, 0x86, 0x91, 0x52, 0x15, 0xf2, 0x50, 0x51, 0xba, 0xd0, 0xb8,
		0x52, 0xb2, 0x3b, 0x7f, 0xa2, 0x3e, 0xeb, 0x75, 0x32, 0x61, 0x25, 0x0a,
		0x3c, 0x73, 0xcb, 0x08, 0xd5, 0x44, 0x54, 0xd4, 0xb3, 0x0a, 0x94, 0x92,
		0x00, 0x5b, 0x23, 0xca, 0xb3, 0x9b, 0x0c, 0xf8, 0x1d, 0xf0, 0xc0, 0xb3,
		0x09, 0x22, 0x42, 0xce, 0x41, 0xd2, 0x89, 0x4f, 0x0a, 0x05, 0x27, 0xc6,
		0x2c, 0x23, 0xce, 0x8c, 0x8d, 0x2b, 0x2d, 0x5b, 0x15, 0x1d, 0x26, 0x64,
		0xa2, 0xa4, 0x3b, 0x6a, 0xfa, 0x1c, 0x9c, 0x6c, 0x28, 0xff, 0xf5, 0x50,
		0xcc, 0x2c, 0x6d, 0xa2, 0x3c, 0x63, 0x42, 0xb5, 0x16, 0x82, 0xbd, 0x9a,
		0x94, 0x4c, 0x46, 0x20, 0x48, 0x83, 0x2f, 0xcc, 0xb3, 0x75, 0x9a, 0x22,
		0xb8, 0x23, 0xbc, 0xc5, 0x25, 0x47, 0xec, 0x1a, 0x37, 0x31, 0x74, 0xd4,
		0x33, 0x6f, 0x88, 0xae, 0x98, 0x88, 0x35, 0xab, 0x6a, 0x8b, 0xfe, 0xce,
		0xd4, 0x2e, 0x08, 0xe8, 0x70, 0x59, 0xea, 0xf3, 0x8f, 0x85, 0x7c, 0x28,
		0x2d, 0x41, 0xef, 0xbf, 0x7a, 0xd1, 0xcf, 0x92, 0xba, 0xa6, 0xf7, 0xe0,
		0x51, 0x0b, 0x70, 0x04, 0x8e, 0x5d, 0x10, 0x42, 0xd6, 0x0e, 0xd1, 0x71,
		0x7c, 0xdd, 0xfd, 0x8c, 0x74, 0x0d, 0x89, 0x5b, 0xba, 0xc3, 0x3a, 0xee,
		0x34, 0xfa, 0x8a, 0x2d, 0xa4, 0xc0, 0x92, 0xe2, 0xbd, 0x18, 0x64, 0xde,
		0x11, 0x57, 0x9c, 0x14, 0xb4, 0x96, 0x1c, 0xab, 0x74, 0x19, 0xd1, 0xa4,
		0x4a, 0x60, 0x45, 0xf4, 0x37, 0x6c, 0x9c, 0x03, 0xf3, 0x91, 0xea, 0x95,
		0x2c, 0x29, 0x1f, 0x32, 0x1d, 0x8b, 0xb6, 0x52, 0x2d, 0xe0, 0x2f, 0xa7,
		0x78, 0xe0, 0x7d, 0x96, 0x86, 0x72, 0xed, 0xfe, 0x4f, 0x0f, 0x2a, 0x49,
		0xcb, 0xfb, 0x71, 0x90, 0x8d, 0xb4, 0x56, 0x36, 0x68, 0x9d, 0x0f, 0x34,
		0xaa, 0x3c, 0x34, 0x23, 0x6e, 0xca, 0xf8, 0x6f, 0xbf, 0x81, 0x0d, 0xf0,
		0x67, 0x1e, 0x3a, 0x9c, 0x85, 0xa6, 0xc1, 0xf7, 0xdf, 0x77, 0xa2, 0x50,
		0x44, 0x84, 0x00, 0xc7, 0xbc, 0xa8, 0xad, 0x55, 0x8b, 0x34, 0x35, 0xb2,
		0xa1, 0x49, 0x8f, 0x3a, 0x66, 0x8b, 0xe3, 0x71, 0xe5, 0xf7, 0x9d, 0x31,
		0x4e, 0x67, 0x68, 0xd0, 0x1b, 0x62, 0x2c, 0x05, 0x02, 0x2e, 0xcb, 0x96,
		0xf9, 0x31, 0x74, 0x96, 0x9f, 0xaa, 0x82, 0x8b, 0x0e, 0xf4, 0x9e, 0x2a,
		0x70, 0x7a, 0x89, 0xa6, 0x64, 0xa0, 0xce, 0x85, 0x3c, 0x61, 0x01, 0x5e,
		0x00, 0xae, 0xcd, 0xbb, 0xa0, 0x4f, 0x6a, 0x9e, 0x88, 0x64, 0x69, 0x1f,
		0x71, 0x78, 0x9b, 0x2f, 0xdc, 0xca, 0x79, 0x8b, 0xe3, 0x54, 0xc3, 0x7a,
		0xf5, 0xc6, 0x3b, 0x21, 0xfe, 0x3c, 0x2b, 0x8f, 0x7b, 0x1d, 0xa2, 0xaf,
		0x57, 0x87, 0x4e, 0x77, 0xa3, 0x70, 0xab, 0x65, 0x03, 0xb7, 0x9f, 0x3f,
		0x3d, 0x6b, 0x74, 0xc0, 0x22, 0xea, 0xfd, 0xfd, 0x36, 0x9c, 0x0d, 0xad,
		0x46, 0x05, 0x2f, 0x4a, 0x3c, 0x26, 0x69, 0x48, 0x68, 0x54, 0x86, 0x5d,
		0x05, 0x2a, 0x62, 0xeb, 0xc4, 0xca, 0xf4, 0x8b, 0x6c, 0x75, 0x7c, 0xe5,
		0x7f, 0xd2, 0xe2, 0x59, 0x32, 0x9b, 0x27, 0x78, 0x37, 0x49, 0xf5, 0xfd,
		0x84, 0x11, 0x5c, 0x16, 0x84, 0x43, 0xc9, 0x34, 0x2d, 0xac, 0xd4, 0xfb,
		0x67, 0x4d, 0xb9, 0xc4, 0x03, 0x2b, 0xa6, 0x3f, 0xca, 0x12, 0x97, 0x1f,
		0x13, 0xbf, 0xd1, 0x90, 0x27, 0x32, 0x23, 0x3b, 0xd2, 0x1a, 0x5b, 0x31,
		0x6d, 0x0d, 0xd5, 0xa9, 0xab, 0xed, 0x57, 0xf8, 0x71, 0xab, 0xb8, 0x24,
		0x25, 0xf6, 0x60, 0xf8, 0x10, 0x00, 0x2c, 0xcc, 0x57, 0x75, 0xa3, 0xa3,
		0x70, 0xf9, 0xce, 0x6e, 0x1c, 0x1b, 0x12, 0x5a, 0x31, 0x38, 0x12, 0x9e,
		0x4f, 0x89, 0xef, 0x32, 0x1f, 0xb7, 0xec, 0xc5, 0xc1, 0xea, 0xdc, 0x8e,
		0x0b, 0x2a, 0x2c, 0xd5, 0x01, 0xbc, 0x69, 0x71, 0xac, 0x8a, 0x7e, 0x7b,
		0x63, 0x05, 0xe0, 0x5f, 0x5c, 0xd2, 0x2d, 0xc1, 0x0f, 0x9b, 0x9e, 0x91,
		0x69, 0x37, 0x0d, 0xb3, 0x38, 0xbe, 0x0d, 0xfa, 0xd3, 0x63, 0x2b, 0xbe,
		0x57, 0x35, 0xc3, 0xcc, 0xf0, 0xf3, 0x29, 0xc6, 0x1b, 0xd9, 0xc7, 0x05,
		0xd3, 0x05, 0x2a, 0xc0, 0x61, 0xef, 0xe0, 0x39, 0xfc, 0x4b, 0x71, 0x7e,
		0x43, 0xe8, 0x84, 0x2c, 0x0d, 0x09, 0x07, 0xfc, 0x32, 0xff, 0xe5, 0x94,
		0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0x50, 0x70, 0x2d, 0xc8, 0x9e, 0x09,
		0x00, 0x00,
		},
		"assets/templates/import.html",
	)
}

func assets_templates_layout_streamstart_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x54, 0x8f,
		0xcd, 0x4a, 0xc5, 0x30, 0x10, 0x85, 0xd7, 0xf7, 0x3e, 0x45, 0xcc, 0xda,
		0x20, 0xee, 0x5c, 0x24, 0x05, 0x11, 0x17, 0x3e, 0x80, 0xe0, 0x76, 0x6e,
		0x72, 0x6c, 0x06, 0xf2, 0x67, 0x3b, 0x6d, 0xe9, 0xdb, 0x1b, 0x4b, 0x17,
		0xde, 0xd9, 0x1c, 0x86, 0x39, 0x1f, 0x7c, 0x63, 0x1f, 0x42, 0xf5, 0xb2,
		0x37, 0xa8, 0x28, 0x39, 0x0d, 0x57, 0xfb, 0x17, 0x2a, 0x51, 0x19, 0x9d,
		0x46, 0xd1, 0xc3, 0xf5, 0x62, 0x23, 0x28, 0xf4, 0xbc, 0xd8, 0x0c, 0x21,
		0xe5, 0x23, 0x4d, 0x33, 0xc4, 0xe9, 0x45, 0xbe, 0xcd, 0x4b, 0x2f, 0xa8,
		0x3e, 0xe7, 0x2d, 0x8a, 0x34, 0x83, 0x9f, 0x85, 0x57, 0xa7, 0xbf, 0xcc,
		0xe7, 0xab, 0x79, 0xab, 0xb9, 0x91, 0xf0, 0x2d, 0x41, 0x2b, 0x5f, 0x8b,
		0xa0, 0x74, 0xf0, 0xe3, 0xdd, 0x21, 0x8c, 0xb8, 0x47, 0x0b, 0x65, 0x38,
		0xbd, 0x32, 0xb6, 0x56, 0x27, 0xf9, 0xd7, 0xde, 0x38, 0x48, 0x74, 0x01,
		0x2b, 0x7b, 0x98, 0x63, 0x79, 0x54, 0x5c, 0x58, 0x98, 0x92, 0x99, 0x3d,
		0x25, 0xb8, 0xe7, 0xc3, 0xf2, 0xe9, 0xd4, 0xb4, 0xb7, 0x1a, 0xf6, 0xfe,
		0x47, 0x9b, 0x30, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x08, 0x83,
		0x71, 0xdd, 0x00, 0x00, 0x00,
		},
		"assets/templates/layout_streamstart.html",
	)
}

func assets_templates_search_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x84, 0x53,
		0xc1, 0x8e, 0x9b, 0x30, 0x10, 0x3d, 0x27, 0x5f, 0x31, 0x72, 0xcf, 0x84,
		0x8d, 0xaa, 0x55, 0xd5, 0x2c, 0x20, 0xed, 0xa9, 0x97, 0x2a, 0x6a, 0x37,
		0xea, 0x07, 0x18, 0x18, 0xc0, 0x92, 0xb1, 0x5d, 0xdb, 0x6c, 0x37, 0x8d,
		0xf6, 0xdf, 0x3b, 0x06, 0x07, 0x48, 0xa4, 0x6a, 0x0f, 0x48, 0xe6, 0xcd,
		0x9b, 0xe7, 0x37, 0x6f, 0x20, 0x73, 0xfe, 0x2c, 0xb1, 0xd8, 0x6e, 0x76,
		0x0e, 0xb9, 0xad, 0xba, 0xa4, 0x43, 0x5e, 0xa3, 0x85, 0x6e, 0x0f, 0x97,
		0xed, 0x66, 0xd3, 0x73, 0xdb, 0x0a, 0x95, 0x78, 0x6d, 0x0e, 0xf0, 0xe5,
		0xc1, 0xbc, 0x3d, 0x11, 0x56, 0x69, 0xa9, 0xed, 0x01, 0x3e, 0x71, 0xce,
		0x9f, 0x16, 0x4a, 0xa9, 0xbd, 0xd7, 0xfd, 0x01, 0x1e, 0x23, 0xab, 0xd1,
		0xca, 0x27, 0x4e, 0xfc, 0x45, 0x82, 0xb0, 0x27, 0xe4, 0x9d, 0xee, 0x30,
		0xdc, 0x77, 0xa3, 0xee, 0x9d, 0xc6, 0x8a, 0xfc, 0xb0, 0xfb, 0x3a, 0xd3,
		0x2d, 0xba, 0x41, 0x7a, 0x07, 0x83, 0x1c, 0x9b, 0xa4, 0x70, 0xc4, 0x0a,
		0x7e, 0x13, 0x7f, 0x36, 0xc4, 0x55, 0x5a, 0x61, 0x68, 0x37, 0xbc, 0xae,
		0x85, 0x6a, 0x13, 0x89, 0x8d, 0x3f, 0xc0, 0xfe, 0x71, 0x74, 0x70, 0x27,
		0x20, 0xc5, 0x7a, 0xa0, 0xab, 0xdb, 0xcf, 0x91, 0x9a, 0xa5, 0x31, 0x87,
		0xcb, 0x05, 0x44, 0x43, 0xc2, 0x1e, 0x76, 0x3f, 0x07, 0xb4, 0x67, 0x78,
		0xa7, 0x62, 0x2d, 0x5e, 0xa1, 0x92, 0xdc, 0xb9, 0x9c, 0x79, 0x7c, 0xf3,
		0x49, 0x85, 0xca, 0x53, 0x46, 0x37, 0x89, 0x31, 0xca, 0x30, 0xeb, 0xf6,
		0xc5, 0x89, 0x6e, 0xf0, 0x3f, 0x9e, 0x8f, 0x59, 0x4a, 0x2f, 0xa4, 0x4b,
		0xbd, 0xa3, 0x2a, 0xaa, 0x3a, 0x68, 0xfd, 0x4f, 0x6c, 0x6c, 0x6f, 0xb4,
		0xed, 0xc1, 0x6a, 0x89, 0x39, 0x0b, 0x47, 0x06, 0xbc, 0xf2, 0x42, 0xab,
		0x9c, 0x31, 0xe8, 0xd1, 0x77, 0xba, 0xce, 0x59, 0x8b, 0x3e, 0x50, 0x01,
		0xd6, 0x3a, 0x81, 0x9c, 0xb4, 0x56, 0x0f, 0x66, 0xaa, 0x51, 0x55, 0xf2,
		0x12, 0x25, 0x50, 0x21, 0x67, 0x93, 0x4d, 0x76, 0x25, 0x77, 0xa2, 0xae,
		0x51, 0xb1, 0xe2, 0x34, 0xc2, 0x59, 0x3a, 0x32, 0xaf, 0x6d, 0x42, 0x99,
		0xc1, 0x43, 0x48, 0x77, 0x72, 0xc7, 0x6e, 0xae, 0xa8, 0x68, 0x4d, 0x64,
		0x0f, 0x46, 0x56, 0x22, 0x5b, 0x06, 0xa2, 0x5e, 0xf4, 0x8d, 0xe4, 0x15,
		0x76, 0x5a, 0x52, 0x18, 0x39, 0x3b, 0x45, 0x70, 0xca, 0x73, 0xce, 0xf2,
		0x95, 0xcb, 0x81, 0xa4, 0x09, 0x9d, 0x21, 0x36, 0x87, 0x03, 0x7c, 0xf0,
		0xba, 0xd1, 0xd5, 0xe0, 0x40, 0xf1, 0x9e, 0x68, 0xbf, 0xe3, 0xa8, 0x53,
		0x88, 0xe1, 0x54, 0x0e, 0xb4, 0x37, 0x15, 0x0d, 0xba, 0xa1, 0xec, 0xc5,
		0x62, 0xb1, 0xf4, 0x0a, 0xe8, 0x49, 0x6a, 0x6c, 0x38, 0x6d, 0x9d, 0x15,
		0x99, 0x33, 0x5c, 0x5d, 0xab, 0xad, 0x3c, 0x9b, 0x4e, 0xd0, 0x04, 0x30,
		0x9f, 0x92, 0xe8, 0xbc, 0xa0, 0xed, 0x13, 0xb3, 0x80, 0xc9, 0x34, 0x2c,
		0x2b, 0x9c, 0xae, 0x0b, 0xab, 0x49, 0x43, 0x00, 0xf3, 0x3e, 0xb7, 0x77,
		0x63, 0x6d, 0x37, 0x11, 0x78, 0x89, 0x1f, 0x5c, 0x80, 0xd6, 0x1b, 0x8a,
		0x1f, 0x62, 0x98, 0x67, 0x93, 0x99, 0xe2, 0xb9, 0xd4, 0x14, 0x33, 0xf5,
		0x48, 0x54, 0xeb, 0x26, 0x88, 0xbc, 0x2c, 0x35, 0x23, 0x73, 0x08, 0x8b,
		0x09, 0xda, 0x96, 0xab, 0x16, 0x6f, 0xe5, 0xa9, 0x2c, 0x45, 0x28, 0xd3,
		0x81, 0x43, 0x67, 0xb1, 0xc9, 0x59, 0x5a, 0x5a, 0xfd, 0xc7, 0x61, 0x1a,
		0xf2, 0xfd, 0xf5, 0xf2, 0x3d, 0xa4, 0xfb, 0x71, 0x0a, 0x81, 0xfc, 0x2d,
		0xbc, 0x4d, 0xf4, 0x98, 0x45, 0x40, 0x8f, 0xb4, 0x05, 0x02, 0xb3, 0x94,
		0xc7, 0x7b, 0x4a, 0x0b, 0x69, 0x3c, 0xae, 0x55, 0xc3, 0x6f, 0xcd, 0x8a,
		0xe5, 0xd2, 0xa8, 0x31, 0x7a, 0x4c, 0x27, 0x93, 0xcb, 0x1f, 0x10, 0xb0,
		0x71, 0xae, 0xeb, 0x5a, 0x43, 0x49, 0x3a, 0x8c, 0x35, 0x53, 0x1c, 0xf5,
		0x6d, 0x0c, 0x4b, 0xeb, 0x7c, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0x3b,
		0x56, 0x44, 0x3c, 0xb6, 0x04, 0x00, 0x00,
		},
		"assets/templates/search.html",
	)
}

func assets_templates_layout_streamend_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb2, 0xd1,
		0x2f, 0x28, 0x4a, 0xb5, 0xe3, 0xe2, 0xb4, 0xd1, 0x4f, 0xca, 0x4f, 0xa9,
		0xb4, 0xe3, 0xb2, 0xd1, 0xcf, 0x28, 0xc9, 0xcd, 0xb1, 0xe3, 0x02, 0x04,
		0x00, 0x00, 0xff, 0xff, 0xe6, 0x83, 0x20, 0xb5, 0x18, 0x00, 0x00, 0x00,
		},
		"assets/templates/layout_streamend.html",
	)
}

func assets_css_default_css() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4a, 0xca,
		0x4f, 0xa9, 0x54, 0xa8, 0xe6, 0xe2, 0x2c, 0x48, 0x4c, 0x49, 0xc9, 0xcc,
		0x4b, 0xd7, 0x2d, 0xc9, 0x2f, 0xb0, 0x52, 0x30, 0x37, 0x28, 0xa8, 0xb0,
		0xe6, 0xaa, 0xe5, 0x02, 0x04, 0x00, 0x00, 0xff, 0xff, 0x56, 0xf0, 0x95,
		0xff, 0x1d, 0x00, 0x00, 0x00,
		},
		"assets/css/default.css",
	)
}


// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"assets/templates/help.html": assets_templates_help_html,
	"assets/templates/import2.html": assets_templates_import2_html,
	"assets/templates/notfound.html": assets_templates_notfound_html,
	"assets/templates/browse.html": assets_templates_browse_html,
	"assets/templates/error.html": assets_templates_error_html,
	"assets/templates/layout.html": assets_templates_layout_html,
	"assets/templates/import.html": assets_templates_import_html,
	"assets/templates/layout_streamstart.html": assets_templates_layout_streamstart_html,
	"assets/templates/search.html": assets_templates_search_html,
	"assets/templates/layout_streamend.html": assets_templates_layout_streamend_html,
	"assets/css/default.css": assets_css_default_css,

}
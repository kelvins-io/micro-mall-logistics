package proto

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

var _proto_micro_mall_logistics_proto_logistics_business_logistics_business_swagger_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4f\x6f\xdc\x44\x14\xbf\xef\xa7\x18\x19\x8e\x55\x12\xc2\x2d\x37\xd7\x35\xcd\x2a\xc9\x6e\xb1\x77\x25\x24\x14\x59\x13\xfb\xed\x66\x2a\x7b\xc6\x99\x3f\x0b\x11\x8a\x44\x6f\xc0\x1d\x2e\x3d\x20\x21\x71\xe0\x02\x3d\x00\x07\x10\x1f\x27\x69\x3f\x06\x1a\xaf\xed\xb5\xbd\xe3\x8d\xe3\x4d\xd5\x46\xaa\xa5\x48\xf1\xcc\xfb\x3f\xef\xf7\xde\x1b\xef\x37\x03\x84\x2c\xf1\x15\x9e\xcf\x81\x5b\x07\xc8\xda\xdf\xd9\xb3\x1e\xe9\x35\x42\x67\xcc\x3a\x40\x7a\x1f\x21\x4b\x12\x19\x83\xde\x4f\x39\x93\x6c\x37\x21\x21\x67\x41\x82\xe3\x38\x88\xd9\x9c\x08\x49\x42\x11\x2c\xb7\x56\xef\x67\x4a\x10\x0a\x42\x18\x96\x76\x32\xda\x4c\x11\x42\xd6\x02\xb8\x20\x8c\x6a\xf1\xf9\xbf\x88\x32\x89\x04\x48\x6b\x80\xd0\x55\x66\x4e\xc8\xa8\x50\x09\x08\xeb\x00\x7d\xb9\xe4\xc2\x69\x1a\x93\x10\x4b\xc2\xe8\xee\x73\xc1\xa8\xa6\x3d\xcd\x68\x53\xce\x22\x15\x76\xa4\xc5\xf2\x5c\xac\xfc\xdc\x5d\x7c\xb2\x32\x77\x37\xe4\x80\x25\x94\xbb\x9a\x9c\x09\x59\x79\xd7\xb1\x53\x49\x82\xf9\xa5\x36\xfe\xfa\xbb\x97\xd7\xff\xfe\xf3\xfa\xfb\xdf\x6e\xfe\x7a\xf1\xe6\xf7\x57\xd7\xff\xfd\x98\x7b\x98\x11\xb2\x14\x78\x66\xc1\x30\xd2\xc4\xc7\x85\x96\xc7\x79\x4c\x7c\xe0\x0b\x12\x42\xe0\x64\x4a\x3d\x08\x19\x8f\xaa\xfc\x1c\x44\xca\xa8\x00\x51\xd3\x8f\x90\xb5\xbf\xb7\xd7\x58\x42\xc8\x8a\x40\x84\x9c\xa4\x32\x8f\xab\x8d\x84\x0a\x43\x10\x62\xa6\x62\x54\x48\xda\xa9\x88\x5f\xfa\x12\x9e\x43\x82\xd7\x84\x21\x64\x7d\xcc\x61\xa6\xe5\x7c\xb4\x1b\xc1\x8c\x50\xa2\xe5\x9a\x0e\xb6\x6a\xbc\x97\xab\xb1\x6a\xc2\xae\x2a\x6f\x57\x55\xfd\x56\x04\x33\xac\x62\x79\xbb\x2f\x14\x29\x0a\x5f\xa7\x10\x4a\x88\x10\x70\xce\x78\xe9\xd2\xb6\x1e\x71\x45\x25\x49\xc0\xd5\x42\x37\xd8\x3d\x30\x78\x60\xa5\x98\xe3\x04\x24\xf0\x55\xe2\x2d\x9f\x86\x3b\x14\x27\x19\x94\xce\x58\x74\xd9\xb4\x97\xd0\xb6\x1d\x0e\x17\x8a\x70\xd0\xb9\x23\xb9\x82\xb7\x7c\x72\x17\x0a\x84\xec\x12\x80\xd3\x4a\x00\x24\x9e\x37\x5d\x6f\x4d\xf3\x95\xec\xd3\x41\x55\x66\x1e\xd0\x06\x10\x2f\x14\x64\x10\x2b\x71\x38\x87\x76\x18\xde\xfc\xfc\xeb\x9b\x3f\x7e\xd9\x16\x86\x9f\x6b\x95\x0f\x14\x85\x15\xdb\x3f\x80\x30\x7f\x5a\x40\xb8\x0a\x5e\xc8\xa2\x35\xcb\x97\x70\x5c\x66\x5f\x3b\x1e\x67\x38\x16\x4d\x40\xca\xcb\x34\x93\x2f\x24\x27\x74\x6e\xbd\x43\xec\xa8\x34\x6a\x36\x31\xb5\x01\x3c\x2f\xff\xbc\xf9\xe9\xd5\x12\x3c\xaf\x7f\xf8\xfb\xe6\xdb\x17\x7d\xc0\x33\xcd\x74\xfa\x52\x2b\x7e\x68\xe0\xa9\xd8\xfe\x01\x3c\xf9\xf3\x20\x3a\x58\xed\xe0\xde\x5d\x03\x2b\x47\xd6\x8a\xb5\xab\x01\xd3\xd0\x78\x59\x92\x30\x5a\xa6\x5a\x05\xa7\x45\x0d\x61\x67\xcf\x21\x94\x65\x38\xf5\x7c\x9b\x02\x97\xa4\x81\x23\x2b\xab\x60\x0d\x64\x75\x0d\x9e\x07\xd2\xd1\xec\xc6\xbc\x48\xc4\xbc\x29\xb7\xad\xbe\x5d\x19\x2b\x52\xb7\x71\x63\x0b\xd7\x99\x92\x81\xe4\x38\x82\x80\xb2\xce\xa6\x3e\xaa\x86\x4e\x71\x92\xdd\x80\xfa\xb2\x06\x39\xad\x91\x9f\x50\x09\xfa\x82\x55\xab\x16\x33\xc6\x13\x2c\xf3\xed\x4f\xf7\x5b\xa4\x2b\x21\x59\xb2\x6e\x59\xe7\xb1\x2e\xe7\x1f\xea\xeb\x9c\x51\xc3\x9c\xb1\x68\xad\x20\x17\x86\x63\xce\x71\x1d\xc8\x16\x91\x90\x34\xe9\xef\x60\xd0\x53\xad\xae\x6e\x4d\xad\x1c\xf5\xca\x9f\xfb\xc0\x8e\x46\x61\xef\x28\xd7\x31\x6c\x8c\x73\x63\xd2\xb8\x6f\x3c\x55\xcf\x79\x8b\x38\x08\xa0\x51\xa0\x44\x3f\x28\x64\xcc\x38\x8a\xfa\x33\xa7\xe7\x8c\x76\x8f\x4d\x93\x5b\x37\xbe\x3e\xcc\x1c\x42\x20\x0b\xe8\xed\x77\xc1\xdf\xd7\xf5\x82\xff\x6e\xde\x77\xcc\x8c\x15\xe0\xb6\xa9\xaf\x3c\x02\x7e\xb7\xc4\xad\xf8\x97\xcf\x0b\xf7\xea\x56\xd9\x97\x97\x53\xe6\x16\xbe\xc5\x6c\xf9\x65\xa8\x8f\x67\x2c\xed\x9d\x34\x21\xa3\x12\x68\x73\x58\xec\xc4\x4a\x44\xc0\x16\xed\x5a\xcf\x18\x8b\x01\xd3\xb6\x6e\x53\x6c\xdf\x39\xe4\xa6\x6b\xe5\x7b\x5d\x72\x85\x4e\x8d\x20\x26\xa2\x35\xca\xf7\xdf\xdf\x1a\x79\xb9\x55\x93\x2b\x86\x32\x43\x90\xf3\xf4\x28\x83\x0c\x54\x25\xb5\xd1\xd5\xf2\xa7\x8e\xe3\xfa\x7e\xf5\xfe\xe5\x7a\xde\xd8\xab\x2e\x4c\x7d\xd7\x0b\x46\xe3\x49\xe0\x7e\x31\xf4\x27\x6b\x3b\x6b\xab\x27\xae\xe7\x1c\xda\xa3\x89\x99\xa7\xdc\x5d\xdb\xf1\x0f\xc7\xcf\xcc\x3c\xd9\xce\xfa\xea\xd1\xb4\x85\xfc\x68\x6a\x5e\xb4\x4f\xc6\xd3\xc2\xae\xd1\x78\xfa\xf4\x70\xcd\x99\xc7\xf6\xb1\x3d\x72\xdc\x16\x92\xd2\xf6\xcd\x64\xb6\xe3\x64\x7a\x8e\xc7\xce\x91\x91\xbd\x20\xd8\x1c\xa0\x82\xca\x9f\xd8\x13\x77\x4d\x58\x66\xee\x46\x41\x35\x0a\xb3\x90\x27\xae\x33\x3c\xb1\x8f\x83\x67\xb6\xe7\xbb\x81\xeb\xd5\xce\x7d\xe2\xd9\x23\xdf\x76\x26\xc3\xf1\x28\xf8\xcc\x1e\x1e\xbb\x4f\x4c\x4e\x2e\xd5\xe6\xeb\xe5\x3d\xa9\x72\xcb\x2d\x73\xec\x96\x3c\x36\xdc\xcc\xb6\xaa\xd7\xbd\x06\xa9\x66\x5d\xe8\x5b\x78\xda\x00\xde\x15\xd4\xa6\xef\x0b\xef\x5f\x15\x35\x3a\x93\xfd\x6a\x73\xa6\x66\x36\xbd\xdc\xc6\x64\x4d\x1f\x28\x1e\xf7\x39\xba\x05\x8e\xd5\x6d\x67\xde\xd6\xf7\x2e\x6f\x3f\xae\xda\xd7\x93\x2d\x5c\x84\x86\x80\xce\xfe\x6d\x4a\xe9\xfe\xb7\xc8\x04\x84\xc0\xf3\x5e\x58\x89\x40\x62\x12\xbf\xb5\x0b\x62\x35\xa5\xba\xf4\xca\x81\xfe\xbb\x1a\xfc\x1f\x00\x00\xff\xff\xa3\x6b\xff\xe8\xb7\x1c\x00\x00")

func proto_micro_mall_logistics_proto_logistics_business_logistics_business_swagger_json() ([]byte, error) {
	return bindata_read(
		_proto_micro_mall_logistics_proto_logistics_business_logistics_business_swagger_json,
		"proto/micro_mall_logistics_proto/logistics_business/logistics_business.swagger.json",
	)
}

var _proto_micro_mall_order_proto_order_business_order_business_swagger_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6f\x6b\xe4\x44\x18\x7f\xbf\x9f\x62\x88\xbe\x94\xeb\x79\x8a\x2f\xfa\xae\xee\x2d\x6d\x39\xae\x2b\xbb\x5d\x50\xa4\x84\x69\xf2\xec\x76\xce\x64\x26\x9d\x99\xd4\x2e\x52\xf0\xd4\x93\x13\xaa\x16\xce\xb3\xde\x59\x04\x45\x0e\x51\xb8\xbe\xd3\xa3\x27\xf8\x65\x9a\xbd\xde\x2b\xbf\x82\x4c\x92\x4d\x93\xd9\x64\x37\x9b\x74\x41\xc1\x85\x85\xcd\xfc\x79\xe6\xf7\xfc\xf9\x3d\xcf\x33\xd9\x8f\x1a\x08\x19\xe2\x43\x3c\x18\x00\x37\x96\x91\x71\xe3\xda\x75\xe3\x35\x35\x46\x68\x9f\x19\xcb\x48\xcd\x23\x64\x48\x22\x1d\x50\xf3\x1e\x67\x92\x2d\xb9\xc4\xe2\xcc\x74\xb1\xe3\x98\x8c\xdb\xc0\xcd\x68\x38\xfa\xbd\xed\x0b\x42\x41\x08\xed\xf1\x5a\xb8\x26\x14\x8e\x90\xb1\x07\x5c\x10\x46\x95\xc8\xf8\x27\xa2\x4c\x22\x01\xd2\x68\x20\x74\x10\x42\xb0\x18\x15\xbe\x0b\xc2\x58\x46\xef\x47\xbb\xb0\xe7\x39\xc4\xc2\x92\x30\xba\x74\x47\x30\xaa\xd6\x6e\x85\x6b\x3d\xce\x6c\xdf\x2a\xb9\x16\xcb\x1d\x71\xa9\xdb\xd2\xde\xeb\x11\xd4\x25\x8b\x03\x96\x90\xcc\xa8\xa5\x4c\xc8\xd4\xb3\xb2\x95\xef\xba\x98\x0f\x15\xf0\xe0\xfe\xf7\xc1\xf3\xb3\x8b\xa7\x3f\x05\x5f\x3e\x8c\xf5\x0a\x97\x30\x0f\x78\x78\xee\xba\xad\x96\xb5\x95\xec\xb7\x63\x2b\x74\x81\xef\x11\x0b\xcc\x66\x78\x54\x38\x95\xde\xca\x41\x78\x8c\x0a\x10\x99\x43\x11\x32\x6e\x5c\xbf\xae\x0d\x21\x64\xd8\x20\x2c\x4e\x3c\x19\x1b\x72\x05\x09\xdf\xb2\x40\x88\xbe\xef\xa0\xb1\xa4\x6b\x29\xf1\x91\x02\xd6\x0e\xb8\x78\x42\x18\x42\xc6\xab\x1c\xfa\x4a\xce\x2b\x4b\x36\xf4\x09\x25\x4a\xae\xee\xc5\x14\xee\x4e\x7c\x82\x91\x91\x73\x90\x7a\x3a\x48\x1f\x6d\xd8\xd0\xc7\xbe\x23\x67\xab\x41\x91\x4f\x61\xdf\x03\x4b\x82\x8d\x80\x73\xc6\x13\x6d\xea\x2a\xc3\x7d\x2a\x89\x0b\x2d\x25\x74\x0a\xee\x46\x8e\x06\x86\x87\x39\x76\x41\x02\xbf\x0c\xb2\xe8\xa3\xa9\x43\xb1\x1b\x52\x65\x9b\xd9\x43\x1d\x2f\xa1\x45\x33\x1c\x76\x7d\xc2\x41\x45\x8c\xe4\x3e\x2c\xce\x69\xbb\x3e\x08\x59\x46\xf7\xad\x94\xee\x12\x0f\x74\xad\x73\xe3\xfa\x52\xee\x56\x23\x2d\x2f\xb6\x63\x8a\x6b\x36\x48\x4c\x9c\x34\xd7\x06\x50\x4c\xb5\x8b\xaf\xfe\x08\xbe\xfe\x36\xa2\xda\xc5\xe9\x93\xd1\xa7\xf7\xe6\x25\xdc\x2a\xc8\x70\xfc\x66\x74\xee\x7f\x89\x73\x59\xe8\xff\xd3\x2e\xfe\x14\xd0\x4e\xee\x9b\x16\xb3\x27\x20\x47\xcc\xdb\xf5\x81\x4f\xa3\x5e\x1f\x3b\x42\xe7\x9e\x1c\x7a\xa1\x60\x21\x39\xa1\x03\x63\xf1\x5c\x49\x0a\x60\xca\x84\x97\xe5\x4a\xa3\x36\x23\x74\x33\x02\x98\x10\x49\x03\x3c\x86\x66\x00\xf5\xdd\x0c\x30\xa3\xb9\xf1\x5e\x9a\x09\xbd\xee\xcd\x31\xa8\x44\xa1\x54\xfc\x84\xcb\x33\x74\xd6\xb1\xb8\x2e\xa3\x49\x7c\xe6\x20\x62\xdb\x77\xc0\x92\x97\x88\x3c\xae\xa8\x2b\x89\xc6\x3d\x23\x74\xa0\xc6\xc6\x32\x4c\xe9\x80\x6c\xaa\xad\xb9\x81\xe4\x8a\x81\x2e\xb3\xc8\xb5\x07\xb9\xc9\x6b\x66\x52\xad\xa1\xb1\x4f\xec\x19\xe0\x32\xa4\xee\x33\xee\xe2\xd0\x27\x84\xca\xb7\xde\xcc\x57\x58\x51\xae\xb4\xc6\xa9\x7d\xd9\xc4\x30\xf7\x76\x0f\x0f\x81\x9b\x96\x43\x80\x4a\x93\x78\xd5\x10\x84\x39\x7b\xa6\x4d\x0a\x36\x6b\x95\x05\x95\x0d\x9f\x74\x81\x98\x8c\x85\xc4\x87\x49\x33\x7c\xfe\xec\xe3\xf3\x67\xbf\x06\x67\xc7\x2f\x1f\x9c\xbd\x78\xfc\x59\xf0\xf0\xf3\xe0\xc1\xdd\x17\x4f\x8f\xa3\xf1\xa8\x5a\xfd\xfd\xe7\x61\xf0\xe4\x93\xd1\x0f\x27\xc1\xd1\xe1\xe8\xf8\xf7\xe0\xf0\x5e\x70\xf4\x5b\xf0\xf3\xe3\x64\x63\x70\xff\xd1\xe8\xe4\x8b\x68\x24\xee\x25\xcb\x46\xdd\x55\x10\x4d\x51\xb6\x8a\xad\x34\xb2\xe7\x07\x60\x9c\x8a\xaf\x90\x75\x05\xb5\xf0\x5f\x6b\x02\x6c\x59\xcc\xa7\x7a\xf5\x2d\x15\xc7\x16\x23\xd4\x94\xd9\xdc\x3e\x07\xb6\xb8\x2e\xe4\x8a\x76\x88\x28\x84\x84\x39\xc7\xd9\x12\x69\x10\x09\xae\xde\x1d\x95\xc4\xd1\xdd\x61\x5e\x2e\xad\x32\xe5\xbe\x84\xe3\xd3\x42\x6a\x78\x5b\xec\x30\xcf\xcc\xcf\x10\x8b\x31\x40\x88\x5b\x59\xe1\x4a\x0c\xb0\xca\x98\x2d\x6a\xe9\xff\x81\x3f\x1f\x27\xd3\x89\x9d\xab\xae\xa5\xc2\x46\xec\x96\xe0\xc0\xfc\xf5\x2d\xee\xf8\xe6\x86\xc3\xc1\xf6\xad\xb9\xaa\x5b\x59\xf7\x74\x2d\xa0\xb0\x9e\x7e\x73\x52\xc5\x45\x92\x71\x30\x49\x56\x0a\x2a\x4d\x38\xb5\x3b\x84\x50\xa6\x82\x8d\xbe\x39\x3d\x7f\xfe\x5d\x70\x72\x36\x7a\x74\x7a\xfe\xd7\x8f\xa3\xbb\xa7\xd3\x8a\x8f\x1e\xcb\x75\x79\xb8\x80\x9e\x67\x81\x49\x73\xa0\x71\x0f\x2d\x38\x69\x44\x5c\xcf\xcf\x17\x29\x58\x42\xc5\x5c\xe5\x70\xd1\xc2\xb6\x4c\xcc\x04\x47\x87\x51\x83\xf3\xf2\xf8\x97\xe0\xec\x38\x69\x7c\xa6\x45\xce\xb8\x31\xaf\x74\x4b\xe9\xf6\x9a\xcd\x56\xb7\x9b\xbe\xa9\xb4\x3a\x9d\x76\x27\x7b\x75\x69\x75\xcc\x8d\xf6\xa6\xd9\x7a\x77\xbd\xbb\x39\x31\x33\x31\x7a\xbb\xd5\x69\xae\xad\x6c\x6c\xe6\xef\x49\x66\x27\x66\xba\x6b\xed\x77\xf2\xf7\x84\x33\x93\xa3\xb7\x7a\x05\xcb\x6f\xf5\xf2\x07\x57\x6e\xb7\x7b\x63\x5c\x1b\xed\xde\xea\xda\xd4\x2b\xd9\xd8\x36\x53\x6c\xaf\x97\xe0\x1a\xac\x8d\x04\x57\xad\x1f\x2e\x70\x6b\x07\x57\x6b\x86\xd4\x5d\xc6\x84\x7d\x8f\xf0\x4a\x47\x53\x26\x49\x7f\x68\xfa\xbc\xb0\xea\x2f\xf0\x42\x54\xfd\x42\xb2\xc9\xb1\x0d\x61\x1e\x28\xbc\x95\x4c\x73\x7c\x52\x0a\x6a\xb8\x7c\x01\x39\xba\x6a\xdd\xc6\x1c\x70\xe5\xe0\xc3\xb6\xcd\x41\x14\xe6\xef\xc2\x92\x3f\x99\xff\xa2\xb4\x37\xbb\x5a\x4e\x38\xaf\x86\x17\x5c\x46\x61\x58\xaf\x5d\x09\xff\x7e\xd9\xf6\xfb\x2b\x74\x58\x07\x89\x5a\x5f\x95\x47\x7b\xd8\xf1\x67\x39\xaf\x20\x9c\xb6\x87\x12\x66\x69\x98\x79\xcb\x58\x43\x45\xd0\x04\x94\xd6\x6f\x5a\x6c\x12\x2a\x61\x90\xf9\xb3\x47\xa7\xcb\x1b\x37\x8a\xd2\xa6\x10\x78\x50\xf1\x4d\x8e\x0a\xbc\x85\x35\x2d\xe9\x90\xca\xef\x53\xb2\x7e\x6a\xa8\xef\x41\xe3\x9f\x00\x00\x00\xff\xff\x8e\x12\xe6\x21\x74\x1c\x00\x00")

func proto_micro_mall_order_proto_order_business_order_business_swagger_json() ([]byte, error) {
	return bindata_read(
		_proto_micro_mall_order_proto_order_business_order_business_swagger_json,
		"proto/micro_mall_order_proto/order_business/order_business.swagger.json",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"proto/micro_mall_logistics_proto/logistics_business/logistics_business.swagger.json": proto_micro_mall_logistics_proto_logistics_business_logistics_business_swagger_json,
	"proto/micro_mall_order_proto/order_business/order_business.swagger.json": proto_micro_mall_order_proto_order_business_order_business_swagger_json,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"proto": &_bintree_t{nil, map[string]*_bintree_t{
		"micro_mall_logistics_proto": &_bintree_t{nil, map[string]*_bintree_t{
			"logistics_business": &_bintree_t{nil, map[string]*_bintree_t{
				"logistics_business.swagger.json": &_bintree_t{proto_micro_mall_logistics_proto_logistics_business_logistics_business_swagger_json, map[string]*_bintree_t{
				}},
			}},
		}},
		"micro_mall_order_proto": &_bintree_t{nil, map[string]*_bintree_t{
			"order_business": &_bintree_t{nil, map[string]*_bintree_t{
				"order_business.swagger.json": &_bintree_t{proto_micro_mall_order_proto_order_business_order_business_swagger_json, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}

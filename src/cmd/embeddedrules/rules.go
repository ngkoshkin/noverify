// Package embeddedrules Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// embeddedrules/rules.php
package embeddedrules

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _embeddedrulesRulesPhp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xcd\x31\x6e\xf4\x20\x14\x04\xe0\x9e\x53\xcc\x2f\x51\xd8\x6e\xdc\xef\x1f\x89\xa4\xdf\x2e\x17\x58\x6c\x3f\xcb\x44\x2c\x20\x78\x68\x8d\xa2\xdc\x3d\xc2\x8e\x15\x57\x5b\xa5\x41\xe2\x69\xe6\x9b\x17\x15\x96\x20\x44\xdf\x75\x02\x1d\x5e\x9d\x37\x2e\x05\x1a\xd9\x78\x87\xb7\xeb\x75\x3b\x5a\xe3\x98\x22\x26\x93\xf4\x60\x49\xa0\xeb\x85\x98\xb3\xdb\x43\x4c\xd1\xe9\x58\xde\xcd\x3d\x58\x33\x97\xa6\xc5\xa7\x00\x36\x0f\xb5\x7c\xd7\x65\x20\x8c\x3e\xdb\x09\x91\x82\xd5\x23\x81\x17\x3a\x7a\x78\x18\x5e\xf0\x91\x13\x43\x8e\xde\x4d\x3f\x2d\x2e\x81\x30\x78\x6f\x4f\xd7\x5e\x60\xff\x41\x81\x63\x26\x5c\x30\x6b\x9b\xe8\xbf\x78\x32\xf8\x88\x86\x09\x3a\xe1\xd6\x54\xae\xdd\x80\xdb\x79\xe5\xdf\xdf\xce\xc8\x15\xea\x02\x59\x8e\x8d\x90\x23\x41\xae\xbf\xf4\x0a\x55\x9f\x9a\x39\xc8\xe7\x9a\xaa\xda\x5e\x36\x29\x11\x37\x72\x6d\xcf\xc6\xd7\x77\x00\x00\x00\xff\xff\x8a\x8a\x47\x68\xc4\x01\x00\x00")

func embeddedrulesRulesPhpBytes() ([]byte, error) {
	return bindataRead(
		_embeddedrulesRulesPhp,
		"embeddedrules/rules.php",
	)
}

func embeddedrulesRulesPhp() (*asset, error) {
	bytes, err := embeddedrulesRulesPhpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "embeddedrules/rules.php", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"embeddedrules/rules.php": embeddedrulesRulesPhp,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"embeddedrules": &bintree{nil, map[string]*bintree{
		"rules.php": &bintree{embeddedrulesRulesPhp, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
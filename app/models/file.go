package models

type File struct {
	File    string
	Size    int64
	Data    []byte
	GeoJSON GeoJSON
}

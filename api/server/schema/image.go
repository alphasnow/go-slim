package schema

type ImageInfo struct {
	Base64 int64
	Size   int64
	Width  int
	Height int
	Ext    string
}

type ImageRule struct {
	Base64    int64
	Size      int64
	Width     int
	Height    int
	Ext       []string
	Ratio     float32
	MinWidth  int
	MinHeight int
}

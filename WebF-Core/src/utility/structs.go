package utility

//Header structure
type Header struct {
	Filetype      uint16
	Titlelength   uint16
	Metalength    uint16
	Wordslength   uint32
	Contentlength uint64
}

//Meta structure
type Meta struct {
	Keylength   uint8
	Key         string
	Valuelength uint16
	Value       string
}

//WebF structure
type WebF struct {
	Header  Header
	Title   string
	Meta    []Meta
	Words   []Words
	Content []uint32
}

//Words structure
type Words struct {
	Length uint8
	Word   []byte
}

//JSONMeta structure
type JSONMeta struct {
	Key   string
	Value string
}

//JSONWebF structure
type JSONWebF struct {
	Title   string
	Meta    []JSONMeta
	Content string
}

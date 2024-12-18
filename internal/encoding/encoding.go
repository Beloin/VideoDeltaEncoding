package encoding

type RawData struct {
	Data []uint8
}

// TODO: Create an stream based output in this result?
// Or some byte array output?
type Result interface {
}

type Encoder interface {
	Encode(old *RawData, new *RawData) (Result, error)
}


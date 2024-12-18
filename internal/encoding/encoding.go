package encoding

type RawData struct {
  Data []uint8
}

type Encoder interface {
	Encode(*RawData)
}

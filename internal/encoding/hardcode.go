package encoding

import (
	"container/list"
	"errors"
)

type HardEncode struct {
	encodeRate float32
}

type Diff struct {
	Index uint64
	Delta int8
}

type HardResult struct {
	// List of Diff
	Diff *list.List
}

func NewHardEncode(encodeRate float32) Encoder {
	return &HardEncode{
		encodeRate,
	}
}

// Hard encode:
// index:new_value
func (h *HardEncode) Encode(old *RawData, newD *RawData) (Result, error) {
	if len(old.Data) != len(newD.Data) {
		return nil, errors.New("both old and new data must be equal in size")
	}

	return h.encode(old, newD), nil
}

func (h *HardEncode) encode(old *RawData, newD *RawData) *HardResult {
	linkedList := list.New()
	oldData := old.Data
	newData := newD.Data
	for i := 0; i < len(old.Data); i++ {
		oldV := oldData[i]
		newV := newData[i]
		diff := int8(newV) - int8(oldV)
		normDiff := norm(oldV) - norm(newV)

		if normDiff > h.encodeRate-1 {
			linkedList.PushBack(&Diff{uint64(i), diff})
		}

	}

	return &HardResult{linkedList}
}

func norm(v uint8) float32 {
	var max_ float32 = 255
	var min_ float32 = 0
	return (float32(v) - min_) / (max_ - min_)
}

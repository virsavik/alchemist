package pagination

import "errors"

type Input struct {
	Index     int
	Size      int
	WithTotal bool
}

func (in Input) IsValid() error {
	if in.Index < 1 {
		return errors.New("page index must be greater than zero")
	}

	if in.Size < 1 {
		return errors.New("page size must be greater than zero")
	}

	return nil
}

func (in Input) ToOffsetLimit() (int, int) {
	limit := in.Size
	offset := in.Size * (in.Index - 1)

	return offset, limit
}

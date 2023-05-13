package pagination

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPagination_ToOffsetLimit(t *testing.T) {
	tcs := map[string]struct {
		input     Input
		expOffset int64
		expLimit  int64
	}{
		"page 1 size 20": {
			input: Input{
				Index: 1,
				Size:  20,
			},
			expOffset: 0,
			expLimit:  20,
		},
		"page 2 size 30": {
			input: Input{
				Index: 2,
				Size:  30,
			},
			expOffset: 30,
			expLimit:  30,
		},
		"page 3 size 50": {
			input: Input{
				Index: 3,
				Size:  50,
			},
			expOffset: 100,
			expLimit:  50,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given
			input := tc.input

			// When
			offset, limit := input.ToOffsetLimit()

			// Then
			require.Equal(t, tc.expOffset, offset)
			require.Equal(t, tc.expLimit, limit)
		})
	}
}

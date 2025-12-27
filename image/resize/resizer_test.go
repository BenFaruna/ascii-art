package resize

import "testing"

func TestCalculateAspectRatio(t *testing.T) {
	cases := []struct {
		width, height int
		aspectRatio   float64
	}{
		{100, 100, 1},
		{200, 100, 2},
		{100, 200, 0.5},
	}
	s := Strategy{}
	for _, c := range cases {
		if got := s.calculateAspectRatio(c.width, c.height); got != c.aspectRatio {
			t.Errorf("calculateAspectRatio(%d, %d) = %f; want %f", c.width, c.height, got, c.aspectRatio)
		}
	}
}

func TestCalculateNewDimensions(t *testing.T) {
	cases := []struct {
		width, height, maxWidth, maxHeight int
		expectedWidth, expectedHeight      int
	}{
		{100, 100, 200, 200, 200, 200},
		{200, 100, 0, 0, 200, 100},
		{100, 200, 0, 100, 50, 100},
		{100, 100, 150, 0, 150, 150},
		{200, 100, 150, 0, 150, 75},
		{100, 200, 0, 150, 75, 150},
	}
	s := Strategy{}
	for _, c := range cases {
		if gotWidth, gotHeight := s.calculateNewDimensions(c.width, c.height, c.maxWidth, c.maxHeight); gotWidth != c.expectedWidth || gotHeight != c.expectedHeight {
			t.Errorf("calculateNewDimensions(%d, %d, %d, %d) = (%d, %d); want (%d, %d)", c.width, c.height, c.maxWidth, c.maxHeight, gotWidth, gotHeight, c.expectedWidth, c.expectedHeight)
		}
	}
}

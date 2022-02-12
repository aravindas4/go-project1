package tested

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.0, 3.0}
	got := rectangle.Perimeter()
	want := 10.0

	if got != want {
		t.Errorf("Got %v, But got %v", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if shape.Area() != want {
			t.Errorf("Got %v, But got %v", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{2.0, 3.0}
		want := 6.0

		checkArea(t, rectangle, want)

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{1}
		want := 3.14

		checkArea(t, circle, want)
	})
}

//Table driven testing
func TestArea2(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{2.0, 3.0}, 6.0},
		{Circle{1}, 3.14},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		t.Logf("%#v", tt.shape)

		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

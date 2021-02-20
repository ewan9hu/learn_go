package main

import (
	"math"
	"testing"
)

type Rectangle struct{
	Width float64
	Height float64
}

func (r Rectangle) Area() float64{
	return  r.Width * r.Height
}

type Circles struct{
	Radius float64
}

func (c Circles) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

type Shap interface{
	Area() float64
}

func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want{
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T){
	assertMessage := func(t *testing.T, got, want float64){
		t.Helper()
		if got != want{
			t.Errorf("want %.2f but got %.2f", want, got)
		}
	}

	checkArea := func(t *testing.T, shap Shap, want float64){
		t.Helper()
		got := shap.Area()
		assertMessage(t, got, want)
	}

	t.Run("rectangle area", func(t *testing.T){
		rectangle := Rectangle{5, 3}
		checkArea(t, rectangle, 15.0)
	})

	t.Run("circles", func(t *testing.T){
		circles := Circles{10.0}
		checkArea(t, circles, 314.16)
	})
}

func Perimeter(rectangle Rectangle) float64{
	return  2*(rectangle.Width + rectangle.Height)
}
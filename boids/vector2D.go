package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x + v2.x, y: v1.y + v2.y }
}

func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x - v2.x, y: v1.y - v2.y }
}

func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x * v2.x, y: v1.y * v2.y }
}

func (v1 Vector2D) AddValue(decimal float64) Vector2D {
	return Vector2D{v1.x + decimal, v1.y + decimal}
}

func (v1 Vector2D) MultiplyValue(decimal float64) Vector2D {
	return Vector2D{v1.x * decimal, v1.y * decimal}
}

func (v1 Vector2D) DivideValue(decimal float64) Vector2D {
	return Vector2D{v1.x / decimal, v1.y / decimal}
}

func (v1 Vector2D) Limit(lower, upper float64) Vector2D {
	return Vector2D{
		x: math.Min(math.Max(v1.x, lower), upper),
		y: math.Min(math.Max(v1.y, lower), upper)}
}

func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x - v2.x, y: 2) + math.Pow(v1.y - v2.y, y: 2))
}
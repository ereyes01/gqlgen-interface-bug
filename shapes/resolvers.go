package shapes

import "context"

type ShapeResolver struct{}

func (r *ShapeResolver) Query_shapes(ctx context.Context) ([]Shape, error) {
	return []Shape{
		&Circle{Radius: 10.0},
		&Rectangle{Length: 1.0, Width: 10.0},
		&Rectangle{Length: 10.0, Width: 10.0},
	}, nil
}

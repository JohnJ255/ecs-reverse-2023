package loader

import (
	"ecs_test_cars/framework"
)

type ImageResource struct {
	Filenames []string
	Rotation  framework.Radian
}

func (f *ImageResource) GetFileNames() []string {
	return f.Filenames
}

func (f *ImageResource) GetBaseAngle() framework.Radian {
	return f.Rotation
}

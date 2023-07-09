package loader

import (
	"ecs_test_cars/ecs/component"
	"ecs_test_cars/framework"
)

var CarFileNames = map[framework.ControlType]*ImageResource{
	framework.Player: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"car1.png",
		},
	},
	framework.Computer: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"car2.png",
		},
	},
}

var TrailerFileNames = map[component.TrailerType]*ImageResource{
	component.TrailerTypeCart: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"trailer1.png",
		},
	},
}

var LevelFileNames = map[int]*ImageResource{
	1: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"level1.png",
		},
	},
	2: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"level1.png",
		},
	},
	3: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"level1.png",
		},
	},
	4: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"level1.png",
		},
	},
	5: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"level1.png",
		},
	},
	6: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"level1.png",
		},
	},
	7: {
		Rotation: framework.AngleRight,
		Filenames: []string{
			"about.png",
		},
	},
}

var SoundsFilenames = map[string]string{
	"background": "background1.mp3",
	"forward":    "forward.mp3",
	"reverse":    "reverse.mp3",
	"win":        "finish_level.mp3",
	"collide":    "collide.mp3",
}

var FontFilenames = map[string]string{
	"default": "EvilEmpire.ttf",
}

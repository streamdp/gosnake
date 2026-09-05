package game

type food struct {
	position   map[coordinate]struct{}
	limit      int
	proportion int
}

func newFood(w, h, limit int) *food {
	return &food{
		position:   make(map[coordinate]struct{}, limit),
		limit:      limit,
		proportion: w * h / limit,
	}
}

func (f *food) add(snake *snake, desk *desk) {
	if len(f.position) < f.limit {
		var newFoodPoint = coordinate{}

	loop:
		for {
			newFoodPoint = desk.getRandPoint()

			for i := range snake.body {
				_, ok := f.position[newFoodPoint]
				if ok || snake.body[i] == newFoodPoint {
					continue loop
				}
			}

			break
		}

		f.position[newFoodPoint] = struct{}{}
	}
}

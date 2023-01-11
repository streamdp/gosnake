package game

type food struct {
	position map[coordinate]struct{}
	limit    int
}

func newFood() *food {
	return &food{
		position: make(map[coordinate]struct{}, 10),
		limit:    10,
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

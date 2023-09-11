package person

type Person interface {
	GetAge() uint
}

func FindMaxAge(persons ...Person) uint {
	if len(persons) == 0 {
		return 0
	}

	max := persons[0].GetAge()
	for _, p := range persons {
		age := p.GetAge()
		if age > max {
			max = age
		}
	}

	return max
}

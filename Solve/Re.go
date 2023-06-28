package Solve

func Re(sl []string, index int) []string {
	return append(sl[:index], sl[index+1:]...)
}

// Re убирает по индексу элемент слайса

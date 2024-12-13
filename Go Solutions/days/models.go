package days

type stone struct {
	number int
	blinks int
}

type farmPlot struct {
	position  [2]int
	letter    string
	neighbors []*farmPlot
	fences    int
}

func (f *farmPlot) determineFences() {
	var fences int
	for _, neighbor := range f.neighbors {
		if neighbor.letter != f.letter {
			fences++
		}
	}
	f.fences = fences + 4 - len(f.neighbors)
}

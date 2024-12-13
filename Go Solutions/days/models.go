package days

import "math"

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

type clawMachine struct {
	aButton        [2]float64
	bButton        [2]float64
	prize          [2]float64
	tokensForPrize int
	aButtonPresses float64
	bButtonPresses float64
}

func (c *clawMachine) calcButtonPresses() {
	c.aButtonPresses = math.Round((c.prize[0] - c.prize[1]*c.bButton[0]/c.bButton[1]) / (c.aButton[0] - c.aButton[1]*c.bButton[0]/c.bButton[1]))
	c.bButtonPresses = math.Round((c.prize[0] - c.prize[1]*c.aButton[0]/c.aButton[1]) / (c.bButton[0] - c.bButton[1]*c.aButton[0]/c.aButton[1]))
	if c.aButtonPresses*c.aButton[0]+c.bButtonPresses*c.bButton[0] == c.prize[0] && c.aButtonPresses*c.aButton[1]+c.bButtonPresses*c.bButton[1] == c.prize[1] {
		c.tokensForPrize = int(3*c.aButtonPresses + c.bButtonPresses)
	}
}

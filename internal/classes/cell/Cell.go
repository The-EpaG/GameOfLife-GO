package cell

type State int

const (
	Dead	State = iota
	Alive 
)

type Cell struct {
	State State
	X     int
	Y     int
}

func (cell *Cell) IsAlive() bool {
	return cell.State == Alive
}

func getNumberOfNeighborsAlive(neighbors []Cell) int {
	var numberOfNeighborsAlive int = 0

	for _, neighbor := range neighbors {
		if neighbor.State == Alive {
			numberOfNeighborsAlive += 1
		}
	}

	return numberOfNeighborsAlive
}


func (cell *Cell) WillSurvive(neighbors []Cell) bool {
	if cell.State == Dead {
		return false
	}

	var numberOfNeighborsAlive int = getNumberOfNeighborsAlive(neighbors)

	if numberOfNeighborsAlive < 2 || numberOfNeighborsAlive > 3 {
		return false
	}

	return true
}

func (cell *Cell) WillRevive(neighbors []Cell) bool {
	if cell.State == Alive {
		return false
	}

	var numberOfNeighborsAlive int = getNumberOfNeighborsAlive(neighbors)

	return numberOfNeighborsAlive == 3
}

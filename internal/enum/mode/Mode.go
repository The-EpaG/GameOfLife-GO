package mode

type Mode int
const (
	Init	Mode = iota
	Start
)

func (w Mode) String() string {
	return [...]string{"init", "start"}[w]
}
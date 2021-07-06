package factory

type Worker struct {
	Name, Position, Company string
	AnnualIncome int
}

const (
	Developer = iota
	Manager
)

func NewWorker(role int) *Worker {
	switch role {
	case Developer:
		return &Worker{"", "developer", "", 60000}
	case Manager:
		return &Worker{"", "manager", "",80000}
	default:
		panic("unsupported role")
	}
}

package inter

type Human struct {
	Name string
	Age int
	Phone string
}

type Student struct {
	Human
	School string
	Loan float32
}

type Employee struct {
	Human
	Company string
	Money float32
}

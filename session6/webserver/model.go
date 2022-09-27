package webserver

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var emps = []Employee{
	{
		ID:       1,
		Name:     "MNC",
		Age:      20,
		Division: "Tech",
	},
}

func GetEmployees() []Employee {
	return emps
}

func CreateEmployee(emp *Employee) []Employee {
	emps = append(emps, *emp)

	return emps
}

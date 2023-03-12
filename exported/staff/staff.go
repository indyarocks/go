package staff

var OverPaidLimit = 75000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) OverPaid() []Employee {
	var overpaid []Employee

	for _, emp := range e.AllStaff {
		if emp.Salary > OverPaidLimit {
			overpaid = append(overpaid, emp)
		}
	}
	return overpaid
}

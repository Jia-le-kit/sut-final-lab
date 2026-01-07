package backend/entity

import(
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestSalary(t *testing.T){
	g := NewGomegaWithT(t)

	employee:= Employee{
		Name: "Jia",
		Salary: 10000,
		EmployeeCode: "HR-1024"
	}
	ok, err := govalidator.ValidateStruct(employee)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
}
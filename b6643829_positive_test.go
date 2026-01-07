package backend/entity

import(
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPost(t *testing.T){
	g := NewGomegaWithT(t)

	employee:= Employee{
		Name: "Jia",
		Salary: 20000,
		EmployeeCode: "HR-1024"
	}
	ok, err := govalidator.ValidateStruct(employee)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
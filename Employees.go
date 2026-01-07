package backend/entity

import(
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

type Employees struct {
	gorm.Model
	Name string `valid:"length(2|80)~Name must be between 2 and 80"`
	Salary float64 `valid:"float,range(15000|200000)~Salary must be between 15000 and 200000"`
	EmployeeCode string `valid:"matches(^[A-Z]{2}+"-"+[0-9]{4})~EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"`
}

func init(){
	govalidator.SetFieldsRequiredByDefault(false)
}
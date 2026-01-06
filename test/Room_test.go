package test

import (
	"practice2/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestRoom(t *testing.T) {
	gomega := NewGomegaWithT(t)

	validData := entity.Room{
		Name:   "NormalRoom",
		Type:   "solo",
		Price:  500.5,
		Price2: 50,
	}
	t.Run("Case 1: ถูกทั้งหมด", func(t *testing.T) {
		v := validData
		ok, err := govalidator.ValidateStruct(v)
		gomega.Expect(ok).To(BeTrue())
		gomega.Expect(err).To(BeNil())
	})

	t.Run("Case 2: Name ว่าง", func(t *testing.T) {
		v := validData
		v.Name = ""
		ok, err := govalidator.ValidateStruct(v)
		gomega.Expect(ok).To(BeFalse())
		gomega.Expect(err.Error()).To(Equal("Name is required"))
	})

	t.Run("Case 3: Name ยาวน้อยกว่า 5 ตัวอักษร", func(t *testing.T) {
		v := validData
		v.Name = "Nor"
		ok, err := govalidator.ValidateStruct(v)
		gomega.Expect(ok).To(BeFalse())
		gomega.Expect(err.Error()).To(Equal("Name must be at least 5 char"))
	})

	t.Run("Case 4: Price น้อยกว่าหรือเท่ากับ 0", func(t *testing.T) {
		v := validData
		v.Price = 0
		ok, err := govalidator.ValidateStruct(v)
		gomega.Expect(ok).To(BeFalse())
		gomega.Expect(err.Error()).To(Equal("Price must be greater than 0"))
	})

	t.Run("Case 5: Price2 ติดลบ", func(t *testing.T) {
		v := validData
		v.Price2 = -0.5
		ok, err := govalidator.ValidateStruct(v)
		gomega.Expect(ok).To(BeFalse())
		gomega.Expect(err.Error()).To(Equal("Price2 cannot be negative"))
	})

	t.Run("Case 6: Type มีตัวเลข", func(t *testing.T) {
		v := validData
		v.Type = "solo50"
		ok, err := govalidator.ValidateStruct(v)
		gomega.Expect(ok).To(BeFalse())
		gomega.Expect(err.Error()).To(Equal("Type must be alphabet"))
	})

	t.Run("Case 7: Type ว่าง",func(t *testing.T) {
		v:= validData
		v.Type = ""
		ok, err := govalidator.ValidateStruct(v)
		gomega.Expect(ok).To(BeFalse())
		gomega.Expect(err.Error()).To(Equal("Type is required"))
	})
}

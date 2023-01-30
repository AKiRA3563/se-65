package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

type Users struct {
	Name		string	`valid:"required~name not blank"`
	Email		string	`valid:"email"`
	PhoneNumber	string
	Password	string
}

func TestUserValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("check name not blank", func(t *testing.T) {
		u := Users{
			Name: 		"", //false
			Email: 		"star@rmail.com",
			PhoneNumber: "0841362233",
			Password: 	"13-46",
		}

		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("name not blank"))
	})

	// t.Run("check email is valid", func(t *testing.T) {
	// 	u := Users{
	// 		Name: "Nakd",
	// 		Email: "starmail",
	// 		PhoneNumber: "0841362233",
	// 		Password: "13-46",
	// 	}

	// 	ok, err := govalidator.ValidateStruct(u)

	// 	g.Expect(ok).NotTo(gomega.BeTrue())

	// 	g.Expect(err).ToNot(gomega.BeNil())

	// 	g.Expect(err.Error()).To(gomega.Equal("Email: not vaild"))

	// })
} 
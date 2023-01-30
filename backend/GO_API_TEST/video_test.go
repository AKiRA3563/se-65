package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name	string	`valid:"required~Name cannot be blank"`
	Url		string	`gorm:"uniqueIndex" valid:"url"`
}

func TestVideoValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("check data is valid", func(t *testing.T) {
		v := Video{
			Name: "C++ for Beginner",
			Url:  "www.youhub.com",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).To(gomega.BeTrue())

		g.Expect(err).To(gomega.BeNil())

		g.Expect(err)
	})

	t.Run("check name cannot be blank", func(t *testing.T) {
		v := Video{
			Name: "",
			Url: "www.youhub.com",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).ToNot(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Name cannot be blank"))
	
	})

	t.Run("check url is valid", func(t *testing.T) {
		v := Video{
			Name: "C++ for Beginner",
			Url: "youhub,com",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).ToNot(gomega.BeTrue())

		g.Expect(err).ToNot(gomega.BeNil())

		g.Expect(err.Error()).To(gomega.Equal("Url: youhub,com does not validate as url"))
	})
}
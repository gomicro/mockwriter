package penname

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestPenName(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Errors", func() {
		g.It("should return the error set", func() {
			errText := "Something went wrong"
			p := &PenName{}
			p.ReturnError(fmt.Errorf(errText))

			_, err := p.Write([]byte("I'm trying to do something"))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(errText))
		})
	})
}

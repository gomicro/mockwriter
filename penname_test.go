package penname

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestPenname(t *testing.T) {
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

	g.Describe("Writing", func() {
		g.It("should record what is written to it", func() {
			t := "Nothing to see here move along"
			p := &PenName{}

			n, err := p.Write([]byte(t))
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(len(t)))
			Expect(p.Written).To(Equal([]byte(t)))
		})

		g.It("should function as a writer", func() {
			t := "Nothing to see here move along"
			p := &PenName{}

			fmt.Fprint(p, t)
			Expect(p.Written).To(Equal([]byte(t)))
		})
	})
}

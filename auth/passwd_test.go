package auth_test

import (
	"testing"

	gc "gopkg.in/check.v1"

	"github.com/johnweldon/go/auth"
)

func Test(t *testing.T) { gc.TestingT(t) }

type PasswordSuite struct{}

var _ = gc.Suite(&PasswordSuite{})

func (s *PasswordSuite) TestPassword(c *gc.C) {
	password := "test.password"

	hash, err := auth.HashPassword(password)
	c.Assert(err, gc.IsNil)

	c.Assert(auth.TestPassword(password, hash), gc.Equals, true)
	c.Assert(auth.TestPassword("T"+password[1:], hash), gc.Not(gc.Equals), true)
}

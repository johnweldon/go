package auth_test

import (
	gc "gopkg.in/check.v1"

	"github.com/johnweldon/go-misc/auth"
)

type RepoSuite struct{}

var _ = gc.Suite(&RepoSuite{})

func (s *RepoSuite) TestRepo(c *gc.C) {
	repo := auth.NewUserRepo()

	_, ok := repo.User("test.user")
	c.Assert(ok, gc.Not(gc.Equals), true)

	hash, err := auth.HashPassword("password")
	c.Assert(err, gc.IsNil)

	err = repo.Save("test.user", hash)
	c.Assert(err, gc.IsNil)

	u, ok := repo.User("test.user")
	c.Assert(ok, gc.Equals, true)

	ok = u.Authenticate("password")
	c.Assert(ok, gc.Equals, true)
}

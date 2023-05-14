package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestPostsHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:6543/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}

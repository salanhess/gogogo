package hello_test

import (
	"fmt"
	. "gopkg.in/check.v1"
	"io"
	"runtime"
	"testing"
)

type MySuite struct {
	scenario string
	dir      string
}

var _ = Suite(&MySuite{
	scenario: "hello bh",
})

func (s *MySuite) SetUpTest(c *C) {
	s.dir = c.MkDir()
	// Use s.dir to prepare some data.
}

func (s *MySuite) TestWithDir2(c *C) {
	// Use the data in s.dir in the test.
	fmt.Println("Senario1:", s.scenario, "\ns.dir:", s.dir)
}

func (s *MySuite) TestWithDir1(c *C) {
	// Use the data in s.dir in the test.
	fmt.Println("Senario2:", s.scenario, "\ns.dir:", s.dir)
}

func (s *MySuite) TestHelloWorld(c *C) {
	fmt.Println("Senario3:", s.scenario, "\ns.dir:", s.dir)
	//c.Assert(42, Equals, "42")
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
	str1 := "hello, I'm there"
	foo := "abc"
	bar := "abc"
	c.Assert(str1, Matches, "hel.*there")
	//c.Assert(err, IsNil)
	c.Assert(foo, Equals, bar, Commentf("#CPUs == %d", runtime.NumCPU()))
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

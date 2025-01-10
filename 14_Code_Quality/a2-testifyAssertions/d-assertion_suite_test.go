import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type MySuite struct {
	suite.Suite
}

func (s *MySuite) SetupTest() {
	// Setup code before each test
}

func (s *MySuite) TearDownTest() {
	// Teardown code after each test
}

func (s *MySuite) TestExample() {
	result := 2 + 2
	s.Equal(4, result, "2 + 2 should equal 4")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}
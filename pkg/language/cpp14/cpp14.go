package cpp14

import (
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/cpp"
)

var Lang = cpp.New("cpp14", "C++ 14", "c++14").(cpp.Cpp)

func init() {
	language.Register("cpp14", Lang)
}
package cpp11

import (
	"github.com/mraron/njudge/pkg/language"
	"github.com/mraron/njudge/pkg/language/cpp"
)

var Lang = cpp.New("cpp11", "C++ 11", "c++11").(cpp.Cpp)

func init() {
	language.Register("cpp11", Lang)
}

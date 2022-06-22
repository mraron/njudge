package cpp11

import (
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/language/cpp"
)

var Lang = cpp.New("cpp11", "C++ 11", "c++11").(cpp.Cpp)

func init() {
	language.Register("cpp11", Lang)
}

package cpp11

import (
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/language/cpp"
)

func init() {
	language.Register("cpp11", cpp.New("cpp11", "C++ 11", "c++11"))
}

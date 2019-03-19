package cpp14

import (
	"github.com/mraron/njudge/utils/language"
	"github.com/mraron/njudge/utils/language/cpp"
)

func init() {
	language.Register("cpp14", cpp.New("cpp14", "C++ 14", "c++14"))
}

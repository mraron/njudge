package testutils

import "flag"

var UseIsolate = flag.Bool("isolate", false, "run isolate integration tests")
var AllLanguages = flag.Bool("all_languages", false, "run tests for all languages")

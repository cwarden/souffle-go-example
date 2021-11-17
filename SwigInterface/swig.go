package SwigInterface

// #include <sqlite3.h>
// #include <sqlite3ext.h>
// #cgo CXXFLAGS: -fopenmp -std=c++17 -D __EMBEDDED_SOUFFLE__ -D USE_SQLITE
// #cgo LDFLAGS: -lsqlite3 -lgomp
import "C"

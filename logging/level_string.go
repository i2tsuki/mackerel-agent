// Code generated by "stringer -type=level level.go"; DO NOT EDIT

package logging

import "fmt"

const _level_name = "TRACEDEBUGINFOWARNINGERRORCRITICAL"

var _level_index = [...]uint8{0, 5, 10, 14, 21, 26, 34}

func (i level) String() string {
	i -= 1
	if i >= level(len(_level_index)-1) {
		return fmt.Sprintf("level(%d)", i+1)
	}
	return _level_name[_level_index[i]:_level_index[i+1]]
}

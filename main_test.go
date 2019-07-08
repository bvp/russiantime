package russiantime

import (
	"testing"
)

func TestParseDateString(t *testing.T) {
	t.Logf("%v", ParseDateString("22 февраля 2018 г. 16:43"))
	t.Logf("%v", ParseDateString("22 февраля 2018 г."))
}

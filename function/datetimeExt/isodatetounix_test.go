package datetimeExt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatDatetime_Eval(t *testing.T) {
	n := &IsoDateToUnix{}

	date, err := n.Eval("2017-04-12T22:15:09")
	assert.Nil(t, err)
	assert.NotNil(t, date)
	assert.Equal(t, int64(1492031709), date)
	fmt.Println(date)
}

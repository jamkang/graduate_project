package Test

import (
	"fmt"
	"pro10/src/app/tool"
	"testing"
)

func TestName(t *testing.T) {
	if value, bo := tool.Province("上海市"); !bo {
		fmt.Println(value)
	} else {
		fmt.Println(value)
	}
}

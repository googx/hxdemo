/*
--------------------------------------------------
 File Name: xxtest.go
 Author: hanxu
 AuthorSite: http://www.googx.top/
 GitSource: https://github.com/googx/linuxShell
 Created Time: 2019-11-4-下午11:15
---------------------说明--------------------------

---------------------------------------------------
*/

package xxtest

import (
	"fmt"
)

type HxTest struct {
	Xname string
	Xage  int8
	Xsex  bool
}

func NewHxTest() *HxTest {
	test := new(HxTest)
	test.Xage = 22
	return test
}

func (this HxTest) Name() string {
	return fmt.Sprintf("hello,%v", this.Xname)
}

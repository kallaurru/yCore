package test

import (
	"github.com/kallaurru/ycore/rule"
	"log"
	"testing"
)

func TestMakeComplexRule(t *testing.T) {
	cr := rule.MakeEmptyComplexRule()
	cr = cr.OnNumbersChanging().OnCaseChanging()
	log.Println(cr)
}

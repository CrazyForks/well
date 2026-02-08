package hookjs

import (
	_ "embed"
	"testing"

	"github.com/dop251/goja"
	"github.com/shynome/err0/try"
	almond "github.com/shynome/goja-almond"
)

//go:embed allow_anyone.js
var tjs string

func TestModule(t *testing.T) {
	vm := goja.New()
	try.To1(almond.Enable(vm))
	deps := try.To1(genDeps([]string{"peers"}))
	try.To1(vm.RunProgram(deps))
	h := try.To1(precompile(vm, tjs))
	t.Log(h)
}

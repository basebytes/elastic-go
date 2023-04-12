package _examples

import (
	"fmt"
	"github.com/basebytes/tools"
	"testing"
)

func TestScript(t *testing.T) {
	script := Specialized.Script(
		Specialized.Script.WithScriptId("stored_script", map[string]interface{}{
			"count": "10",
		}),
	)
	fmt.Println(tools.Encode(script))
}

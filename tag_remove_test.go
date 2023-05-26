package uitools_test

import (
	"strings"
	"testing"

	"github.com/cdvelop/uitools"
)

// quita solo 2 etiquetas principio y final <scrip></script>
func Test_RemoveTagStartedEnd(t *testing.T) {
	//se esperan 6 lineas de regreso
	text := `

	<scrip>
	function otherFun(x) {

	}


	function DataForm(form) {
	hola 
	soy 
	linea 
	cuatro 
	cinco 
	seis 

	}
	</scrip>
	


	`
	uitools.RemoveTagStartdEnd(&text, "scrip")
	lines := strings.Split(string(text), "\n")
	if len(lines) != 14 {
		t.Fail()
	}

}

func TestRemoveTagStartdEnd(t *testing.T) {
	input := `<script>
			function keyboard(t) {
				let key = t.which || t.keyCode;
				if (key == 13) {
					t.preventDefault();
					// console.log('key enter save', key);
					functionM(t);
					keyEnterOn = false;
				}
			};
		</script>`
	expectedOutput := `function keyboard(t) {
				let key = t.which || t.keyCode;
				if (key == 13) {
					t.preventDefault();
					// console.log('key enter save', key);
					functionM(t);
					keyEnterOn = false;
				}
			};`

	uitools.RemoveTagStartdEnd(&input, "script")

	if input != expectedOutput {
		t.Errorf("Expected output:\n%s\n\nActual output:\n%s", expectedOutput, input)
	}
}

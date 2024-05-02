package main

import (
	"MiniGO/checker"
	"MiniGO/compiler"
	"MiniGO/compiler/parser"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
	"strings"
)

func main() {

	myApp := app.New()
	codeWindow := myApp.NewWindow("Editor de código")
	errorWindow := myApp.NewWindow("Errores de compilación")

	// Textarea para el código
	codeTextArea := widget.NewMultiLineEntry()
	codeTextArea.SetPlaceHolder("Escribe tu código aquí")

	// Función para obtener el número de líneas del texto
	getLineNumber := func(text string) string {
		lines := strings.Split(text, "\n")
		lineNumber := ""
		for i := 1; i <= len(lines); i++ {
			lineNumber += strconv.Itoa(i) + "\n"
		}
		return lineNumber
	}

	// Columna de números para enumerar las filas
	lineNumberTextArea := widget.NewMultiLineEntry()
	lineNumberTextArea.SetText(getLineNumber(codeTextArea.Text))
	lineNumberTextArea.Disable()

	// Botón para compilar
	compileButton := widget.NewButton("Compilar", func() {

		// Input string to tokenize
		input := antlr.NewInputStream(codeTextArea.Text)

		// Create the lexer and stream
		lexer := parser.NewgoScanner(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewgoParser(stream)

		errorListener := compiler.NewMyErrorListener()
		p.RemoveErrorListeners()
		p.AddErrorListener(errorListener)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errorListener)

		tree := p.Root()

		myChecker := checker.NewTypeChecker()
		tree.Accept(myChecker)

		var errorMessages []string

		if len(errorListener.ErrorMsgs) > 0 {
			fmt.Println("Errores encontrados de sintaxis:")
			for _, errMsg := range errorListener.ErrorMsgs {
				fmt.Println(errMsg)
				errorMessages = append(errorMessages, errMsg)
			}
		} else {
			fmt.Println("Compilación terminada sin errores")
		}

		if myChecker.HasErrors() {
			fmt.Println("Errores encontrados de tipo:")
			for _, errMsg := range myChecker.GetErrors() {
				fmt.Println(errMsg)
				errorMessages = append(errorMessages, string(errMsg))
			}
		} else {
			fmt.Println("Compilación terminada sin errores")
		}

		// Unir todas las letras del mensaje de error en una sola cadena
		errorMessage := strings.Join(errorMessages, "")

		errorTextArea := widget.NewMultiLineEntry()
		errorTextArea.SetText(errorMessage)

		if len(errorMessages) > 0 {
			errorWindowContent := container.NewVBox(
				widget.NewLabel("Errores en la compilación:"),
				container.New(layout.NewGridWrapLayout(fyne.NewSize(400, 300)), errorTextArea),
			)
			errorWindow.SetContent(errorWindowContent)
			errorWindow.Show()
		}

		fmt.Println("Compilación terminada")
	})

	// Función para actualizar los números de línea
	updateLineNumbers := func() {
		lineNumberTextArea.SetText(getLineNumber(codeTextArea.Text))
	}

	// Escuchar cambios en el texto del código
	codeTextArea.OnChanged = func(text string) {
		updateLineNumbers()
	}

	// Ajustar el ancho del lineNumberTextArea
	lineNumberTextArea.Resize(fyne.NewSize(10, 0)) // Ancho de 50px (ajusta según necesites)

	// Contenedor para el área de texto del código y la columna de números
	codeContainer := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		lineNumberTextArea,
		codeTextArea,
	)

	mainContainer := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, compileButton, nil, nil),
		codeContainer,
		compileButton,
	)

	codeWindow.SetContent(mainContainer)
	codeWindow.Resize(fyne.NewSize(800, 500)) // Tamaño más grande
	codeWindow.ShowAndRun()

}

/*
import (
	parser "MiniGO/compiler/generated"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	myApp := app.New()
	codeWindow := myApp.NewWindow("Editor de código")
	errorWindow := myApp.NewWindow("Errores de compilación")

	// Textarea para el código
	codeTextArea := widget.NewMultiLineEntry()
	codeTextArea.SetPlaceHolder("Escribe tu código aquí")

	// Botón de compilación
	compileButton := widget.NewButton("Compilar", func() {
		// Capturar el código del textarea
		input := antlr.NewInputStream(codeTextArea.Text)
		lexer := parser.NewgoScanner(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewgoParser(stream)

		// Capturar errores
		errorListener := &ErrorListener{}
		p.RemoveErrorListeners()
		p.AddErrorListener(errorListener)

		// Intentar analizar el código
		_ = p.Root()

		// Mostrar errores en la ventana de errores
		if len(errorListener.errors) > 0 {
			var errorMessages []fyne.CanvasObject
			for _, err := range errorListener.errors {
				errorMessages = append(errorMessages, widget.NewLabel(err))
			}
			errorWindowContent := container.NewVBox(
				widget.NewLabel("Errores de compilación:"),
				container.NewVScroll(container.NewVBox(errorMessages...)),
			)
			errorWindow.SetContent(errorWindowContent)
			errorWindow.Show()
		} else {
			// Si no hay errores, limpiar la ventana de errores
			errorWindow.Hide()
			fmt.Println("Compilación exitosa.")
		}
	})

	// Contenedor principal
	mainContainer := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, compileButton, nil, nil),
		codeTextArea,
		compileButton,
	)

	codeWindow.SetContent(mainContainer)
	codeWindow.Resize(fyne.NewSize(800, 500)) // Tamaño más grande
	codeWindow.ShowAndRun()
}

// ErrorListener para capturar los errores de ANTLR
type ErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func (e *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	errorMessage := fmt.Sprintf("Error en línea %d, columna %d: %s", line, column, msg)
	e.errors = append(e.errors, errorMessage)
}*/

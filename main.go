package main

import (
	parser "MiniGO/Compiler"
	"bytes"
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
			var errorBuffer bytes.Buffer
			for _, err := range errorListener.errors {
				errorBuffer.WriteString(err + "\n")
			}
			errorWindowContent := container.NewVBox(
				widget.NewLabel("Errores de compilación:"),
				widget.NewLabel(errorBuffer.String()),
			)
			errorWindow.Resize(codeWindow.Canvas().Size()) // Establecer el tamaño de la ventana de errores
			errorWindow.SetContent(errorWindowContent)
			errorWindow.Show()
		} else {
			// Si no hay errores, limpiar la ventana de errores
			errorWindow.Hide()
		}
	})

	// Contenedor principal con BorderLayout
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
	errorMessage := antlr.NewDefaultErrorStrategy().GetTokenErrorDisplay(offendingSymbol.(antlr.Token))
	e.errors = append(e.errors, errorMessage)
}

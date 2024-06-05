package main

import (
	"MiniGO/compiler"
	"MiniGO/compiler/parser"
	"MiniGO/encoder"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"os"
	"os/exec"
)

//------------------------Main con interfaz para correr el type checker------------------------
/*func main() {

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

}*/

//--------------------------------------------------------------------------

func main() {
	fileName := "Prueba.txt"
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	goScanner := parser.NewgoScanner(input)
	tokens := antlr.NewCommonTokenStream(goScanner, antlr.TokenDefaultChannel)
	goParser := parser.NewgoParser(tokens)

	//Manejo de errores con el errorListener de antlr(se comenta para usar el errorListener personalizado)
	//goParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	goErrorListener := compiler.NewMyErrorListener()
	goScanner.RemoveErrorListeners()
	goScanner.AddErrorListener(goErrorListener)
	goParser.RemoveErrorListeners()
	goParser.AddErrorListener(goErrorListener)

	tree := goParser.Root()

	if !goErrorListener.HasErrors() {

		encoder := encoder.NewEncoder()
		module := encoder.Visit(tree)
		fmt.Println(module)
		runModule(module.(*ir.Module))

	} else {
		fmt.Println("Compilation failed")
		fmt.Println(goErrorListener.String())
	}

}

func runModule(module *ir.Module) {
	// Escribir el módulo LLVM en un archivo
	f, err := os.Create("module.ll")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer f.Close()
	if _, err := module.WriteTo(f); err != nil {
		fmt.Println("Error al escribir el módulo:", err)
		return
	}

	// Compilar el módulo LLVM a código objeto
	cmd := exec.Command("clang", "", "module.ll", "-o", "module.exe")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al compilar el módulo:", err)
		return
	}

	fmt.Println("El archivo ejecutable .exe ha sido generado correctamente.")
}

//ejecute el programa
/*cmd = exec.Command("module.exe")
var out bytes.Buffer
cmd.Stdout = &out
if errors.Is(cmd.Err, exec.ErrDot) {
	cmd.Err = nil
}
if err := cmd.Run(); err != nil {
	log.Fatal(err)
	fmt.Println("Error al ejecutar el comando:", err)
	return
}

fmt.Println("Salida de la consola:", out.String())*/

/*
// Funcion para la sincronizacion de archivos
var fileMutex sync.Mutex

func saveToFileAsync(filename string, content string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	go func() {
		// Verificar si el archivo existe
		if _, err := os.Stat(filename); err == nil {
			// El archivo existe, se puede reescribir
			err = ioutil.WriteFile(filename, []byte(content), 0644)
			if err != nil {
				log.Fatalf("Error saving file: %v", err)
			}
			log.Printf("File %s has been rewritten successfully", filename)
		} else if os.IsNotExist(err) {
			// El archivo no existe, se puede crear y escribir en él
			err := ioutil.WriteFile(filename, []byte(content), 0644)
			if err != nil {
				log.Fatalf("Error saving file: %v", err)
			}
			log.Printf("File %s has been created and content saved", filename)
		} else {
			log.Fatalf("Error checking file existence: %v", err)
		}
	}()
}
func openEditorWindow(a fyne.App) {
	// Crear una nueva ventana para el editor
	w := a.NewWindow("Code Editor")
	w.Resize(fyne.NewSize(450, 350))

	// Contenido de la nueva ventana
	editorContent := container.NewVBox()

	// Añadir un cuadro de texto al contenido
	codeEntry := widget.NewMultiLineEntry()
	codeEntry.Resize(fyne.NewSize(400, 300))
	editorContent.Add(codeEntry)

	// Añadir un botón para guardar el código en un archivo
	saveButton := widget.NewButton("Save", func() {
		code := codeEntry.Text

		// Crear una nueva ventana para el nombre del archivo
		b := a.NewWindow("Enter File Name")
		b.Resize(fyne.NewSize(200, 100))

		// Crear un campo de entrada para el nombre del archivo
		nameEntry := widget.NewEntry()
		nameEntry.SetPlaceHolder("Enter file name")

		// Crear un botón para guardar el archivo
		saveFileButton := widget.NewButton("Save File", func() {
			filename := nameEntry.Text

			// Guardar el código en el archivo especificado por el usuario
			saveToFileAsync(filename+".txt", code)

			// Cerrar la ventana de entrada de nombre de archivo
			b.Close()
		})

		// Crear un contenedor para organizar los elementos en la ventana de nombre de archivo
		nameContent := container.NewVBox(
			nameEntry,
			saveFileButton,
		)

		// Establecer el contenido de la ventana de nombre de archivo
		b.SetContent(nameContent)

		// Mostrar la ventana de nombre de archivo
		b.Show()
	})

	editorContent.Add(saveButton)

	// Añadir un botón para compilar al contenido
	compileButton := widget.NewButton("Compile", func() {
		//ventanita para solicitar el nombre del archivo a compilar
		m := a.NewWindow("Enter File Name to Compile")
		m.Resize(fyne.NewSize(200, 100))
		nameEntryFile := widget.NewEntry()
		nameEntryFile.SetPlaceHolder("Enter file name to compile")
		//Agregar boton de compilacion
		compileFileButton := widget.NewButton("Compile File", func() {
			filename := nameEntryFile.Text
			errors := parseCode(filename+".txt", w)
			m.Close()
			//Dividir los errores en errores individuales
			errorList := strings.Split(errors, "\n")
			//Mostrar un cuadro de diálogo para cada error
			for _, err := range errorList {
				if err != "" {
					dialog.ShowInformation("Compilation Error", err, w)
				}
			}
		})
		nameContent := container.NewVBox(
			nameEntryFile,
			compileFileButton,
		)

		// Establecer el contenido de la ventana de nombre del archivo a compilar
		m.SetContent(nameContent)

		// Mostrar la ventana de nombre del archivo a compilar
		m.Show()
	})
	editorContent.Add(compileButton)

	// Establecer el contenido de la nueva ventana
	w.SetContent(editorContent)

	// Mostrar la nueva ventana
	w.Show()
}

func readFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func parseCode(filename string, w fyne.Window) string {
	if filename == "" {
		dialog.ShowInformation("Error", "File name cannot be empty", w)
		return ""
	}
	if filename == "" {
		dialog.ShowInformation("Error", "File name cannot be empty", w)
		return ""
	}

	// Abrir el archivo para leer
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		return ""
	}

	// Crear el lexer y el parser
	lexer := parser.NewgoScanner(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewgoParser(stream)

	// Agregar un error listener al parser
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	goErrorListener := compiler.NewMyErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(goErrorListener)
	p.RemoveErrorListeners()
	p.AddErrorListener(goErrorListener)

	// Llamar al método de inicio de regla del parser
	p.Root()

	if !goErrorListener.HasErrors() {
		//mostrar en la gui que la compilacion fue exitosa
		//crear una ventana para mostrar el mensaje
		dialog.ShowInformation("Compilation", "Succesful Compilation", w)
		fmt.Println("Succesful Compilation")
		return ""
	} else {
		//mostrar en la gui que la compilacion fue fallida
		//crear una ventana para mostrar el mensaje
		dialog.ShowInformation("Compilation", "Failed Compilation", w)
		//mostrar cada error en la gui en un cuadro de dialogo
		dialog.ShowInformation("Compilation Error", goErrorListener.String(), w)
		fmt.Println("Failed Compilation")
		return goErrorListener.String()
	}
}*/

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

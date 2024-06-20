package main

import (
	webview "seb.cl/mermaid/webview"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var mermaidCode string

	if len(os.Args) < 2 {
		mermaidCode = ""
	} else {
		arg := os.Args[1]
		if _, err := os.Stat(arg); err == nil {
			mermaidCodeBytes, err := ioutil.ReadFile(arg)
			if err != nil {
				log.Fatalf("Error al leer el archivo: %v", err)
			}
			mermaidCode = string(mermaidCodeBytes)
		} else {
			mermaidCode = arg
		}
	}

	contenidoHTML := fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
    <body style="text-align:center">
        <div class="mermaid" style="color:white">
            %s
        </div>
        <div>
            <input type="file" id="fileInput" style="display: none;" accept=".mmd,.mermaid,.txt"/><br><br>
            <button onclick="document.getElementById('fileInput').click();" id="btn">Seleccionar archivo Mermaid</button>
        </br>
        <script>
						const container = document.querySelector('.mermaid');
						function insertarScript() {
								var script = document.createElement('script');
								script.src = 'https://cdnjs.cloudflare.com/ajax/libs/mermaid/10.9.1/mermaid.min.js';
								script.id = 'mermaidScript';
								script.onload=()=>mermaid.init();
								document.head.appendChild(script);
						}
						if (container.innerText.length>5) insertarScript();
            document.getElementById('fileInput').addEventListener('change', function(event) {
                const file = event.target.files[0];
                if (file) {
                    const reader = new FileReader();
                    reader.onload = function(e) {
                        const content = e.target.result;
                        container.innerHTML = content.replace(/\n/g, ';').trim();
												if (window.mermaidScript) {
													mermaid.render('mermaidContainer', container.innerText).then(svg=>container.innerHTML=svg.svg);
												}
												else insertarScript();
                    };
                    reader.readAsText(file);
                }
            });
        </script>
    </body>
    </html>
    `, strings.Replace(mermaidCode, "\n", ";", -1))

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Visualizador de Mermaid por Seb")
	w.SetSize(800, 600, 0)
	w.Navigate("data:text/html," + contenidoHTML)
	w.Run()
}

# GO Mermaid

Este es un visualizador simple de diagramas Mermaid que muestra un diagrama en una ventana web utilizando la biblioteca [webview](https://github.com/webview/webview_go).

## Instalación

```
go run mermaid.go
```

## Compilacion (opcional)

   - Para Linux o macOS:

     ```sh
     go build mermaid.go
     ```

   - Para Windows:

     ```sh
     go build -ldflags -H=windowsgui mermaid.go
     ```

   - Para Windows (desde Linux):

     ```sh
     GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build mermaid.go
     ```

## Uso

1. Ejecuta el programa sin argumentos para abrir un diálogo de selección de archivos y seleccionar un archivo que contenga el código Mermaid, o ejecútalo con un argumento que sea el código Mermaid directamente:

   ```sh
   ./mermaid          # Para abrir el diálogo de selección de archivos
   ./mermaid archivo.txt  # Para usar un archivo específico
   ./mermaid "código Mermaid"  # Para usar el código Mermaid directamente
   ```


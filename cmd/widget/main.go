package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	webview "github.com/webview/webview_go"
)

const internalPort = "18080"

func main() {
	rootDir := findRootDir()
	log.Printf("[Clippy] Directorio raíz: %s", rootDir)

	// Arrancar servidor HTTP interno en goroutine
	go startServer(rootDir)
	time.Sleep(500 * time.Millisecond)

	// Obtener tamaño de pantalla (plataforma-específico)
	sw, sh := getScreenSize()

	// Crear ventana WebView
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("Clippy")
	w.SetSize(sw, sh, webview.HintFixed)

	// Inyectar clase widget-mode ANTES de que cargue el DOM
	// Esto oculta el panel de debug y activa fondo transparente
	w.Init(`
		document.addEventListener('DOMContentLoaded', function() {
			document.body.classList.add('widget-mode');
		}, true);
	`)

	// Navegar al frontend
	url := fmt.Sprintf("http://localhost:%s/", internalPort)
	w.Navigate(url)

	// Aplicar propiedades de escritorio (sin bordes, transparente, siempre-abajo)
	// Debe llamarse ANTES de Run() mientras la ventana existe pero no está visible
	setupDesktopWindow(w, sw, sh)

	log.Printf("[Clippy] Widget Desktop iniciado en %s (%dx%d)", runtime.GOOS, sw, sh)
	w.Run()
}

// findRootDir localiza el directorio raíz del proyecto buscando la carpeta "frontend".
func findRootDir() string {
	// 1. Intentar relativo al ejecutable (producción)
	if exe, err := os.Executable(); err == nil {
		dir := filepath.Dir(exe)
		if filepath.Base(dir) == "bin" {
			dir = filepath.Dir(dir)
		}
		if _, err := os.Stat(filepath.Join(dir, "frontend")); err == nil {
			return dir
		}
	}

	// 2. Buscar hacia arriba desde el directorio actual (desarrollo)
	cwd, _ := os.Getwd()
	dir := cwd
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(filepath.Join(dir, "frontend")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return cwd
}

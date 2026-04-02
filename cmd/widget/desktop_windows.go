//go:build windows

package main

import (
	"log"
	"syscall"
	"unsafe"

	webview "github.com/webview/webview_go"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetSystemMetrics = user32.NewProc("GetSystemMetrics")
	procSetWindowLongPtr = user32.NewProc("SetWindowLongPtrW")
	procGetWindowLongPtr = user32.NewProc("GetWindowLongPtrW")
	procSetWindowPos     = user32.NewProc("SetWindowPos")
	procShowWindow       = user32.NewProc("ShowWindow")
	procSetLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")
)

const (
	GWL_EXSTYLE       = ^uintptr(19) // -20
	GWL_STYLE         = ^uintptr(15) // -16

	WS_POPUP          = 0x80000000
	WS_EX_LAYERED     = 0x00080000
	WS_EX_TOOLWINDOW  = 0x00000080 // Oculta de Alt+Tab y taskbar
	WS_EX_NOACTIVATE  = 0x08000000 // No roba el foco
	WS_EX_TRANSPARENT = 0x00000020 // Click-through en áreas transparentes

	HWND_BOTTOM   = uintptr(1) // Z-order: debajo de todo
	SWP_NOSIZE    = 0x0001
	SWP_NOMOVE    = 0x0002
	SWP_NOACTIVATE= 0x0010
	SWP_SHOWWINDOW= 0x0040

	SM_CXSCREEN = 0
	SM_CYSCREEN = 1

	LWA_ALPHA    = 0x00000002
	LWA_COLORKEY = 0x00000001
	SW_SHOW      = 5
)

// setupDesktopWindow configura la ventana de la WebView como widget de escritorio en Windows.
func setupDesktopWindow(w webview.WebView, screenW, screenH int) {
	hwnd := uintptr(unsafe.Pointer(w.Window()))
	if hwnd == 0 {
		log.Println("[Widget/Windows] WARN: No se pudo obtener el HWND")
		return
	}

	// Leer el estilo extendido actual
	exStyle, _, _ := procGetWindowLongPtr.Call(hwnd, GWL_EXSTYLE)

	// Añadir: ToolWindow (desaparece de Alt+Tab/taskbar), Layered (transparencia), NoActivate
	newExStyle := exStyle | WS_EX_TOOLWINDOW | WS_EX_LAYERED | WS_EX_NOACTIVATE
	procSetWindowLongPtr.Call(hwnd, GWL_EXSTYLE, newExStyle)

	// Para click-through en áreas transparentes, añadir WS_EX_TRANSPARENT.
	// Nota: esto hace TODA la ventana click-through. Clippy intercepta los suyos via JS.
	newExStyle = newExStyle | WS_EX_TRANSPARENT
	procSetWindowLongPtr.Call(hwnd, GWL_EXSTYLE, newExStyle)

	// Mover al fondo del Z-order (debajo de todas las ventanas normales)
	procSetWindowPos.Call(
		hwnd,
		HWND_BOTTOM,
		0, 0,
		uintptr(screenW), uintptr(screenH),
		SWP_NOACTIVATE|SWP_SHOWWINDOW,
	)

	// Transparencia por alpha (255 = completamente opaco para el área de Clippy)
	// El CSS maneja la transparencia real del fondo
	procSetLayeredWindowAttributes.Call(hwnd, 0, 255, LWA_ALPHA)

	log.Printf("[Widget/Windows] Ventana desktop configurada: %dx%d, toolwindow, hwnd_bottom, transparente", screenW, screenH)
}

// getScreenSize obtiene las dimensiones de la pantalla principal en Windows.
func getScreenSize() (int, int) {
	w, _, _ := procGetSystemMetrics.Call(SM_CXSCREEN)
	h, _, _ := procGetSystemMetrics.Call(SM_CYSCREEN)
	if int(w) <= 0 || int(h) <= 0 {
		log.Println("[Widget/Windows] WARN: Usando dimensiones fallback 1920x1080")
		return 1920, 1080
	}
	return int(w), int(h)
}

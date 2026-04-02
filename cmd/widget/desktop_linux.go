//go:build linux

package main

// #cgo pkg-config: gtk+-3.0 webkit2gtk-4.0
// #include <gtk/gtk.h>
// #include <gdk/gdk.h>
// #include <webkit2/webkit2.h>
//
// // Configura la ventana GTK como widget de escritorio:
// // sin decoraciones, debajo de todo, skip-taskbar, fondo transparente.
// void setup_gtk_desktop_window(void* gtkWindowPtr, int w, int h) {
//     GtkWindow* win = GTK_WINDOW(gtkWindowPtr);
//
//     // Sin decoraciones (sin barra de título)
//     gtk_window_set_decorated(win, FALSE);
//
//     // Tipo DESKTOP: pegada al escritorio, debajo de todo
//     gtk_window_set_type_hint(win, GDK_WINDOW_TYPE_HINT_DESKTOP);
//
//     // Saltar en la barra de tareas y en el paginador de escritorios
//     gtk_window_set_skip_taskbar_hint(win, TRUE);
//     gtk_window_set_skip_pager_hint(win, TRUE);
//
//     // Mantener siempre debajo de otras ventanas
//     gtk_window_set_keep_below(win, TRUE);
//
//     // Activar RGBA visual para transparencia real
//     GdkScreen* screen = gtk_window_get_screen(win);
//     GdkVisual* visual = gdk_screen_get_rgba_visual(screen);
//     if (visual != NULL) {
//         gtk_widget_set_visual(GTK_WIDGET(win), visual);
//     }
//     gtk_widget_set_app_paintable(GTK_WIDGET(win), TRUE);
//
//     // Tamaño: pantalla completa
//     gtk_window_set_default_size(win, w, h);
//     gtk_window_move(win, 0, 0);
// }
//
// // Obtiene tamaño de la pantalla principal.
// void get_screen_dimensions(int* outW, int* outH) {
//     GdkDisplay* display = gdk_display_get_default();
//     GdkMonitor* monitor = gdk_display_get_primary_monitor(display);
//     GdkRectangle geometry;
//     gdk_monitor_get_geometry(monitor, &geometry);
//     *outW = geometry.width;
//     *outH = geometry.height;
// }
import "C"

import (
	"log"
	"unsafe"

	webview "github.com/webview/webview_go"
)

// setupDesktopWindow configura la ventana WebView como widget de escritorio en Linux.
func setupDesktopWindow(w webview.WebView, screenW, screenH int) {
	// w.Window() devuelve el puntero al GtkWindow en Linux
	gtkWin := w.Window()
	if gtkWin == nil {
		log.Println("[Widget/Linux] WARN: No se pudo obtener el GtkWindow")
		return
	}

	C.setup_gtk_desktop_window(
		unsafe.Pointer(gtkWin),
		C.int(screenW),
		C.int(screenH),
	)
	log.Printf("[Widget/Linux] Ventana desktop configurada: %dx%d, sin decoraciones, debajo de todo", screenW, screenH)
}

// getScreenSize obtiene las dimensiones de la pantalla principal.
func getScreenSize() (int, int) {
	var w, h C.int
	C.get_screen_dimensions(&w, &h)
	if int(w) <= 0 || int(h) <= 0 {
		log.Println("[Widget/Linux] WARN: Usando dimensiones fallback 1920x1080")
		return 1920, 1080
	}
	return int(w), int(h)
}

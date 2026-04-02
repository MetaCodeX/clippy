package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Emotion representa pitch y speed para SAM TTS
type Emotion struct {
	Pitch string
	Speed string
}

var emotionProfiles = map[string]Emotion{
	"base":       {"44", "82"},
	"feliz":      {"34", "65"},
	"enojado":    {"58", "70"},
	"triste":     {"60", "110"},
	"asustado":   {"30", "72"},
	"confundido": {"48", "95"},
}

// startServer arranca el servidor HTTP interno que sirve el frontend y el TTS.
func startServer(rootDir string) {
	mux := http.NewServeMux()

	// Frontend estático
	frontendDir := filepath.Join(rootDir, "frontend")
	mux.Handle("/", http.FileServer(http.Dir(frontendDir)))

	// Endpoint SAM TTS
	samPath := filepath.Join(rootDir, "SAM", "sam")
	mux.HandleFunc("/habla", func(w http.ResponseWriter, r *http.Request) {
		handleHabla(w, r, samPath)
	})

	// Assets (GIFs, etc.)
	assetsDir := filepath.Join(rootDir, "assets")
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetsDir))))

	addr := ":" + internalPort
	log.Printf("[Clippy/Server] Escuchando en localhost%s (frontend: %s)", addr, frontendDir)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("[Clippy/Server] Error fatal: %v", err)
	}
}

// handleHabla sintetiza voz con SAM y devuelve el WAV.
func handleHabla(w http.ResponseWriter, r *http.Request, samPath string) {
	texto := r.URL.Query().Get("texto")
	if texto == "" {
		http.Error(w, "Falta el parámetro 'texto'", http.StatusBadRequest)
		return
	}

	emocionReq := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("emocion")))
	profile, exists := emotionProfiles[emocionReq]
	if !exists {
		profile = emotionProfiles["base"]
	}

	textoStr := strings.TrimSpace(texto)
	tempFileName := fmt.Sprintf("clippy_%d.wav", time.Now().UnixNano())
	tempFilePath := filepath.Join(os.TempDir(), tempFileName)

	cmd := exec.Command(samPath,
		"-pitch", profile.Pitch,
		"-speed", profile.Speed,
		"-mouth", "141",
		"-throat", "145",
		"-wav", tempFilePath,
		textoStr,
	)

	if err := cmd.Run(); err != nil {
		log.Printf("[Clippy/TTS] Error SAM: %v", err)
		http.Error(w, "Error generando voz", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFilePath)

	file, err := os.Open(tempFilePath)
	if err != nil {
		log.Printf("[Clippy/TTS] Error abriendo WAV: %v", err)
		http.Error(w, "Error leyendo audio", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "audio/wav")
	if _, err = io.Copy(w, file); err != nil {
		log.Printf("[Clippy/TTS] Error enviando audio: %v", err)
	}
}

package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	http.HandleFunc("/SAPI4/SAPI4", handleSapi4)
	http.HandleFunc("/SAPI4/VoiceLimitations", handleLimitations)

	port := "8081"
	log.Printf("Iniciando proxy de SAPI4 local (Balabolka Console) sobre puerto %s...", port)
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleLimitations(w http.ResponseWriter, r *http.Request) {
    // Retornamos un JSON genérico si se necesita, aunque el backend prinicpal de momento no lo consulta.
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`[{"voice": "Adult Male #1, Spanish (TruVoice)"}, {"voice": "Adult Female #1, Spanish (TruVoice)"}]`))
}

func handleSapi4(w http.ResponseWriter, r *http.Request) {
	// Leer parámetros de tetyys
	text := r.URL.Query().Get("text")
	voice := r.URL.Query().Get("voice") // Ej: "Adult Male #1, Spanish (TruVoice)"
	pitchStr := r.URL.Query().Get("pitch") // Tetyys pitch: 0-200, Balcon: -10 to +10, pero Balcon soporta SAPI4 directo si no enviamos mod. O podemos mapear.
	speedStr := r.URL.Query().Get("speed") // Tetyys speed: 0-200+, Balcon: -10 to +10.

	if text == "" {
		http.Error(w, "text is required", http.StatusBadRequest)
		return
	}

	if voice == "" || voice == "Adult Male #1, American English (TruVoice)" {
		// Por si acaso recibimos la petición legacy, usamos la voz en español que instalamos de TruVoice.
        // Balcon.exe nombra las voces L&H según lo informe el sistema.
		voice = "Julio" // L&H TruVoice Spanish male, la clásica. ("Carmen" es la female).
	}

    // Ignoramos la escala de pitch de balcon.exe ya que a valores altos (-p 8) cuelga
    // la llamada COM de L&H TruVoice creando dialogs de error Win32 que bloquean Xvfb.
    // También omitimos enviar el tag \Pit= fijo en SAPI4 porque este anula las inflexiones
    // y entonaciones naturales de puntuaciones como `!` o `?`.
    // Delegaremos la velocidad y el pitch a conversiones de SoX puras en Linux nativo.
    taggedText := text

	// Nombre del archivo de salida
	hash := fmt.Sprintf("%x", md5.Sum([]byte(text+voice+pitchStr+speedStr)))
	outFile := filepath.Join("/tmp", hash+".wav")
	
	// Ejecutar balcon.exe via wine de forma acelerada
	cmd := exec.Command("wine", "balcon.exe",
		"-n", voice,
		"-t", taggedText,
		"-w", outFile)
		
	// Inyectar envs para saltar el driver de ALSA que traba a Wine por 1.5 segundos
	cmd.Env = append(os.Environ(), "WINEDEBUG=-all", "ALSA_CONFIG_PATH=/dev/null", "DISPLAY=:99")
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error ejecutando balcon.exe: %v. Output: %s", err, string(output))
		http.Error(w, "Error generando audio = "+err.Error()+"\n"+string(output), http.StatusInternalServerError)
		return
	}

	// Leer el archivo resultante final
	audioData, err := ioutil.ReadFile(outFile)
	if err != nil {
		log.Printf("Error leyendo archivo de audio: %v", err)
		http.Error(w, "Error leyendo WAV final", http.StatusInternalServerError)
		return
	}

	// Eliminar temporales
	_ = os.Remove(outFile)

	// Enviar el WAV al cliente
	w.Header().Set("Content-Type", "audio/wav")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(audioData)
}

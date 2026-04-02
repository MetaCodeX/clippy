# 📐 Especificación de Layout — Widget Clippy (V34)

> Documento de referencia para mantener y/o reimplementar el sistema de 
> posicionamiento de la burbuja de diálogo en cualquier otro entorno 
> (Go, Python, Qt, webview, etc.)

---

## 1. El GIF / Wrapper

| Propiedad         | Valor         |
|-------------------|---------------|
| Tamaño del wrapper | **160 × 160 px** |
| GIF usado         | `assets/gifs/idle_movimiento.gif` |
| El GIF llena      | 100% del wrapper (object-fit: contain) |
| Anclado en pantalla | `position: fixed`, centrado con `transform: translate(-50%, -50%)` |
| Clamping del drag | ±80px desde los bordes de la pantalla (mitadClip = 80) |

---

## 2. Sistema de 9 Cuadrantes (Espejo)

El wrapper se divide en zonas. Cuando Clippy entra en una zona, 
la burbuja salta al cuadrante **opuesto**.

```
Umbrales de activación:
  Horizontal: 35% de screen_width desde cada lado
  Vertical:   35% de screen_height desde cada lado
```

### Tabla de cuadrantes

| Clippy en...      | Clase CSS aplicada  | Burbuja aparece en... |
|-------------------|---------------------|-----------------------|
| Centro            | `pos-right`         | Derecha               |
| Borde izquierdo   | `pos-right`         | Derecha               |
| Borde derecho     | `pos-left`          | Izquierda             |
| Borde superior    | `pos-bottom`        | Abajo                 |
| Borde inferior    | `pos-top`           | Arriba                |
| Esquina top-left  | `pos-bottom-right`  | Abajo-derecha         |
| Esquina top-right | `pos-bottom-left`   | Abajo-izquierda       |
| Esquina bot-left  | `pos-top-right`     | Arriba-derecha        |
| Esquina bot-right | `pos-top-left`      | Arriba-izquierda      |

---

## 3. Offsets de la Burbuja (CSS position: absolute relativa al wrapper)

### Cardinales (4 lados rectos)

| Cuadrante    | CSS                                          | Gap real |
|--------------|----------------------------------------------|----------|
| `pos-top`    | `bottom: 168px; left: 50%; translateX(-50%)` | 8px      |
| `pos-bottom` | `top: 168px; left: 50%; translateX(-50%)`    | 8px      |
| `pos-left`   | `right: 168px; top: 50%; translateY(-50%)`   | 8px      |
| `pos-right`  | `left: 168px; top: 50%; translateY(-50%)`    | 8px      |

> **Gap cardinal = wrapper size + 8px = 160 + 8 = 168px**

### Esquinas (4 diagonales) ← **VALORES CALIBRADOS V34**

| Cuadrante          | CSS                            |
|--------------------|--------------------------------|
| `pos-bottom-right` | `top: 160px; left: 70px`       |
| `pos-bottom-left`  | `top: 160px; right: 70px`      |
| `pos-top-right`    | `bottom: 160px; left: 70px`    |
| `pos-top-left`     | `bottom: 160px; right: 70px`   |

```
Eje vertical (top/bottom): 160px  → burbuja queda justo debajo/arriba del GIF
Eje horizontal (left/right): 70px → flecha señala el cuerpo de Clippy
```

---

## 4. La Flecha (::before rotado 45°)

```
Tamaño del elemento: 14 × 14 px
Transform: rotate(45deg)  ← SIEMPRE en el base rule
Offset de posición: -8px  (la mitad del elemento + 1px de borde)
```

### Mapeo de bordes → dirección de la punta

Un cuadrado 14×14 rotado 45° CW en CSS (y↓):

| Esquina del cuadrado | Se convierte en punta | Bordes visibles              |
|----------------------|-----------------------|------------------------------|
| TL corner            | Punta ARRIBA ↑        | `border-top + border-left`   |
| TR corner            | Punta DERECHA →       | `border-top + border-right`  |
| BR corner            | Punta ABAJO ↓         | `border-right + border-bottom` |
| BL corner            | Punta IZQUIERDA ←     | `border-bottom + border-left` |

### Posición de la flecha por cuadrante

| Cuadrante          | Flecha apunta | position CSS                        | Bordes activos                    |
|--------------------|---------------|-------------------------------------|-----------------------------------|
| `pos-top`          | ↓ Abajo       | `bottom: -8px; left: 50%; ml:-7px`  | `border-right + border-bottom`    |
| `pos-bottom`       | ↑ Arriba      | `top: -8px; left: 50%; ml:-7px`     | `border-top + border-left`        |
| `pos-left`         | → Derecha     | `right: -8px; top: 50%; mt:-7px`    | `border-top + border-right`       |
| `pos-right`        | ← Izquierda   | `left: -8px; top: 50%; mt:-7px`     | `border-bottom + border-left`     |
| `pos-bottom-right` | ↑ (izq)       | `top: -8px; left: 8px`              | `border-top + border-left`        |
| `pos-bottom-left`  | ↑ (der)       | `top: -8px; right: 8px`             | `border-top + border-left`        |
| `pos-top-right`    | ↓ (izq)       | `bottom: -8px; left: 8px`           | `border-right + border-bottom`    |
| `pos-top-left`     | ↓ (der)       | `bottom: -8px; right: 8px`          | `border-right + border-bottom`    |

---

## 5. Tipografía y Estética

| Propiedad         | Valor                          |
|-------------------|--------------------------------|
| Fuente            | `TahomaXP` (local .otf) → fallback: `Tahoma, Segoe UI, Arial` |
| Archivo fuente    | `frontend/fonts/tahoma.otf`    |
| Font size burbuja | `22px`                         |
| Line-height       | `1.25`                         |
| Font-smooth       | `never` (pixel-art retro XP)   |
| Color fondo       | `#ffffe1` (amarillo clásico Clippy) |
| Border burbuja    | `1px solid #000`               |
| Border-radius     | `8px`                          |
| Box-shadow        | `2px 2px 0px rgba(0,0,0,0.4)` |

---

## 6. Compatibilidad con Go/Python (WebView)

**Sí, todo esto funciona perfectamente en el contexto Go+Python+WebView.**

El backend Go ya sirve el frontend estático:
```
/frontend/index.html
/frontend/styles.css
/frontend/fonts/tahoma.otf
/assets/gifs/idle_movimiento.gif
```

Para la versión final del widget embebido:
- **Go** sirve el HTTP server (ya implementado en `cmd/backend/main.go`)
- **WebView** (ej. `webview/webview` para Go, o `CEF` / `pywebview` para Python) 
  renderiza el `index.html` como ventana nativa sin bordes
- **Python** puede conectarse via HTTP al backend Go para enviar texto/emociones
- El CSS/JS funciona igual que en Chrome — el motor de renderizado es el mismo (Blink/WebKit)

### Stack completo previsto:
```
[Python Brain] → HTTP → [Go Backend] → sirve HTML+CSS+JS
                                      ↓
                              [WebView Window]
                              (ventana nativa sin bordes,
                               transparente, siempre encima)
```

El widget puede hacerse **sin borde de ventana** y con **fondo transparente** 
usando los flags del WebView, para que solo se vea el GIF y la burbuja flotando 
sobre el escritorio — exactamente como el Clippy original de Office 97.

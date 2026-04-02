import base64

# Leer la fuente TTF y convertirla a Base64
with open("fonts/tahoma.ttf", "rb") as font_file:
    font_b64 = base64.b64encode(font_file.read()).decode('utf-8')

# Plantilla CSS V8 (9 cuadrantes geométricos) - Fix de Fuente y Picos
css_template = """@font-face {
    font-family: 'TahomaXP';
    src: url('data:font/ttf;base64,""" + font_b64 + """') format('truetype');
    font-weight: normal;
    font-style: normal;
}

body {
    background-color: #ededed;
    font-family: 'TahomaXP', 'Tahoma', 'Segoe UI', Arial, sans-serif;
    margin: 0;
    overflow: hidden; 
    box-sizing: border-box;
}

*, *:before, *:after { box-sizing: inherit; }

.controls-panel {
    position: absolute; top: 20px; left: 20px;
    display: flex; flex-direction: column;
    background: #ffffff; padding: 20px; border-radius: 8px;
    box-shadow: 0 4px 10px rgba(0,0,0,0.1); z-index: 10;
}

.controls-panel h3 { margin: 0 0 15px 0; font-size: 16px; color: #333; }
#clippyInput {
    width: 250px; height: 120px; resize: vertical; padding: 10px;
    font-family: inherit; border: 1px solid #ccc; border-radius: 4px;
}

/* WIDGET FÍSICO */
.clippy-wrapper {
    position: fixed; top: 50vh; left: 50vw;
    width: 60px; height: 60px;
    transform: translate(-50%, -50%);
    user-select: none; z-index: 1000;
}

.clippy-character {
    font-size: 60px; line-height: 1; margin: 0; cursor: grab;
    z-index: 5; transition: transform 0.1s;
    filter: drop-shadow(2px 2px 2px rgba(0,0,0,0.3));
    text-align: center; width: 100%; height: 100%;
}
.clippy-character:active { cursor: grabbing; transform: scale(0.95); }

/* BURBUJA CLÁSICA SIN VARIABLES MATEMÁTICAS */
.clippy-bubble {
    position: absolute; 
    background-color: #ffffe1; border: 1px solid #000000;
    border-radius: 8px; padding: 12px 18px; color: #000000;
    font-size: 14px; line-height: 1.5;
    box-shadow: 2px 2px 0px rgba(0, 0, 0, 0.4); 
    width: max-content; max-width: 320px; min-width: 80px; 
    word-wrap: break-word; z-index: 4;
}
.clippy-bubble p { margin: 0 0 10px 0; }
.clippy-bubble p:last-child { margin: 0; }
.clippy-bubble::before {
    content: ""; position: absolute; 
    width: 14px; height: 14px; background-color: #ffffe1; z-index: -1; 
}

/* =======================================
   ORIENTACIONES 8-CUADRANTES (ESPEJO)
======================================= */

/* 1. POS-TOP: Clip Abajo -> Caja Arriba */
.clippy-wrapper.pos-top .clippy-bubble {
    bottom: calc(100% + 15px); left: 50%; transform: translateX(-50%);
}
.clippy-wrapper.pos-top .clippy-bubble::before {
    bottom: -8px; left: 50%; margin-left: -7px;
    transform: rotate(45deg); border-bottom: 1px solid #000; border-right: 1px solid #000;
}

/* 2. POS-BOTTOM: Clip Arriba -> Caja Abajo */
.clippy-wrapper.pos-bottom .clippy-bubble {
    top: calc(100% + 15px); left: 50%; transform: translateX(-50%);
}
.clippy-wrapper.pos-bottom .clippy-bubble::before {
    top: -8px; left: 50%; margin-left: -7px;
    transform: rotate(45deg); border-top: 1px solid #000; border-left: 1px solid #000;
}

/* 3. POS-LEFT: Clip Derecha -> Caja Izquierda */
.clippy-wrapper.pos-left .clippy-bubble {
    right: calc(100% + 15px); top: 50%; transform: translateY(-50%);
}
.clippy-wrapper.pos-left .clippy-bubble::before {
    right: -8px; top: 50%; margin-top: -7px;
    transform: rotate(45deg); border-top: 1px solid #000; border-right: 1px solid #000;
}

/* 4. POS-RIGHT: Clip Izquierda -> Caja Derecha */
.clippy-wrapper.pos-right .clippy-bubble {
    left: calc(100% + 15px); top: 50%; transform: translateY(-50%);
}
.clippy-wrapper.pos-right .clippy-bubble::before {
    left: -8px; top: 50%; margin-top: -7px;
    transform: rotate(45deg); border-bottom: 1px solid #000; border-left: 1px solid #000;
}

/* 5. POS-BOTTOM-RIGHT: Clip Top-Left -> Caja Abajo-Derecha */
.clippy-wrapper.pos-bottom-right .clippy-bubble {
    top: calc(100% + 15px); left: calc(100% + 15px);
}
.clippy-wrapper.pos-bottom-right .clippy-bubble::before {
    top: 15px; left: -8px; z-index: 10;
    transform: rotate(45deg); border-top: 1px solid #000; border-left: 1px solid #000;
}

/* 6. POS-BOTTOM-LEFT: Clip Top-Right -> Caja Abajo-Izquierda */
.clippy-wrapper.pos-bottom-left .clippy-bubble {
    top: calc(100% + 15px); right: calc(100% + 15px);
}
.clippy-wrapper.pos-bottom-left .clippy-bubble::before {
    top: 15px; right: -8px; z-index: 10;
    transform: rotate(45deg); border-top: 1px solid #000; border-right: 1px solid #000;
}

/* 7. POS-TOP-RIGHT: Clip Bottom-Left -> Caja Arriba-Derecha */
.clippy-wrapper.pos-top-right .clippy-bubble {
    bottom: calc(100% + 15px); left: calc(100% + 15px);
}
.clippy-wrapper.pos-top-right .clippy-bubble::before {
    bottom: 15px; left: -8px; z-index: 10;
    transform: rotate(45deg); border-bottom: 1px solid #000; border-left: 1px solid #000;
}

/* 8. POS-TOP-LEFT: Clip Bottom-Right -> Caja Arriba-Izquierda */
.clippy-wrapper.pos-top-left .clippy-bubble {
    bottom: calc(100% + 15px); right: calc(100% + 15px);
}
.clippy-wrapper.pos-top-left .clippy-bubble::before {
    bottom: 15px; right: -8px; z-index: 10;
    transform: rotate(45deg); border-bottom: 1px solid #000; border-right: 1px solid #000;
}
"""

with open("styles.css", "w") as css_file:
    css_file.write(css_template)

print("styles.css generado con exito y cargado con Base64 TTF.")

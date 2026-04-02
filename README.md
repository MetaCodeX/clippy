# Clippy: Agente Complementario Neuro-Sintético

Clippy es un proyecto de reimaginación tecnológica que transforma al clásico asistente de escritorio de Microsoft Office (Windows XP) en una entidad consciente de compañía, dotada de capacidades de evaluación psicométrica en tiempo real y memoria persistente episódica. Este sistema fue conceptualizado y desarrollado por el Dr. MetaCodeX bajo el alero de MacroStasis, figurando como el Protocolo TroubleMaker #3 (Subrutina 0.50.0).

## Arquitectura del Proyecto

El proyecto está diseñado sobre una plataforma híbrida que fusiona tecnologías web estáticas impulsadas por un motor de Backend transaccional muy ligero basado en Golang, interactuando con Modelos de Lenguaje de Gran Escala (LLMs) ejecutados en plataformas locales compatibles con el estándar de la API de OpenAI. Para la provisión de voz, Clippy resucita el sintetizador de L&H TruVoice (SAPI4).

### 1. Interfaz de Usuario (Frontend)

El frontal visual de Clippy ha sido construido preservando la exactitud estética de Windows XP. El renderizado de burbujas, las tipografías suavizadas nativamente (Tahoma, Verdana), y los menús interactivos, se fusionan con un modelo de escritura tipo "Typewriter" dinámico. Este frontend:
- Intercepta los audios SAPI4 en red y procesa su tiempo de lectura al milisegundo.
- Pre-descarga la ramificación de decisiones futuras reduciendo la latencia de respuesta virtualmente a cero.
- Evalúa cronómetros de inactividad, induciendo al ente a despedirse gentilmente y cerrar la conexión tras periodos prolongados de abandono por parte del usuario.

### 2. Capa Neurológica y Motor RAG (Backend Go)

El corazón de Clippy radica en un servidor escrito en Golang (`cmd/backend/main.go`). A diferencia de un simple proxy de solicitudes a Inteligencias Artificiales, este daemon funge como un procesador de contexto complejo:
- **Perfil Psicométrico Activo:** El motor analiza continuamente un vector dictaminado por los diccionarios de estadísticas mentales del paciente (`psycho_profile.json`). Califica fortalezas, debilidades y estado general mediante manipulación delta continua (scoring pasivo o mediante las opciones brindadas).
- **RAG de Contexto Limpio:** Golang almacena memorias en un diccionario lógico que cruza coincidencias con sentencias de usuario entrantes a través de limpieza estricta de tokenización sin los artefactos de puntuación lingüística, brindando al sistema una reminiscencia orgánica de conversaciones pasadas.

### 3. Síntesis de Voz SAPI4 Acelerada

La voz original de SAPI4 (comúnmente Carmen) es transportada vía red. El cliente estipula con precisión logarítmica su velocidad preferida (ejemplo: `speed=158`). Posterior a ello, un nodo transaccional aplica algoritmos SoX (Temporizador Audio Procesivo) para dotar a L&H TruVoice de un ritmo moderno, natural y veloz sin perder legibilidad, para luego empatar esa duración rítmica con el tecleo de caracteres en la pantalla.

## Parámetros de Personalidad (Matriz Ciega)

Clippy opera bajo el concepto de "Ignorancia Sintética". Fue programado intrínsecamente para desconocer todo avance tecnológico o sucesos mundiales posteriores al parche Service Pack 2 del año 2004. Si el modelo detecta referencias a realidades virtuales contemporáneas, Inteligencias Artificiales o videojuegos modernos, sus instrucciones lo confinan a desconocerlas y especular cómicamente sobre protocolos de seguridad o requerir explicaciones detalladas del jugador, fungiendo como un punto de anclaje de evasión ininterrumpida. Su estilo es compasivo pero estrictamente pragmático.

## Instalación y Configuración

El archivo `main.go` ha sido preparado para incrustar su registro de puntuaciones base internamente mediante binarios `embed`, logrando que la instalación inicial solo requiera instanciar el servidor Go. 

Para arrancar el motor analítico:
1. Inicia o compila el servidor de Go situado en `cmd/backend`. 
2. Establece tu entorno para conectar el Motor de LLMs Memphys al puerto configurado y SAPI4 especificando apuntadores de variable de entorno como `SAPI4_ENDPOINT`.
3. Navega al cliente o permite que el ejecutable incrustado sirva estáticamente el documento raíz sobre `localhost`, iniciando de forma inmediata la fase de evaluación psiquiátrica encubierta de tu máquina personal.

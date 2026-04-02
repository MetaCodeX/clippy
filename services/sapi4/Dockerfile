FROM --platform=linux/386 debian:bullseye-slim

ENV DEBIAN_FRONTEND=noninteractive
ENV WINEARCH=win32
ENV WINEPREFIX=/root/.wine

# Instalar dependencias esenciales
RUN apt-get update && apt-get install -y \
    wine \
    xvfb \
    wget \
    unzip \
    golang \
    p7zip-full \
    && rm -rf /var/lib/apt/lists/*

# Configurar Wine Headless
RUN xvfb-run --server-args="-screen 0 1024x768x24" wineboot --init && sleep 5

# Crear directorio de trabajo
WORKDIR /app

# Descargar y configurar el motor LTS SAPI4 (Spchapi)
RUN wget -q "https://raw.githubusercontent.com/TETYYS/SAPI4/master/spchapi.exe" -O spchapi.exe && \
    (xvfb-run --server-args="-screen 0 1024x768x24" wine spchapi.exe /q || true) && \
    sleep 5

# Descargar e instalar las voces TruVoice (Español) desde Archive.org
RUN wget -q "https://archive.org/download/TextToSpeechVoices/mike9012.neocities.org/Lernout%20and%20Hauspie%20Truvoice%20Voices.7z" -O voices.7z && \
    7z e voices.7z "*lhttsspe.exe" -r && \
    (xvfb-run --server-args="-screen 0 1024x768x24" wine lhttsspe.exe /q || true) && \
    sleep 5 && \
    rm voices.7z lhttsspe.exe spchapi.exe

# Descargar e instalar Balabolka Console (balcon.exe) como wrapper de SAPI4
RUN wget -q "https://community.chocolatey.org/api/v2/package/balcon" -O bconsole.zip && \
    unzip -j bconsole.zip "tools/*" -d . && \
    rm -rf bconsole.zip

# Añadir los archivos del servidor Go
COPY . /app/
RUN go build -o server main.go

EXPOSE 8081

# Script de arranque: Xvfb permanente y wineserver persistente para latencia de ~300ms.
CMD sh -c "Xvfb :99 -screen 0 1024x768x24 & sleep 1 && export DISPLAY=:99 && wineserver -p & sleep 1 && ./server"

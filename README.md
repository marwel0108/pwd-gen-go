# 🟩 Generador de Contraseñas Seguras (CLI) pwd-gen

**Foco:** Sintaxis estricta, manejo de memoria básico (`slices`/`runes`) y control de flujo sin excepciones.

### 📋 Checklist de Requerimientos

* [X] **[RF-1.1]** El sistema debe ser una herramienta CLI funcional ejecutable mediante la terminal.
* [X] **[RF-1.2]** Debe aceptar las banderas: `--length` (int, por defecto 12), `--no-digits` (bool), `--no-special` (bool).
* [X] **[RF-1.3]** Validar que `--length` esté en un rango de 8 a 64; abortar en `stderr` si no cumple.
* [X] **[RF-1.4]** El output limpio debe enviarse a `stdout` para permitir redirección de comandos (`> out.txt`).

### 🛠️ Objetivos Técnicos (Conceptos a dominar)

* [X] Inicialización del entorno con `go mod init`.
* [X] Implementación de la librería externa `github.com/spf13/cobra`.
* [X] Uso de un slice de runas (`[]rune`) en lugar de strings crudos para asegurar soporte UTF-8.
* [X] Retorno explícito de errores con el patrón de firmas `func (args) (string, error)`.
* [X] Definir un makefile simple para hacer el proceso de `build` más rápido. 

### 🏁 Definition of Done (DoD)

> El comando `pwd-gen --length 16 --no-digits` genera una contraseña válida en la terminal y el código no contiene un solo bloque simulado de try/catch.
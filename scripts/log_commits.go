package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

/*
"fmt": Aquest paquet proporciona funcions per a la formatació d'entrada i sortida, similar a printf i scanf en C. S'utilitza habitualment per imprimir missatges a la consola (amb funcions com fmt.Println, fmt.Printf) i per llegir dades de l'entrada estàndard.

"os": Aquest paquet ofereix una interfície independent del sistema operatiu per a funcionalitats del sistema operatiu. Permet interactuar amb el sistema subjacent, com ara manipular fitxers i directoris, obtenir informació de l'entorn, i gestionar processos. Algunes funcions comunes inclouen os.Args (arguments de la línia de comanda), os.Getenv (variables d'entorn), os.Mkdir (crear directoris), os.Remove (eliminar fitxers o directoris), etc.

"os/exec": Aquest subpaquet del paquet os proporciona funcionalitats per executar comandaments externs al sistema operatiu. Permet al teu programa Go executar altres programes o scripts i interactuar amb la seva entrada, sortida i errors. La funció principal aquí és exec.Command, que crea un objecte per representar un comandament a executar.

"path/filepath": Aquest paquet implementa utilitats per manipular noms de fitxers de manera compatible amb les convencions del sistema operatiu. Proporciona funcions per analitzar, construir i manipular camins de fitxers, tenint en compte les diferències entre sistemes operatius (per exemple, l'ús de / o \ com a separador de directoris). Funcions comunes inclouen filepath.Join (unir components de camins), filepath.Base (obtenir el nom base d'un camí), filepath.Dir (obtenir el directori d'un camí), etc.

"time": Aquest paquet proporciona funcionalitats per treballar amb el temps, incloent la gestió d'instants en el temps, duracions, rellotges i temporitzadors. Permet realitzar operacions com obtenir l'hora actual (time.Now), esperar durant un cert període (time.Sleep), formatar i analitzar dates i hores, i realitzar càlculs amb duracions.
*/

func main() {
	// Obtener los últimos 3 commits
	cmd := exec.Command("git", "log", "-n", "3", "--pretty=format:%h - %an, %ar : %s")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error ejecutando git log: %v\n", err)
		os.Exit(1)
	}

	// Crear directorio log/ en la raíz del repo (no en scripts/)
	logDir := filepath.Join("..", "log")  // 👈 Cambio clave
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		// Li donem els permisos al directori
		err = os.Mkdir(logDir, 0755)
		if err != nil {
			fmt.Printf("Error creando directorio %s: %v\n", logDir, err)
			os.Exit(1)
		}
	}

	// Generar nombre de archivo, amb una "mascara"
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	logFile := filepath.Join(logDir, fmt.Sprintf("commits_%s.txt", currentTime))

	// Escribir archivo
	content := fmt.Sprintf("Últimos 3 commits del repositorio:\n\n%s", string(out))
	err = os.WriteFile(logFile, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error escribiendo en %s: %v\n", logFile, err)
		os.Exit(1)
	}

	fmt.Printf("Archivo de log creado en: %s\n", logFile)
}

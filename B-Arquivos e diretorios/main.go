package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	createArchive()
	createDirectory()
	rename()
	copyToSpecificLocation()
	moveFile()
	metadata()
	deleting()
	truncade()
	appendText()
	changeFilePermissions()
}

func createArchive() {
	emptyFile, err := os.Create("texto.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}

func createDirectory() {
	_, err := os.Stat("pastaCriada")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("pastaCriada", 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
}

func rename() {
	oldName := "texto.txt"
	newName := "renomeada.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

func copyToSpecificLocation() {

	sourceFile, err := os.Open("C:/Users/Igor/Desktop/Escola/Golang do Zero/B-Arquivos e diretorios/renomeada.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("C:/Users/Igor/Desktop/Escola/Golang do Zero/B-Arquivos e diretorios/pastaCriada/renomeada.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)
}

func moveFile() {
	oldLocation := "C:/Users/Igor/Desktop/Escola/Golang do Zero/B-Arquivos e diretorios/renomeada.txt"
	newLocation := "C:/Users/Igor/Desktop/Escola/Golang do Zero/B-Arquivos e diretorios/pastaCriada/movida.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func metadata() {
	fileStat, err := os.Stat("empty.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()
}

func deleting() {
	err := os.Remove("C:/Users/Igor/Desktop/Escola/Golang do Zero/B-Arquivos e diretorios/pastaCriada/renomeada.txt")
	if err != nil {
		log.Fatal(err)
	}
}

// Reduz o tamanho do arquivo
func truncade() {
	err := os.Truncate("empty.txt", 100)

	if err != nil {
		log.Fatal(err)
	}
}

func appendText() {
	message := "Add this content at end"
	filename := "empty.txt"

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", message)
}

func changeFilePermissions() {
	// Test File existence.
	_, err := os.Stat("empty.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File exist.")

	// Change permissions Linux.
	err = os.Chmod("empty.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	// Change file ownership.
	err = os.Chown("empty.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Change file timestamps.
	addOneDayFromNow := time.Now().Add(24 * time.Hour)
	lastAccessTime := addOneDayFromNow
	lastModifyTime := addOneDayFromNow
	err = os.Chtimes("empty.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}

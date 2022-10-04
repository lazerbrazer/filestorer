package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//declare folder path. Currently, just 'output' folder
	folderPath := "output"

	//variable with some dummy data
	someTextToSave := `{"data":"resting"}`

	//check if the folder does not exist
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {

		//if the folder really doesnt exist create it usign mkdir()
		//q?:ModePerm=511 read the perm codes later
		errFolder := os.Mkdir(folderPath, os.ModePerm)
		if errFolder != nil {
			log.Fatalf("Problem making folder %v", errFolder)
		}
	}

	//create a file with pass parameter of file location
	filePtr, err := os.Create("output/scan_result.json")
	//immieately declare defer command that is postponed, it would run in the end (LIFO)
	defer filePtr.Close()
	if err != nil {
		log.Fatalf("Problem creating file file: %v", err)
	}

	//write dummy data from variable to file. Transform string data of variable
	//slice of bytes ([]byte), wrapping variable within slice of byte.
	//because it required for Write() method
	//Write() method requires bytes data for input, not string
	_, errOfSaving := filePtr.Write([]byte(someTextToSave))
	if errOfSaving != nil {
		log.Fatalf("problem saving file: %v", errOfSaving)
	}

	//that's it
	fmt.Print("Finalle!!")
}

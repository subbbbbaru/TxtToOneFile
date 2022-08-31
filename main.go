package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func main() {

	checkDirPtr := flag.String("dir", ".", "Find file from directory. By default, the search occurs in the current directory.")
	typeFilePtr := flag.String("type", "", `Find file with some mask. Example ("*.md"). By default is not set.`)
	findFilePtr := flag.String("print", "y", `Only find file in directory and print it on console. Type y/n. By default is set "y".`)
	outFileNamePtr := flag.String("out", "", `Output file name. If not set then print to console.`)
	// flag.Set("dir", ".")
	// flag.Set("type", "*.dart")
	flag.Parse()

	if *typeFilePtr == "" {
		log.Fatal("Error type file mask!!!")
	}

	files, err := WalkMatch(*checkDirPtr, *typeFilePtr)
	if err != nil {
		log.Fatal()
	}

	if len(strings.Fields(*outFileNamePtr)) == 0 {

		fmt.Println("Output file name not set!")

		if *findFilePtr == "n" {
			log.Fatalln(`Wrong combine flag. You must set name of file or print="y"!`)
		}

		for _, file := range files {
			if *findFilePtr == "y" {
				fi, _ := os.Stat(file)

				fmt.Println(file + " " + fi.Mode().String())

			}
<<<<<<< HEAD
=======
			contents, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(contents))
>>>>>>> 307d8de9d1e46badae19a65f683c451db84a08f8
		}

	} else {
		f, errorCreate := os.OpenFile(*outFileNamePtr, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if errorCreate != nil {
			log.Fatalln(errorCreate)

		}
		defer f.Close()

		for _, file := range files {

			if *findFilePtr == "y" {
				fmt.Println(file)
				contents, err := ioutil.ReadFile(file)
				if err != nil {
					log.Fatalln(err)
				}

				fmt.Fprintln(f, "\n"+file+"\n"+string(contents))
			} else if *findFilePtr == "n" {

				// fileStat, _ := os.Stat(f.Name())
				// if fileStat.Size() > int64(1024*1024) {

				// 	fmt.Println(fileStat.Size())
				// }

				contents, err := ioutil.ReadFile(file)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Fprintln(f, "\n"+file+"\n"+string(contents))
			}
		}
	}

}

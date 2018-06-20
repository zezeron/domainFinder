package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

type IoController struct {
	InPut  string
	OutPut string
}

func (io IoController) GetFolderList(val string) ([]string, error) {
	var returnValue = []string{}

	files, err := ioutil.ReadDir(val)
	// ./ nerdeysen onu arar
	//if err != nil {
	//	//fmt.Println("Dosya Bulunamadi")
	//	//fmt.Println("soyledigin ", files)
	//}

	for _, f := range files {

		returnValue = append(returnValue, f.Name())

	}

	return returnValue, err
}
func (io IoController) WriteFile(val string, row []string) {

	f, _ := os.OpenFile(val, os.O_APPEND|os.O_WRONLY, 0644)

	for _, ttt := range row {

		f.WriteString(ttt + "\n")

	}
	defer f.Close()
}




// buraci dic create eder privatetir

func (io IoController) createDic(val int) {
	//
	if val == 1 {
		os.MkdirAll(io.InPut, os.ModePerm)
	} else {
		os.MkdirAll(io.OutPut, os.ModePerm)
	}
}
func (io IoController) testCreateFile(val string, lval string, nm string, ny []string) {
	//	var returnVal bool

	if val == "yes" || val == "" {

		var file, err = os.Create("./" + lval + "/" + nm)
		if err != nil {
			fmt.Println("Dosya Olusturulamadi...")
		}
		defer file.Close()

		io.WriteFile("./"+lval+"/"+nm, ny)

		//returnVal = true

	} else {
		fmt.Println("lutfen " + lval + " folderiniza domainler yazan bir dosya yerlestiriniz")
		//returnVal = false
	}
	//return returnVal
}

func (io IoController) GetFolderListOrCreateFolder(val int) []string {
	var returnVal = []string{}
	var err error
	var path string
	var pathInt int

	if val == 1 {

		path = io.InPut
		pathInt = 1

	} else {
		path = io.OutPut
		pathInt = 2

	}

	returnVal, err = io.GetFolderList(path)

	//fmt.Println("========", returnVal)

	if err != nil {

		io.createDic(pathInt)
	}
	return returnVal

}

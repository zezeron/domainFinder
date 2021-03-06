package main

import (
	"strings"
	"os"
	"fmt"
	"bufio"
)

type filesScanner struct {
	folderList []string
}

func (fs filesScanner) HaydiScan() int {
	var scanCount = 0

	for _, scanRow := range fs.folderList {
		t := strings.TrimSpace(scanRow)
		e := strings.HasSuffix(t, ".txt")
		if e == true {
			scanCount++
		}
		//print(e)
		//println(t)
		//println(t)
	}

	return scanCount
}

type FileCheckandScan struct {
	//extension strings
	path string
}

//
func (fc FileCheckandScan) ExtensionControlAndScan(txtNames []string) []string {
	//fmt.Println("=========ssssss",txtNames)
	//[domains.txt test.txt]
	var returnVal = []string{}

	for _, txtName := range txtNames {

		txtName = fc.path + "/" + txtName
		b := fc.txtScanner(txtName)

		for _, a := range b {
			//fmt.Println("sasasasasassa",b)
			//fmt.Println("iiiiiiii", a)

			returnVal = append(returnVal, a)
		}

	}
	return returnVal
}

// buraya txt name leri vericem oda bana icindekileri array donucek
func (fc FileCheckandScan) txtScanner(val string) []string {

	var returnVal = []string{}

	OpenedFile, error := os.Open(val)
	if error != nil {
		fmt.Println("Dosya Acilmadi")
	}
	defer OpenedFile.Close()
	fileScanner := bufio.NewScanner(OpenedFile)

	for fileScanner.Scan() {

		b := fileScanner.Text()

		a := strings.Split(b, " ")

		for _, c := range a {
			//fmt.Println("=======",c)
			e := strings.TrimSpace(c)
			if !strings.ContainsAny(e, "@?") && strings.ContainsAny(e, ".") && !strings.HasPrefix(e, ".") {
				returnVal = append(returnVal, e)

			}

		}
		// 👉🏻🤟🏻
		// burada  ins.tx
		//              domains.txtptbLustr,Users/tolgaozen/_GOPROJECT/src/domainFinder/
		//   domains.txtptbNustr
		//                     domains.txt
		//                                 @? @ @ @
		//                                         DSDB ` @ @ @      bunun cikmasini engelledim
		//fmt.Println("bbbbbbb", b)

	}
	return returnVal

}

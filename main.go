package main

import (
	"fmt"
	"time"
	"strings"
)


var inputFolderPath string
var outputFolderPath string
var standartPath string


func PathlerVarKacTxtvar() []string {

	var returnval = []string{}

	ioa := IoController{}
	folderPathFiller()

	FolderControl, _ := ioa.GetFolderList(inputFolderPath)

	Foldercheck := 0

	for _, inputcheck := range FolderControl {

		if strings.HasPrefix(inputcheck, ".") {

		} else {

			Foldercheck++
		}

	}

	//gopre.Pre(len(inputFolderControl) == 0)
	if Foldercheck == 0 {
		//burada input folder icerisine test dosyasi sorulacak
		var answer string
		fmt.Println("Input Folderinizda herhangi bir dosya yok ornek bir dosya olusturmamizi ister misiniz? ")
		fmt.Scanln(&answer)
		ioa.testCreateFile(answer, inputName, "tester.txt", []string{"deneme.com", "deneme1.com", "deneme2.com", "ojnjkdjfidjfi.com"})

	}

	inputFolderControl2, _ := ioa.GetFolderList(inputFolderPath)

	Foldercheck2 := 0

	for _, inputcheck2 := range inputFolderControl2 {

		if strings.HasPrefix(inputcheck2, ".") {

		} else {

			// en son haline bakilsin diye actim
			//gopre.Pre(inputFolderPath)
			// pathler var demek
			//adam kullanmissa buradan devam ediyor
			returnval = append(returnval, inputcheck2)
			Foldercheck2++
		}

	}

	virusScanProgram := filesScanner{inputFolderControl2[1:]}
	virusScanProgram.HaydiScan()
	// txt sayisi lazim olursa kullan

	return returnval
}

func domainFor(UygunDomains []string) {

	var ty IoController
	for i := 0; i < len(UygunDomains); i++ {
		//for _, uygundomain := range UygunDomains {

		//a := time.Date(2018, time.May, 5, 23, 55, 0, 0, time.Local)
		tn := time.Now()
		//ts := tn.String()
		//fmt.Println(ts)
		a := tn.Format("01-02-2006")
		// bu splitten daha iyi time konusunda iki kere cikmasinin sebebi iki kere donmesi
		//fmt.Println(a)
		//Now := strings.Split(ts, " ")
		//fmt.Println(Now[0])

		ty.testCreateFile("yes", outputName, a+".txt", []string{"UYGUN DOMAINLER LISTESI (" + a + ")"})
		ty.WriteFile(outputFolderPath+"/"+a+".txt", UygunDomains)

		//fmt.Printf("(%v) %v ye eklendi\n", uygundomain, outputName+a+".txt")
	}

}

func main() {
	flagInitialize()


	returnIsFolder := checkInputOutputFolder()
	//fmt.Println(returnIsFolder)
	if returnIsFolder == true {
		fmt.Println("Tekrar Hosgeldiniz")
		//folderPathFiller()
		//PathlerVarInputtaKacTxtvar()
		//hasTxt :=
		// burada inputta kac tane txt var o yazar

		//gopre.Pre(hasTxt)

	} else {

		fmt.Println("Pathinizi giriniz:[girmezseniz bulundugunuz dizin kullanilacaktir]")
		fmt.Scanln(&standartPath)

		folderPathFiller()

		io := IoController{InPut: inputFolderPath, OutPut: outputFolderPath}

		io.GetFolderListOrCreateFolder(1)
		io.GetFolderListOrCreateFolder(2)

		//fmt.Println(inPut, outPut)

	}

	inputtxtNameArray := PathlerVarKacTxtvar()
	//fmt.Println("gggggg", inputtxtNameArray)
	fc := FileCheckandScan{inputFolderPath}
	a := fc.ExtensionControlAndScan(inputtxtNameArray)

	if CountVal == 50 {
		fmt.Println("bir Count belirleyiniz:")
		fmt.Scanln(&CountVal)

	}

	dm := DomainQuest{a, CountVal}

	dm.DomainQuestioningCreateAndWriteFolder()

	//domainFor(UygunDomains)

	// tol := []string{"metromedya"}
	//domainFor(tol)

	//go domainFor(UygunDomains)

	//fmt.Println(";;;;;;;;;;;;",ty)

}

//fmt.Println("======", UygunDomains)
//fmt.Println("======", a) = [hedef.in wehgwyrwyerewfwfb.com metromedya.com deneme.com deneme1.com]
//var ask string
//fmt.Println("Path i kendiniz mi girmek istersiniz yoksa bulundugunuz yerde biz mi olusturalim? (me or ... )")
//fmt.Scanln(&ask)
//fmt.Println("=======",folder)

//if ask == "me" {

//"/Users/tolgaozen/_GOPROJECT/src/domainFinder"

//	if len(inputFolderPath) == 0 {
//		//flagten gelebilirdi
//		inputFolderPath = standartPath + "/" + inputName
//	}
//
//	if len(outputFolderPath) == 0 {
//
//		outputFolderPath = standartPath + "/" + outputName
//	}
//
//	//} else {
//
//	if len(inputFolderPath) == 0 {
//
//		inputFolderPath = "./" + inputName
//	}
//
//	if len(outputFolderPath) == 0 {
//
//		outputFolderPath = "./" + outputName
//	}
//	//}
//

//if len(inputFolderPath) == 0 && len(outputFolderPath) == 0 {
//
//}
//fc := FileCheckandScan{".txt"}
//
//inputString := strings.Join(inPut, " ")
////fmt.Println("===========",outputString)
//
//otst := strings.Count(inputString, "test.txt")
//
//if otst == 0 {
//
//	var do string
//
//	fmt.Println("Input a ornek dosya olusturup yazdirmamizi isterminiz ? (yes or ... )")
//
//	fmt.Scanln(&do)
//
//	io.createFile(do)
//}
//
//domainsArray := fc.ExtensionControlAndScan(inPut, inputFolderPath)
//
////fmt.Println("=======", domainsArray)
//
//dm := DomainQuest{domainsArray, CountVal}
//uygunDomains := dm.DomainQuestioning()
//
//for _, uygunDomain := range uygunDomains {
//
//	for _, val := range outPut {
//
//		if strings.HasSuffix(val, fc.extension) {
//			val = "/" + val
//			val = outputFolderPath + val
//
//			io.WriteFile(val, uygunDomain)
//
//		}
//	}
//
//}

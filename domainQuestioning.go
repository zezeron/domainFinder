package main

import (
	"strings"
	"fmt"
	"os/exec"
	"matrixPrint"
)

type DomainQuest struct {
	domains []string
	count   int
}

func (dm DomainQuest) DomainQuestioningCreateAndWriteFolder() {
	var returnVal = []string{}

	var val = []string{}

	for _, row := range dm.domains {

		// splitting head => g++ parts => rest of the command
		parts := strings.Fields(row)

		// [dfsdsf.com]

		head := "whois"

		//parts = parts[1:len(parts)]

		out, err := exec.Command(head, parts...).Output()
		//  whois  , dfsdsf.com

		if err != nil {
			fmt.Printf("%s", err)

		}

		val := append(val, string(out))
		aVall := dm.lookCountAndPrint(val, row)

		for _, vall := range aVall {

			returnVal = append(returnVal, vall)

			domainFor(returnVal)
		}

	}
	//return returnVal
}
func (dm DomainQuest) lookCountAndPrint(searchResults []string, row string) []string {

	var uygunReturnVal = []string{}

	for _, searchResult := range searchResults {

		findMatch := strings.Count(searchResult, "No match for domain")

		if findMatch >= 1 {

			if len(row) >= dm.count {
				//mRow := row + "...Uygun (uzun domain listeye eklenmeyecek) ...❌"
				spaceWrite(row, "✔︎ count:❌", 100)
				//fmt.Println(row, "...Uygun (uzun domain listeye eklenmeyecek) ...❌")
				//matrix.MPrint(mRow)
			} else {

				spaceWrite(row, "✔︎", 100)
				//matrix.MPrint(mRowU)
				//fmt.Println(row, "...Uygun listeye ekleniyor ...🔥︎")
				uygunReturnVal = append(uygunReturnVal, row)
			}

		} else {

			if len(row) >= dm.count {

				spaceWrite(row, "❌ count:❌", 100)
				//matrix.MPrint(mRowD)
				//fmt.Println(row, "...❌ (uzun domain)")

			} else {

				//mRowDD := row + "...❌ "
				spaceWrite(row, "❌ ", 100)
				//matrix.MPrint(mRowDD)
				//fmt.Println(row, "...❌ ")
				//if len(row) <= 8 {
				//
				//	tSReturnVal = append(tSReturnVal, row)
				//
			}
		}
	}

	return uygunReturnVal

}
func spaceWrite(row string, figure string, spaceNum int) {

	a := strings.Split(figure, "")
	b := strings.Split(row, "")

	figureSayi := 0

	for i := 0; i < len(a); i++ {
		//fmt.Print("======",i)
		figureSayi++
	}

	//for i, _ := range a {
	//
	//	figureSayi++
	//
	//	fmt.Sprint(i)
	//
	//	//sayi = sayi + 1
	//
	//}

	//fmt.Println("==", figureSayi)

	rowSayi := 0

	for i := 0; i < len(b); i++ {
		//fmt.Print("======",i)
		rowSayi++
	}

	//for i, _ := range b {
	//
	//	rowSayi++
	//
	//	fmt.Sprint(i)
	//
	//}

	//if rowSayi > spaceNum{
	//
	//
	//
	//
	//
	//
	//
	//
	//}

	//fmt.Println("======", rowSayi)

	toplam := figureSayi + rowSayi

	ppp := strings.Repeat(" ", spaceNum-toplam)

	//smartBlank := row + ppp + figure

	matrix.MPrint(row)
	fmt.Print(ppp)
	matrix.MPrint(figure + "\n")
}

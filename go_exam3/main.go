package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"go-tutorial/go_exam3/api"
	"log"
	"sync"
)

func main() {

	idRange := flag.Int("id_range", 0, "")
	empId := flag.String("employee_id", "", "")

	flag.Parse()

	if *idRange == 0 {
		log.Fatalf("error id_range is require.")
	}

	if *idRange > 500 {
		log.Fatalf("error id_range must less or equal than 500")
	}

	if *empId == "" {
		log.Fatalf("error employee_id is require.")
	}

	var file *xlsx.File
	var sheet *xlsx.Sheet

	file = xlsx.NewFile()
	sheet, err := file.AddSheet("gobasic_exam3")
	if err != nil {
		log.Fatalln("error new sheet", err)
	}

	st := sheet.AddRow()
	st.AddCell().SetValue("EmployeeId")
	st.AddCell().SetValue("RoutineId")
	st.AddCell().SetValue("ID")
	st.AddCell().SetValue("AlbumId")
	st.AddCell().SetValue("Title")
	st.AddCell().SetValue("Url")
	st.AddCell().SetValue("ThumbnailUrl")
	st.AddCell().SetValue("Error")

	phChan := make(chan api.Photo)
	defer close(phChan)

	var wg sync.WaitGroup
	wg.Add(*idRange)

	go func() {
		for v := range phChan {
			if v.Error != nil {
				logrus.Errorln("call ", v.ApiUrl, " error ", v.Error)
				wg.Done()
			}
			st := sheet.AddRow()
			st.AddCell().SetValue(*empId)
			st.AddCell().SetValue(v.RoutineId)
			st.AddCell().SetValue(v.Id)
			st.AddCell().SetValue(v.AlbumId)
			st.AddCell().SetValue(v.Title)
			st.AddCell().SetValue(v.Url)
			st.AddCell().SetValue(v.ThumbnailUrl)
			st.AddCell().SetValue(v.Error)
			wg.Done()
		}
	}()

	for i := 0; i < *idRange; i++ {
		go api.CallPhotoApi(fmt.Sprintf("%s%d", "https://jsonplaceholder.typicode.com/photos/", i+1), fmt.Sprintf("%d", i+1), phChan)
	}

	wg.Wait()
	err = file.Save(fmt.Sprintf("%s.xlsx", *empId))
	if err != nil {
		fmt.Printf(err.Error())
	}
}

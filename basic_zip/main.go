package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	tutorialPath := "./basic_zip"

	file1 := fmt.Sprintf("%s/%s", tutorialPath, "sample_119_row.csv")
	file2 := fmt.Sprintf("%s/%s", tutorialPath, "sample_1109_row.csv")

	files := []string{file1, file2}

	output := fmt.Sprintf("%s/%s", tutorialPath, "./sample.zip")

	if err := ZipFiles(output, files); err != nil {
		panic(err)
	}
	fmt.Println("Zipped File:", output)
}

func ZipFiles(filename string, files []string) error {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Project Path :", path)

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

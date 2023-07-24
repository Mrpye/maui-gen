package lib

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

// zip  the folder
func ZipFolder(ZipFile string, zipFolder string) {
	file, err := os.Create(ZipFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		/*log.Print(path)
		parts := strings.Split(path, string(os.PathSeparator))
		cleaned := parts[1:]
		path = strings.Join(cleaned, string(os.PathSeparator))*/
		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	err = filepath.Walk(zipFolder, walker)
	if err != nil {
		panic(err)
	}

	defer file.Close()

}

func MakeDir(filename string) error {
	file := path.Dir(filename)
	return os.MkdirAll(file, os.ModePerm)
}

func GetFileNameFromIndex(repo_path string, string_index string) (string, error) {
	files, err := ioutil.ReadDir(repo_path)
	if err != nil {
		return "", err
	}
	index, err := strconv.Atoi(string_index)
	if err != nil {
		return "", err
	}
	if index >= len(files) {
		return "", errors.New("index out of range")
	}
	return files[index].Name(), nil
}

// Helper function to create a directly and save string
func SaveStringToFile(filename string, content string) error {
	if content != "" {
		file := path.Dir(filename)
		err := os.MkdirAll(file, os.ModePerm)

		if err != nil {
			return err
		}
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		if err != nil {
			return err
		}
		f.WriteString(string(content))
	}
	return nil
}

// download file from url
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file := path.Dir(filepath)
	err = os.MkdirAll(file, os.ModePerm)
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Copy file
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// Reads a file to string
func ReadFileToString(sourceFile string) (string, error) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return "", err
	}
	return string(input), nil
}

// read file to string
func ReadFile(sourceFile string) string {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return ""
	}
	return string(input)
}

// file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func DirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	source_info, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, source_info.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		source_file_pointer := source + "/" + obj.Name()

		destination_file_pointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(source_file_pointer, destination_file_pointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			_, err = CopyFile(source_file_pointer, destination_file_pointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

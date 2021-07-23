package knows

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Persistor interface {
	Write(string, []byte) error
	Read(string) ([]byte, error)
	Update(string, []byte) error
	ProcessAll(func([]byte) error) error
}

type FileSystemPersistor struct {
	baseDir string
}

func NewFileSystemPersistor(baseDir string) *FileSystemPersistor {
	return &FileSystemPersistor{
		baseDir: baseDir,
	}
}

func (f *FileSystemPersistor) Write(uuid string, data []byte) error {
	return ioutil.WriteFile(fmt.Sprintf("%s/%s.json", f.baseDir, uuid), data, 0644)
}

func (f *FileSystemPersistor) Read(uuid string) ([]byte, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.json", f.baseDir, uuid))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f *FileSystemPersistor) Update(uuid string, data []byte) error {
	err := os.Remove(fmt.Sprintf("%s/%s.json", f.baseDir, uuid))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fmt.Sprintf("%s/%s.json", f.baseDir, uuid), data, 0644)
}

func (f *FileSystemPersistor) ProcessAll(fn func([]byte) error) error {

	files, err := os.ReadDir(f.baseDir)
	if err != nil {
		return err
	}

	for _, fl := range files {
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", f.baseDir, fl.Name()))
		if err != nil {
			return err
		}

		err = fn(data)
		if err != nil {
			return err
		}
	}

	return nil
}

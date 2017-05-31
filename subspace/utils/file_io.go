package utils

import (
	"io/ioutil"
	"os"
	"bufio"
)

func ReadFromFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, err
	}
	return data, nil
}

func WriteToFile(path string, data string) error {
	file, err := os.Create(path)
	defer file.Close()

	if nil != err {
		return err
	}

	bufferWriter := bufio.NewWriter(file)
	if _, e := bufferWriter.WriteString(data); nil != e {
		return e
	}

	if e := bufferWriter.Flush(); nil != e {
		return e
	}

	return nil
}
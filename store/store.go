package store

import (
	"log"
	"os"
	"path"
)

type Store struct {
	rootPath string
}

func Manager() *Store {
	return &Store{
		rootPath: "./public/",
	}
}

func (s Store) CreateFile(fileName string) error {
	file, err := os.Create(path.Join(s.rootPath, fileName))
	if err != nil {
		log.Println("[CreateFile] Error Create")
		return err
	}
	err = file.Close()
	if err != nil {
		log.Println("[CreateFile] Error Close")
		return err
	}

	return nil
}

func (s Store) WriteStringFile(fileName string, text string) error {
	file, err := os.OpenFile(path.Join(s.rootPath, fileName), os.O_APPEND, 0644)
	if err != nil {
		log.Println("[WriteStringFile] Error Open")
		return err
	}
	_, err = file.WriteString(text)
	if err != nil {
		log.Println("[WriteStringFile] Error WriteString")
		return err
	}

	err = file.Close()
	if err != nil {
		log.Println("[WriteStringFile] Error Close")
		return err
	}

	return nil
}

package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}

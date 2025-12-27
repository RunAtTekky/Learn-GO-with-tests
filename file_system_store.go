package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	score := 0
	for _, player := range f.GetLeague() {
		if player.Name == name {
			score = player.Wins
			break
		}
	}
	return score
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}

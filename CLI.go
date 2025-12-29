package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()

	winner := extractWinner(reader.Text())
	cli.playerStore.RecordWin(winner)
}

func extractWinner(input string) string {
	return strings.ReplaceAll(input, " wins", "")
}

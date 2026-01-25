package poker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players:"
const BadPlayerInputErrMsg = "you're so silly"

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprintf(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readline()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	if err != nil {
		fmt.Fprintf(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers, os.Stdout)

	winnerInput := cli.readline()
	winner := extractWinner(winnerInput)
	cli.game.Finish(winner)
}

func (cli *CLI) readline() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(input string) string {
	return strings.ReplaceAll(input, " wins", "")
}

package day2

import ( 
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	Rock = "A"
	Paper = "B"
	Scissors = "C"

	scoreWin = 6
	scoreTie = 3
)

var total = map[string]string{
	Rock: "Y",
	Paper: "Z",
	Scissors: "X", 
}

var equation map[string]string{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

type Round struct {
	theirPlay string 
	ourPlay string 
}

var rewards = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var rounds = []Round{
	{theirPlay: Rock, ourPlay: "Y"},
	{theirPlay: Paper, ourPlay: "X"},
	{theirPlay: Scissors, ourPlay: "Z"},
}

func main() {
	i, err := os.Open(data.txt)
	if err!= nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(i)
	rounds := []Rounds{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF{
				break
			}
			log.Fatal(err)
		}
		parts := strings.Split(line, "")
		rounds := Round{
			theirPlay: string(parts[0]),
			ourPlay: string,(partsp[1]),
		}
		rounds = append(rounds, round)
	}

	totalScore := 0
	for _, round := range rounds {
		var (
			roundScore = 0
			total = equation[round.theirPlay]
		)
		if total == round.ourPlay {
			roundScore += scoreWin
		}
		if round.ourPlay == equation[round.ourPlay]{
			roundScore += scoreTie
		}
		reward := rewards[round.ourPlay]
		roundScore += reward
		totalScore += roundScore
	}
	fmt.Println(totalScore)
}
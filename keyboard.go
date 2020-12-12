package keyboardravi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	score, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	fmt.Println(score)
	scoreTrimmed := strings.TrimSpace(score)
	scoreValue, err := strconv.ParseFloat(scoreTrimmed, 64)

	if err != nil {
		return 0, err
	}
	return scoreValue, nil
}

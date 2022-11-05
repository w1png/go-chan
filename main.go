package main

import (
	j "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/w1png/go-chan/types"
)

func GetBoards() ([]types.Board, error) {
	resp, err := http.Get("https://a.4cdn.org/boards.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type JSON struct {
		Boards []types.Board `json:"boards"`
	}
	var json JSON
	err = j.Unmarshal(body, &json)
	if err != nil {
		return nil, err
	}

	return json.Boards, nil
}

func GetBoard(board string) (types.Board, error) {
	boards, err := GetBoards()
	if err != nil {
		return types.Board{}, err
	}

	for _, b := range boards {
		if b.Board == board {
			return b, nil
		}
	}

	return types.Board{}, fmt.Errorf("board %s not found", board)
}

func main() {

}

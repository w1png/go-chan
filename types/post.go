package types

type Post struct {
	Thread

	BoardFlag interface{} `json:"board_flag"`
	FlagName  string      `json:"flag_name"`
}

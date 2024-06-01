package clients

type GuessRequest struct {
	Guess string `json:"guess"`
}

type Wordle struct {
	Guess         string `json:"guess"`
	WasCorrect    bool   `json:"was_correct"`
	CharacterInfo []struct {
		Char    string `json:"char"`
		Scoring struct {
			InWord     bool `json:"in_word"`
			CorrectIdx bool `json:"correct_idx"`
		} `json:"scoring"`
	} `json:"character_info"`
}

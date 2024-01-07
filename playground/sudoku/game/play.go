package game

func (game *game) play() error {
	for i := 0; i < 9; i++ {
		turn, err := game.board.MakeTurn(i+1, i, i, '1')
		if err != nil {
			return err
		}
		game.turns = append(game.turns, &turn)
	}
	return nil
}

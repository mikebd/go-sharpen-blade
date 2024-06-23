In this live coding interview, I was asked to evaluate the state of a tic-tac-toe game.

The state of the board is represented by a comma separated string of characters.
The characters are either `x`, `X`, `o`, `O`, or ` `.  The input is in 
[row-major order](https://en.wikipedia.org/wiki/Row-_and_column-major_order).

The output should be a string indicating the state of the game:
  * `"Error"`
  * `"Insufficient Data"`
  * `"Tie"` (this solution ignores eventual ties)
  * `"O Win"`
  * `"X Win"`

The code here is a retrospective of the approach I wish I had taken.

Unfortunately I had never previously done 2D arrays in Go, my solution in the interview
was not elegant.  A related thread: [What is a concise way to create a 2D slice in Go?](https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go).

Running tests:
```shell
> go test ./interview/ticTacToe/... -count=1 -cover
ok  	go-sharpen-blade/interview/ticTacToe	0.002s	coverage: 90.9% of statements
```

I am also working on a Sudoku solver, [here](https://github.com/mikebd/go-sharpen-blade/tree/master/playground/sudoku)...
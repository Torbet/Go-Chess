package main

type Piece byte 
var values = map[Piece]int{'r': -500, 'b': -320, 'n': -300, 'p': -100, 'q': -900, 'k': -2200, 'R': 500, 'B': 320, 'N': 300, 'P': 100, 'Q': 900, 'K': 2200}

func (p Piece) Value() int {
  return values[p]
}


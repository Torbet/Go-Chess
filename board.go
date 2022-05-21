package main

import (
  "fmt"
  "strings"
  "unicode"
)

type Board [8][8]Piece

func FEN(fen string) (b Board) {
  parts := strings.Split(fen, " ")
  rows := strings.Split(parts[0], "/")

  for i, row := range rows {
    shift:=0
    for j, piece := range row {
      if !unicode.IsDigit(piece) {
        b[i][j+shift] = Piece(piece)
      } else {
        r := int(piece-'0')
        for n:=0;n<r;n++ {
          b[i][j+shift+n] = Piece('·')
        }
        shift+=r-1
      }
    }
  }
  return b
}

func (b Board) Display() {
  for _, row := range b {
    for _, piece := range row {
     fmt.Printf("%c ", piece) 
    }
    fmt.Println()
  }
}

func (b *Board) Move(from C, to C) {
  if isValidC(to) && isValidC(from) {
    b[to.y][to.x] = b[from.y][from.x]
    b[from.y][from.x] = ' '
  }
}

type C struct {
  x int
  y int
}

//struct State {}

func main() {
  b := FEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2")
  b.Display()
}

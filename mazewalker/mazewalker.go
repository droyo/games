package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"log"
	"math/rand"

	"image/color"
	_ "image/png"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("")
	pixelgl.Run(run)
}

func run() {
	const spriteSize = 64
	cfg := pixelgl.WindowConfig{
		Title:  "Mazewalker",
		Bounds: pixel.R(0, 0, 640, 640),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	worldsheet, err := pictureFromFile("roguelikeSheet_transparent.png")
	if err != nil {
		log.Fatal(err)
	}
	charsheet, err := pictureFromFile("roguelikeChar_transparent.png")
	if err != nil {
		log.Fatal(err)
	}

	bg := pixel.NewBatch(&pixel.TrianglesData{}, worldsheet)
	fg := pixel.NewBatch(&pixel.TrianglesData{}, charsheet)

	var (
		bgSprites = make([]*pixel.Sprite, NumCellTypes)
		hero      = pixel.NewSprite(charsheet, pixel.R(16*0+0, 16*0+0, 16*1+0, 16*1+0))
	)

	bgSprites[Wall] = pixel.NewSprite(worldsheet, pixel.R(16*1+1, 16*1+1, 16*2+1, 16*2+1))
	bgSprites[Open] = pixel.NewSprite(worldsheet, pixel.R(16*8+8, 16*8+8, 16*9+8, 16*9+8))
	bgSprites[Goal] = pixel.NewSprite(worldsheet, pixel.R(16*10+10, 16*1+1, 16*11+10, 16*2+1))

	board := NewBoard(24, 24)
	buildMaze(board)

	// Scale board to screen
	scale := pixel.V(
		win.Bounds().W()/hero.Frame().W()/float64(board.W),
		win.Bounds().H()/hero.Frame().H()/float64(board.H))

	mat := pixel.IM.ScaledXY(pixel.ZV, scale)

	win.SetMatrix(mat)

	pos := pixel.ZV
	drawBoard(bg, board, bgSprites)
	drawHero(fg, hero, pos)

	bg.Draw(win)
	fg.Draw(win)

	for !win.Closed() && !win.JustPressed(pixelgl.KeyEscape) {
		if board.Cell(pos) == Goal {
			pos = pixel.ZV
			buildMaze(board)
			drawBoard(bg, board, bgSprites)
			drawHero(fg, hero, pos)
			bg.Draw(win)
			fg.Draw(win)
		}
		if move := keyDirection(win); move != pixel.ZV {
			if newpos := pos.Add(move); board.Cell(newpos) != Wall {
				pos = newpos
				drawHero(fg, hero, pos)
				bg.Draw(win)
				fg.Draw(win)
			} else {
				fmt.Println("Can't go that way")
			}
		}
		win.Update()
	}
}

func keyDirection(win *pixelgl.Window) pixel.Vec {
	if win.JustPressed(pixelgl.KeyUp) || win.Repeated(pixelgl.KeyUp) {
		return pixel.V(0, 1)
	} else if win.JustPressed(pixelgl.KeyDown) || win.Repeated(pixelgl.KeyDown) {
		return pixel.V(0, -1)
	} else if win.JustPressed(pixelgl.KeyLeft) || win.Repeated(pixelgl.KeyLeft) {
		return pixel.V(-1, 0)
	} else if win.JustPressed(pixelgl.KeyRight) || win.Repeated(pixelgl.KeyRight) {
		return pixel.V(1, 0)
	}
	return pixel.ZV
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Generate a maze with recursive backtracking
func buildMaze(b *Board) {
	// Seed data
	for i := range b.C {
		if rand.Float32() > 0.8 {
			b.C[i] = Open
		}
	}
	for i := 0; i < 300; i++ {
		b.MapGroup(mazectric)
	}
	for i := 0; i < max(b.W, b.H); i++ {
		b.Set(i, 0, Open)
		b.Set(i, b.H-1, Open)
		b.Set(0, i, Open)
		b.Set(b.W-1, i, Open)
	}
	b.Set(b.W-1, b.H-1, Goal)
}

func mazectric(group [9]Cell) Cell {
	n := 0
	for i, c := range group {
		if i != 4 && c == Open { // self
			n++
		}
	}
	switch group[4] {
	case Wall:
		if n == 3 {
			return Open
		}
	case Open:
		if n > 0 && n < 5 {
			return Open
		}
	}
	return Wall
}

func drawHero(fg *pixel.Batch, sprite *pixel.Sprite, pos pixel.Vec) {
	fg.Clear()
	box := pixel.Rect{}.ResizedMin(sprite.Frame().Size())
	box = box.Moved(pos.ScaledXY(box.Size()))
	sprite.Draw(fg, pixel.IM.Moved(box.Center()))
}

func drawBoard(bg *pixel.Batch, b *Board, sprite []*pixel.Sprite) {
	bg.Clear()
	b.MapCell(func(c Cell, x, y int) Cell {
		coords := pixel.R(float64(x), float64(y),
			float64(x+1), float64(y+1))

		coords.Min = coords.Min.ScaledXY(sprite[c].Frame().Size())
		coords.Max = coords.Max.ScaledXY(sprite[c].Frame().Size())
		sprite[c].Draw(bg, pixel.IM.Moved(coords.Center()))

		return c
	})
}

type Cell uint8

const (
	Wall Cell = iota
	Open
	Goal
	NumCellTypes
)

type Board struct {
	W, H int
	C    []Cell
}

func NewBoard(width, height int) *Board {
	return &Board{W: width, H: height, C: make([]Cell, width*height)}
}

func (b *Board) MapCell(fn func(Cell, int, int) Cell) {
	for i, c := range b.C {
		b.C[i] = fn(c, i%b.W, i/b.W)
	}
}

func (b *Board) MapGroup(fn func([9]Cell) Cell) {
	newBoard := make([]Cell, len(b.C))
	for i := range b.C {
		x, y := i%b.W, i/b.W
		newBoard[i] = fn([...]Cell{
			b.Get(x-1, y-1), b.Get(x, y-1), b.Get(x+1, y-1),
			b.Get(x-1, y), b.Get(x, y), b.Get(x+1, y),
			b.Get(x-1, y+1), b.Get(x, y+1), b.Get(x+1, y+1),
		})
	}
	b.C = newBoard
}

func (b *Board) Set(x, y int, new Cell) {
	i := x + (y * b.W)
	if i >= 0 && i < len(b.C) {
		b.C[i] = new
	}
}

func (b *Board) Get(x, y int) Cell {
	if x >= 0 && x < b.W && y >= 0 && y < b.H {
		i := x + (y * b.W)
		if i >= 0 && i < len(b.C) {
			return b.C[i]
		}
	}
	return Wall
}

func (b *Board) Cell(v pixel.Vec) Cell {
	return b.Get(int(v.X), int(v.Y))
}

func pictureFromFile(filename string) (pixel.Picture, error) {
	var r io.Reader
	if data, ok := BuiltinAssets[filename]; ok {
		r = bytes.NewReader(data)
	} else {
		return nil, fmt.Errorf("no builtin asset %s", filename)
	}
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	// Clear the background color
	img = newColorMask(img, colornames.White)
	return pixel.PictureDataFromImage(img), nil
}

// Makes a single color transparent
type colorMask struct {
	C color.Color
	I image.Image
}

func newColorMask(src image.Image, c color.Color) colorMask {
	return colorMask{I: src, C: src.ColorModel().Convert(c)}
}

func (m colorMask) ColorModel() color.Model { return m.I.ColorModel() }
func (m colorMask) Bounds() image.Rectangle { return m.I.Bounds() }
func (m colorMask) At(x, y int) color.Color {
	if c := m.I.At(x, y); c != m.C {
		return c
	}
	return color.Transparent
}

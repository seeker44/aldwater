package main

import (
	"aldwater/floor"
	"aldwater/player"
	"errors"
	"image/color"
	"log"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

//30x30
var (
	cols               = 31
	rows               = 31
	fontSize   float64 = 24
	normalFont font.Face
	width      int
	height     int
	gameMap    *floor.Floor
)

var p = player.Player{
	X:    4,
	Y:    4,
	Char: "@",
}

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72

	normalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	width = cols*int(fontSize) + 20
	height = rows*int(fontSize) + 20
}

type Game struct {
	pressed []ebiten.Key
}

func (g *Game) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("ended by player")
	}
	p.HandleMovement(gameMap)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, row := range gameMap.Area {
		for _, tile := range row {
			text.Draw(screen, tile.Char, normalFont, tile.Posx, tile.Posy, tile.Color)
		}
	}
	text.Draw(screen,
		p.Char,
		normalFont,
		gameMap.Area[p.Y][p.X].Posx,
		gameMap.Area[p.Y][p.X].Posy,
		color.White)

	text.Draw(screen, strconv.Itoa(int(ebiten.CurrentTPS())), normalFont, 24, 700, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	gameMap = floor.New(cols, rows, int(fontSize))
	p.StartingPosition(gameMap)

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Aldwater")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

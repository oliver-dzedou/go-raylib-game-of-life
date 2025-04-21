package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "math/rand/v2"

type Grid struct {
	rows     int32
	columns  int32
	cellSize int32
	cells    [][]int32
}

func InitGrid(width int32, height int32, cellSize int32) Grid {
	rows := height / cellSize
	columns := width / cellSize
	cells := make([][]int32, rows)
	for r := range cells {
		cells[r] = make([]int32, columns)
	}

	grid := Grid{
		rows:     height / cellSize,
		columns:  width / cellSize,
		cellSize: cellSize,
		cells:    cells,
	}
	return grid
}

func (g *Grid) Draw() {
	for row := range g.rows {
		for column := range g.columns {
			position := g.cells[row][column]
			color := rl.NewColor(55, 55, 55, 255)
			if position == 1 {
				color = rl.RayWhite
			}
			rl.DrawRectangle(column*g.cellSize, row*g.cellSize, g.cellSize-1, g.cellSize-1, color)
		}
	}
}

func (g *Grid) SetCell(row int32, column int32, value int32) {
	if g.IsWithinBounds(row, column) {
		g.GetCells()[row][column] = value
	}
}

func (g *Grid) ToggleCell(row int32, column int32) {
	if g.GetCell(row, column) == 1 {
		g.SetCell(row, column, 0)
	} else {
		g.SetCell(row, column, 1)
	}
}

func (g *Grid) GetCell(row int32, column int32) int32 {
	if g.IsWithinBounds(row, column) {
		return g.GetCells()[row][column]
	}
	return 0
}

func (g *Grid) GetRows() int32 {
	return g.rows
}

func (g *Grid) GetColumns() int32 {
	return g.columns
}

func (g *Grid) GetCells() [][]int32 {
	return g.cells
}

func (g *Grid) IsWithinBounds(row int32, column int32) bool {
	return row >= 0 && row < g.GetRows() && column >= 0 && column < g.GetColumns()
}

func (g *Grid) FillRandom() {
	for row := range g.GetRows() {
		for column := range g.GetColumns() {
			random := rand.Int32N(5)
			if random == 4 {
				g.SetCell(row, column, 1)
			} else {
				g.SetCell(row, column, 0)
			}
		}
	}
}

func (g *Grid) Clear() {
	for row := range g.GetRows() {
		for column := range g.GetColumns() {
			g.SetCell(row, column, 0)
		}
	}
}

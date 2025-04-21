package main

var NEIGHBOUR_OFFSET [8][2]int32 = [8][2]int32{
	{-1, 0},  // Directly above
	{1, 0},   // Directly below
	{0, -1},  // To the left
	{0, 1},   // To the right
	{-1, -1}, // Diagonal upper left
	{-1, 1},  // Diagonal upper right
	{1, -1},  // Diagonal lower left
	{1, 1},   // Diagonal lower right
}

type Simulation struct {
	Grid     Grid
	TempGrid Grid
	Paused   bool
}

func InitSimulation(width int32, height int32, cellSize int32) Simulation {
	return Simulation{
		Grid:     InitGrid(width, height, cellSize),
		TempGrid: InitGrid(width, height, cellSize),
		Paused:   true,
	}
}

func (s *Simulation) Draw() {
	s.Grid.Draw()
}

func (s *Simulation) SetCell(row int32, column int32, value int32) {
	s.Grid.SetCell(row, column, value)
}

func (s *Simulation) CountLiveNeighbours(row int32, column int32) int32 {
	var liveNeighbours int32 = 0

	for i := range NEIGHBOUR_OFFSET {
		neighbourRow := (row + NEIGHBOUR_OFFSET[i][0] + s.Grid.GetRows()) % s.Grid.GetRows()
		neighbourColumn := (column + NEIGHBOUR_OFFSET[i][1] + s.Grid.GetColumns()) % s.Grid.GetColumns()
		liveNeighbours += s.Grid.GetCell(neighbourRow, neighbourColumn)
	}
	return liveNeighbours
}

func (s *Simulation) Update() {
	if s.Paused {
		return
	}
	for row := range s.Grid.rows {
		for column := range s.Grid.columns {
			liveNeighbours := s.CountLiveNeighbours(row, column)
			cellValue := s.Grid.GetCell(row, column)
			if cellValue == 1 {
				if liveNeighbours > 3 || liveNeighbours < 2 {
					s.TempGrid.SetCell(row, column, 0)
				} else {
					s.TempGrid.SetCell(row, column, 1)
				}
			} else {
				if liveNeighbours == 3 {
					s.TempGrid.SetCell(row, column, 1)
				} else {
					s.TempGrid.SetCell(row, column, 0)
				}
			}
		}
	}
	s.Grid, s.TempGrid = s.TempGrid, s.Grid
}

func (s *Simulation) IsPaused() bool {
	return s.Paused
}

func (s *Simulation) Pause() {
	s.Paused = true
}

func (s *Simulation) Resume() {
	s.Paused = false
}

func (s *Simulation) Clear() {
	s.Grid.Clear()
}

func (s *Simulation) FillRandom() {
	s.Grid.FillRandom()
}

func (s *Simulation) ToggleCell(row int32, column int32) {
	if !s.IsPaused() {
		return
	}
	s.Grid.ToggleCell(row, column)
}

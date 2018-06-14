package ByteFrequencyTable

import "sort"

// class Table
// the class store and return frequency table
//
type Table struct {
	frequencyMap [255] byte
}

func (table *Table) AddValue(value byte) {
	table.frequencyMap[value] = table.frequencyMap[value] + 1
}

func (table *Table) GetFrequencyTable() []Cell {
	var cells []Cell

	for i := 0; i != len(table.frequencyMap); i++ {
		if table.frequencyMap[i] == 0 {
			continue
		}
		cell := Cell{int(table.frequencyMap[i]), byte(i)}
		cells = append(cells, cell)
	}

	sort.Slice(cells, func(i, j int) bool {
		return cells[i].Weight > cells[j].Weight
	})

	return cells
}
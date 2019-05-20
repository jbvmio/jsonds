package jsonds

import (
	"encoding/json"
	"fmt"
)

// TableData contains the datapoints for a TableDataResponse.
type TableData struct {
	Columns    []TagKey        `json:"columns"`
	Rows       [][]interface{} `json:"rows"`
	Type       string          `json:"type"`
	columnSize int
}

// NewTableData returns a new TableData struct.
func NewTableData(columnSize int) TableData {
	td := TableData{
		Columns:    make([]TagKey, 0, columnSize),
		Type:       `table`,
		columnSize: columnSize,
	}
	return td
}

// InsertColumn inserts Column Values.
func (t *TableData) InsertColumn(textVal, typeVal string) {
	t.Columns = append(t.Columns, TagKey{
		Type: typeVal,
		Text: textVal,
	})
}

// InsertRow inserts Row Values in the column order received.
// Number of args should match the columnSize of the table.
func (t *TableData) InsertRow(rowVals ...interface{}) error {
	if len(rowVals) != t.columnSize {
		return fmt.Errorf(`number of Row elements do not match the number of Table Columns`)
	}
	t.Rows = append(t.Rows, rowVals)
	return nil
}

// TableResponse contains all the information needed to render a TimeSeries event.
type TableResponse struct {
	Data []TableData
}

// RespType satisfies the QueryResponse interface and returns the response type.
func (r TableResponse) RespType() ResponseType {
	return RespTable
}

// MarshalJSON provides JSON marshalling for a TimeSeriesResponse.
func (r TableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Data)
}

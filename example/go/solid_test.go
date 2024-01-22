package example

import "testing"

func TestSOLID(t *testing.T) {
	const value string = "new data"
	example := New(value)
	t.Run("Test Read", func(t *testing.T) {
		data := example.Read()
		if data != "new data" {
			t.Errorf("error, value not valid %v", data)
		}
	})

	t.Run("Test Write", func(t *testing.T) {
		example.Write("another data")
		data := example.Read()
		if data != "another data" {
			t.Errorf("error, value not valid %v", data)
		}
	})
}

package tools

import (
	"fmt"
	"time"
)

func FechaSQLServer() string {
	t := time.Now()
	// Le doy formato a la fecha 
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

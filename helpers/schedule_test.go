package helpers

import (
	"fmt"
	"testing"
	"time"
)


func TestSchedule(t *testing.T) {
	today := time.Now()
	fmt.Println(ParseCron(Semanal, today))

}
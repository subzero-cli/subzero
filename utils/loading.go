package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

func AsyncTaskLoading(task func(), description string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " " + description
	s.Start()

	func() {
		time.Sleep(3 * time.Second)
		task()
		s.Stop()
	}()
}

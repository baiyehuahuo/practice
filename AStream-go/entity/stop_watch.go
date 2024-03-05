package entity

import "time"

type StopWatch struct {
	StartTime   time.Time
	ElapsedTime time.Duration
	Running     bool
}

func (s *StopWatch) Start() {
	// Start the stopwatch, ignore if running.
	if !s.Running {
		s.StartTime = time.Now().Add(-s.ElapsedTime)
		s.Running = true
	}
}

func (s *StopWatch) Pause() {
	// Stop the stopwatch, ignore if already paused
	if s.Running {
		s.ElapsedTime = time.Now().Sub(s.StartTime)
		s.Running = false
	}
}

func (s *StopWatch) Reset() {
	// Reset the stopwatch
	s.StartTime = time.Now()
	s.ElapsedTime = 0.0
}

func (s *StopWatch) Time() time.Duration {
	// return: elapsed time
	if s.Running {
		s.ElapsedTime = time.Now().Sub(s.StartTime)
	}
	return s.ElapsedTime
}

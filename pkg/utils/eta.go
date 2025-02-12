package utils

import (
    "fmt"
    "time"
)

type ETACalculator struct {
    startTime time.Time
    samples   []float64
    size      int
}

func NewETACalculator() *ETACalculator {
    return &ETACalculator{
        startTime: time.Now(),
        samples:   make([]float64, 10), // Keep last 10 samples
        size:      0,
    }
}

func (e *ETACalculator) Update(progress int) (float64, time.Duration) {
    now := time.Now()
    elapsed := now.Sub(e.startTime).Seconds()

    // Calculate current speed
    speed := float64(progress) / elapsed

    // Update samples
    e.samples[e.size%len(e.samples)] = speed
    if e.size < len(e.samples) {
        e.size++
    }

    // Calculate average speed
    var total float64
    for i := 0; i < e.size; i++ {
        total += e.samples[i]
    }
    avgSpeed := total / float64(e.size)

    // Calculate ETA
    remaining := float64(100 - progress)
    seconds := remaining / avgSpeed
    eta := time.Duration(seconds * float64(time.Second))

    return avgSpeed, eta
}

func (e *ETACalculator) TotalTime() time.Duration {
    return time.Since(e.startTime)
}

func FormatETA(eta time.Duration) string {
    if eta <= 0 {
        return ".."
    }

    minutes := int(eta.Minutes())
    seconds := int(eta.Seconds()) % 60

    if minutes > 0 {
        return fmt.Sprintf("%dm%ds", minutes, seconds)
    }
    return fmt.Sprintf("%ds", seconds)
}
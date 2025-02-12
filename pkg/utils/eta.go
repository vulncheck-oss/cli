package utils

import (
    "fmt"
    "time"
)

type ETACalculator struct {
    startTime    time.Time
    lastUpdate   time.Time
    lastProgress int
    samples      []float64
    sampleSize   int
}

func NewETACalculator() *ETACalculator {
    now := time.Now()
    return &ETACalculator{
        startTime:  now,
        lastUpdate: now,
        sampleSize: 10, // Keep last 10 speed samples
    }
}

func (e *ETACalculator) Update(progress int) (speed float64, eta time.Duration) {
    now := time.Now()
    elapsed := now.Sub(e.lastUpdate).Seconds()
    
    if elapsed > 0 {
        // Calculate speed for this sample
        progressDelta := progress - e.lastProgress
        currentSpeed := float64(progressDelta) / elapsed

        // Add to samples
        e.samples = append(e.samples, currentSpeed)
        if len(e.samples) > e.sampleSize {
            e.samples = e.samples[1:] // Keep last N samples
        }

        // Calculate average speed from samples
        var totalSpeed float64
        for _, s := range e.samples {
            totalSpeed += s
        }
        avgSpeed := totalSpeed / float64(len(e.samples))

        // Calculate ETA based on average speed
        if avgSpeed > 0 {
            remainingProgress := 100 - progress
            secondsLeft := float64(remainingProgress) / avgSpeed
            eta = time.Duration(secondsLeft * float64(time.Second))
        }

        e.lastUpdate = now
        e.lastProgress = progress
        speed = avgSpeed
    }

    return speed, eta
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
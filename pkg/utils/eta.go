package utils

import (
    "fmt"
    "time"
)

type ETACalculator struct {
    startTime    time.Time
    lastUpdate   time.Time
    lastProgress int
    speeds       []float64
    sampleSize   int
}

func NewETACalculator() *ETACalculator {
    now := time.Now()
    return &ETACalculator{
        startTime:    now,
        lastUpdate:   now,
        sampleSize:   5, // Number of recent speeds to average over
        speeds:       make([]float64, 0),
    }
}

func (e *ETACalculator) Update(progress int) (speed float64, eta time.Duration) {
    now := time.Now()
    elapsed := now.Sub(e.lastUpdate).Seconds()

    if elapsed > 0 {
        // Calculate the speed since the last update
        progressDelta := progress - e.lastProgress
        currentSpeed := float64(progressDelta) / elapsed

        // Add the current speed to the list of speeds
        e.speeds = append(e.speeds, currentSpeed)
        if len(e.speeds) > e.sampleSize {
            e.speeds = e.speeds[1:] // Keep only the last sampleSize speeds
        }

        // Calculate the average speed
        var totalSpeed float64
        for _, s := range e.speeds {
            totalSpeed += s
        }
        speed = totalSpeed / float64(len(e.speeds))

        // Calculate the remaining time
        if speed > 0 {
            remainingProgress := 100 - progress
            secondsLeft := float64(remainingProgress) / speed
            eta = time.Duration(secondsLeft * float64(time.Second))
        }

        e.lastUpdate = now
        e.lastProgress = progress
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
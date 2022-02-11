package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, count := range birdsPerDay {
		total += count
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekStart := (week - 1) * 7
	weekEnd := weekStart + 7
	birdsAtWeek := birdsPerDay[weekStart:weekEnd]
	return TotalBirdCount(birdsAtWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for day := range birdsPerDay {
		if day%2 == 0 {
			birdsPerDay[day] += 1
		}
	}

	return birdsPerDay
}

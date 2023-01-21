// Package raindrops create a raindrop sound.
package raindrops

import "strconv"

// approvedRaindrops initializes the sounds.
func approvedRaindrops() map[int]string {
	sounds := make(map[int]string)

	sounds[3] = "Pling"
	sounds[5] = "Plang"
	sounds[7] = "Plong"

	return sounds
}

// Convert converts a number to a raindrop sound.
func Convert(number int) string {
	sounds := approvedRaindrops()
	sound := ""

	for i := 3; i <= 7; i++ {
		if number%i == 0 {
			sound += sounds[i]
		}
	}

	if sound == "" {
		sound = strconv.Itoa(number)
	}

	return sound
}

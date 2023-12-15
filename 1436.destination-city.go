/*
 * @lc app=leetcode id=1436 lang=golang
 *
 * [1436] Destination City
 */

package main

// @lc code=start
func destCity(paths [][]string) string {
	validCities := make(map[string]bool)

	for _, path := range paths {
		validCities[path[0]] = false

		if _, ok := validCities[path[1]]; !ok {
			validCities[path[1]] = true
		}
	}

	for city, isDestination := range validCities {
		if isDestination {
			return city
		}
	}

	return ""
}

// @lc code=end

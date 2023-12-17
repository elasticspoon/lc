package main

import "testing"

type FoodTest struct {
	action string
	rating int
	arg    string
	want   string
}

func TestFoodRatings(t *testing.T) {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	genre := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	rating := []int{9, 12, 8, 15, 14, 7}

	t.Run("test case 1", func(t *testing.T) {
		ratings := Constructor(foods, genre, rating)

		tests := []FoodTest{
			{"h", 0, "korean", "kimchi"},
			{"h", 0, "japanese", "ramen"},
			{"c", 16, "sushi", ""},
			{"h", 0, "japanese", "sushi"},
			{"c", 16, "ramen", ""},
			{"h", 0, "japanese", "ramen"},
		}

		testFoods(&ratings, tests, t)
	})
}

func testFoods(ratings *FoodRatings, tests []FoodTest, t *testing.T) {
	for i, tt := range tests {
		switch tt.action {
		case "h":
			got := ratings.HighestRated(tt.arg)
			if got != tt.want {
				t.Errorf("%d: HighestRated(%s) = %s, want %s", i, tt.arg, got, tt.want)
			}
		case "c":
			ratings.ChangeRating(tt.arg, tt.rating)
		}
	}
}

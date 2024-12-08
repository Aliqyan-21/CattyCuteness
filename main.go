package main

import (
	"math"
	"math/rand"
	"net/http"
	"path/filepath"
	"sort"

	"github.com/aliqyan-21/hawkwing"
)

type CatRating struct {
	Image  string
	Rating int
}

var catRatings = make(map[string]int)
var catImages []string
var currentCats = [2]string{}

const K = 30 // K factor for ELO ratings

func main() {
	app := hawkwing.Init()

	hawkwing.LoadTemplates("templates")
	app.LoadStatic("/static/", "./static")

	initCatImages("./static/cats/")

	app.AddRoute("GET", "/", showCats)
	app.AddRoute("POST", "/vote", vote)

	hawkwing.Start("localhost", "5000", app)
}

func initCatImages(catDir string) {
	files, err := filepath.Glob(filepath.Join(catDir, "*.jpg"))
	if err != nil {
		panic("Failed to read cat images: " + err.Error())
	}
	for _, file := range files {
		relativePath := filepath.Base(file)
		catImages = append(catImages, relativePath)
		catRatings[relativePath] = 1200
	}
}

func randomCat(exclude map[string]bool) string {
	if len(catImages) == 0 {
		return "" // No images available
	}
	for {
		cat := catImages[rand.Intn(len(catImages))]
		if !exclude[cat] {
			return cat
		}
	}
}

func showCats(w http.ResponseWriter, r *http.Request) {
	if currentCats[0] == "" && currentCats[1] == "" {
		currentCats[0] = randomCat(nil)
		currentCats[1] = randomCat(map[string]bool{currentCats[0]: true})
	}

	var sortedRatings []CatRating
	for cat, rating := range catRatings {
		sortedRatings = append(sortedRatings, CatRating{Image: cat, Rating: rating})
	}

	sort.Slice(sortedRatings, func(i, j int) bool {
		return sortedRatings[i].Rating > sortedRatings[j].Rating
	})

	hawkwing.RenderHTML(w, "home.html", map[string]interface{}{
		"Cat1":    currentCats[0],
		"Cat2":    currentCats[1],
		"Ratings": sortedRatings,
	})
}

func vote(w http.ResponseWriter, r *http.Request) {
	selectedCat := r.FormValue("selectedCat")

	if selectedCat != "" {
		var winner string
		var loser string

		if selectedCat == currentCats[0] {
			winner = currentCats[0]
			loser = currentCats[1]
		} else {
			winner = currentCats[1]
			loser = currentCats[0]
		}

		updateRatings(winner, loser)

		exclude := map[string]bool{
			currentCats[0]: true,
			currentCats[1]: true,
		}

		currentCats[0] = randomCat(exclude)
		currentCats[1] = randomCat(map[string]bool{currentCats[0]: true})

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func updateRatings(winner, loser string) {
	winnerRating := float64(catRatings[winner])
	loserRating := float64(catRatings[loser])

	expectedWinner := 1 / (1 + math.Pow(10, (loserRating-winnerRating)/400))
	expectedLoser := 1 / (1 + math.Pow(10, (winnerRating-loserRating)/400))

	catRatings[winner] += int(K * (1 - expectedWinner))
	catRatings[loser] += int(K * (0 - expectedLoser))
}

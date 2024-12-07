package main

import (
	"fmt"
	"github.com/aliqyan-21/hawkwing"
	"math/rand"
	"net/http"
	"path/filepath"
	"sort"
)

type CatRating struct {
	Image string
	Votes int
}

var catRatings = make(map[string]int)
var catImages []string
var currentCats = [2]string{}

func main() {
	app := hawkwing.Init()

	hawkwing.LoadTemplates("templates")
	app.LoadStatic("/static/", "./static")

	initCatImages("./static/cats/")

	app.AddRoute("GET", "/", showCats)
	app.AddRoute("POST", "/vote", vote)

	hawkwing.Start("0.0.0.0", "5000", app)
}

func initCatImages(catDir string) {
	files, err := filepath.Glob(filepath.Join(catDir, "*.jpg"))
	if err != nil {
		panic("Failed to read cat images: " + err.Error())
	}
	for _, file := range files {
		relativePath := filepath.Base(file)
		catImages = append(catImages, relativePath)
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
		sortedRatings = append(sortedRatings, CatRating{Image: cat, Votes: rating})
	}

	sort.Slice(sortedRatings, func(i, j int) bool {
		return sortedRatings[i].Votes > sortedRatings[j].Votes
	})

	fmt.Println("Cat Ratings:", catRatings)       // Debugging line
	fmt.Println("Sorted Ratings:", sortedRatings) // Debugging line

	hawkwing.RenderHTML(w, "home.html", map[string]interface{}{
		"Cat1":    currentCats[0],
		"Cat2":    currentCats[1],
		"Ratings": sortedRatings,
	})
}

func vote(w http.ResponseWriter, r *http.Request) {
	selectedCat := r.FormValue("selectedCat")

	if selectedCat != "" {
		catRatings[selectedCat]++
	}

	exclude := map[string]bool{
		currentCats[0]: true,
		currentCats[1]: true,
	}

	currentCats[0] = randomCat(exclude)
	currentCats[1] = randomCat(map[string]bool{currentCats[0]: true})

	http.Redirect(w, r, "/", http.StatusFound)
}

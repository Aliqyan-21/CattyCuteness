package main

import (
	"math/rand"
	"net/http"
	"path/filepath"

	"github.com/aliqyan-21/hawkwing"
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

	// Initialize the list of cat images from the directory
	initCatImages("./static/cats/")

	// Home route to show the two random cats for voting
	app.AddRoute("GET", "/", showCats)

	// Vote route to capture the user's vote and update the ratings
	app.AddRoute("POST", "/vote", vote)

	hawkwing.Start("localhost", "5000", app)
}

// initCatImages scans the given directory for image files and populates the catImages slice.
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

// randomCat picks a random cat image that is not in the `exclude` set.
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

// showCats renders the homepage with two random cats for voting
func showCats(w http.ResponseWriter, r *http.Request) {
	if currentCats[0] == "" && currentCats[1] == "" {
		currentCats[0] = randomCat(nil)
		currentCats[1] = randomCat(map[string]bool{currentCats[0]: true})
	}

	hawkwing.RenderHTML(w, "home.html", map[string]interface{}{
		"Cat1":    currentCats[0],
		"Cat2":    currentCats[1],
		"Ratings": catRatings,
	})
}

// vote handles the voting logic, updates the rating in memory
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

package movieinfo

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/search"
	"html/template"
	"net/http"
)

type Movie struct {
	Title   string
	Summary string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/new-movie", handleNewMovie)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	err := tpl.ExecuteTemplate(res, "index", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}

func handleNewMovie(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)

	type newMovieModel struct {
		CreatedID string
	}
	model := &newMovieModel{}

	if req.Method == "POST" {
		title := req.FormValue("title")
		summary := req.FormValue("summary")

		index, err := search.Open("movies")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		movie := &Movie{
			Title:   title,
			Summary: summary,
		}

		id, err := index.Put(ctx, "", movie)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		model.CreatedID = id
	}

	err := tpl.ExecuteTemplate(res, "new-movie", model)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}

func handleSearch(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	query := req.FormValue("q")

	type searchModel struct {
		Query  string
		Movies []Movie
	}
	model := &searchModel{
		Query: query,
	}

	index, err := search.Open("movies")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	iterator := index.Search(ctx, query, nil)
	for {
		var movie Movie
		_, err := iterator.Next(&movie)
		if err == search.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		model.Movies = append(model.Movies, movie)
	}

	err = tpl.ExecuteTemplate(res, "search", model)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}

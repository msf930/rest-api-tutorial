package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var articleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"title":   &graphql.Field{Type: graphql.String},
			"desc":    &graphql.Field{Type: graphql.String},
			"content": &graphql.Field{Type: graphql.String},
		},
	},
)

var articles = []map[string]string{
	{
		"title":   "Hello World",
		"desc":    "First article",
		"content": "This is content",
	},
}

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"articles": &graphql.Field{
			Type: graphql.NewList(articleType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return articles, nil
			},
		},
	},
})

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	log.Println("Server running on :8082")
	log.Fatal(http.ListenAndServe("0.0.0.0:8082", nil))
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/graphql-go/handler"
// )

// type Article struct {
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// type Articles []Article

// func allArticles(w http.ResponseWriter, r *http.Request) {
// 	articles := Articles{
// 		{Title: "Test Title", Desc: "Test Description", Content: "Test Content"},
// 	}
// 	json.NewEncoder(w).Encode(articles)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Welcome to the HomePage!")
// }

// // CORS middleware
// func enableCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

// // Presign handler
// func presignHandler(w http.ResponseWriter, r *http.Request) {
// 	filename := r.URL.Query().Get("filename")
// 	contentType := r.URL.Query().Get("contentType")

// 	if filename == "" || contentType == "" {
// 		http.Error(w, "missing filename or contentType", http.StatusBadRequest)
// 		return
// 	}

// 	url, err := getPresignedURL(filename, contentType)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"url": url,
// 	})
// }

// func main() {
// 	mux := http.NewServeMux()

// 	// REST endpoints
// 	mux.HandleFunc("/", homePage)
// 	mux.HandleFunc("/articles", allArticles)
// 	mux.HandleFunc("/presign", presignHandler)

// 	// GraphQL handler
// 	h := handler.New(&handler.Config{
// 		Schema:   &schema, // your graphql schema
// 		Pretty:   true,
// 		GraphiQL: true,
// 	})
// 	mux.Handle("/graphql", h)

// 	log.Println("Server running on http://localhost:8082")
// 	// log.Fatal(http.ListenAndServe(":8082", enableCORS(mux)))
// 	log.Fatal(http.ListenAndServe("0.0.0.0:8082", nil))
// }

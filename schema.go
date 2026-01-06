package main

import "github.com/graphql-go/graphql"

var articles = Articles{
	{
		Title:   "Test Title",
		Desc:    "Test Description",
		Content: "Test Content",
	},
}

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

// ---------- QUERY ----------
var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"articles": &graphql.Field{
				Type: graphql.NewList(articleType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return articles, nil
				},
			},
		},
	},
)

// ---------- MUTATION ----------
var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createArticle": &graphql.Field{
				Type: articleType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"desc": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					article := Article{
						Title:   p.Args["title"].(string),
						Desc:    p.Args["desc"].(string),
						Content: p.Args["content"].(string),
					}

					articles = append(articles, article)
					return article, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)

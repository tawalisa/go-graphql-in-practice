package graphql

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-grapohql-in-practice/graphql/mysql"
	"net/http"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operationName"`
	Variables map[string]interface{} `json:"variables"`
}

type MyGraphql struct{}

var ScoresType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Scores",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"score": &graphql.Field{
				Type: graphql.Int,
			},
			"calculateDate": &graphql.Field{
				Type: graphql.String,
			},
			"scoreGrade": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"scores": &graphql.Field{
				Type:        ScoresType,
				Description: "Get Scores by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						score, err := mysql.GetScoreByID(id)
						return score, err
					}
					return nil, nil
				},
			},

			//"list": &graphql.Field{
			//	Type:        graphql.NewList(productType),
			//	Description: "Get product list",
			//	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			//		return products, nil
			//	},
			//},
		},
	})

func init() {
	log.Info("init graphQL module")
	initGraphqlSchema()
}

var schema graphql.Schema

func initGraphqlSchema() {
	schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
}

func (m MyGraphql) RunwithMyEcho(myecho *echo.Echo) {

	log.Info("graphQL integrate echo start")
	myecho.POST("/graphql", func(c echo.Context) error {

		var p postData
		if err := json.NewDecoder(c.Request().Body).Decode(&p); err != nil {
			return c.NoContent(400)
		}
		result := graphql.Do(graphql.Params{
			Context:        c.Request().Context(),
			Schema:         schema,
			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		})

		return c.JSON(http.StatusOK, result)
	})

	myecho.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, Echo!",
		})
	})

}
func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

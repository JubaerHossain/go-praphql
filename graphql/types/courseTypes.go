package types

import (
	"errors"
	"fmt"
	"lms/model"
	"lms/query"
	"reflect"

	"github.com/graphql-go/graphql"
)

var CourseType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Course",
	Description: "Course Type",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				fmt.Println(p.Info.FieldName)

				fmt.Println("Source")
				fmt.Println(p.Source)
				fmt.Println(reflect.TypeOf(p.Source))

				x := p.Source.(model.Course)

				p.Args["id"] = x.User
				user, err := query.FindByID(reflect.TypeOf(model.User{}), "users", p)
				if err != nil {
					fmt.Println(err)
					return nil, errors.New("no data found")
				}
				return user, nil
			},
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

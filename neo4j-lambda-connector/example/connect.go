package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

//Neo4jConnector : simple function to create a node in neo4j referance: https://github.com/neo4j/neo4j-go-driver
func Neo4jConnector(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	neo4jURL := os.Getenv("NEO4J_URL")
	neo4jUsername := os.Getenv("NEO4J_USERNAME")
	neo4jPassword := os.Getenv("NEO4J_PASSWORD")
	var (
		driver  neo4j.Driver
		session neo4j.Session
		result  neo4j.Result
		err     error
	)
	// connect to the graph database instance
	if driver, err = neo4j.NewDriver(neo4jURL, neo4j.BasicAuth(neo4jUsername, neo4jPassword, ""), func(c *neo4j.Config) {
		c.Encrypted = false
	}); err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	// create a session to write to the database
	session, err = driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	defer session.Close()
	// execute a simple query
	params := map[string]interface{}{
		"name": "Bob",
		"age":  25,
	}
	query := "CREATE (n:Person { name: $name ,age: $age}) RETURN n.name, n.age"

	result, err = session.Run(query, params)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	// print out results
	for result.Next() {
		fmt.Printf("Created Person with Name = '%s' and age = '%d' \n", result.Record().GetByIndex(0).(string), result.Record().GetByIndex(1).(int64))
	}

	err = result.Err()
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}
	// send a success response
	response := map[string]string{"message": "Created Node"}
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(jsonResponse), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Neo4jConnector)
}

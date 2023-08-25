package factories

import "github.com/iagomaia/sample-go-api/internal/infra/repositories"

func GetMongoClient() *repositories.MongoClient {
	return new(repositories.MongoClient)
}

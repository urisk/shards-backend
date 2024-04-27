package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"shards-backend/models"
)

const (
	shardTypesCollection = "shard_types"
)

// FillNewsTypes <function>
// is used to fill news types in DB
func (s Service) fillShardTypes() error {
	c := s.Collection(shardTypesCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	// we cast it to array of interfaces because mongo driver asks so
	var shardTypes []interface{} // []models.NewsType
	shardTypes = append(
		shardTypes,
		models.TypeDream,
		models.TypeTarot,
		models.TypeMeditation,
		models.TypeRitual,
	)

	_, err := c.InsertMany(
		ctx,
		shardTypes,
		s.InsertManyOptionsOrdered(),
	)

	return err
}

// ShardTypesGet <function>
// is used to get news types
func (s Service) ShardTypesGet() ([]models.ShardType, error) {
	c := s.Collection(shardTypesCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	var shardTypes []models.ShardType
	cur, err := c.Find(ctx, bson.D{})
	if err != nil {
		return []models.ShardType{}, err
	}

	for cur.Next(ctx) {
		var n models.ShardType
		err = cur.Decode(&n)
		if err != nil {
			return []models.ShardType{}, err
		}

		shardTypes = append(shardTypes, n)
	}

	if err = cur.Err(); err != nil {
		return []models.ShardType{}, err
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return shardTypes, nil
}


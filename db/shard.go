package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"shards-backend/models"
)

const (
	shardCollection = "shard"
)

// ShardInsert <function>
// is used to insert an array of news to DB
func (s Service) ShardInsert(shard []models.Shard) error {
	c := s.Collection(shardCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	// we cast it to array of interfaces because mongo driver asks so
	iShard := []interface{}{} // []models.Shard
	for _, n := range shard {
		iShard = append(iShard, n)
	}
	_, err := c.InsertMany(
		ctx,
		iShard,
		s.InsertManyOptionsOrdered(),
	)

	return err
}

// ShardGet <function>
// is used to get shard with given count
func (s Service) ShardGet(count int64) ([]models.Shard, error) {
	c := s.Collection(shardCollection)
	ctx, cncl := s.CTX(5)
	defer cncl()

	var shard []models.Shard
	options := options.Find()
	cur, err := c.Find(
		ctx,
		bson.D{},
		options.SetSort(bson.D{{"time_added", -1}}), // desc sort by time_added
		options.SetLimit(count),                     // limit to <count>
	)
	if err != nil {
		return []models.Shard{}, err
	}

	for cur.Next(ctx) {
		var n models.Shard
		err = cur.Decode(&n)
		if err != nil {
			return []models.Shard{}, err
		}

		shard = append(shard, n)
	}

	if err = cur.Err(); err != nil {
		return []models.Shard{}, err
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return shard, nil
}


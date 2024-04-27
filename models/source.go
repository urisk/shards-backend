package models

// Shard <model>
// is used to describe article model.
type Shard struct {
	ID         string     `json:"_id" bson:"_id"`
	Title      string     `json:"title" bson:"title"`
	TimeAdded  int64      `json:"time_added" bson:"time_added"`
	Shard      string     `json:"shard" bson:"shard"`
	ShardType  ShardType  `json:"shards_type" bson:"shards_type"`
	Author 	   User       `json:"author" bson:"author"`
}

package necessarystructs

// All the structs that are used would be available here.
// Consider this as a Blueprint for all the messages

type StructUser struct {
	Username string `bson:"Username"`
	Role     string `bson:"Role"`
	Token    string `bson:"Token"`
	Validity int64  `bson:"Validity"`
}

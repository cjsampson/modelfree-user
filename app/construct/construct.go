package construct

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

func NewUserStore() (UserStore, nil) {

}

func NewSubStorage(config conf.SubStorage) (*substore.SubDB, error) {

	sess, err := NewSession()
	if err != nil {
		return nil, failure.Wrap(err, "NewSession failed")
	}

	tbl, err := dynamo.NewTable(config.SubDBTable, substore.HashKey, substore.SortKey)
	if err != nil {
		return nil, failure.Wrap(err, "dynamo.NewTable failed")
	}

	db := dynamo.NewClient(dynamodb.New(sess), tbl)
	return substore.NewSubDB(db), nil
}

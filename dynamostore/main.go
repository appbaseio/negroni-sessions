package dynamostore

import (
	gSessions "github.com/appbaseio/sessions"
	dynstore "github.com/denizeren/dynamostore"
	nSessions "github.com/appbaseio/negroni-sessions"
)

//New returns a new Dynamodb store
func New(accessKey string, secretKey string, tableName string, region string, keyPairs ...[]byte) (nSessions.Store, error) {
	store, err := dynstore.NewDynamoStore(accessKey, secretKey, tableName, region, keyPairs...)

	if err != nil {
		return nil, err
	}
	return &dynamoStore{store}, nil
}

type dynamoStore struct {
	*dynstore.DynamoStore
}

func (c *dynamoStore) Options(options nSessions.Options) {
	c.DynamoStore.Options = &gSessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
		SameSite: options.SameSite,
	}
}

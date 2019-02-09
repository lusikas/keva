package keva

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// DefaultRegion as Frankfurt
const DefaultRegion = "eu-central-1"

// DynamoItem ...
type DynamoItem struct {
	Key   string      `dynamo:"key"`
	Value interface{} `dynamo:"value,omitempty"`
}

// Get the environment variable or default value
func getEnv(key, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}

// Client ...
type Client struct {
	Con   *dynamo.DB
	Table dynamo.Table
}

// New creates new client
func New(tablename string) *Client {
	return NewWithConfig(tablename, &aws.Config{})
}

// NewWithConfig creates new client but with your own config
func NewWithConfig(tablename string, config *aws.Config) *Client {
	if config.Region == nil {
		config.WithRegion(getEnv("AWS_REGION", DefaultRegion))
	}

	sess, err := session.NewSession(config)
	if err != nil {
		log.Fatalf("error creating AWS session: %v", err)
	}
	svc := dynamo.New(sess)
	return &Client{
		Con:   svc,
		Table: svc.Table(tablename),
	}

}

func (c *Client) getByKey(key string) (DynamoItem, error) {
	var result DynamoItem
	err := c.Table.Get("key", key).One(&result)
	return result, err
}

// Get value of key (works fine with strings, bools and float64s)
func (c *Client) Get(key string) interface{} {
	req, err := c.getByKey(key)
	if err == dynamo.ErrNotFound {
		return ""
	} else if err != nil {
		log.Fatalf("error getting key %s: %v", key, err)
	}
	return req.Value
}

// Set any value to key
func (c *Client) Set(key string, value interface{}) error {
	return c.Table.Put(DynamoItem{Key: key, Value: value}).Run()
}

// Delete item from table
func (c *Client) Delete(key string) error {
	return c.Table.Delete("key", key).Run()
}

// GetSlice gets a slice of interfaces (just to make things bit easier... simple
// types like strings, ints and floats work nice with it)
func (c *Client) GetSlice(key string) []interface{} {
	req, err := c.getByKey(key)
	if err == dynamo.ErrNotFound {
		return make([]interface{}, 0)
	} else if err != nil {
		log.Fatalf("error getting key %s: %v", key, err)
	}
	return req.Value.([]interface{})
}

// GetStringMap gets a map[string]string (just to make things bit easier...)
func (c *Client) GetStringMap(key string) map[string]string {
	req, err := c.getByKey(key)
	res := map[string]string{}
	if err == dynamo.ErrNotFound {
		return res
	} else if err != nil {
		log.Fatalf("error getting key %s: %v", key, err)
	}
	intr := req.Value.(map[string]interface{})
	for i, item := range intr {
		res[i] = item.(string)
	}
	return res
}

// GetFloatMap gets a map[string]float64 (just to make things bit easier...)
func (c *Client) GetFloatMap(key string) map[string]float64 {
	req, err := c.getByKey(key)
	res := map[string]float64{}
	if err == dynamo.ErrNotFound {
		return res
	} else if err != nil {
		log.Fatalf("error getting key %s: %v", key, err)
	}
	intr := req.Value.(map[string]interface{})
	for i, item := range intr {
		res[i] = item.(float64)
	}
	return res
}

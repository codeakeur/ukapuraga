package tests

import (
	"fmt"
	"gomies/app/gateway/persistence/postgres"
	"log"

	"github.com/ory/dockertest"
	"github.com/pkg/errors"
)

type Container struct {
	config   postgres.Config
	database *Database
	pool     *dockertest.Pool
	resource *dockertest.Resource
}

func (c *Container) create() error {
	c.database.Name = c.config.Name

	pool, err := dockertest.NewPool("")
	if err != nil {
		return errors.Wrap(err, "the docker pool connection could not be established")
	}

	if err := pool.Client.Ping(); err != nil {
		return errors.Wrap(err, "could not contact docker pool")
	}

	resource, err := pool.Run("postgres", "14-alpine", []string{
		fmt.Sprintf("POSTGRES_USER=%s", c.config.User),
		fmt.Sprintf("POSTGRES_PASSWORD=%s", c.config.Password),
		fmt.Sprintf("POSTGRES_DB=%s", c.config.Name),
	})
	if err != nil {
		return errors.Wrap(err, "the docker container could not be created")
	}

	c.pool = pool
	c.resource = resource
	c.config.Port = c.resource.GetPort("5432/tcp")
	err = resource.Expire(360)
	if err != nil {
		return errors.Wrap(err, "could not set an expiration time")
	}

	if err := c.pool.Retry(c.database.connect); err != nil {
		c.teardown()
		return errors.Wrap(err, "an error occurred when creating and connecting to the database")
	}

	return nil

}

func (c *Container) teardown() {
	err := c.pool.Purge(c.resource)
	if err != nil {
		log.Printf("Error purging database container resources: %v", err)
	}
}
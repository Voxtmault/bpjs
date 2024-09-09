package services

import (
	"context"
	"database/sql"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/storage"
)

type ExampleService struct {
	Con *sql.DB // Service usually hold a connection to the database, but you can modify this as you need

	// One of the modifications you can do is Dependency Injection
	Injection interfaces.DependencyInjection
}

// Make sure that ExampleService implement the Example interface
var _ interfaces.Example = &ExampleService{}

func (s *ExampleService) SampleGet(ctx context.Context, query any) ([]*any, error) {
	// Implement the logic here
	return nil, nil
}

func (s *ExampleService) SampleCreate(ctx context.Context, createObj any) error {
	// Implement the logic here

	// ctx / Context is usefull for cancellation, deadlines, and values that go across API boundaries and between processes.
	// For example if you have a variable to control the timeout of the request, you can use the context to pass the timeout value

	// No need to assign new db connection, just use the existing one

	// Whenever you need to modify something in the database, it is recommended to  use the transaction method
	// to ensure the data integrity and consistency
	tx, err := s.Con.BeginTx(ctx, nil)
	if err != nil {
		// Whenever there is an error, make sure to Rollback the transaction before returning the error
		tx.Rollback()
		return eris.Wrap(err, string(storage.MariaDBErrorsBeginTx))
	}

	// Logic here

	// Remember to commit the transaction before exiting the function
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return eris.Wrap(err, string(storage.MariaDBErrorsCommitTx))
	}

	// You don't need to close the DB connection, it will be returned to the pool automatically

	return nil
}

func (s *ExampleService) SampleUpdate(ctx context.Context, updateObj any) error {
	// Implement the logic here
	return nil
}

func (s *ExampleService) SampleDelete(ctx context.Context, deleteObj any) error {
	// Implement the logic here
	return nil
}

func (s *ExampleService) SampleInjectionUsage(ctx context.Context) error {
	// Do Something

	// You can call the methods in the Injected obj
	s.Injection.InjectionSample()

	return nil
}

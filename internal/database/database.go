package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	GetAllProperties() []Property
	GetProperty(id int64) (Property, error)
	AddProperty(property Property) (Property, error)
	UpdateProperty(property Property) (Property, error)
	DeleteProperty(id int64) error

	GetAllClients() []Client
	GetClient(id int64) (Client, error)
	AddClient(property Client) (Client, error)
	UpdateClient(property Client) (Client, error)
	DeleteClient(id int64) error

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *sql.DB
}

var (
	dburl      = os.Getenv("BLUEPRINT_DB_URL")
	dbInstance *service
)

func Run() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	queries := New(dbInstance.db)
	clients, err := queries.ListClients(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	stats["clients"] = strconv.Itoa(len(clients))

	queries.CreateClient(context.Background(), CreateClientParams{FirstName: "Michael", LastName: "Harris", EmailAddress: "michael@me.com", PhoneNumber: "555-555-5555"})
	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

func (s *service) GetAllProperties() []Property {
	queries := New(s.db)
	properties, err := queries.ListProperties(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return properties
}

func (s *service) GetProperty(id int64) (Property, error) {
	queries := New(s.db)
	property, err := queries.GetProperty(context.Background(), id)
	if err != nil {
		log.Print(err)
		return Property{}, err
	}

	return property, nil
}

func (s *service) AddProperty(property Property) (Property, error) {
	queries := New(s.db)
	property, err := queries.CreateProperty(context.Background(), CreatePropertyParams{
		AddressLine1:   property.AddressLine1,
		AddressLine2:   property.AddressLine2,
		City:           property.City,
		Region:         property.Region,
		PropertyTypeID: property.PropertyTypeID,
		PropertySize:   property.PropertySize,
		BlockSize:      property.BlockSize,
		NumBedrooms:    property.NumBedrooms,
		NumBathrooms:   property.NumBathrooms,
		NumCarspaces:   property.NumCarspaces,
		Description:    property.Description,
	})
	if err != nil {
		log.Print(err)
		return property, err
	}
	return property, nil
}

func (s *service) UpdateProperty(property Property) (Property, error) {
	queries := New(s.db)

	property, err := queries.UpdateProperty(context.Background(), UpdatePropertyParams{
		ID:             property.ID,
		AddressLine1:   property.AddressLine1,
		AddressLine2:   property.AddressLine2,
		City:           property.City,
		Region:         property.Region,
		PropertyTypeID: property.PropertyTypeID,
		PropertySize:   property.PropertySize,
		BlockSize:      property.BlockSize,
		NumBedrooms:    property.NumBedrooms,
		NumBathrooms:   property.NumBathrooms,
		NumCarspaces:   property.NumCarspaces,
		Description:    property.Description,
	})
	if err != nil {
		log.Print(err)
		return property, err
	}

	return property, nil
}

func (s *service) DeleteProperty(id int64) error {
	queries := New(s.db)

	err := queries.DeleteProperty(context.Background(), id)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (s *service) GetAllClients() []Client {
	queries := New(s.db)
	users, err := queries.ListClients(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return users
}

func (s *service) GetClient(id int64) (Client, error) {
	queries := New(s.db)
	property, err := queries.GetClient(context.Background(), id)
	if err != nil {
		log.Print(err)
		return Client{}, err
	}

	return property, nil
}

func (s *service) AddClient(client Client) (Client, error) {
	queries := New(s.db)
	client, err := queries.CreateClient(context.Background(), CreateClientParams{
		FirstName:    client.FirstName,
		LastName:     client.LastName,
		EmailAddress: client.EmailAddress,
		PhoneNumber:  client.PhoneNumber,
	})
	if err != nil {
		log.Print(err)
		return client, err
	}
	return client, nil
}

func (s *service) UpdateClient(client Client) (Client, error) {
	queries := New(s.db)

	client, err := queries.UpdateClient(context.Background(), UpdateClientParams{
		FirstName:    client.FirstName,
		LastName:     client.LastName,
		EmailAddress: client.EmailAddress,
		PhoneNumber:  client.PhoneNumber,
	})
	if err != nil {
		log.Print(err)
		return client, err
	}

	return client, nil
}

func (s *service) DeleteClient(id int64) error {
	queries := New(s.db)

	err := queries.DeleteClient(context.Background(), id)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	return s.db.Close()
}

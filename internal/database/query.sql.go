// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createClient = `-- name: CreateClient :one
INSERT INTO client (
  first_name, last_name, email_address, phone_number
) VALUES (
  ?, ?, ?, ?
)
RETURNING id, first_name, last_name, email_address, phone_number
`

type CreateClientParams struct {
	FirstName    string
	LastName     string
	EmailAddress string
	PhoneNumber  string
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, createClient,
		arg.FirstName,
		arg.LastName,
		arg.EmailAddress,
		arg.PhoneNumber,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.EmailAddress,
		&i.PhoneNumber,
	)
	return i, err
}

const createContract = `-- name: CreateContract :one
INSERT INTO contract (
  property_id, listing_type_id, contract_document, responsible_employee_id, client_id, contract_status_id, signed_date, start_date, end_date
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING id, property_id, listing_type_id, contract_document, responsible_employee_id, client_id, contract_status_id, signed_date, start_date, end_date
`

type CreateContractParams struct {
	PropertyID            int64
	ListingTypeID         int64
	ContractDocument      string
	ResponsibleEmployeeID int64
	ClientID              int64
	ContractStatusID      int64
	SignedDate            time.Time
	StartDate             time.Time
	EndDate               time.Time
}

func (q *Queries) CreateContract(ctx context.Context, arg CreateContractParams) (Contract, error) {
	row := q.db.QueryRowContext(ctx, createContract,
		arg.PropertyID,
		arg.ListingTypeID,
		arg.ContractDocument,
		arg.ResponsibleEmployeeID,
		arg.ClientID,
		arg.ContractStatusID,
		arg.SignedDate,
		arg.StartDate,
		arg.EndDate,
	)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.PropertyID,
		&i.ListingTypeID,
		&i.ContractDocument,
		&i.ResponsibleEmployeeID,
		&i.ClientID,
		&i.ContractStatusID,
		&i.SignedDate,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO employee (
  first_name, last_name, start_date, end_date, job_title
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING id, first_name, last_name, start_date, end_date, job_title
`

type CreateEmployeeParams struct {
	FirstName string
	LastName  string
	StartDate time.Time
	EndDate   sql.NullTime
	JobTitle  string
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error) {
	row := q.db.QueryRowContext(ctx, createEmployee,
		arg.FirstName,
		arg.LastName,
		arg.StartDate,
		arg.EndDate,
		arg.JobTitle,
	)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.StartDate,
		&i.EndDate,
		&i.JobTitle,
	)
	return i, err
}

const createFeature = `-- name: CreateFeature :one
INSERT INTO feature (feature_name)
VALUES (?)
RETURNING id, feature_name
`

func (q *Queries) CreateFeature(ctx context.Context, featureName string) (Feature, error) {
	row := q.db.QueryRowContext(ctx, createFeature, featureName)
	var i Feature
	err := row.Scan(&i.ID, &i.FeatureName)
	return i, err
}

const createListing = `-- name: CreateListing :one
INSERT INTO listing (
  property_id, listing_status_id, listing_type_id, price, created_date
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING id, property_id, listing_status_id, listing_type_id, price, created_date
`

type CreateListingParams struct {
	PropertyID      int64
	ListingStatusID int64
	ListingTypeID   int64
	Price           int64
	CreatedDate     time.Time
}

func (q *Queries) CreateListing(ctx context.Context, arg CreateListingParams) (Listing, error) {
	row := q.db.QueryRowContext(ctx, createListing,
		arg.PropertyID,
		arg.ListingStatusID,
		arg.ListingTypeID,
		arg.Price,
		arg.CreatedDate,
	)
	var i Listing
	err := row.Scan(
		&i.ID,
		&i.PropertyID,
		&i.ListingStatusID,
		&i.ListingTypeID,
		&i.Price,
		&i.CreatedDate,
	)
	return i, err
}

const createOffer = `-- name: CreateOffer :one
INSERT INTO offer (
  client_id, property_id, offer_status_id, offer_amount
) VALUES (
  ?, ?, ?, ?
)
RETURNING id, client_id, property_id, offer_status_id, offer_amount
`

type CreateOfferParams struct {
	ClientID      int64
	PropertyID    int64
	OfferStatusID int64
	OfferAmount   int64
}

func (q *Queries) CreateOffer(ctx context.Context, arg CreateOfferParams) (Offer, error) {
	row := q.db.QueryRowContext(ctx, createOffer,
		arg.ClientID,
		arg.PropertyID,
		arg.OfferStatusID,
		arg.OfferAmount,
	)
	var i Offer
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.PropertyID,
		&i.OfferStatusID,
		&i.OfferAmount,
	)
	return i, err
}

const createProperty = `-- name: CreateProperty :one
INSERT INTO property (
  address_line1, address_line2, city, region, property_type_id, property_size, block_size, num_bedrooms, num_bathrooms, num_carspaces, description
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING id, address_line1, address_line2, city, region, property_type_id, property_size, block_size, num_bedrooms, num_bathrooms, num_carspaces, description
`

type CreatePropertyParams struct {
	AddressLine1   string
	AddressLine2   sql.NullString
	City           string
	Region         string
	PropertyTypeID int64
	PropertySize   int64
	BlockSize      int64
	NumBedrooms    int64
	NumBathrooms   int64
	NumCarspaces   int64
	Description    sql.NullString
}

func (q *Queries) CreateProperty(ctx context.Context, arg CreatePropertyParams) (Property, error) {
	row := q.db.QueryRowContext(ctx, createProperty,
		arg.AddressLine1,
		arg.AddressLine2,
		arg.City,
		arg.Region,
		arg.PropertyTypeID,
		arg.PropertySize,
		arg.BlockSize,
		arg.NumBedrooms,
		arg.NumBathrooms,
		arg.NumCarspaces,
		arg.Description,
	)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.Region,
		&i.PropertyTypeID,
		&i.PropertySize,
		&i.BlockSize,
		&i.NumBedrooms,
		&i.NumBathrooms,
		&i.NumCarspaces,
		&i.Description,
	)
	return i, err
}

const createPropertyFeature = `-- name: CreatePropertyFeature :exec
INSERT INTO property_feature (
  property_id, feature_id
) VALUES (
  ?, ?
)
`

type CreatePropertyFeatureParams struct {
	PropertyID sql.NullInt64
	FeatureID  sql.NullInt64
}

func (q *Queries) CreatePropertyFeature(ctx context.Context, arg CreatePropertyFeatureParams) error {
	_, err := q.db.ExecContext(ctx, createPropertyFeature, arg.PropertyID, arg.FeatureID)
	return err
}

const createPropertyType = `-- name: CreatePropertyType :one
INSERT INTO property_type (description)
VALUES (?)
RETURNING id, description
`

func (q *Queries) CreatePropertyType(ctx context.Context, description string) (PropertyType, error) {
	row := q.db.QueryRowContext(ctx, createPropertyType, description)
	var i PropertyType
	err := row.Scan(&i.ID, &i.Description)
	return i, err
}

const deleteClient = `-- name: DeleteClient :exec
DELETE FROM client
WHERE id = ?
`

func (q *Queries) DeleteClient(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteClient, id)
	return err
}

const deleteContract = `-- name: DeleteContract :exec
DELETE FROM contract
WHERE id = ?
`

func (q *Queries) DeleteContract(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteContract, id)
	return err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE id = ?
`

func (q *Queries) DeleteEmployee(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, id)
	return err
}

const deleteFeature = `-- name: DeleteFeature :exec
DELETE FROM feature
WHERE id = ?
`

func (q *Queries) DeleteFeature(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFeature, id)
	return err
}

const deleteListing = `-- name: DeleteListing :exec
DELETE FROM listing
WHERE id = ?
`

func (q *Queries) DeleteListing(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteListing, id)
	return err
}

const deleteOffer = `-- name: DeleteOffer :exec
DELETE FROM offer
WHERE id = ?
`

func (q *Queries) DeleteOffer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOffer, id)
	return err
}

const deleteProperty = `-- name: DeleteProperty :exec
DELETE FROM property
WHERE id = ?
`

func (q *Queries) DeleteProperty(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProperty, id)
	return err
}

const deletePropertyFeature = `-- name: DeletePropertyFeature :exec
DELETE FROM property_feature
WHERE property_id = ? AND feature_id = ?
`

type DeletePropertyFeatureParams struct {
	PropertyID sql.NullInt64
	FeatureID  sql.NullInt64
}

func (q *Queries) DeletePropertyFeature(ctx context.Context, arg DeletePropertyFeatureParams) error {
	_, err := q.db.ExecContext(ctx, deletePropertyFeature, arg.PropertyID, arg.FeatureID)
	return err
}

const deletePropertyType = `-- name: DeletePropertyType :exec
DELETE FROM property_type
WHERE id = ?
`

func (q *Queries) DeletePropertyType(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePropertyType, id)
	return err
}

const getClient = `-- name: GetClient :one
SELECT id, first_name, last_name, email_address, phone_number FROM client
WHERE id = ? LIMIT 1
`

func (q *Queries) GetClient(ctx context.Context, id int64) (Client, error) {
	row := q.db.QueryRowContext(ctx, getClient, id)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.EmailAddress,
		&i.PhoneNumber,
	)
	return i, err
}

const getContract = `-- name: GetContract :one
SELECT id, property_id, listing_type_id, contract_document, responsible_employee_id, client_id, contract_status_id, signed_date, start_date, end_date FROM contract
WHERE id = ? LIMIT 1
`

func (q *Queries) GetContract(ctx context.Context, id int64) (Contract, error) {
	row := q.db.QueryRowContext(ctx, getContract, id)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.PropertyID,
		&i.ListingTypeID,
		&i.ContractDocument,
		&i.ResponsibleEmployeeID,
		&i.ClientID,
		&i.ContractStatusID,
		&i.SignedDate,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const getEmployee = `-- name: GetEmployee :one
SELECT id, first_name, last_name, start_date, end_date, job_title FROM employee
WHERE id = ? LIMIT 1
`

func (q *Queries) GetEmployee(ctx context.Context, id int64) (Employee, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.StartDate,
		&i.EndDate,
		&i.JobTitle,
	)
	return i, err
}

const getFeature = `-- name: GetFeature :one
SELECT id, feature_name FROM feature
WHERE id = ? LIMIT 1
`

func (q *Queries) GetFeature(ctx context.Context, id int64) (Feature, error) {
	row := q.db.QueryRowContext(ctx, getFeature, id)
	var i Feature
	err := row.Scan(&i.ID, &i.FeatureName)
	return i, err
}

const getListing = `-- name: GetListing :one
SELECT id, property_id, listing_status_id, listing_type_id, price, created_date FROM listing
WHERE id = ? LIMIT 1
`

func (q *Queries) GetListing(ctx context.Context, id int64) (Listing, error) {
	row := q.db.QueryRowContext(ctx, getListing, id)
	var i Listing
	err := row.Scan(
		&i.ID,
		&i.PropertyID,
		&i.ListingStatusID,
		&i.ListingTypeID,
		&i.Price,
		&i.CreatedDate,
	)
	return i, err
}

const getOffer = `-- name: GetOffer :one
SELECT id, client_id, property_id, offer_status_id, offer_amount FROM offer
WHERE id = ? LIMIT 1
`

func (q *Queries) GetOffer(ctx context.Context, id int64) (Offer, error) {
	row := q.db.QueryRowContext(ctx, getOffer, id)
	var i Offer
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.PropertyID,
		&i.OfferStatusID,
		&i.OfferAmount,
	)
	return i, err
}

const getProperty = `-- name: GetProperty :one
SELECT id, address_line1, address_line2, city, region, property_type_id, property_size, block_size, num_bedrooms, num_bathrooms, num_carspaces, description FROM property
WHERE id = ? LIMIT 1
`

func (q *Queries) GetProperty(ctx context.Context, id int64) (Property, error) {
	row := q.db.QueryRowContext(ctx, getProperty, id)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.Region,
		&i.PropertyTypeID,
		&i.PropertySize,
		&i.BlockSize,
		&i.NumBedrooms,
		&i.NumBathrooms,
		&i.NumCarspaces,
		&i.Description,
	)
	return i, err
}

const getPropertyFeature = `-- name: GetPropertyFeature :one
SELECT property_id, feature_id FROM property_feature
WHERE property_id = ? AND feature_id = ? LIMIT 1
`

type GetPropertyFeatureParams struct {
	PropertyID sql.NullInt64
	FeatureID  sql.NullInt64
}

func (q *Queries) GetPropertyFeature(ctx context.Context, arg GetPropertyFeatureParams) (PropertyFeature, error) {
	row := q.db.QueryRowContext(ctx, getPropertyFeature, arg.PropertyID, arg.FeatureID)
	var i PropertyFeature
	err := row.Scan(&i.PropertyID, &i.FeatureID)
	return i, err
}

const getPropertyType = `-- name: GetPropertyType :one
SELECT id, description FROM property_type
WHERE id = ? LIMIT 1
`

func (q *Queries) GetPropertyType(ctx context.Context, id int64) (PropertyType, error) {
	row := q.db.QueryRowContext(ctx, getPropertyType, id)
	var i PropertyType
	err := row.Scan(&i.ID, &i.Description)
	return i, err
}

const listClients = `-- name: ListClients :many
SELECT id, first_name, last_name, email_address, phone_number FROM client
ORDER BY first_name
`

func (q *Queries) ListClients(ctx context.Context) ([]Client, error) {
	rows, err := q.db.QueryContext(ctx, listClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.EmailAddress,
			&i.PhoneNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listContracts = `-- name: ListContracts :many
SELECT id, property_id, listing_type_id, contract_document, responsible_employee_id, client_id, contract_status_id, signed_date, start_date, end_date FROM contract
ORDER BY signed_date
`

func (q *Queries) ListContracts(ctx context.Context) ([]Contract, error) {
	rows, err := q.db.QueryContext(ctx, listContracts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Contract
	for rows.Next() {
		var i Contract
		if err := rows.Scan(
			&i.ID,
			&i.PropertyID,
			&i.ListingTypeID,
			&i.ContractDocument,
			&i.ResponsibleEmployeeID,
			&i.ClientID,
			&i.ContractStatusID,
			&i.SignedDate,
			&i.StartDate,
			&i.EndDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEmployees = `-- name: ListEmployees :many
SELECT id, first_name, last_name, start_date, end_date, job_title FROM employee
ORDER BY first_name
`

func (q *Queries) ListEmployees(ctx context.Context) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, listEmployees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.StartDate,
			&i.EndDate,
			&i.JobTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFeatures = `-- name: ListFeatures :many
SELECT id, feature_name FROM feature
ORDER BY feature_name
`

func (q *Queries) ListFeatures(ctx context.Context) ([]Feature, error) {
	rows, err := q.db.QueryContext(ctx, listFeatures)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feature
	for rows.Next() {
		var i Feature
		if err := rows.Scan(&i.ID, &i.FeatureName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listListings = `-- name: ListListings :many
SELECT id, property_id, listing_status_id, listing_type_id, price, created_date FROM listing
ORDER BY created_date
`

func (q *Queries) ListListings(ctx context.Context) ([]Listing, error) {
	rows, err := q.db.QueryContext(ctx, listListings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Listing
	for rows.Next() {
		var i Listing
		if err := rows.Scan(
			&i.ID,
			&i.PropertyID,
			&i.ListingStatusID,
			&i.ListingTypeID,
			&i.Price,
			&i.CreatedDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOffers = `-- name: ListOffers :many
SELECT id, client_id, property_id, offer_status_id, offer_amount FROM offer
ORDER BY offer_amount
`

func (q *Queries) ListOffers(ctx context.Context) ([]Offer, error) {
	rows, err := q.db.QueryContext(ctx, listOffers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Offer
	for rows.Next() {
		var i Offer
		if err := rows.Scan(
			&i.ID,
			&i.ClientID,
			&i.PropertyID,
			&i.OfferStatusID,
			&i.OfferAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProperties = `-- name: ListProperties :many
SELECT id, address_line1, address_line2, city, region, property_type_id, property_size, block_size, num_bedrooms, num_bathrooms, num_carspaces, description FROM property
ORDER BY address_line1
`

func (q *Queries) ListProperties(ctx context.Context) ([]Property, error) {
	rows, err := q.db.QueryContext(ctx, listProperties)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Property
	for rows.Next() {
		var i Property
		if err := rows.Scan(
			&i.ID,
			&i.AddressLine1,
			&i.AddressLine2,
			&i.City,
			&i.Region,
			&i.PropertyTypeID,
			&i.PropertySize,
			&i.BlockSize,
			&i.NumBedrooms,
			&i.NumBathrooms,
			&i.NumCarspaces,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPropertyTypes = `-- name: ListPropertyTypes :many
SELECT id, description FROM property_type
ORDER BY description
`

func (q *Queries) ListPropertyTypes(ctx context.Context) ([]PropertyType, error) {
	rows, err := q.db.QueryContext(ctx, listPropertyTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PropertyType
	for rows.Next() {
		var i PropertyType
		if err := rows.Scan(&i.ID, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateClient = `-- name: UpdateClient :exec
UPDATE client
SET first_name = ?,
    last_name = ?,
    email_address = ?,
    phone_number = ?
WHERE id = ?
`

type UpdateClientParams struct {
	FirstName    string
	LastName     string
	EmailAddress string
	PhoneNumber  string
	ID           int64
}

func (q *Queries) UpdateClient(ctx context.Context, arg UpdateClientParams) error {
	_, err := q.db.ExecContext(ctx, updateClient,
		arg.FirstName,
		arg.LastName,
		arg.EmailAddress,
		arg.PhoneNumber,
		arg.ID,
	)
	return err
}

const updateContract = `-- name: UpdateContract :exec
UPDATE contract
SET property_id = ?,
    listing_type_id = ?,
    contract_document = ?,
    responsible_employee_id = ?,
    client_id = ?,
    contract_status_id = ?,
    signed_date = ?,
    start_date = ?,
    end_date = ?
WHERE id = ?
`

type UpdateContractParams struct {
	PropertyID            int64
	ListingTypeID         int64
	ContractDocument      string
	ResponsibleEmployeeID int64
	ClientID              int64
	ContractStatusID      int64
	SignedDate            time.Time
	StartDate             time.Time
	EndDate               time.Time
	ID                    int64
}

func (q *Queries) UpdateContract(ctx context.Context, arg UpdateContractParams) error {
	_, err := q.db.ExecContext(ctx, updateContract,
		arg.PropertyID,
		arg.ListingTypeID,
		arg.ContractDocument,
		arg.ResponsibleEmployeeID,
		arg.ClientID,
		arg.ContractStatusID,
		arg.SignedDate,
		arg.StartDate,
		arg.EndDate,
		arg.ID,
	)
	return err
}

const updateEmployee = `-- name: UpdateEmployee :exec
UPDATE employee
SET first_name = ?,
    last_name = ?,
    start_date = ?,
    end_date = ?,
    job_title = ?
WHERE id = ?
`

type UpdateEmployeeParams struct {
	FirstName string
	LastName  string
	StartDate time.Time
	EndDate   sql.NullTime
	JobTitle  string
	ID        int64
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployee,
		arg.FirstName,
		arg.LastName,
		arg.StartDate,
		arg.EndDate,
		arg.JobTitle,
		arg.ID,
	)
	return err
}

const updateFeature = `-- name: UpdateFeature :exec
UPDATE feature
SET feature_name = ?
WHERE id = ?
`

type UpdateFeatureParams struct {
	FeatureName string
	ID          int64
}

func (q *Queries) UpdateFeature(ctx context.Context, arg UpdateFeatureParams) error {
	_, err := q.db.ExecContext(ctx, updateFeature, arg.FeatureName, arg.ID)
	return err
}

const updateListing = `-- name: UpdateListing :exec
UPDATE listing
SET property_id = ?,
    listing_status_id = ?,
    listing_type_id = ?,
    price = ?,
    created_date = ?
WHERE id = ?
`

type UpdateListingParams struct {
	PropertyID      int64
	ListingStatusID int64
	ListingTypeID   int64
	Price           int64
	CreatedDate     time.Time
	ID              int64
}

func (q *Queries) UpdateListing(ctx context.Context, arg UpdateListingParams) error {
	_, err := q.db.ExecContext(ctx, updateListing,
		arg.PropertyID,
		arg.ListingStatusID,
		arg.ListingTypeID,
		arg.Price,
		arg.CreatedDate,
		arg.ID,
	)
	return err
}

const updateOffer = `-- name: UpdateOffer :exec
UPDATE offer
SET client_id = ?,
    property_id = ?,
    offer_status_id = ?,
    offer_amount = ?
WHERE id = ?
`

type UpdateOfferParams struct {
	ClientID      int64
	PropertyID    int64
	OfferStatusID int64
	OfferAmount   int64
	ID            int64
}

func (q *Queries) UpdateOffer(ctx context.Context, arg UpdateOfferParams) error {
	_, err := q.db.ExecContext(ctx, updateOffer,
		arg.ClientID,
		arg.PropertyID,
		arg.OfferStatusID,
		arg.OfferAmount,
		arg.ID,
	)
	return err
}

const updateProperty = `-- name: UpdateProperty :exec
UPDATE property
SET address_line1 = ?,
    address_line2 = ?,
    city = ?,
    region = ?,
    property_type_id = ?,
    property_size = ?,
    block_size = ?,
    num_bedrooms = ?,
    num_bathrooms = ?,
    num_carspaces = ?,
    description = ?
WHERE id = ?
`

type UpdatePropertyParams struct {
	AddressLine1   string
	AddressLine2   sql.NullString
	City           string
	Region         string
	PropertyTypeID int64
	PropertySize   int64
	BlockSize      int64
	NumBedrooms    int64
	NumBathrooms   int64
	NumCarspaces   int64
	Description    sql.NullString
	ID             int64
}

func (q *Queries) UpdateProperty(ctx context.Context, arg UpdatePropertyParams) error {
	_, err := q.db.ExecContext(ctx, updateProperty,
		arg.AddressLine1,
		arg.AddressLine2,
		arg.City,
		arg.Region,
		arg.PropertyTypeID,
		arg.PropertySize,
		arg.BlockSize,
		arg.NumBedrooms,
		arg.NumBathrooms,
		arg.NumCarspaces,
		arg.Description,
		arg.ID,
	)
	return err
}

const updatePropertyType = `-- name: UpdatePropertyType :exec
UPDATE property_type
SET description = ?
WHERE id = ?
`

type UpdatePropertyTypeParams struct {
	Description string
	ID          int64
}

func (q *Queries) UpdatePropertyType(ctx context.Context, arg UpdatePropertyTypeParams) error {
	_, err := q.db.ExecContext(ctx, updatePropertyType, arg.Description, arg.ID)
	return err
}
-- name: GetProperty :one
SELECT * FROM property
WHERE id = ? LIMIT 1;

-- name: ListProperties :many
SELECT * FROM property
ORDER BY address_line1;

-- name: CreateProperty :one
INSERT INTO property (
  address_line1, address_line2, city, region, property_type_id, property_size, block_size, num_bedrooms, num_bathrooms, num_carspaces, description
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateProperty :exec
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
WHERE id = ?;

-- name: DeleteProperty :exec
DELETE FROM property
WHERE id = ?;

-- name: GetPropertyType :one
SELECT * FROM property_type
WHERE id = ? LIMIT 1;

-- name: ListPropertyTypes :many
SELECT * FROM property_type
ORDER BY description;

-- name: CreatePropertyType :one
INSERT INTO property_type (description)
VALUES (?)
RETURNING *;

-- name: UpdatePropertyType :exec
UPDATE property_type
SET description = ?
WHERE id = ?;

-- name: DeletePropertyType :exec
DELETE FROM property_type
WHERE id = ?;

-- name: GetListing :one
SELECT * FROM listing
WHERE id = ? LIMIT 1;

-- name: ListListings :many
SELECT * FROM listing
ORDER BY created_date;

-- name: CreateListing :one
INSERT INTO listing (
  property_id, listing_status_id, listing_type_id, price, created_date
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateListing :exec
UPDATE listing
SET property_id = ?,
    listing_status_id = ?,
    listing_type_id = ?,
    price = ?,
    created_date = ?
WHERE id = ?;

-- name: DeleteListing :exec
DELETE FROM listing
WHERE id = ?;

-- name: GetFeature :one
SELECT * FROM feature
WHERE id = ? LIMIT 1;

-- name: ListFeatures :many
SELECT * FROM feature
ORDER BY feature_name;

-- name: CreateFeature :one
INSERT INTO feature (feature_name)
VALUES (?)
RETURNING *;

-- name: UpdateFeature :exec
UPDATE feature
SET feature_name = ?
WHERE id = ?;

-- name: DeleteFeature :exec
DELETE FROM feature
WHERE id = ?;

-- name: GetPropertyFeature :one
SELECT * FROM property_feature
WHERE property_id = ? AND feature_id = ? LIMIT 1;

-- name: CreatePropertyFeature :exec
INSERT INTO property_feature (
  property_id, feature_id
) VALUES (
  ?, ?
);

-- name: DeletePropertyFeature :exec
DELETE FROM property_feature
WHERE property_id = ? AND feature_id = ?;

-- name: GetClient :one
SELECT * FROM client
WHERE id = ? LIMIT 1;

-- name: ListClients :many
SELECT * FROM client
ORDER BY first_name;

-- name: CreateClient :one
INSERT INTO client (
  first_name, last_name, email_address, phone_number
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateClient :exec
UPDATE client
SET first_name = ?,
    last_name = ?,
    email_address = ?,
    phone_number = ?
WHERE id = ?;

-- name: DeleteClient :exec
DELETE FROM client
WHERE id = ?;

-- name: GetEmployee :one
SELECT * FROM employee
WHERE id = ? LIMIT 1;

-- name: ListEmployees :many
SELECT * FROM employee
ORDER BY first_name;

-- name: CreateEmployee :one
INSERT INTO employee (
  first_name, last_name, start_date, end_date, job_title
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateEmployee :exec
UPDATE employee
SET first_name = ?,
    last_name = ?,
    start_date = ?,
    end_date = ?,
    job_title = ?
WHERE id = ?;

-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE id = ?;

-- name: GetOffer :one
SELECT * FROM offer
WHERE id = ? LIMIT 1;

-- name: ListOffers :many
SELECT * FROM offer
ORDER BY offer_amount;

-- name: CreateOffer :one
INSERT INTO offer (
  client_id, property_id, offer_status_id, offer_amount
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateOffer :exec
UPDATE offer
SET client_id = ?,
    property_id = ?,
    offer_status_id = ?,
    offer_amount = ?
WHERE id = ?;

-- name: DeleteOffer :exec
DELETE FROM offer
WHERE id = ?;

-- name: GetContract :one
SELECT * FROM contract
WHERE id = ? LIMIT 1;

-- name: ListContracts :many
SELECT * FROM contract
ORDER BY signed_date;

-- name: CreateContract :one
INSERT INTO contract (
  property_id, listing_type_id, contract_document, responsible_employee_id, client_id, contract_status_id, signed_date, start_date, end_date
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateContract :exec
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
WHERE id = ?;

-- name: DeleteContract :exec
DELETE FROM contract
WHERE id = ?;


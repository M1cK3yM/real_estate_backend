-- Table property
CREATE TABLE property (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    address_line1 VARCHAR NOT NULL,
    address_line2 VARCHAR,
    city VARCHAR NOT NULL,
    region VARCHAR NOT NULL,
    property_type_id INTEGER NOT NULL,
    property_size INTEGER NOT NULL,
    block_size INTEGER NOT NULL,
    num_bedrooms INTEGER NOT NULL,
    num_bathrooms INTEGER NOT NULL,
    num_carspaces INTEGER NOT NULL,
    description VARCHAR,
    FOREIGN KEY (property_type_id) REFERENCES property_type(id)
);

-- Table property_type
CREATE TABLE property_type (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description VARCHAR NOT NULL
);

-- Table listing_type
CREATE TABLE listing_type (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description VARCHAR NOT NULL
);

-- Table listing_status
CREATE TABLE listing_status (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description VARCHAR NOT NULL
);

-- Table feature
CREATE TABLE feature (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    feature_name VARCHAR NOT NULL
);

-- Table property_feature
CREATE TABLE property_feature (
    property_id INTEGER,
    feature_id INTEGER,
    PRIMARY KEY (property_id, feature_id),
    FOREIGN KEY (property_id) REFERENCES property(id),
    FOREIGN KEY (feature_id) REFERENCES feature(id)
);

-- Table listing
CREATE TABLE listing (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    property_id INTEGER NOT NULL,
    listing_status_id INTEGER NOT NULL,
    listing_type_id INTEGER NOT NULL,
    price INTEGER NOT NULL,
    created_date DATE NOT NULL,
    FOREIGN KEY (property_id) REFERENCES property(id),
    FOREIGN KEY (listing_status_id) REFERENCES listing_status(id),
    FOREIGN KEY (listing_type_id) REFERENCES listing_type(id)
);

-- Table employee
CREATE TABLE employee (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    job_title VARCHAR NOT NULL
);

-- Table property_employee
CREATE TABLE property_employee (
    property_id INTEGER,
    employee_id INTEGER,
    role_type_id INTEGER,
    start_date DATE NOT NULL,
    end_date DATE,
    PRIMARY KEY (property_id, employee_id, role_type_id),
    FOREIGN KEY (property_id) REFERENCES property(id),
    FOREIGN KEY (employee_id) REFERENCES employee(id),
    FOREIGN KEY (role_type_id) REFERENCES role_type(id)
);

-- Table role_type
CREATE TABLE role_type (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description VARCHAR NOT NULL
);

-- Table inspection
CREATE TABLE inspection (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    property_id INTEGER NOT NULL,
    inspection_datetime DATETIME NOT NULL,
    responsible_employee_id INTEGER NOT NULL,
    FOREIGN KEY (property_id) REFERENCES property(id),
    FOREIGN KEY (responsible_employee_id) REFERENCES employee(id)
);

-- Table client
CREATE TABLE client (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email_address VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL
);

-- Table client_property_interest
CREATE TABLE client_property_interest (
    client_id INTEGER,
    property_id INTEGER,
    PRIMARY KEY (client_id, property_id),
    FOREIGN KEY (client_id) REFERENCES client(id),
    FOREIGN KEY (property_id) REFERENCES property(id)
);

-- Table client_inspection
CREATE TABLE client_inspection (
    client_id INTEGER,
    inspection_id INTEGER,
    PRIMARY KEY (client_id, inspection_id),
    FOREIGN KEY (client_id) REFERENCES client(id),
    FOREIGN KEY (inspection_id) REFERENCES inspection(id)
);

-- Table offer
CREATE TABLE offer (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    property_id INTEGER NOT NULL,
    offer_status_id INTEGER NOT NULL,
    offer_amount INTEGER NOT NULL,
    FOREIGN KEY (client_id) REFERENCES client(id),
    FOREIGN KEY (property_id) REFERENCES property(id),
    FOREIGN KEY (offer_status_id) REFERENCES offer_status(id)
);

-- Table offer_status
CREATE TABLE offer_status (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description VARCHAR NOT NULL
);

-- Table contract
CREATE TABLE contract (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    property_id INTEGER NOT NULL,
    listing_type_id INTEGER NOT NULL,
    contract_document VARCHAR NOT NULL,
    responsible_employee_id INTEGER NOT NULL,
    client_id INTEGER NOT NULL,
    contract_status_id INTEGER NOT NULL,
    signed_date DATE NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    FOREIGN KEY (property_id) REFERENCES property(id),
    FOREIGN KEY (listing_type_id) REFERENCES listing_type(id),
    FOREIGN KEY (responsible_employee_id) REFERENCES employee(id),
    FOREIGN KEY (client_id) REFERENCES client(id),
    FOREIGN KEY (contract_status_id) REFERENCES contract_status(id)
);

-- Table contract_status
CREATE TABLE contract_status (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description VARCHAR NOT NULL
);


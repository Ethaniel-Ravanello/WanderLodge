-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    "firstName" VARCHAR(255),
    lastName VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    phoneNumber VARCHAR(50) UNIQUE,
    roles VARCHAR(50),
    password VARCHAR(255) NOT NULL,
);

CREATE TABLE IF NOT EXISTS listings
(
    id SERIAL PRIMARY KEY,
    hostId INTEGER,
    title VARCHAR(255),
    description TEXT,
    location VARCHAR(255),
    address VARCHAR(255),
    maxPeople INTEGER,
    pricePerNight INTEGER,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    approvalStatus VARCHAR(50),
    CONSTRAINT fk_host_id FOREIGN KEY (hostId) REFRENCE user(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS bookings
(
    id SERIAL PRIMARY KEY,
    guestId INTEGER,
    listingId INTEGER,
    persons INTEGER,
    startDate TIMESTAMP,
    endDate TIMESTAMP,
    status VARCHAR(50)
    CONSTRAINT fk_guest_id FOREIGN KEY (hostId) REFRENCE user(id) ON DELETE CASCADE
    CONSTRAINT fk_list_id FOREIGN KEY (hostId) REFRENCE listing(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS approvals
(
    id SERIAL PRIMARY KEY NOT NULL,
    approvalTypeId INTEGER,
    approvalType VARCHAR(50),
    approverId INTEGER,
    status VARCHAR(50),
    createdAt TIMESTAMP,
    updatedAt TIMESTAMP
    CONSTRAINT fk_approval_approver FOREIGN KEY (approverId) REFERENCES user(id) ON DELETE SET NULL
);
-- +migrate StatementEnd


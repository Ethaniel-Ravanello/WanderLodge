-- +migrate Up
-- +migrate StatementBegin

-- Rename tables to use camelCase
ALTER TABLE users RENAME TO "users";
ALTER TABLE listings RENAME TO "listings";
ALTER TABLE bookings RENAME TO "bookings";
ALTER TABLE approvals RENAME TO "approvals";

-- Rename columns to enforce camelCase
ALTER TABLE "users" RENAME COLUMN firstname TO "firstName";
ALTER TABLE "users" RENAME COLUMN lastname TO "lastName";
ALTER TABLE "users" RENAME COLUMN phonenumber TO "phoneNumber";

ALTER TABLE "listings" RENAME COLUMN hostid TO "hostId";
ALTER TABLE "listings" RENAME COLUMN maxpeople TO "maxPeople";
ALTER TABLE "listings" RENAME COLUMN pricepernight TO "pricePerNight";
ALTER TABLE "listings" RENAME COLUMN createdat TO "createdAt";
ALTER TABLE "listings" RENAME COLUMN approvalstatus TO "approvalStatus";

ALTER TABLE "bookings" RENAME COLUMN guestid TO "guestId";
ALTER TABLE "bookings" RENAME COLUMN listingid TO "listingId";
ALTER TABLE "bookings" RENAME COLUMN startdate TO "startDate";
ALTER TABLE "bookings" RENAME COLUMN enddate TO "endDate";

ALTER TABLE "approvals" RENAME COLUMN approvaltypeid TO "approvalTypeId";
ALTER TABLE "approvals" RENAME COLUMN approvaltype TO "approvalType";
ALTER TABLE "approvals" RENAME COLUMN approverid TO "approverId";
ALTER TABLE "approvals" RENAME COLUMN createdat TO "createdAt";
ALTER TABLE "approvals" RENAME COLUMN updatedat TO "updatedAt";

-- +migrate StatementEnd

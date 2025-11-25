-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE Bab7Servo(
    code INTEGER PRIMARY KEY,
    StatusServo INTEGER
);

-- +migrate StatementEnd
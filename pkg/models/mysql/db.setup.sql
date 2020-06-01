CREATE DATABASE IF NOT EXISTS bcnchess CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE bcnches;

CREATE TABLE tournaments (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
  	location VARCHAR(255) NOT NULL,
    tournamentDate DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    matchTimeStart DATETIME NOT NULL, 
    matchTimeEnd DATETIME NOT NULL,
    additionalInformation TEXT NOT NULL,
    isOnline TINYINT NOT NULL DEFAULT 0,
  	timeControl VARCHAR(100) NOT NULL,
  	tournamentType VARCHAR(100) NOT NULL,
  	rated TINYINT NOT NULL DEFAULT 0,
  	poster VARCHAR(200) NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_tournaments ON tournaments(created);

INSERT INTO tournaments (title, location, matchTimeStart, matchTimeEnd, additionalInformation, isOnline, timeControl, tournamentType, rated, poster, created, expires) VALUES (
    'cool tournament one',
    'https://lichess.org/pEQoPXt5F2mz',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY),
  	'additional information about 1',
  	1,
  	'3+2',
  	'Round Robin',
  	1,
  	"./ui/static/image/logo.png",
  	UTC_TIMESTAMP(),
  	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 1 MINUTE)
);

INSERT INTO tournaments (title, location, matchTimeStart, matchTimeEnd, additionalInformation, isOnline, timeControl, tournamentType, rated, poster, created, expires) VALUES (
    'cool tournament two',
    'https://lichess.org/pEQoPXt5F2mz',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY),
  	'additional information about 2',
  	1,
  	'5+0',
  	'Swiss',
  	0,
  	"./ui/static/image/logo.png",
  	UTC_TIMESTAMP(),
  	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 1 MINUTE)
);



CREATE USER IF NOT EXISTS 'chess-web'@'localhost' IDENTIFIED BY 'password';

GRANT SELECT, INSERT, UPDATE ON bcnchess.* TO 'chess-web'@'localhost';

CREATE TABLE users(
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  firstName VARCHAR(255) NOT NULL,
  lastName VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password CHAR(60) NOT NULL,
  club VARCHAR(255) NOT NULL,
  eloStandard VARCHAR(255) NOT NULL DEFAULT "na",
  eloRapid VARCHAR(255) NOT NULL DEFAULT "na",
  lichessUserName VARCHAR(255) NOT NULL DEFAULT "na",
  chesscomUserName VARCHAR(255) NOT NULL DEFAULT "na",
  created DATETIME NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_us_email UNIQUE (email);




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
  	rated BOOLEAN NOT NULL DEFAULT false,
  	poster VARCHAR(200) NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_tournaments_created ON tournaments(created);

CREATE TABLE users(
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  firstName VARCHAR(255) NOT NULL,
  lastName VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password CHAR(60) NOT NULL,
  club VARCHAR(255) NOT NULL,
  eloStandard VARCHAR(255) NOT NULL DEFAULT "na",
  eloRapid VARCHAR(255) NOT NULL DEFAULT "na",
  lichessUsername VARCHAR(255) NOT NULL DEFAULT "na",
  chesscomUsername VARCHAR(255) NOT NULL DEFAULT "na",
  created DATETIME NOT NULL,
  active BOOLEAN NOT NULL DEFAULT true 
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users (firstName, lastName, email, password, club, eloStandard, eloRapid, lichessUsername, chesscomUsername, created) VALUES (
    'John',
    'Doe',
    'john@example.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    'Congres C.E',
    '1700',
    '1700',
    'na',
    'na',
    '2018-12-23 17:25:22'
); 

-- CREATE DATABASE test_bcnchess CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
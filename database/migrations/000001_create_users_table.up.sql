CREATE TABLE IF NOT EXISTS users(
   id VARCHAR (30) PRIMARY KEY,
   username VARCHAR (30) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL
);
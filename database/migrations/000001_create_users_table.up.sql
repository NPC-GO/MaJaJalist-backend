CREATE TABLE IF NOT EXISTS users(
   id VARCHAR (30) PRIMARY KEY,
   nickName VARCHAR (30) NOT NULL,
   passWord VARCHAR (50) NOT NULL,
   email VARCHAR (100) UNIQUE NOT NULL,
   avatar VARCHAR (256) NOT NULL DEFAULT 'https://i.stack.imgur.com/34AD2.jpg',
   userLevel INT NOT NULL DEFAULT 0 ,
   special boolean NOT NULL DEFAULT FALSE ,
   token VARCHAR (30) NOT NULL UNIQUE ,
   emailVerified boolean NOT NULL DEFAULT FALSE ,
   lastLogin TIMESTAMP
);
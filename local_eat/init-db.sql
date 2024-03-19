CREATE TABLE producers (
  id      INT AUTO_INCREMENT PRIMARY KEY,
  name    VARCHAR(60) NOT NULL,
  picture VARCHAR(60),
  created DATE
);

CREATE TABLE users (
  username VARCHAR(60) PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  email    VARCHAR(60) NOT NULL,
  age      INT,
  gender   VARCHAR(1),
  address VARCHAR(60),
  locality VARCHAR(60),
  cellphone VARCHAR(60),
  CHECK (gender IN ('M', 'F'))
);

INSERT INTO producers(name, picture, created) VALUES 
('Alain', 'Alain.jpg', NOW()),
('Jacques', 'Jacques.jpg', NOW()),
('Marc', 'Marc.jpg', NOW()),
('Mirco', 'Mirco.jpg', NOW()),
('Tom', 'Tom.jpg', NOW());

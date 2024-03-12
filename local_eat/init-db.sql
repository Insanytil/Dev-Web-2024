CREATE TABLE producers (
  id      INT AUTO_INCREMENT PRIMARY KEY,
  name    VARCHAR(60) NOT NULL,
  picture VARCHAR(60),
  created DATE
);

INSERT INTO producers(name, picture, created) VALUES 
('Alain', 'Alain.jpg', NOW()),
('Jacques', 'Jacques.jpg', NOW()),
('Marc', 'Marc.jpg', NOW()),
('Mirco', 'Mirco.jpg', NOW()),
('Tom', 'Tom.jpg', NOW());

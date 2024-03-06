CREATE TABLE technologies (
  name    VARCHAR(255),
  details VARCHAR(255)
);
insert into technologies values (
  'Go', 'An open source programming language that makes it easy to build simple and efficient software.'
);
insert into technologies values (
  'JavaScript', 'A lightweight, interpreted, or just-in-time compiled programming language with first-class functions.'
);
insert into technologies values (
  'MySQL', 'A powerful, open source object-relational database'
);

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

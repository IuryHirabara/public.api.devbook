insert into users (name, nick, email, password)
values
('Usuario 1', 'usuario_1', 'usuario1@gmail.com', '$2a$10$W8T.romWx4m16nkPo9fL2OzrAu0n5k7JZ1CWdCaZv0ijVb577G1ZG'),
('Usuario 2', 'usuario_2', 'usuario2@gmail.com', '$2a$10$W8T.romWx4m16nkPo9fL2OzrAu0n5k7JZ1CWdCaZv0ijVb577G1ZG'),
('Usuario 3', 'usuario_3', 'usuario3@gmail.com', '$2a$10$W8T.romWx4m16nkPo9fL2OzrAu0n5k7JZ1CWdCaZv0ijVb577G1ZG'),
('Usuario 4', 'usuario_4', 'usuario4@gmail.com', '$2a$10$W8T.romWx4m16nkPo9fL2OzrAu0n5k7JZ1CWdCaZv0ijVb577G1ZG'),
('Usuario 5', 'usuario_5', 'usuario5@gmail.com', '$2a$10$W8T.romWx4m16nkPo9fL2OzrAu0n5k7JZ1CWdCaZv0ijVb577G1ZG');

insert into followers (userId, followerId)
values
(1,2),
(2,1),
(1,3),
(3,1),
(4,5),
(5,1);

insert into publications (title, content, authorId)
values
('Publicação do usuário 1', 'Essa é a publicação do usuário 1', 1),
('Publicação do usuário 2', 'Essa é a publicação do usuário 2', 2),
('Publicação do usuário 3', 'Essa é a publicação do usuário 3', 3),
('Publicação do usuário 4', 'Essa é a publicação do usuário 4', 4),
('Publicação do usuário 5', 'Essa é a publicação do usuário 5', 5);
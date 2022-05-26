insert into usuarios (nome, nick, email, senha)
values
("usuario1", "usuario1", "usuario1@gmail.com", "$2a$10$W9NOlZp/RcUxxhtpDD6RUe0NmOg/rrmKihahjForJirdk3QhEmlki"),
("usuario2", "usuario2", "usuario2@gmail.com", "$2a$10$W9NOlZp/RcUxxhtpDD6RUe0NmOg/rrmKihahjForJirdk3QhEmlki"),
("usuario3", "usuario3", "usuario3@gmail.com", "$2a$10$W9NOlZp/RcUxxhtpDD6RUe0NmOg/rrmKihahjForJirdk3QhEmlki"),
("usuario4", "usuario4", "usuario4@gmail.com", "$2a$10$W9NOlZp/RcUxxhtpDD6RUe0NmOg/rrmKihahjForJirdk3QhEmlki"),
("usuario5", "usuario5", "usuario5@gmail.com", "$2a$10$W9NOlZp/RcUxxhtpDD6RUe0NmOg/rrmKihahjForJirdk3QhEmlki");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(1, 3),
(1, 4),
(2, 3),
(2, 4),
(2, 5),
(3, 4),
(3, 5),
(3, 1),
(4, 5),
(4, 1),
(4, 2),
(5, 1),
(5, 2),
(5, 3);
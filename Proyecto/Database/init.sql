CREATE TABLE IF NOT EXISTS hotels (
	id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(350) NOT NULL UNIQUE,
    descripcion TEXT,
    email VARCHAR(150) NOT NULL UNIQUE,
    cant_hab INT NOT NULL,
    amenities VARCHAR(1000)
);


CREATE TABLE IF NOT EXISTS clientes (
	id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(350) NOT NULL,
    last_name VARCHAR(250) NOT NULL,
    user_name VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
);

CREATE TABLE IF NOT EXISTS admins (
	id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(350) NOT NULL,
    last_name VARCHAR(250) NOT NULL,
    user_name VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
);

CREATE TABLE IF NOT EXISTS reservas (
	id INT PRIMARY KEY AUTO_INCREMENT,
    hotel_id INT,
    cliente_id INT,
    anio_inicio VARCHAR(10),
    anio_final VARCHAR(10),
    mes_inicio VARCHAR(10),
    mes_final VARCHAR(10),
    dia_inicio VARCHAR(10),
    dia_final VARCHAR(10),
    dias VARCHAR(2),
    FOREIGN KEY (cliente_id) REFERENCES clientes(id),
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);


CREATE TABLE IF NOT EXISTS imagens (
	id INT PRIMARY KEY AUTO_INCREMENT,
    url VARCHAR(500) NOT NULL,
    hotel_id INT,
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)	
);


INSERT INTO hotels (nombre, descripcion, email, cant_hab, amenities) VALUES
('Hotel ABC', 'Un hotel de lujo ubicado en el corazón de la ciudad. Ofrecemos habitaciones cómodas y elegantes con vistas panorámicas, un restaurante gourmet y un spa de clase mundial.', 'info@hotelabc.com', 100, 'piscina, gimnasio, servicio de habitaciones, wifi gratis'),
('Hotel XYZ', 'Un hotel boutique con un ambiente moderno y sofisticado. Nuestras habitaciones cuentan con diseño vanguardista y comodidades de alta gama. Disfrute de nuestro bar en la azotea con vistas panorámicas.', 'info@hotelxyz.com', 50, 'bar en la azotea, centro de negocios, servicio de lavandería, transporte al aeropuerto'),
('Hotel 123', 'Un acogedor hotel familiar cerca de la playa. Ofrecemos habitaciones amplias y luminosas, un restaurante con menú infantil y actividades para niños. Ideal para unas vacaciones en familia.', 'info@hotel123.com', 80, 'piscina para niños, parque infantil, servicio de niñera, wifi gratis'),
('Hotel PQR', 'Un retiro tranquilo en medio de la naturaleza. Nuestras cabañas rústicas brindan una experiencia única con vistas al bosque y acceso a senderos para caminar. Disfrute de nuestro spa de bienestar y cocina orgánica.', 'info@hotelpqr.com', 20, 'spa, senderismo, restaurante orgánico, estacionamiento gratuito'),
('Hotel LMN', 'Un hotel histórico con encanto en el casco antiguo. Nuestras habitaciones están decoradas con elegancia y ofrecen vistas a los monumentos históricos. Disfrute de nuestro bar de cócteles vintage y de paseos por la ciudad.', 'info@hotellmn.com', 40, 'bar de cócteles, visitas guiadas, servicio de conserjería, servicio de despertador'),
('Hotel DEF', 'Un hotel de negocios con todas las comodidades necesarias. Nuestras habitaciones están equipadas con escritorios y acceso a internet de alta velocidad. Ofrecemos salas de reuniones y un centro de fitness 24/7.', 'info@hoteldef.com', 60, 'centro de negocios, sala de reuniones, gimnasio, servicio de lavandería');


INSERT INTO clientes (name, last_name, user_name, password, email) VALUES
('Tomas', 'Garbellotto', 'user1', '1234', 'tomasgarbellotto@gmail.com'), 
('Facundo', 'Gazzera', 'user2', '1234', 'facundogazzera@gmail.com'); 

INSERT INTO admins (name, last_name, user_name, password, email) VALUES
('Tomas', 'Garbellotto', 'user1', '1234', 'tomasgarbellotto@gmail.com'), 
('Facundo', 'Gazzera', 'user2', '1234', 'facundogazzera@gmail.com');


INSERT INTO reservas (hotel_id, cliente_id, anio_inicio, anio_final, mes_inicio, mes_final, dia_inicio, dia_final, dias) VALUES
(1, 1, '2023', '2023', '07', '07', '15', '20', '5'),
(2, 1, '2023', '2023', '08', '08', '10', '15', '5'),
(3, 2, '2023', '2023', '09', '09', '10', '15', '5'),
(4, 2, '2023', '2023', '10', '10', '05', '10', '5'),
(5, 2, '2023', '2023', '11', '11', '15', '20', '5');

INSERT INTO imagens (url, hotel_id) VALUES
('Imagenes/hotel1.jpg', 1),
('Imagenes/hotel2.jpg', 2),
('Imagenes/hotel3.jpg', 3),
('Imagenes/hotel4.jpg', 4),
('Imagenes/hotel5.jpg', 5);
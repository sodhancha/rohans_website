CREATE TABLE books (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    isbn char(14) NOT NULL,
    title varchar(255) NOT NULL,
    author varchar(255) NOT NULL,
    price decimal(5,2) NOT NULL
);

INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);
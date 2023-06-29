DROP TABLE IF EXISTS Categories;
DROP TABLE IF EXISTS Products;

CREATE TABLE IF NOT EXISTS Categories (
    Id INTEGER NOT NULL PRIMARY KEY,
    Name TEXT
);

CREATE TABLE IF NOT EXISTS Products (
  Id INTEGER NOT NULL PRIMARY KEY,
  Name TEXT,
  Category INTEGER,
  Price decimal(8,2),
    CONSTRAINT CatRef FOREIGN KEY(Category) REFERENCES Categories (Id)
);

INSERT INTO Categories (Id, Name) VALUES
  (1, "Mobile"), (2, "Accessories"), (3, "Headphones"), (4, "Books");
INSERT INTO Products (Id, Name, Category, Price) VALUES
    (1, "Apple Watch", 2, 799 ),
    (2, "iPhone", 1, 3999),
    (3, "AirPod", 3, 299),
    (4, "Distributed Systems", 4, 35)
--{ I. MySQL }--

--/ 1. Structure of Tables /--
CREATE TABLE `Product` (
  `Id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `Name` varchar(45) NOT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Id_UNIQUE` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `Sale` (
  `Id` int(10) unsigned NOT NULL,
  `ProductId` int(10) unsigned NOT NULL,
  `ProductQuantity` int(10) unsigned NOT NULL,
  `Time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `id_UNIQUE` (`Id`),
  KEY `fk_Sale_ProductId_idx` (`ProductId`),
  CONSTRAINT `fk_Sale_ProductId` FOREIGN KEY (`ProductId`) REFERENCES `product` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--/ 2. Sample Data /--
INSERT INTO `test`.`Product` (`Name`) VALUES ('Молоко');
INSERT INTO `test`.`Product` (`Name`) VALUES ('Кефир');
INSERT INTO `test`.`Sale` (`ProductId`, `ProductQuantity`) VALUES ('1', '5');
INSERT INTO `test`.`Sale` (`ProductId`, `ProductQuantity`) VALUES ('1', '10');

--/ 3. Queries /--

--/ 3.1. Get Ids of all Products which are not used in the 'Sales' Table /--
--/ Method: "NOT IN" /--
SELECT DISTINCT p.Id
FROM test.Product AS p
WHERE p.Id NOT IN 
(
	SELECT DISTINCT s.ProductId
	FROM test.Sale AS s
);

--/ 3.2. Get Ids of all Products which are not used in the 'Sales' Table /--
--/ Method: "NOT EXISTS" /--
SELECT DISTINCT p.Id
FROM test.Product AS p
WHERE NOT EXISTS
(
	SELECT s.ProductId
	FROM test.Sale AS s
    WHERE s.ProductId = p.Id
);

--{ II. Microsoft SQL Server }--

--/ 1. Structure of Tables /--
CREATE TABLE Product
(
	Id int NOT NULL IDENTITY(1,1) PRIMARY KEY,
	Name varchar(45) NOT NULL,
);

CREATE TABLE Sale
(
  Id int NOT NULL IDENTITY(1,1) PRIMARY KEY,
  ProductId int NOT NULL REFERENCES Product(Id),
  ProductQuantity int NOT NULL,
  Time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--/ 2. Sample Data /--
INSERT INTO Product (Name) 
VALUES ('Молоко');
INSERT INTO Product (Name) 
VALUES ('Кефир');
INSERT INTO Sale (ProductId, ProductQuantity) 
VALUES (1, 5);
INSERT INTO Sale (ProductId, ProductQuantity) 
VALUES (1, 10);

--/ 3.1. Get Ids of all Products which are not used in the 'Sales' Table /--
--/ Method: "NOT IN" /--
SELECT DISTINCT p.Id
FROM Product AS p
WHERE p.Id NOT IN
(
	SELECT DISTINCT s.ProductId
	FROM Sale AS s
);

--/ 3.2. Get Ids of all Products which are not used in the 'Sales' Table /--
--/ Method: "NOT EXISTS" /--
SELECT DISTINCT p.Id
FROM Product AS p
WHERE NOT EXISTS
(
	SELECT s.ProductId
	FROM Sale AS s
    WHERE s.ProductId = p.Id
);

--/ 3.3. Get Ids of all Products which are not used in the 'Sales' Table /--
--/ Method: "EXCEPT" /--
SELECT DISTINCT p.Id
FROM Product AS p
EXCEPT
(
	SELECT DISTINCT s.ProductId
	FROM Sale AS s
);

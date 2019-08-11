package a

const TableNameFull = `public."TestA"`

const QueryCreateTable = `CREATE TABLE ` + TableNameFull + `
(
    "Id"   serial NOT NULL PRIMARY KEY,
    "Name" text   NOT NULL
);`

const QueryFillTable = `INSERT INTO ` + TableNameFull + ` ("Id", "Name")
VALUES 
	(1, 'Aaa'),
	(2, 'Bbb'),
	(3, 'Ccc'),
	(4, 'John'),
	(5, 'Jack'),
	(6, 'Gojoko'),
	(7, 'Fojo'),
	(8, 'Koho');
`

const QuerySelectAll = `SELECT * 
FROM ` + TableNameFull + `;`

const QueryCreateFunction = `CREATE FUNCTION add_name_clear(name text)
RETURNS integer
LANGUAGE sql
AS $$
	----------------------------------------------------------------------------
	--|	This is a demonstrative Function which inserts a Name into the Table |-- 
	--| which is cleared before the Insertion.                               |--
	----------------------------------------------------------------------------
    DELETE FROM ` + TableNameFull + ` WHERE TRUE;
    INSERT INTO ` + TableNameFull + ` ("Name")
    VALUES (name)
    RETURNING "Id";
$$;
`

const QueryFormatDumb = `SELECT "Id", "Name"
FROM ` + TableNameFull + `
WHERE
	("Name" ILIKE '%%%s%%');`

const QueryCreateCreator = `CREATE FUNCTION create_object(src text)
RETURNS bool
LANGUAGE 'plpgsql'
AS $$
BEGIN
	----------------------------------------------------------------------------
	--|	This is a demonstrative Function which uses the dynamic SQL.         |--
	--|	Returns 'true' on a successful Query Execution.                      |--
	----------------------------------------------------------------------------
    EXECUTE src;
	RETURN true;
EXCEPTION WHEN OTHERS THEN
	RETURN false;
END
$$;`

const QueryFinalization = `DROP TABLE ` + TableNameFull + `;
DROP FUNCTION create_object(text);
DROP FUNCTION add_name_clear(text);`

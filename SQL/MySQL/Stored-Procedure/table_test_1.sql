CREATE DEFINER=`test`@`localhost` PROCEDURE `table_test_1`(
    IN columnValueFirst INT UNSIGNED,
    IN columnValueLast INT UNSIGNED
)
BEGIN
	DECLARE LineCounter INT UNSIGNED DEFAULT columnValueFirst;
	
    /* Re-Create the Table */
	DROP TABLE IF EXISTS `tableX`;
	CREATE TABLE `tableX`
    (
		`id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
        `columnA` INT UNSIGNED NOT NULL
    )
    ENGINE = InnoDB;
    
    /* Insert Data */
    START TRANSACTION;
	WHILE LineCounter <= columnValueLast DO
		INSERT INTO `tableX`
        (
			`columnA`
        )
        VALUES
        (
			LineCounter
        );
		SET LineCounter = LineCounter + 1;
	END WHILE;
	COMMIT;
END
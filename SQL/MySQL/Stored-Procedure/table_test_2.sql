CREATE DEFINER=`test`@`localhost` PROCEDURE `table_test_2`(
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
    SET @cmd = CONCAT('INSERT INTO `tableX` (`columnA`) VALUES ');
    WHILE LineCounter <= columnValueLast DO
		IF LineCounter < columnValueLast THEN 
			SET @cmd = CONCAT(@cmd, '(',LineCounter,'),');
		ELSE
			SET @cmd = CONCAT(@cmd, '(',LineCounter,');');
		END IF;
		SET LineCounter = LineCounter + 1;
	END WHILE;
    
    /* SELECT @cmd; */
    START TRANSACTION;
    PREPARE stmt FROM @cmd;
    EXECUTE stmt;
    DEALLOCATE PREPARE stmt;
	COMMIT;
END

/* This version (v2) works ~2x Times slower than v1 */

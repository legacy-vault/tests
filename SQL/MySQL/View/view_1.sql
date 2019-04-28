CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `test`@`localhost` 
    SQL SECURITY DEFINER
VIEW `topics_activity` AS
    SELECT 
        `topics`.`id` AS `id`,
        `topics`.`name` AS `topic`,
        (`topics`.`views` / `topics`.`messages`) AS `view_reply_ratio`
    FROM
        `topics`
    WHERE
        (`topics`.`messages` > 0)
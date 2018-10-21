<?php

// logic/db.php

// Version: 1.0.
// Date:    2017-05-24.
// Author:  McArcher.

function db_connect()
{
    global $db_host, $db_username, $db_pwd, $db_dbname, $db_port, $db_socket, $db;

    $db = new mysqli($db_host, $db_username, $db_pwd, $db_dbname, $db_port);
    
    if ($db->connect_errno)
    {
        echo 'Failed to connect to database: ' . $db->connect_error;
        exit();
    }
}

function db_close()
{
    global $db;
    
    mysqli_close($db);
}

function db_create_tables()
{
    global $db, $br;
    global $table_articles, $col_articles_id, $col_articles_date;
    global $col_articles_themes_ids, $col_articles_title, $col_articles_description;
    global $col_articles_text, $col_articles_date_idx;
    
    $cmd = 
    "CREATE TABLE `$table_articles` (
    `$col_articles_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `$col_articles_date` INT UNSIGNED NOT NULL,
    `$col_articles_themes_ids` TEXT NOT NULL,
    `$col_articles_title` TINYTEXT NOT NULL,
    `$col_articles_description` TEXT NOT NULL,
    `$col_articles_text` MEDIUMTEXT NOT NULL,
    PRIMARY KEY (`$col_articles_id`),
    UNIQUE INDEX `ID_UNIQUE` USING BTREE (`$col_articles_id` ASC),
    INDEX `$col_articles_date_idx` USING BTREE (`$col_articles_date` ASC));";
    
    $res = $db->query($cmd);
    
    if (!$res)
    {
        echo 'Error creating table: ' . $db->error . $br;
        exit();
    }
}

function db_insert_article($Title, $ShortDescription, $Text, $Themes_IDs)
{
    global $db, $arrays_delimiter_in_str;
    global $table_articles, $col_articles_date, $col_articles_themes_ids;
    global $col_articles_title, $col_articles_description, $col_articles_text;
    
    $timestamp = time();
    $Title = $db->real_escape_string($Title);
    $ShortDescription = $db->real_escape_string($ShortDescription);
    $Text = $db->real_escape_string($Text);
    $Themes_IDs_str = implode($arrays_delimiter_in_str, $Themes_IDs);
    $Themes_IDs_str = $db->real_escape_string($Themes_IDs_str);
    
    $cmd =
    "INSERT INTO `$table_articles`
    ($col_articles_date, $col_articles_themes_ids, $col_articles_title, 
    $col_articles_description, $col_articles_text)
    VALUES ($timestamp, '$Themes_IDs_str', '$Title', '$ShortDescription', '$Text');";
    
    $row = $db->query($cmd);
    
    if ($row)
    {
        $row_id = $db->insert_id;
        return $row_id;
    }
    else
    {
        exit('Error: (' . $db->errno . ') ' . $db->error);
        exit();
    }
}

function db_get_article_desc($id)
{
    global $db, $table_articles, $col_articles_id, $col_articles_date, $col_articles_themes_ids;
    global $col_articles_title, $col_articles_description;
    
    $cmd =
    "SELECT `$col_articles_id`, `$col_articles_date`, `$col_articles_themes_ids`, 
    `$col_articles_title`, `$col_articles_description` 
    FROM `$table_articles` 
    WHERE `$col_articles_id` = $id;";
    
    $res = $db->query($cmd);
    
    if ($res->num_rows != 1)
    {
        return NULL;
    }
    
    if ( $row = $res->fetch_assoc() )
    {
        return $row;
    }
    else
    {
        return NULL;
    }
}

function db_get_article_text($id)
{
    global $db, $table_articles, $col_articles_id, $col_articles_date;
    global $col_articles_themes_ids, $col_articles_title, $col_articles_text;
    
    $cmd = "SELECT `$col_articles_id`, `$col_articles_date`, `$col_articles_themes_ids`, 
    `$col_articles_title`, `$col_articles_text` 
    FROM `$table_articles` WHERE `$col_articles_id` = $id;";
    
    $res = $db->query($cmd);
    
    if ($res->num_rows != 1)
    {
        return NULL;
    }
    
    if ( $row = $res->fetch_assoc() )
    {
        return $row;
    }
    else
    {
        return NULL;
    }
}

?>

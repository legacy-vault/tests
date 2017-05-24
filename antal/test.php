<?php

// test.php

// logic/theme.php

// Version: 1.0.
// Date:    2017-05-24.
// Author:  McArcher.

include_once 'conf/config.php';
include_once $path_logic_article;
include_once $path_logic_theme;
include_once $path_logic_db;
include_once $path_logic_visitor;
include_once $path_logic_init;

// Header
include_once $path_html_header;

/* BEGINNING OF MAIN WORKING AREA */


//db_create_tables();   // Create working Table in DataBase
//apcu_clear_cache();   // Clear Cache (APCu)
/*
add_theme("First Theme");
add_theme("Theme #2");
add_theme("Theme #3");
add_article("Article #1", "No Description.", "This is a Test!", [2]);
add_article("Article #2", "No Description.", "This is a Test!", [1,3]);
add_article("Article #3", "No Description.", "This is a Test!", [1,2,3]);
add_article("Article #4", "No Description.", "This is a Test!", [2,8]);
*/

echo "Example of Usage of the Module <b>'Themes List'</b>." . $br;
// Use Module: Themes List
mod_themes_list();
echo $br;

echo "Example of Usage of the Module <b>'All Articles in all Themes List'</b>." . $br;
// Use Module: All Articles in all Themes List
mod_articles_list_all();
echo $br;

echo "Example of Usage of the Module <b>'All Articles of one Theme List'</b>." . $br;
// Use Module: All Articles of one Theme List
$id = 2; // simulate GET variable from server to use it in next module
mod_articles_of_theme($id);
echo $br;

echo "Example of Usage of the Module <b>'One Article'</b>." . $br;
$id = 30; // simulate GET variable from server to use it in next module
// Use Module: One Article
mod_article_show($id);
echo $br;

echo "Example of Usage of the Module <b>'Visitors Count'</b>." . $br;
// Use Module: Visitors Count
mod_visitors_txt();

//apcu_clear_cache();   // Clear Cache (APCu)


/* END OF MAIN WORKING AREA */

// Footer
include_once $path_html_footer;

include_once $path_logic_final;

?>

<?php

// conf/config.php

// Version: 1.0.
// Date:    2017-05-24.
// Author:  McArcher.

// Useful Constants
$nl = "\r\n";
$br = "<br>$nl";
$nbsp = '&nbsp;';
$arrays_delimiter_in_str = ' ';
$date_format = 'j F Y';

// Variables
$db = NULL;

// APCu Variables
$Articles = array();  // Keys (Indices) are IDs
$Articles_Count = 0;
$Themes = array();  // Keys (Indices) are IDs
$Themes_Count = 0;
$Visitors_All = 0;
$Visitors_Unique = 0;
$Visitors_LAT = 0;

// Locations & Paths
$path_separator = '/';
$ext_bak = '.bak';

$dir_css = 'css';
$path_css_main = $dir_css . $path_separator . 'main.css';

$dir_data = 'data';
$path_data_themes = $dir_data . $path_separator . 'themes.dat';
$path_data_themes_bak = $path_data_themes . $ext_bak;
$path_data_visitors = $dir_data . $path_separator . 'visitors.dat';
$path_data_visitors_bak = $path_data_visitors . $ext_bak;

$dir_html = 'html';
$path_html_footer = $dir_html . $path_separator . 'footer.php';
$path_html_header = $dir_html . $path_separator . 'header.php';

$dir_logic = 'logic';
$path_logic_article = $dir_logic . $path_separator . 'article.php';
$path_logic_db = $dir_logic . $path_separator . 'db.php';
$path_logic_final = $dir_logic . $path_separator . 'final.php';
$path_logic_init = $dir_logic . $path_separator . 'init.php';
$path_logic_theme = $dir_logic . $path_separator . 'theme.php';
$path_logic_visitor = $dir_logic . $path_separator . 'visitor.php';

// Names of Variables (for APCu)
$Articles_vn = 'Articles';
$Articles_Count_vn = 'Articles_Count';
$Themes_vn = 'Themes';
$Themes_Count_vn = 'Themes_Count';
$Visitors_All_vn = 'Visitors_All';
$Visitors_Unique_vn = 'Visitors_Unique';
$Visitors_LAT_vn = 'Visitors_LAT'; // Last Access Time to File

// DataBase Parameters
$db_host = 'localhost';
$db_username = 'antal';
$db_pwd = 'antal';
$db_dbname = 'antal';
$db_port = '3306';
$db_socket = '';

// DataBase Tables
$table_articles = 'Articles';

// DataBase Columns
$idx_prefix = 'idx_';
$col_articles_id = 'ID';
$col_articles_date = 'Date';
$col_articles_date_idx = $idx_prefix . $col_articles_date;
$col_articles_title = 'Title';
$col_articles_description = 'ShortDescription';
$col_articles_text = 'Text';
$col_articles_themes_ids = 'Themes_IDs';

// HTML CSS Styles
$css_table_container = 'container';
$css_td_h1 = 'h1';
$css_td_h1l = 'h1l';
$css_td_h1s = 'h1s';
$css_td_h1r = 'h1r';
$css_td_h2l = 'h2l';
$css_td_h2s = 'h2s';
$css_td_h2r = 'h2r';
$css_td_h3 = 'h3';
$css_span_tag = 'tag';

// Cookies
$cookie_name = 'clid';
$cookie_expire_delta =  60*60*24*365; // Seconds
$cookie_path = '/';
$cookie_domain = '';
$cookie_secure = false;
$cookie_http = true;
$cookie_file_update_interval = 60; // Seconds

?>

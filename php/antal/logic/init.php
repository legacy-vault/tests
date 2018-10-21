<?php

// logic/init.php

// Version: 1.0.
// Date:    2017-05-24.
// Author:  McArcher.

// Themes
init_themes();

// Visitors
init_visitors();
calc_visitors();
update_visitors_cache();
refresh_visitors_file();

// DataBase
db_connect();

?>

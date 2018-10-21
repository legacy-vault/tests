<?php

// logic/visitor.php

// Version: 1.0.
// Date:    2017-05-24.
// Author:  McArcher.

/*
 * Checks if $Visitors_All & $Visitors_Unique exist in Cache.
 * If exists, reads them from Cache.
 * If no, reads them from File and stores into Cache.
 */
function init_visitors()
{
    global $Visitors_All, $Visitors_All_vn, $Visitors_Unique, $Visitors_Unique_vn;
    global $Visitors_LAT, $Visitors_LAT_vn;
    
    if (    apcu_exists($Visitors_All_vn) && apcu_exists($Visitors_Unique_vn) && 
            apcu_exists($Visitors_LAT_vn) )
    {
        $Visitors_All = apcu_fetch($Visitors_All_vn);
        $Visitors_Unique = apcu_fetch($Visitors_Unique_vn); 
        $Visitors_LAT = apcu_fetch($Visitors_LAT_vn); 
    }
    else
    {
        read_visitors(); 
        update_visitors_cache();
        update_visitors_lat();
    }
}

/*
 * Reads $Visitors_All, $Visitors_Unique, $Visitors_LAT from File to Object.
 */
function read_visitors()
{
    global $Visitors_All, $Visitors_All_vn, $Visitors_Unique, $Visitors_Unique_vn, $nl;
    global $Visitors_LAT_vn, $Visitors_LAT, $path_data_visitors;
    
    $file = fopen($path_data_visitors, 'r');
    
    // Read $Visitors_All
    $line = fgets($file);
    $line = trim($line);
    if ( is_numeric($line) )
    {
        $Visitors_All = intval($line);
    }
    
    // Read $Visitors_Unique
    $line = fgets($file);
    $line = trim($line);
    if ( is_numeric($line) )
    {
        $Visitors_Unique = intval($line);
    }
    
    // Read $Visitors_LAT
    $line = fgets($file);
    $line = trim($line);
    if ( is_numeric($line) )
    {
        $Visitors_LAT = intval($line);
    }
    
    fclose($file);
}

/*
 * Updates Last Access Time to Vistors File.
 * This is done to calculate Interval between last File update and 
 * this Moment and decide whether to update File or not.
 * The Function is re-reading the File from beginning, instead of reading from Memory.
 * Such way is a bit odd, but it more safe.
 * While such updates are done very rarely, it is not hitting the Performance.
 */
function update_visitors_lat()
{
    global $path_data_visitors, $path_data_visitors_bak, $nl;
    global $Visitors_LAT_vn, $Visitors_LAT;
    
    $t1 = '';
    $t2 = '';
    
    // Read old Contents of File
    $file = fopen($path_data_visitors, 'r');
    // Read $Visitors_All
    $t1 = fgets($file);
    // Read $Visitors_Unique
    $t2 = fgets($file);
    fclose($file);
    
    // Calculate LAT
    $now = time();
    $now_str = $now . $nl;
    
    // Make Back-up Copy
    copy($path_data_visitors, $path_data_visitors_bak);
    
    // Update LAT in File
    $file = fopen($path_data_visitors, 'w');
    fwrite($file, $t1); // $t1 already has $nl
    fwrite($file, $t2); // $t2 already has $nl
    fwrite($file, $now_str);
    fclose($file);
    
    // Delete Back-up Copy
    unlink($path_data_visitors_bak);
    
    // Store LAT in Cache & Object
    apcu_store($Visitors_LAT_vn, $now);
    $Visitors_LAT = $now;
}

/*
 * Calculates $Visitors_All & $Visitors_Unique.
 */
function calc_visitors()
{
    global $cookie_expire_delta, $cookie_name, $cookie_path, $cookie_domain, $cookie_secure;
    global $cookie_http, $Visitors_All, $Visitors_Unique;
    
    if( !isset($_COOKIE[$cookie_name]) )
    {
        // New Client
        $rn = random_bytes(8);
        $cookie_val = base64_encode($rn);
        $cookie_expire = time() + $cookie_expire_delta;
        
        setcookie(  $cookie_name, $cookie_val,  $cookie_expire, $cookie_path, $cookie_domain, 
                    $cookie_secure, $cookie_http);
        
        $Visitors_All++;
        $Visitors_Unique++;
    }
    else
    {
        // Existing Client
        //$cookie_val = $_COOKIE[$cookie_name];
        
        $Visitors_All++;
    }
}

/*
 * Stores $Visitors_All & $Visitors_Unique into Cache.
 */
function update_visitors_cache()
{
    global $Visitors_All, $Visitors_All_vn, $Visitors_Unique, $Visitors_Unique_vn, $nl;
    global $Visitors_LAT_vn, $Visitors_LAT;
    
    apcu_store($Visitors_All_vn, $Visitors_All);
    apcu_store($Visitors_Unique_vn, $Visitors_Unique);
    apcu_store($Visitors_LAT_vn, $Visitors_LAT);
}

/*
 * Periodically refreshes visitors' File.
 * This function is always gets run.
 * This function itself decides whether to update L.A.T. to the File or not.
 */
function refresh_visitors_file()
{
    global $Visitors_All, $Visitors_Unique, $Visitors_LAT, $Visitors_LAT_vn, $nl;
    global $path_data_visitors, $path_data_visitors_bak;
    global $cookie_file_update_interval;
    
    // Check if LAT was too long ago
    $now = time();
    if ( ($now - $Visitors_LAT) < $cookie_file_update_interval)
    {
        return;
    }
    
    $lat_str = $now . $nl;
    
    // Create Reserved Copy before re-writing the File!
    copy($path_data_visitors, $path_data_visitors_bak);

    $file = fopen($path_data_visitors, 'w');
    
    // Write $Visitors_All
    fwrite($file, $Visitors_All . $nl);
    
    // Write $Visitors_Unique
    fwrite($file, $Visitors_Unique . $nl);
    
    // Write L.A.T.
    fwrite($file, $lat_str);
    
    fclose($file);
    
    // Delete Reserved Copy
    unlink($path_data_visitors_bak);
    
    // Syncronize LAT with Cache and Object.
    $Visitors_LAT = $now;
    apcu_store($Visitors_LAT_vn, $Visitors_LAT);
}

function mod_visitors_txt()
{
    global $Visitors_All, $Visitors_Unique, $br, $nl;
    global $css_table_container, $css_td_h2l, $css_td_h2r;
    
    echo "<table class='$css_table_container'>$nl";
    
    echo    "<tr><td class='$css_td_h2l'>Total Visitors</td>" . 
            "<td class='$css_td_h2r'>" . $Visitors_All . '</td></tr>' . $nl;
    echo    "<tr><td class='$css_td_h2l'>Unique Visitors</td>" . 
            "<td class='$css_td_h2r'>" . $Visitors_Unique . '</td></tr>' . $nl;
    
    echo "</table>$nl";
}

?>

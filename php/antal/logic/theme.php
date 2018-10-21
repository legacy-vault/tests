<?php

// logic/theme.php

// Version: 1.0.
// Date:    2017-05-24.
// Author:  McArcher.

// Class
class Theme
{
    public $Articles_Count;
    public $Articles_IDs;
    public $ID;     // Starts from 1
    public $Name;    
    
    public function __construct()
    {
        $this->Articles_Count = 0;
        $this->Articles_IDs = array();
        $this->ID = 0;
        $this->Name = '';        
    }
}

// Functions

/*
 * Adds a Theme to Themes Array.
 */
function create_theme($Name, $ID = 0)
{
    global $Themes, $Themes_Count, $br;
    
    $t = new Theme();
    $Themes_Count++;
    
    if ($ID == 0)
    {
        $t->ID = $Themes_Count;
    }
    else
    {
        $t->ID = $ID;
    }
    
    if ( array_key_exists($t->ID, $Themes) )
    {
        echo 'Error: index (' . $t->ID . ") already exists in Themes array.$br";
        $Themes_Count--;
        exit();
    }
    
    $t->Name = $Name;
    
    $Themes[ $t->ID ] = $t;
}

/*
 * Checks if $Themes & $Themes_Count exist in Cache.
 * If exists, reads them from Cache.
 * If no, reads them from File and stores into Cache.
 */
function init_themes()
{
    global $Themes_vn, $Themes_Count_vn, $Themes, $Themes_Count;
    
    if ( apcu_exists($Themes_vn) && apcu_exists($Themes_Count_vn) )
    {
        $Themes = apcu_fetch($Themes_vn);
        $Themes_Count = apcu_fetch($Themes_Count_vn); 
    }
    else
    {
        read_themes();        
        apcu_store($Themes_vn, $Themes);
        apcu_store($Themes_Count_vn, $Themes_Count);
    }
}

/*
 * Reads $Themes & $Themes_Count from File.
 */
function read_themes()
{
    global $Themes, $Themes_Count, $Themes_Count_vn, $path_data_themes, $nl;
    global $arrays_delimiter_in_str;
    
    function themes_line_check($str)
    {
        if ($str === false)
        {
            echo "Themes File is broken.$nl";
            exit();
        }
    }
    
    $file = fopen($path_data_themes, 'r');
    
    // Read Themes_Count
    $line = fgets($file);
    themes_line_check($line);
    $line = trim($line);
    $count = (int)$line;
    
    // Read Themes
    for ($i = 0; $i < $count; $i++)
    {
        // Theme's ID
        $id = fgets($file);
        themes_line_check($id);
        $id = trim($id);
        
        // Theme's Name
        $name = fgets($file);
        themes_line_check($name);
        $name = trim($name);
        
        // Theme's Articles_Count
        $ac = fgets($file);
        themes_line_check($ac);
        $ac = trim($ac);
        
        // Theme's Articles_IDs (Array)
        $aids_str = fgets($file);
        themes_line_check($aids_str);
        $aids_str = trim($aids_str);
        $aids = explode($arrays_delimiter_in_str, $aids_str);
        
        // Modify Objects: Themes, Themes_Count.
        create_theme($name, $id);
        $Themes[ $id ]->Articles_Count = $ac;
        $Themes[ $id ]->Articles_IDs = $aids;
    }
    
    fclose($file);
}

/*
 * Adds a new $Theme Object into $Themes Array,
 * Writes new $Themes Array to File,
 * Updates Cache.
 */
function add_theme($Name, $ID = 0)
{
    create_theme($Name, $ID);
    refresh_themes_file();
    update_themes_cache();
}

/*
 * Writes Contents of $Themes Array into File.
 */
function refresh_themes_file()
{
    global $Themes, $Themes_Count, $nl, $path_data_themes, $path_data_themes_bak;
    
    // Create Reserved Copy before re-writing the File!
    copy($path_data_themes, $path_data_themes_bak);

    $file = fopen($path_data_themes, 'w');
    
    // Write Themes_Count
    fwrite($file, $Themes_Count . $nl);
    
    // Write Themes
    $j = 0;
    for ($i = 0; $i < $Themes_Count; $i++)
    {
        $j = $i + 1;
        
        // Theme's ID
        fwrite($file, $Themes[$j]->ID . $nl);
        
        // Theme's Name
        fwrite($file, $Themes[$j]->Name . $nl);
        
        // Theme's Articles_Count
        $ac = $Themes[$j]->Articles_Count;
        fwrite($file, $ac . $nl);
        
        // Theme's Articles_IDs (Array)
        foreach ($Themes[$j]->Articles_IDs as $key => $val)
        {
            fwrite($file, $val . ' ');
        }
        fwrite($file, $nl);
    }
    
    fclose($file);
    
    // Delete Reserved Copy
    unlink($path_data_themes_bak);
}

/*
 * Stores $Themes & $Themes_Count into Cache.
 */
function update_themes_cache()
{
    global $Themes_vn, $Themes_Count_vn, $Themes, $Themes_Count, $nl;
    
    apcu_store($Themes_vn, $Themes);
    apcu_store($Themes_Count_vn, $Themes_Count);
}

/*
 * Lists Themes.
 */
function mod_themes_list()
{
    global $Themes, $Themes_Count, $nl;
    global $css_table_container, $css_td_h1;
    
    echo "<table class='$css_table_container'>$nl";
    foreach ($Themes as $key => $val)
    {
        echo "<tr><td class='$css_td_h1'>" . $val->Name . "</td></tr>$nl";
    }
    echo "</table>$nl";
}

?>

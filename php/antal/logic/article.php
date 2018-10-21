<?php

// logic/article.php

// Version: 1.0.
// Date:    2017-05-24.
// Author:  McArcher.

// Class
class Article
{
    public $Date;
    public $ID;     // Starts from 1
    public $ShortDescription;
    public $Text;
    public $Themes_IDs;
    public $Title;
    
    public function __construct()
    {
        $this->Date = time();
        $this->ID = 0;
        $this->ShortDescription = '';
        $this->Text = '';
        $this->$Themes_IDs = array();
        $this->Title = '';
    }
}

// Functions

/*
 * Creates and adds an Article to Articles Array.
 * Used for loading Objects from existing File.
 * Caching is done separately (not here).
 */
function create_article($Title, $ShortDescription, $Text, $ID = 0)
{
    global $Articles, $Articles_Count;
    
    $a = new Article();
    $Articles_Count++;
    
    if ($ID == 0)
    {
        $a->ID = $Articles_Count;
    }
    else
    {
        $a->ID = $ID;
    }
    
    if ( !array_key_exists($a->ID, $Articles) )
    {
        echo 'Error: index (' . $a->ID . ") already exists in Articles array.$br";
        $Articles_Count--;
        exit();
    }
    
    $a->ShortDescription = $ShortDescription;
    $a->Title = $Title;
    $a->Text = $Text;
    
    $Articles[ $a->ID ] = $a;
}

/*
 * Binds an Article with Theme.
 */
function add_article_to_theme($ThemeID, $ArticleID)
{
    global $Themes, $br;
    
    if ( !array_key_exists($ThemeID, $Themes) )
    {
        echo "Error: non-existing index ($ThemeID) in Themes array.$br";
        exit();
    }
    
    if ( !in_array($ArticleID, $Themes[ $ThemeID ]->Articles_IDs) )
    {
        $Themes[ $ThemeID ]->Articles_IDs[] = $ArticleID;
        $Themes[ $ThemeID ]->Articles_Count++;
    }
}

/*
 * Puts an Article to DataBase, refreshes the File, updates Cache.
 * Used for creating Objects on-fly.
 */
function add_article($Title, $ShortDescription, $Text, $Themes_IDs)
{
    global $Themes, $nl;
    global $br;
    
    // Themes_IDs is set?
    $size = count($Themes_IDs);
    if ($size == 0)
    {
        echo "Error: empty Themes_IDs array.$br";
        exit();
    }
    
    // Check if parent Themes exist
    foreach ($Themes_IDs as $key => $ThemeID)
    {
        if ( !array_key_exists($ThemeID, $Themes) )
        {
            echo "Error: non-existing index ($ThemeID) in Themes array.$br";
            exit();
        }
    }
    
    // Insert Article into DB
    $ArticleID = db_insert_article($Title, $ShortDescription, $Text, $Themes_IDs);
    
    // Connect Article with all the Themes mentioned in $Themes_IDs Array
    foreach ($Themes_IDs as $key => $val)
    {
        add_article_to_theme($val, $ArticleID);
    }
    
    // Refresh Themes File
    refresh_themes_file();
    
    // Update Themes Cache
    update_themes_cache();
}

/*
 * Lists all Articles in all Themes.
 */
function mod_articles_list_all()
{
    global $Themes, $nl, $date_format, $nbsp;
    global $css_table_container, $css_td_h1l, $css_td_h1s, $css_td_h1r;
    global $css_td_h2l, $css_td_h2s, $css_td_h2r, $css_td_h3;
    
    echo "<table class='$css_table_container'>$nl";
    
    foreach ($Themes as $key => $val)
    {
        // Theme's Name & Articles_Count
        echo    "<tr><td class='$css_td_h1l'>" . $val->Name . 
                "</td><td class='$css_td_h1s'> </td>" . 
                "<td class='$css_td_h1r'>[" . $val->Articles_Count . 
                "]</td></tr>$nl";
        
        // Articles of a Theme
        foreach ($val->Articles_IDs as $art_id)
        {
            $art = db_get_article_desc($art_id);
            
            if ( $art == NULL )
            {
                echo "<tr><td colSpan='3'>";
                echo "Error. Article with ID = $art_id is not found.";
                echo "</td></tr>$nl";
                continue;
            }
            
            $date_str = date($date_format, $art["Date"]);
            $date_str = str_replace(' ', $nbsp, $date_str);
            
            echo    "<tr><td class='$css_td_h2l'>" . $art["Title"] . 
                    "</td><td class='$css_td_h2s'> </td>" . 
                    "<td class='$css_td_h2r'>" . $date_str . "</td></tr>$nl" . 
                    "<tr><td colSpan='3' class='$css_td_h3'>" . 
                    $art["ShortDescription"] . "</td></tr>$nl";
        }        
    }
    
    echo "</table>$nl";
}

/*
 * Lists all Articles of the selected Theme.
 */
function mod_articles_of_theme($theme_id)
{
    global $Themes, $nl, $date_format, $nbsp;
    global $css_table_container, $css_td_h1l, $css_td_h1s, $css_td_h1r;
    global $css_td_h2l, $css_td_h2s, $css_td_h2r, $css_td_h3;
    
    // chk////////////
    
    $thm = $Themes[ $theme_id ];
    
    echo "<table class='$css_table_container'>$nl";
    
    // Theme's Name & Articles_Count
    echo    "<tr><td class='$css_td_h1l'>" . $thm->Name . 
            "</td><td class='$css_td_h1s'> </td>" . 
            "<td class='$css_td_h1r'>[" . $thm->Articles_Count . 
            "]</td></tr>$nl";
    
    // Articles of a Theme
    foreach ($thm->Articles_IDs as $art_id)
    {
        $art = db_get_article_desc($art_id);
        
        if ( $art == NULL )
        {
            echo "<tr><td colSpan='3'>";
            echo "Error. Article with ID = $art_id is not found.";
            echo "</td></tr>$nl";
            continue;
        }
        
        $date_str = date($date_format, $art["Date"]);
        $date_str = str_replace(' ', $nbsp, $date_str);
        
        echo    "<tr><td class='$css_td_h2l'>" . $art["Title"] . 
                "</td><td class='$css_td_h2s'> </td>" . 
                "<td class='$css_td_h2r'>" . $date_str . "</td></tr>$nl" . 
                "<tr><td colSpan='3' class='$css_td_h3'>" . 
                $art["ShortDescription"] . "</td></tr>$nl";
    }        
    
    
    echo "</table>$nl";
}

/*
 * Shows an Article.
 */
function mod_article_show($article_id)
{
    global $nl, $date_format, $nbsp, $arrays_delimiter_in_str, $Themes;
    global $css_span_tag, $css_table_container, $css_td_h1;
    global $css_td_h2l, $css_td_h2s, $css_td_h2r, $css_td_h3;
    
    
    echo "<table class='$css_table_container'>$nl";
    
    $art = db_get_article_text($article_id);
    
    if ( $art == NULL )
    {
        echo "<tr><td colSpan='3'>";
        echo "Error. Article with ID = $art_id is not found.";
        echo "</td></tr>$nl";
        return;
    }
    
    $date_str = date($date_format, $art["Date"]);
    $date_str = str_replace(' ', $nbsp, $date_str);
    
    $theme_list = '';
    $Themes_IDs = explode($arrays_delimiter_in_str, $art["Themes_IDs"]);
    foreach ( $Themes_IDs as $val)
    {
        $theme_list = $theme_list . "<span class='$css_span_tag'>" . $Themes[$val]->Name . '</span> ';
    }
    $theme_list = rtrim($theme_list);

    echo    "<tr><td colSpan='3' class='$css_td_h1'>" . $art["Title"] . "</td></tr>$nl" . 
            "<tr><td class='$css_td_h2l'>" . $theme_list . 
            "</td><td class='$css_td_h2s'> </td>" . 
            "<td class='$css_td_h2r'>" . $date_str . "</td></tr>$nl" . 
            "<tr><td colSpan='3' class='$css_td_h3'><br>" . 
            $art["Text"] . "<br><br></td></tr>$nl";
    
    echo "</table>$nl";
}

?>

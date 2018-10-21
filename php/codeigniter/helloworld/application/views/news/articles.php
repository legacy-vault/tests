<?php

$output = '<h1>' . $title . '</h1>';

foreach ($news_items as $news_item)
{
	$id = $news_item['id'];
	$alias = $news_item['alias'];
	
	$output .=	'<h2>' . 
					'<a href="' . $alias . '" target="_blank">' . 
						$news_item['title'] . 
					'</a>' . 
				'</h2>';
}

echo $output;

?>

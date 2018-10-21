<?php

$output = '<h1>' . $title . '</h1>';

foreach ($news_items as $news_item)
{
	$id = $news_item['id'];
	$alias = $news_item['alias'];
	$href = 'edit/' . $alias;
	
	$output .=	'<h2>' . 
					'<a href="' . $href . '" target="_blank">' . 
						'News Article: ' . $news_item['title'] . 
					'</a>' . 
				'</h2>';
}

echo $output;

?>

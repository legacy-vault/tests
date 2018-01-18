<?php

echo validation_errors();

$article_alias = $news_item['alias'];
$article_content = $news_item['content'];
$article_title = $news_item['title'];
$article_id = $news_item['id'];

$output = '<h1>' . $title . '</h1>';

$output .= form_open($form_submit_url);

$output .=<<<HEREDOC
	<label for="title">Title</label>
	<input type="input" name="title" value="$article_title"/><br />
	<label for="content">Content</label>
	<textarea name="content">$article_content</textarea><br />
	<label for="alias">Alias</label>
	<input type="input" name="alias" value="$article_alias"/><br />
	<label for="id_hint">ID</label>
	<input type="input" disabled="disabled" name="id_hint" value="$article_id"/><br />
	<input type="hidden" name="id" value="$article_id"/><br />
	<input type="submit" name="submit" value="Edit News Item" />
</form>
HEREDOC;

echo $output;

?>

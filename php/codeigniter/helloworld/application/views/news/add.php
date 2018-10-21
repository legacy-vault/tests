<?php

echo validation_errors();

$output = '<h1>' . $title . '</h1>';

$output .= form_open($form_submit_url);

$output .=<<<HEREDOC
	<label for="title">Title</label>
	<input type="input" name="title" /><br />
	<label for="content">Content</label>
	<textarea name="content"></textarea><br />
	<input type="submit" name="submit" value="Submit News Item" />
</form>
HEREDOC;

echo $output;

?>

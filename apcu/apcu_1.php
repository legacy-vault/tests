<?php

echo '<pre>';
print_r(apcu_cache_info());
echo '</pre>';

$test = NULL;
$test_default = 'Hello, World!';

if ( apcu_exists('test') )
{
    $test = apcu_fetch('test');
}
else
{
    apcu_store('test', $test_default);
    $test = apcu_fetch('test');
}

print_r( $test );

//---------------------------------------------

$run_count = NULL;
$run_count_default = 0;
$run_count_name = 'run_count';

if ( apcu_exists($run_count_name) )
{
    apcu_inc($run_count_name);
    $run_count = apcu_fetch($run_count_name);
}
else
{
    apcu_store($run_count_name, $run_count_default);
    $run_count = apcu_fetch($run_count_name);
}

print_r( $run_count );

?>

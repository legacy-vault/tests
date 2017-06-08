<?php
//
// array_sort_compar_1.php
// Author: McArcher.
// Date: 2017-06-08.
// Version: 0.1.
//
//------------------------------------------------------------------------------
//
// Comparison between two ways of sorting complicated arrays.
//
// Way 1: sorting by a 'array_multisort' function with a helping array.
// Way 2: sorting by a 'usort' function with a comparison function.
//
// Results: 
// on my machine it shows that the first way is faster (11.95 vs 16.64 seconds).
//
//------------------------------------------------------------------------------

Class Person
{
    public $FirstName;
    public $LastName;
    public $Age;

    public function __construct ($fn, $ln, $ag)
    {
        $this->FirstName = $fn;
        $this->LastName = $ln;
        $this->Age = $ag;
    }
}

const BRNL = "<br>\r\n";

//------------------------------------------------------------------------------

main();

//------------------------------------------------------------------------------

function main()
{
    $PersonsArray = array();
    $test_iterations = 1000 * 1000; // 1 M
    $test_i = 0;


    // Test # 1
    echo 'Test # 1' . BRNL;
    $time_1 = microtime(TRUE);

    for ($test_i = 1; $test_i <= $test_iterations; $test_i++)
    {
        data_init($PersonsArray);
        //var_dump($PersonsArray); //
        test_1($PersonsArray);
    }

    $time_2 = microtime(TRUE);
    $time_dur = $time_2 - $time_1;
    echo 'Duration: ' . $time_dur . ' seconds.' . BRNL;
    echo 'Result: ' . BRNL; //
    var_dump($PersonsArray); //

    // Test # 2
    echo 'Test # 2' . BRNL;
    $time_1 = microtime(TRUE);

    for ($test_i = 1; $test_i <= $test_iterations; $test_i++)
    {
        data_init($PersonsArray);
        //var_dump($PersonsArray); //
        test_2($PersonsArray);
    }

    $time_2 = microtime(TRUE);
    $time_dur = $time_2 - $time_1;
    echo 'Duration: ' . $time_dur . ' seconds.' . BRNL;
    echo 'Result: ' . BRNL; //
    var_dump($PersonsArray); //
}

//------------------------------------------------------------------------------

function data_init(&$PersonsArray)
{
    $p1 = new Person('Василий', 'Пупкин', 10);
    $p2 = new Person('John', 'Smith', 30);
    $p3 = new Person('Fernando', 'Alonso', 20);
    $p4 = new Person('Никола', 'Тесла', 25);
    $p5 = new Person('Frodo', 'Baggins', 35);
    $p6 = new Person('Иван', 'Петров', 28);

    $PersonsArray = [$p1, $p2, $p3, $p4, $p5, $p6];
}

//------------------------------------------------------------------------------

function test_1(&$PersonsArray)
{
    $tmpArray = array();

    foreach ($PersonsArray as $v)
    {
        $tmpArray[] = $v->Age;
    }

    array_multisort($tmpArray, SORT_ASC, SORT_NUMERIC, $PersonsArray);
}

//------------------------------------------------------------------------------

function test_2(&$PersonsArray)
{
    usort($PersonsArray, 'comparator');
}

//------------------------------------------------------------------------------

function comparator($obj_a, $obj_b)
{
    return ( $obj_a->Age <=> $obj_b->Age );
}

//------------------------------------------------------------------------------

?>

<?php
    
    // soap_1.php

    
echo "<html>\r\n<body>";


// Create a Client
$soap_URI = "http://www.webservicex.net/ConvertSpeed.asmx?WSDL";
$soap_client = new SoapClient($soap_URI);


// List available Functions
$soap_response = $soap_client->__getFunctions();
var_dump($soap_response);
echo "<br><br>\r\n";


// Run 'ConvertSpeed' Function
$soap_func = 'ConvertSpeed';
$soap_params = 
    [
        "speed" => "100",
        "FromUnit" => "milesPerhour",
        "ToUnit" => "kilometersPerhour"
    ];
$soap_response = $soap_client->$soap_func($soap_params);
var_dump($soap_response);
echo "<br><br>\r\n";


// Get Result from Object
$result = $soap_response->ConvertSpeedResult;
print_r($result);
echo "<br><br>\r\n";


echo "</body>\r\n</html>";

?>

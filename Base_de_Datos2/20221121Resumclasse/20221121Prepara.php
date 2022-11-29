<?php
// 20221121Prepara.php?usuari=Jaume&id=1&nounom=Antoni&nouemail=antoni%40gmail.com&modifica=Enviar
// 20221121Prepara.php?usuari=Jaume&id=1&nounom=&nouemail=antoni%40gmail.com&modifica=Enviar
$lusuari = $_GET['usuari'];
$nom= $_GET['nounom'];
$mail = $_GET['nouemail'];
$id = $_GET['id'];

// UPDATE persona SET nom    = 'xxxxxx',
//                    email  = 'xxxxxxx',
//                    darrer = 'xxxxx'
// WHERE idpers = xxxxxxx

$cadena = "UPDATE persona SET nom = '".$nom."',";
$cadena = $cadena."correu = '".$mail."',";
$cadena = $cadena."darrer = '".$lusuari."'";
$cadena = $cadena." WHERE idpers=".$id;
// falta controlar si nom o email estan en blanc

echo $cadena;

$con = mysqli_connect("localhost","root","") or die("Error Mysql");
$db = mysqli_select_db($con,"20221121persones") or die("Error a DB");
mysqli_query($con,$cadena);
echo "<br>Update realitzat amb ¿exit?";
?>



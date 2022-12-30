/*
Els professors de l’assignatura de Bases de Dades II disposen d’una base de dades per calcular les notes
finals de l’assignatura. Aquesta BD té una taula amb els atributs que contenen les dades personals de
l’alumne, les seves notes d’avaluació i una nota final, apart d’un atribut de comentaris:
Alumne(núm_matrícula, dni, nom, llinatges, nota_pràctica_1, nota_pràctica_2, nota_examen, nota_final)

Tenint en compte que la nota de l’assignatura es calcula de la manera següent:
    1. La nota final es calcula amb la mitjana ponderada de les notes de les pràctiques i la nota de l’examen:
    un 15% la PR1, un 35% la PR2 i un 50% la nota de l’examen.
    2. La nota final serà el resultat del càlcul anterior. Si el resultat és inferior a 5, la nota serà la nota
    calculada. En cas de superar o igualar el 5, la nota final serà la nota calculada, excepte en el cas en
    què alguna de les proves d’avaluació no superi el mínim de 5: en aquest cas es posarà un 4,5 de nota
    final.
2) Es demana que generin un procediment de bases de dades que accedeixi a la taula Alumne
i calculi les notes finals de l’assignatura, modificant la nota de la taula.
*/

DELIMITER //
CREATE procedure calculoFinal()
BEGIN
    DECLARE ldni, lnotaPractica1, lnotaPractica2, lnotaExamen, notaP FLOAT;
    DECLARE hihaerror INT DEFAULT 0;
    DECLARE varcursor CURSOR FOR
        SELECT dni, notaPractica1, notaPractica2, notaExamen 
        FROM Alumne;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET hihaerror = 1;
    OPEN varcursor;
    loopA: loop
        FETCH varcursor INTO ldni, lnotaPractica1, lnotaPractica2, lnotaExamen;
        IF hihaerror THEN
            LEAVE loopA;
        END IF;
        SET notaP = 0.15*lnotaPractica1 + 0.35*lnotaPractica2 + 0.5*lnotaExamen;
        if notaP < 5.0 THEN
            UPDATE alumne SET notaFinal = notaP WHERE dni = ldni;
        ELSE
            IF lnotaPractica1 < 5.0 OR lnotaPractica2 < 5.0 OR lnotaExamen < 5.0 THEN
                UPDATE alumne SET notaFinal = 4.5 WHERE dni = ldni;
            ELSE
                UPDATE alumne SET notaFinal = notaP WHERE dni = ldni;
            END IF;
        END IF;
    END LOOP;
    CLOSE varcursor;
END

/*
3. Generi el codi dels programes HTML i PHP necessaris per fer l’alta de la taula alumne
(demanant el dni i el nom complet). Poden obviar les sentències de connexió a la base de dades posant
un comentari en el seu lloc.
*/

-- FORMULARIO
<html>
    <body>
        <form action="regist.php" method="GET">
            DNI: <input name="dni">
            Nombre: <input name="nom">
            Llinatge: <input name="llinatge">
            <input type="submit" name="Registrar">
        </form>
    </body>
</html>

-- CODIGO PHP PARA EL INSERT
Establecemos la conexión;
$dni = $_GET["dni"];
$nom = $_GET["nom"];
$llinatge = $_GET["llinatge"];
$query = "INSERT INTO alumne ";
$query.= "SET dni='".$dni."', nom='".$nom."', llinatge='".$llinatge."';";
$insert = mysql_query($conexio, $query);

/*
4. Expliqui el concepte de motor d’emmagatzemament i quins dels dos motors principals
empraria per a cada una de les taules anteriors, justificant els motius. Poden fer suposicions.

Motor de almacenamiento: proceso que se encarga, en última instancia de almacenar, manipular
y recuperar la información de una tabla

Los dos tipos principales que podemos identificar son InnoDB y MyISAM
InnoDB: motor transaccional, que controla la concurrencia, es más segura, mejor rendimiento
para hacer INSERTS y UPDATES. Conveniente para pocas consultas
MyISAM: motor no-transaccional, no control concurrente de los satos, consultas rápidas,
minimiza el espacio necesario para la BD en el disco. Conveniente para muchas consultas.

AEROPUERTO -> MyISAM: se realizan muy pocas inserciones, ya que para ello se han de añadir
aeropuertos mñas allá de los creados. Muchas consultas para relacionar nombre-código
VUELO -> MyISAM: consultar los vuelos desde las reservas.
PASAJERO -> MyISAM: muchas consultas para relacionar al pasajero con la reserva.
RESERVA -> InnoDB: creación de muchas reservas diarias, es esencial la concurrencia. Es
la que une todos los datos de las demás tablas.
*/

/*
5. Generi el codi de les instruccions SELECT d’SQL que permeten realitzar les següents
consultes:
a) Mostrar la llista de tots els vols i el total de reserves que s’han fet del vol.
b) Mostrar la llista de persones (indicant l’atribut idpass) que han reservat un vol amb sortida des de
l’aeroport de codi PMI amb arribada al de codi BIO
*/

-- a)
SELECT vol.idvol AS VUELO, COUNT(reserva.vol) as RESERVAS FROM vol
JOIN reserva
ON reserva.vol = vol.idvol
GROUP BY vol.idvol

-- b)
SELECT reserva.passatger FROM reserva
JOIN
(SELECT vol.idvol as numvuelo FROM vol
WHERE vol.origen = 'PMI' and vol.desti = 'BIO') codiVol
ON reserva.vol = codiVol.numvuelo
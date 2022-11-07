DELIMITER //
CREATE FUNCTION factorial(valor int)
RETURNS int
BEGIN
    DECLARE  resultat INT;
    SET resultat = 1;
    SET i = 1;
    WHILE i < resultat
    BEGIN
        SET resultat = resultat * i;
        SET i = i+1;
    END
    RETURNS resultat;
END
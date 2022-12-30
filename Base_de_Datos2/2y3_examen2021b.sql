/*
2. Generin una vista no materialitzada que, per a cada persona, calculi la mitjana entre el risc derivat per
categoria professional i el máxim risc derivat de les seves doléncies. La taula s'ha de dir CONSULTA i ha de
contenir els atributs: IDPER (id de la persona), MITJANA (resultat de la mitjana anterior), MAXRISC (rise
máxim per doléncies) i ANYPER (any de naixement de la persona). Si la persona no pateix cap dolencia, llavors
el MAXRISC se suposa que és 0.
*/

CREATE VIEW CONSULTA AS
SELECT persona.idPer as IDPER, (categoria.RISCCAT + IFNULL(MAX(peligroMal.MalRisc), 0))/2 as MITJANA,
IFNULL(MAX(peligroMal.MalRisc), 0) as MAXRISC, persona.anyPer as ANYPER FROM persona
JOIN categoria
ON categoria.idCat = persona.catper
LEFT JOIN 
    (SELECT dolencies.idper as persD, malatia.riscMal as MalRisc FROM dolencies
    JOIN malatia
    ON dolencies.idmal = malatia.idMal) as peligroMal
ON persona.idPer = peligroMal.persD
GROUP BY persona.idPer

/*
3. A partir de la vista anterior (si no s'ha resolt se pot suposar que existeix la vista de nom CONSULTA),
es demana que programi un PROCEDURE o FUNCTION d'SQL que calculi i modifiqui l'atribut RISCPER de
cada registre de PERSONA segons l'algorisme següent: és igual al risc mitjà que retorna la vista CONSULTA
(atribut MITJANA), excepte en els casos en qué l'edat de la persona és superior o igual a 80 anys o el risc máxim
per malalties és igual a 10 (atribut MAXRISC). En ambdues excepcions el RISCPER será igual a 10.
*/

DELIMITER //
CREATE procedure calculoRisc()
BEGIN
    DECLARE lidper, lmitjana, lmaxrisc, lanyper INT;
    DECLARE hihaerror INT DEFAULT 0;
    DECLARE varcursor CURSOR FOR
        SELECT IDPER, MITJANA, MAXRISC, ANYPER
        FROM CONSULTA;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET hihaerror = 1;
    OPEN varcursor;
    loopC: loop
        FETCH varcursor INTO lidper, lmitjana, lmaxrisc, lanyper;
        IF hihaerror THEN
            LEAVE loopC;
        END IF;
        IF (YEAR(CURDATE()) - lanyper) >= 80 OR lmaxrisc = 10 THEN
            UPDATE persona SET RISCPER = 10 WHERE idper = lidper;
        ELSE
            UPDATE persona SET RISCPER = lmitjana WHERE idper = lidper;
        END IF;
    END LOOP;
    CLOSE varcursor;
END;


/*
4. Relacioni, expliqui o valori (avantatges o inconvenients) la utilització de PROCEDURES d'SQL en
front a aplicacions amb accés remot per Web quant a: execució de tasques programades, interacció amb l'usuari,
control d'accés a informació de les taules per part dels programadors de l”aplicació remota, processament de
grans volums de dades, i depuració d'errades.
*/
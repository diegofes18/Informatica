DELIMITER //
CREATE FUNCTION doble(valor INT)
RETURNS INT
BEGIN
   DECLARE resultat INT;
   SET resultat=2*valor;
   RETURN resultat;
END;
------> SELECT doble(5)
------> SELECT doble(idprop) FROM propietari

Crear una function que calculi el factorial
d'un número passat per argument
----> SELECT factorial(5)

DELIMITER //
CREATE FUNCTION factorial(valor INT)
RETURNS INT
BEGIN
   DECLARE resultat INT default 1;
   if valor = 0 THEN
    set resultat = 1;
   else 
     while valor > 0 DO
       set resultat = resultat * valor;
       set valor = valor -1;
     END while;
    end if;
   RETURN resultat;
END;
---> select factorial(3)

Crear una function que, donats 3 paràmetres, retorni el valor d'enmig
select mig(a,b,c)

delimiter //
create mig (a int, b int, c int)
returns int
BEGIN
   if a < B 
      if b < c 
         set resultat = B;
      else
         if c < a 
            set resultat = a;
         else
            set resultat = c;
         end if;
      end if;
    else ------> b < a
      ...
    end if;
    return resultat;
end;



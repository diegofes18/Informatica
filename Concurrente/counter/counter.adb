with Ada.Text_IO; use Ada.Text_IO;

procedure Counter is
   n: Integer := 0;
   pragma Volatile(n);
   task type Tasca_Comptador;
   task body Tasca_Comptador is
      temp: integer;
   begin
      for i in 1..5000000 loop
         temp := n;
         n := temp + 1;
      end loop;
      Put_Line("----->> n és " & Integer'Image(N));
   end Tasca_Comptador;

begin
   declare
     p, q : Tasca_Comptador;
   begin
     null;
   end;
   Put_Line("El valor de n és " & Integer'Image(N));
end Counter;

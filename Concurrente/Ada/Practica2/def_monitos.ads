with Ada.Strings.Bounded;
package def_monitor is

    protected type monitor is
        entry smoker_sol(nom: in Ada.Strings.Unbounded.Unbounded_String);
        entry nsmoker_sol(nom: in Ada.Strings.Unbounded.Unbounded_String);
        procedure leave(nom: in Ada.Strings.Unbounded.Unbounded_String);

    private
        type sala is record
          tipo: integer := 0;
          free: integer := 3;
          table_names : array(1..3) of Ada.Strings.Unbounded.Unbounded_String := ("null", "null", "null");
        end record;

        type arraySala is array(1..3) of sala;
        salas := arraySala;
        free: integer:= 3;
        no_smoke: integer:= 0;
        smoke: integer:= 0;

    end monitor;

end def_monitor;
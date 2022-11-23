with Ada.Strings.Bounded;

package body def_monitor is
    protected body monitor is
    entry smoker_sol(nom: in Ada.Strings.Unbounded.Unbounded_String) when (free>0) or (smoke>0) is 
    begin

    end smoker_sol;

    entry nsmoker_sol(nom: in Ada.Strings.Unbounded.Unbounded_String) when (free>0) or (no_smoke>0) is 
    begin

    end nsmoker_sol;

    procedure leave(nom: in Ada.Strings.Unbounded.Unbounded_String) is 
    begin

    end leave;

    end monitor;
end def_monitor;
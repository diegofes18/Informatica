import java_cup.runtime.*;

%%

%unicode
%cup

nl		=  \n|\r|\r\n
intNumber	=  [1-9][0-9]*

%%

{intNumber}	{System.out.println("NUMBER:"+yytext());return new Symbol(sym.NUMBER);}


"+"		{System.out.println("PLUS");return new Symbol(sym.PLUS);}
"-"		{System.out.println("MINUS");return new Symbol(sym.MINUS);}


"("		{System.out.println("OPEN BRACKET");return new Symbol(sym.OBRACKET);}
")"		{System.out.println("CLOSE BRACKET");return new Symbol(sym.CBRACKET);}

"="		{System.out.println("EQUAL");return new Symbol(sym.EQUAL);}

{nl}|" " 	{;}

.		{System.out.println("Error:" + yytext());}


%{
/**
 * Assignatura 21780 Compiladors
 * Estudis de Grau en Informàtica
 * Professor: Pere Palmer
 *
 * Compilació: flex -+ -olexic0.cpp < lexic0.l
 */
 
#include <math.h>
#include <stdlib.h>
#include <stdio.h>

int yylval;
extern yyerror(const char *);

#include <stdio.h>
#include <iostream>
#include <fstream>

%}

suma	\+
prod	\*
lparen	\(
rparen	\)
blanc	[ \t\r\n]+
valor	0|[1-9][0-9]*

%%
{blanc}			{ /* No fer res */ }
{suma}			{ cout << "Suma : " << this->YYText() << endl; }
{prod}			{ cout << "Prod : " << this->YYText() << endl; }
{lparen}		{ cout << "Lpar : " << this->YYText() << endl; }
{rparen}		{ cout << "Rpar : " << this->YYText() << endl; }
{valor}			{ cout << "Valor: " << this->YYText() << endl; }

.				{ cout << "Sobra: \"" << this->YYText() << "\"" << endl; }
%%
#ifndef yywrap
int yywrap() {
		return 1;
}
#endif 


// Programa principal.
int main(int argc, char*argv[]) {
	if (argc < 2) {
		cerr << "Indica un fitxer amb les dades d'entrada" << endl;
		exit(1);
	} 
	try {
		ifstream *in = new ifstream(argv[1]);
		
		yyFlexLexer scanner(in);
		
		scanner.yylex();   // <- El mètode d'invocació per començar a parsejar el document

		in->close();
		delete in;
	} catch (std::ifstream::failure e) {
		cerr << "Error processant el fitxer d'entrada " << endl;
		cerr << e.what() << endl;
	}
}

%{
/**
 * Compilació:
  flex -+ -oprova.cpp < prova.l 
  gcc prova.cpp -o prova.exe  -I"directori on s'ha instal·lat FLex\include"
  ***/

#include <stdio.h>
#include <iostream>
#include <fstream>
%}


%%
[0-9]+   { cout << "Enter: " << this->YYText() << endl; }
[^0-9]+   {/* No s'ha de fer res */}

%%
#ifndef yywrap
int yywrap() {
		return 1;
}
#endif 

int main(int argc, char*argv[]) {
	if (argc < 2) {
		cerr << "Indica un fitxer amb les dades d'entrada" << endl;
		exit(1);
	}
	try {
		ifstream *in = new ifstream(argv[1]);
		
		yyFlexLexer parser(in);
		
		parser.yylex();   // <- El mètode d'invocació per començar a parsejar el document

		in->close();
		delete in;
	} catch (std::ifstream::failure e) {
		cerr << "Error processant el fitxer d'entrada " << endl;
		cerr << e.what() << endl;
	}
}

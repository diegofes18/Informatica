	ORG    $600
x	DS.L 1
y	DS.L 1
mayor	DS.L 1
t0	DS.L 1
uno	DS.L 1
dos	DS.L 1
SL	DS.L 1 *Variable auxiliar para los saltos de linea de la consola

	ORG    $1000

START:
	MOVE.L #0,SL
	JMP E_MAIN
devuelveElMayor:
	MOVE.L 8(A7),y
	MOVE.L 12(A7),x
	MOVEM.L D0-D2,-(A7)
	 ;GE x,y,E_1
	MOVE.L (x),D0
	MOVE.L (y),D1
	CMP.L  D1,D0
	BGE E_1
	MOVE.L #0,t0
	JMP E_2
E_1:
	MOVE.L #-1,t0
E_2:
	MOVE.L (t0), D0
	MOVE.L #0, D1
	CMP.L  D0,D1
	BEQ E_3
	MOVE.L (x),mayor
	JMP E_4
E_3:
	MOVE.L (y),mayor
E_4:
	MOVE.L #mayor,A0
	MOVEM.L (A7)+,D0-D2
	RTS
E_MAIN:
	MOVE.L #15,uno
	MOVE.L #35,dos
	MOVE.L (uno),-(A7)
	MOVE.L (dos),-(A7)
	MOVE.L A0,-(A7)
	JSR devuelveElMayor
	MOVE.L (A0),x
	MOVE.L (A7)+,A0
	MOVE.L (A7)+,dos
	MOVE.L (A7)+,uno
	MOVE.L (x),-(A7)
	JSR imprimeix
	MOVE.L (A7)+,(x)
	ADD.L #1,SL
	JMP FIN$68K
imprimeix:
	MOVEM.L D0-D1,-(A7)
	MOVE.L (SL),D1
	MOVE.B #11,D0
	TRAP #15
	MOVE.L 12(A7),D1
	MOVE.B #3,D0
	TRAP #15
	MOVEM.L (A7)+,D0-D1
	RTS
imprimeixBooleano
	MOVEM.L D0-D1/A1,-(A7)
	MOVE.L (SL),D1
	MOVE.B #11,D0
	TRAP #15
	MOVE.L 16(A7),D1
	TST.L D1
	BEQ fals
	LEA true,A1
	JMP finb
fals
	LEA false,A1
finb
	MOVE.B #14,D0
	TRAP #15
	MOVEM.L (A7)+,D0-D1/A1
	RTS
true DC.B 'cert',0
false DC.B 'fals',0
	DS.W 0
escriure
	MOVEM.L D0-D1/A1,-(A7)
	MOVE.L (SL),D1
	MOVE.B #11,D0
	TRAP #15
	LEA write,A1
	MOVE.B #14,D0
	TRAP #15
	MOVE.L  #(13)<<8,D1
	OR.W (SL),D1
	MOVE.B #11,D0
	TRAP #15
	MOVE.B #4,D0
	TRAP #15
	MOVE.L D1,D2
	MOVEM.L (A7)+,D0-D1/A1
	RTS
write DC.B 'escriu aqui: ',0
	DS.W 0
neg:
	MOVEM.L D0,-(A7)
	MOVE.L 12(A7),D0
	CMP.L #0,D0
	BEQ CONVERTIR
	MOVE.L #0,D0
	JMP FINNEGARVALOR
CONVERTIR
	MOVE.L #-1,D0
FINNEGARVALOR
	MOVE.L  D0,D1
	MOVE.L (A7)+,D0
	RTS
FIN$68K
	END START

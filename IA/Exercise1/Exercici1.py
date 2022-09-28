import csv


def prRed(skk): print("\033[91m {}\033[00m".format(skk))


def prGreen(skk): print("\033[92m {}\033[00m".format(skk))


def prYellow(skk): print("\033[93m {}\033[00m".format(skk))


def prLightPurple(skk): print("\033[94m {}\033[00m".format(skk))


def prPurple(skk): print("\033[95m {}\033[00m".format(skk))


def prCyan(skk): print("\033[96m {}\033[00m".format(skk))


def prLightGray(skk): print("\033[97m {}\033[00m".format(skk))


def prBlack(skk): print("\033[98m {}\033[00m".format(skk))

class Persona:
    def __init__(self, name: str, age: int, sex: str, heigh: int):
        self.name = name
        self.age = age
        self.sex = sex
        self.heigh = heigh

    def __str__(self):
        return f"{self.name}, {self.age}, {self.sex}, {self.heigh}"


class Estudiante(Persona):
    def __init__(self, id: int, nota: float, name: str, age: int, sex: str, heigh: int):
        super().__init__(name, age, sex, heigh)
        self.id = id
        self.nota = nota

    def hasPassed(self):
        if self.nota > 5:
            return True
        else:
            return False

    def __str__(self) -> str:
        return f"{super().__str__()}, {self.id}, {self.nota}"
    def hasPassed(self):
        if self.nota > 5:
            return True
        else:
            return False

    def toList(self):
        return (self.name,self.age,self.sex,self.heigh,self.id,self.nota)
    def __str__(self) -> str:
        return super().__str__()

class Profesor(Persona):
    def __init__(self, valoracion:float, capacidad:int,name:str, age: int, sex: str, heigh: int):
        super().__init__(name, age, sex, heigh)
        self.valoracion = valoracion
        self.capacidad = capacidad
    
    def moreClasses(self):
        if self.capacidad > 200:
            return True
        else:
            return False
    
    def __str__(self) -> str:
        return super().__str__()

        
class Grupo:
    def __init__(self, name:str):
        self.name = name


    def set_prof(self, prof):
        self.prof = prof

    def set_listest(self, listaest):
        self.listest = listaest

    def __str__(self) -> str:
        return f"Nombre: {self.name}, "

def toCSV(listaest):
    outputfile = 'estudiantes.csv'
    f = open(outputfile, 'w')
    writer = csv.writer(f)
    principio = ("NOM", "EDAT","SEXE", "ALTURA", "ID", "NOTA")
    writer.writerow(principio)
    for i in listaest:
        writer.writerow(i.toList())

    f.close()

def main():
    global nameGrup, listaestk
    numest = int(input("Introduce el numero de estudiantes: "))
    numprof = int(input("Introduce el numero de profesores: "))
    estudiantes = []
    profesores = []


    for i in range(numest):
        prCyan(f"Estudiante {i+1}: ")
        nom = (input(f"Introduce el nombre de estudiante {i+1}: "))
        age = int(input(f"Introduce la edad del estudiante {i+1}: "))
        sex = input(f"Introduce sex del estudiante {i+1}: ")
        heigh = int(input(f"Introduce heigh del estudiante {i+1}:"))
        nota = float(input(f"Introduce la nota media del estudiante {i+1}:"))
        estudiantes.append(Estudiante(i,nota,nom,age, sex,heigh))



    for j in range(numprof):
        prCyan(f"Profesor {j+1}: ")
        nom = (input(f"Introduce el nombre del profesor {j+1}: "))
        age = int(input(f"Introduce la edad del profesor {j+1}: "))
        sex = input(f"Introduce sex del profesor {j+1}: ")
        heigh = int(input(f"Introduce heigh del profesor {j+1}:"))
        valor = float(input(f"Introduce la valoracion media del profesor {j+1}: "))
        capacidad = int(input(f"Introduce la capacidad media del profesor {j+1}: "))
        profesores.append(Profesor(valor,capacidad,nom,age,sex,heigh))
    
    numgrup = int(input("\n       Introduce el numero de grupos: "))
    strgrupos = []

    for k in range(numgrup):
        nameGrup = input(f"Introduce el nombre del grupo {i+1}:")
        strgrupos.append(nameGrup)

    prGreen("\n 1 ===============Ahora vamos a asignar los profesores=================")
    listgrup = []
    for k in range(numgrup):
        for h in profesores:
            print(f"    Profesor numero {i+1}. {h}: ")

        listgrup.append(Grupo(strgrupos[k]))
        profnum = int(input(f"Cual de estos profesores es el del grupo {listgrup[k]} (introduzca el numero)"))

        listgrup[k].set_prof(profesores[profnum-1])

    prGreen("\n 2 ===============Ahora vamos a asignar los Estudiantes=================")
    alumnosgrupo = []
    prYellow("\n 2.1 -----------CUANTOS----------------")
    for i in range(numgrup):
        x = int(input(f"Cuantos alumnos en el grupo {i+1} ?(máximo {len(estudiantes)}): "))
        alumnosgrupo.append(x)

    prYellow("\n 2.1 -----------QUIENES----------------")
    for i in range(numgrup):
        prPurple(f"GRUPO {strgrupos[i]}")
        for j in range(numest):
            print(f"    Estudiante {j+1}: {estudiantes[j]}")
        for k in range(alumnosgrupo[i]):
            listaestk = []
            j_est = int(input(f"Selecciona un estudiante de los {alumnosgrupo[i]-k} restantes: "))
            #Crear el estudiante en la lista vacia y añadirla
            listaestk.append(estudiantes[j_est-1])
        listgrup[i].set_listest(listaestk)

    toCSV(estudiantes)

if __name__ == '__main__':
    main()


class Persona:
    def __init__(self, name:str, age: int, sex: str, heigh: int):
        self.name = name
        self.age = age
        self.sex = sex
        self.heigh = heigh

    def __str__(self):
        return self.name


class Estudiante(Persona):
    def __init__(self,id:int,nota:float, name:str, age: int, sex: str, heigh: int):
        super().__init__( name, age, sex, heigh )
        self.id = id
        self.nota = nota
    
    def hasPassed(self):
        if self.nota > 5:
            return True
        else:
            return False
    
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
    def __init__(self, name:str, listest = None, prof = None):
        self.name = name
        self.listest = listest
        self.prof = prof

    def set_prof(self, prof:Profesor):
        self.prof = prof

    def set_listest(self, listaest:Estudiante):
        self.listest = listaest

    def __str__(self) -> str:
        return f"Nombre: {self.name}, profesor:{self.prof}"

def main():
    global nameGrup
    numest = int(input("Introduce el numero de estudiantes: "))
    numprof = int(input("Introduce el numero de profesores: "))
    estudiantes = [numest]
    profesores = [numprof]


    for i in range(numest):
        print(f"Estudiante {i+1}: ")
        nom = (input(f"Introduce el nombre de estudiante {i+1}: "))
        age = int(input(f"Introduce la edad del estudiante {i+1}: "))
        sex = input(f"Introduce sex del estudiante {i+1}: ")
        heigh = int(input(f"Introduce heigh del estudiante {i+1}:"))
        nota = float(input(f"Introduce la nota media del estudiante {i+1}:"))
        estudiantes.append(Estudiante(i,nota,nom,age, sex,heigh))

    for j in range(numprof):
        print(f"Profesor {j+1}: ")
        nom = (input(f"Introduce el nombre del profesor {j+1}: "))
        age = int(input(f"Introduce la edad del profesor {j+1}: "))
        sex = input(f"Introduce sex del profesor {j+1}: ")
        heigh = int(input(f"Introduce heigh del profesor {j+1}:"))
        valor = float(input(f"Introduce la valoracion media del profesor {j+1}: "))
        capacidad = int(input(f"Introduce la capacidad media del profesor {j+1}: "))
        profesores.append(Profesor(valor,capacidad,nom,age,sex,heigh))
    
    numgrup = int(input("Introduce el numero de grupos: "))
    strgrupos = [numgrup]

    for k in range(numgrup):
        nameGrup = input(f"Introduce el nombre del grupo {i+1}:")
        strgrupos.append(nameGrup)

    print("\n 1 ===============Ahora vamos a asignar los profesores=================")
    listgrup = [numgrup]
    for k in range(numgrup):
        for h in profesores:
            print(f"    Profesor numero {i+1}. {h}: ")

        profnum = int(input(f"Cual de estos profesores es el del grupo {listgrup[k]} (introduzca el numero)"))
        listgrup.append(Grupo(nameGrup[k],profesores[profnum-1]))

    print("\n 2 ===============Ahora vamos a asignar los Estudiantes=================")
    alumnosgrupo = [numgrup]

    for i in range(len(alumnosgrupo)):
        alumnosgrupo[i] = int(input(f"Cuantos alumnos en el grupo {i+1}"))

    print("\n 2.1 -----------Ahora vamos a asignar que grupos----------------")
    for i in range(numgrup):
        for j in range(numest):
            print(f"    Estudiante {j+1}: {estudiantes[j]}")
        for k in range(alumnosgrupo[i]):
            listaestk = [alumnosgrupo[i]]
            j_est = int(input(f"Selecciona un estudiante de los {alumnosgrupo[i]-k} restantes: "))
            #Crear el estudiante en la lista vacia y añadirla
            listaestk.append(estudiantes[j_est-1])
        listgrup[i].set_listest(listaestk)


#Falta csv


if __name__ == '__main__':
    main()


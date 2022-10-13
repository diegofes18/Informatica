from ia_2022 import entorn
from quiques.agent import Barca, Estat
from quiques.entorn import AccionsBarca


class BarcaAmplada(Barca):
    def __init__(self):
        super(BarcaAmplada, self).__init__()
        self.__oberts = None
        self.__tancats = None
        self.__accions = None

    def actua(self, percep: entorn.Percepcio) -> entorn.Accio | tuple[entorn.Accio, object]:
        inicial = Estat(percep.to_dict())
        self.oberts = []
        self.tancats = []

        self.oberts.append(inicial)

        while self.oberts:
            actual = self.oberts.pop(0)
            print(actual)
            if Estat.es_meta(actual):
                pass

            else:
                successors = actual.genera_fill()
                self.tancats.append(actual)
                self.oberts.append(successors)

        pass

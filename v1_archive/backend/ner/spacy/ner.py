from __future__ import unicode_literals, print_function

import plac
import spacy


TEXTS = [
    "Location d'un T2  bis de 45 m2 situé rue Mouneyra à BORDEAUX",
    "Situé au centre de Caudéran, proche de tous les commerces, la résidence dispose d'un parc arboré"
]


@plac.annotations(
    model=("Model to load (needs parser and NER)", "positional", None, str)
)
def main(model="fr_core_news_sm"):
    nlp = spacy.load(model)
    locsentences = []
    for text in TEXTS:
        doc = nlp(text)
        locinfo = ""
        loc_infos = (loc for loc in doc.ents if loc.label_ == "LOC")
        for locel in loc_infos:
            locinfo += locel.text + " "
        locsentences.append(locinfo)
    print(locsentences)

if __name__ == "__main__":
    plac.call(main)

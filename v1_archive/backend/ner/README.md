# Renty - AI improved real estate rent offers description reading

## Goals

We'll use the python [spaCy](https://spacy.io/) [EntityRecognizer](https://spacy.io/api/entityrecognizer) tool to 
get real estate named entities recognized in rent offers descriptions.

The available french langage pre-trained model (namely "fr_core_news_md") is already 
able to recognize some usefull entities for our use case :

- FAC : Abrev for facilities, like buildings, airports, highways, bridges, etc. 

- GPE : Location infos like Countries, cities or states.

- LOC : Other location informations, like street name 

- PRODUCT : Objects (Not services.)

- TIME : Times smaller than a day.

- MONEY : Monetary values, including unit.

- QUANTITY : Measurements, as of weight or distance

But we must first annotate new samples and train spacy with them to get these new entities :

- EMAIL : email addresses

- PHONE : phone numbers (offen found in rent offer description even if they are also available by api call)

- SURFACE : a rent offer volume indication


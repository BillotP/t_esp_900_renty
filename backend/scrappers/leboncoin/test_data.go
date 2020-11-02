package main

import "github.com/BillotP/renty/backend/lib/v2/models"

var fooo = models.RentOffer{
	Title: []models.Label{
		{
			Value:  "à louer Bureau calme, meublé",
			Locale: "fr",
		},
	},
	Source: models.Source{
		URL: "https://www.leboncoin.fr/bureaux_commerces/1853777496.htm/",
	},
	Description: []models.Label{
		{
			Value: `
			Bureau entier meublé à louer, ce n'est pas un Coworking.

			à coté de l'arrêt du Tram C - Camille Godard et à 600m de la place des Chartrons.
			Nous proposons à la location un bureau meublé de 44m2 situé au RDC d'un immeuble en pierres Bordelaise.
			Immeuble d'un seul étage.

			Etat neuf, travaux de rénovation réalisés en Juillet 2019.
			L'emplacement est très calme.

			Locaux aménagés actuellement avec:
			- 4 Bureaux.
			- 1 grande Table de réunion en bois massif avec 6 chaises.
			- 1 grande TV fixée au mur.
			- 1 coin canapé
			- 1 espace de rangement / coin cuisine avec micro onde et point d'eau.
			- WC séparés
			- Parquet massif au sol
			- volets roulants électriques
			- barreaux aux fenetres
			- belle hauteur sous plafond
			- fibre optique

			Bureaux très calme et très lumineux

			- Disponibilité : immédiate

			Loyer 1200 euros / mois
			`,
		},
	},
	Price: []models.Price{
		{
			Value: 1150,
		},
	},
	Assets: []models.Asset{
		{
			URL:     "https://img6.leboncoin.fr/ad-large/e9ee344a5fa53820ffd16893f65a9e5babdd7b5f.jpg",
			Type:    "PICT",
			Storage: "EXT",
		},
	},
	Location: models.Location{
		GeoJSON: models.GeoJSON{
			Type: "Feature",
			Geometry: models.Geometry{
				Type:        "Point",
				Coordinates: []float64{-0.617077, 44.856399},
			},
			Properties: models.Properties{
				Name:    "Rue Paul Doumer",
				Country: "France",
			},
		},
	},
	Offeror: models.Offeror{
		Name: "Thierry",
		Phone: &models.Phone{
			Value: "0676818921",
		},
		Type: "PART",
	},
}

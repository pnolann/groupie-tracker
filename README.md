# Groupie-Tracker

### 1. Objectif du projet

Groupie Tracker est un projet scolaire. Il a pour but de réaliser une application web nommée "Tracker de Tigres", permettant ainsi une consultation détaillée sur différents groupes et artistes de musique. On peut aussi retrouver une section sur les différents concerts via une API externe. L'application centralise les différentes données dans le but d'offrir une visualisation rapide sur différentes informations.

### 2. Comment lancer le projet

* **Prérequis** : Avoir Go installé.
* **Installation** : Téléchargez et extrayez le fichier zip du projet.
* **Lancement** : Ouvrez un terminal dans le dossier racine du projet et exécutez la commande :
  `go run main.go`
* **Accès** : Ouvrez votre navigateur à l'adresse suivante : `http://localhost:8080`.

### 3. Les routes principales et leurs fonctions

* **`/` (HomeHandler)** : Page d'accueil avec une barre de recherche rapide.
* **`/artists` (ArtistsHandler)** : Affiche la liste complète des différents artistes présents sur le site. On peut également y retrouver les différents filtres disponibles.
* **`/artist` (ArtistDetailHandler)** : Affiche les détails sur un artiste (ID demandé en paramètre). Se trouve également une carte interactive avec les différents concerts.
* **`/static/`** : Sert les fichiers statiques tels que le CSS et les images de fond.

### 4. Fonctionnalités implémentées

**Fonctionnalités obligatoires :**

* **Consommation d'API** : Récupération des données depuis : `https://groupietrackers.herokuapp.com/api`.
* **Affichage de liste** : Visualisation des différents artistes, avec leurs informations clés.
* **Pages de détails** : Présentation complète des groupes, avec image, nom, année de création, premier album, membres.
* **Recherche serveur** : Filtrages des différents groupes utilisant les requêtes GET : `/artists`.
* **Cartographie interactive** : Utilisation de la bibliothèque Leaflet pour afficher les lieux de concert sur une carte du monde.
* **Géocodage dynamique** : Utilisation de l'API Nominatim (OpenStreetMap) pour convertir les noms de lieux en coordonnées géographiques réelles.
* **Filtrage client (temps réel)** : Système de recherche dynamique en JavaScript sur la page de liste pour un affichage instantané sans rechargement.

### 5. Gestion de projet

Le projet est structuré selon une architecture modulaire pour faciliter la maintenance :

* **Modèles (`models/`)** : Structures des différentes données de l'API, gérées en Go et JSON.
* **Handlers (`handlers/`)** : Les différentes structures logiques, gérées par différentes routes.
* **Templates (`templates/`)** : Utilisation du moteur de rendu `html/template` de Go pour injecter dynamiquement les données dans le HTML.

**Lien GitHub :** https://github.com/pnolann/groupie-tracker

**Lien GoogleSheet :** https://docs.google.com/spreadsheets/d/1103yORsS1U36EOlT5XPl5mn-OlZj5yeQDB_LNwFPOSU/edit?usp=sharing

# Groupie Tracker - IGDB

Un site web en Go qui utilise l'API IGDB (Internet Game Database) pour afficher les catalogues de jeux vidéo de différents studios : **SEGA**, **Nintendo**, **Ubisoft** et **Level-5**. Le but de ce projet est de s'entraîner à utiliser les API et à créer des sites web en Go. De plus il permet de faire decouvrir des jeux vidéo au utilisateurs.

## Fonctionnalités

- **Catalogue par studio** — Parcourir les jeux de SEGA, Nintendo, Ubisoft et Level-5
- **Fiche détaillée** — Afficher les détails d'un jeu
- **Filtres** — Filtrer par catégorie, genre et plateforme (les filtres se combinent)
- **Pagination** — Navigation par pages (20 jeux par page)
- **Recherche** — Rechercher un jeu par nom via la barre de recherche
- **Favoris** — Ajouter/supprimer des jeux en favoris
- **Page 404 stylisée** — Page d'erreur thème rétro-gaming avec effet glitch
- **Mini-jeu Dino** — Jeu du dinosaure accessible depuis la page 404(petit bonus)

## Technologies Utiliser

 **Go**  Serveur backend
 **HTML/CSS** Templates et mise en page 
 **JavaScript** Mini-jeu Dino. Puis animation de la page index 
 **API IGDB** Source de données des jeux 
 **JSON** Stockage local des favoris 

## Structure du projet

```
SIte_Api_IGBD/
├── src/
│   ├── cmd/main.go           
│   ├── controllers/           
│   ├── models/                
│   ├── routers/               
│   ├── services/              
│   ├── helpers/               
│   └── templates/             
├── templates/                 
├── assets/
│   ├── css/                   
│   ├── js/                    
│   └── images/                
├── favorites.json             
└── go.mod                     
```

## Lancement

### Prérequis

- **Go** installé
- Clés API IGDB (Client ID et Token)

### Installation et exécution

```bash
# Cloner le projet
git clone https://github.com/Zyonnn13/Site-IGBD.git
cd SIte_Api_IGBD

# Lancer le serveur
go run ./src/cmd/main.go
```

Le serveur démarre sur **http://localhost:8080**

## Pages disponibles

| Route | Description |
|---|---|
| `/` | Page d'accueil |
| `/sega` | Catalogue SEGA |
| `/nintendo` | Catalogue Nintendo |
| `/ubisoft` | Catalogue Ubisoft |
| `/level-5` | Catalogue Level-5 |
| `/game?id=XXX` | Détails d'un jeu |
| `/search?q=XXX` | Résultats de recherche |
| `/favorites` | Page des favoris |
| `/dino` | Mini-jeu Dino |

##  Auteurs

- **Clément BELMONDO** — Ynov B1 Informatique

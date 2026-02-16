# Groupie Tracker - IGDB

Un site web en Go qui utilise l'API IGDB (Internet Game Database) pour afficher les catalogues de jeux vidéo de différents studios : **SEGA**, **Nintendo**, **Ubisoft** et **Level-5**.

## Fonctionnalités

- **Catalogue par studio** — Parcourir les jeux de SEGA, Nintendo, Ubisoft et Level-5
- **Fiche détaillée** — Afficher les détails d'un jeu (description, note, date de sortie, DLCs)
- **Filtres cumulatifs** — Filtrer par catégorie, genre et plateforme (les filtres se combinent)
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

## 📁 Structure du projet

```
SIte_Api_IGBD/
├── src/
│   ├── cmd/main.go            # Point d'entrée
│   ├── controllers/           # Logique des handlers HTTP
│   ├── models/                # Structures de données (Game, Genre, Platform...)
│   ├── routers/               # Définition des routes
│   ├── services/              # Appels API IGDB et gestion des favoris
│   ├── helpers/               # Fonctions utilitaires
│   └── templates/             # Chargement des templates Go
├── templates/                 # Fichiers HTML (templates Go)
├── assets/
│   ├── css/                   # Feuilles de style
│   ├── js/                    # Scripts JavaScript
│   └── images/                # Images et sprites
├── favorites.json             # Données des favoris
└── go.mod                     # Module Go
```

## 🚀 Lancement

### Prérequis

- **Go** installé
- Clés API IGDB (Client ID et Token)

### Installation et exécution

```bash
# Cloner le projet
git clone 
cd SIte_Api_IGBD

# Lancer le serveur
go run ./src/cmd/main.go
```

Le serveur démarre sur **http://localhost:8080**

## 📄 Pages disponibles

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

## 👥 Auteurs

- **Belmonte** — Ynov B1 Informatique

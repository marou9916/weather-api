# API Météo

## Introduction  
Ce projet est une API météo simple développée en Go. Elle récupère les données météo via l'API Visual Crossing, les met en cache avec Redis et les expose à travers des endpoints RESTful. L'objectif est de renforcer les compétences en gestion de cache et d'intégration de services externes tout en appliquant les bonnes pratiques en développement backend.

## Objectifs du projet  
- Intégrer un système de cache avec Redis.  
- Gérer les données météo via une API externe.  
- Implémenter des mécanismes d'expiration du cache.  
- Normaliser la gestion des clés dans Redis.  
- Appliquer les principes de bonnes pratiques en développement backend.

## Fonctionnalités  
- Récupération des données météo via Visual Crossing.  
- Mise en cache des résultats avec Redis.  
- Normalisation des requêtes pour éviter les problèmes de casse.  
- Test d’expiration des données avec réduction du délai d'expiration à 10 secondes.

## Technologies utilisées  
- **Langage** : Go  
- **Cache** : Redis  
- **API externe** : Visual Crossing  
- **Framework HTTP** : Gin  
- **Documentation API** : Swagger

## Installation et exécution  
*Prérequis*  
- Go (version 1.20 ou supérieure recommandée)  
- Redis installé et en fonctionnement sur votre machine.

*Étapes*  
1. Cloner le projet :
   ```bash
   git clone https://github.com/tonutilisateur/api-meteo.git
   cd api-meteo

2. Installer les dépendances :
    go mod tidy

3. Exécuter le projet :
    go run .

4. L'API sera disponible à l'adresse http://localhost:8080.

## Structure du projet
```
.
├── cache
│   └── redis.go          # Gestion de Redis, cache et expiration des données
├── cmd
│   └── main.go           # Point d'entrée de l'application
├── configs
│   └── redis.go          # Configuration des paramètres Redis
├── controllers
│   └── weather_handler.go # Gestion des endpoints relatifs à la météo
├── docs
│   └── README.md         # Documentation du projet
├── go.mod                # Dépendances du projet
├── go.sum                # Vérification des dépendances
├── models
│   └── weatherdata.go    # Modèle des données météo
├── routes
│   └── router.go         # Configuration des routes de l'API
└── services
    └── visual_crossing.go # Intégration de l'API Visual Crossing pour récupérer les données météo
```

## Compétences renforcées

Intégration de services externes via API.

Mise en place d'un cache avec Redis et gestion de l'expiration.

Normalisation des données pour éviter les erreurs de casse dans Redis.

Documentation API avec Swagger.

## Améliorations futures

Ajouter des tests unitaires.

Support de géolocalisation pour adapter l’API à différents pays.

Mise en place d’une gestion d'erreur améliorée.

## Contributions

Les contributions sont les bienvenues ! Ouvrez une issue ou une pull request pour proposer des améliorations ou signaler des bugs.

## Licence

Ce projet est sous licence MIT. Vous êtes libre de l'utiliser, de le modifier et de le distribuer conformément aux termes de cette licence.


---

Tu n'as plus qu'à remplacer "tonutilisateur" par ton nom d'utilisateur GitHub et à l'utiliser dans ton projet !

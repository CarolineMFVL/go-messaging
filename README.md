# go-messaging

## Description

**go-messaging** est une application backend écrite en Go permettant de gérer un système de messagerie en temps réel avec authentification, gestion des threads et WebSocket. Elle utilise PostgreSQL comme base de données et GORM comme ORM.

## Fonctionnalités

- Gestion des threads de discussion
- Messagerie en temps réel via WebSocket
- Migration automatique de la base de données

## Prérequis

- [Go](https://golang.org/) 1.22 ou supérieur
- [Docker](https://www.docker.com/) et [Docker Compose](https://docs.docker.com/compose/)
- PostgreSQL

## Installation

1. Clonez le repository :

   ```bash
   git clone https://github.com/votre-utilisateur/go-messaging.git
   cd go-messaging

   Installez les dépendances Go :
   ```

Configurez les variables d'environnement dans docker-compose.yml ou via un fichier .env.

Lancer l'application
Avec Docker
Démarrez les services avec Docker Compose :

`docker-compose up --build`

1. Créez le rôle/utilisateur dans PostgreSQL
   Dans un terminal, lancez :

`psql -h localhost -p 5432 -U postgres -d nls_db`

postgres
Puis, dans le shell psql :

````CREATE ROLE cmf WITH LOGIN PASSWORD 'test1234';
ALTER ROLE cmf CREATEDB;
GRANT ALL ON SCHEMA public TO cmf;```

L'application sera disponible sur http://localhost:4000.

En local
Lancez PostgreSQL et configurez les variables d'environnement nécessaires (PG_HOST, PG_USER, PG_PASSWORD, PG_DB, PG_PORT).

Lancez l'application :

`go run main.go`

Documentation API
La documentation Swagger est générée automatiquement. Pour la générer, utilisez la commande suivante :

`make open-api`

Tests
Pour exécuter les tests, utilisez la commande suivante :

`go test ./...`


````

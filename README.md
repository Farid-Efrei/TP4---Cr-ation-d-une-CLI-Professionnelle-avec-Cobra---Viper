# TP4 — Mini-CRM CLI (Go, Cobra, Viper, GORM/SQLite)

Description  
Mini-CRM CLI est un gestionnaire de contacts professionnel en ligne de commande, écrit en Go. Ce projet illustre les bonnes pratiques de développement Go modernes : architecture en packages découplés, injection de dépendances via interfaces, CLI professionnelle avec Cobra, configuration externe avec Viper, et persistance multi-backends (Memory/JSON/GORM+SQLite).

Il s'agit du TP4 final d'initiation à Go, transformant le simple script du TP1 en une application robuste et configurable.

## Fonctionnalités principales

- **CRUD complet des contacts** : Ajouter, Lister, Mettre à jour, Supprimer des contacts.
- **Multi-backends de stockage** : Basculer entre 3 types de persistance sans recompiler.
- **CLI professionnelle** : Interface Cobra avec sous-commandes et flags intuitifs.
- **Configuration externe** : Fichier YAML pour personnaliser le comportement.
- **Architecture découplée** : Interfaces, injection de dépendances, packages modulaires.
- **Thread-safety** : Accès concurrent sécurisé aux données.
- **Timestamps automatiques** : Suivi des dates de création et modification.

## Architecture du projet

```
TP4/
├── cmd/                    # Commandes CLI (Cobra)
│   ├── root.go            # Commande racine + configuration
│   ├── add.go             # Ajouter un contact
│   ├── get_all.go         # Lister tous les contacts
│   ├── update.go          # Mettre à jour un contact
│   └── delete.go          # Supprimer un contact
├── internal/
│   ├── models/
│   │   └── contact.go     # Modèle Contact (tags GORM/JSON)
│   └── storage/
│       ├── storage.go     # Interface Storer
│       ├── memory.go      # Stockage en mémoire
│       ├── json.go        # Stockage fichier JSON
│       └── gorm.go        # Stockage SQLite via GORM
├── config/
│   └── config.go          # Factory pour créer le store (Viper)
├── config.yaml            # Configuration (type de stockage)
├── data/                  # Dossier de données (auto-créé)
├── main.go                # Point d'entrée
├── go.mod                 # Dépendances Go
└── README.md              # Cette documentation
```

## Types de stockage disponibles

### 1. Memory Store (`type: "memory"`)

- **Usage** : Tests, développement, données temporaires.
- **Avantages** : Ultra-rapide, thread-safe.
- **Inconvénients** : Données perdues à la fermeture.

### 2. JSON Store (`type: "json"`)

- **Usage** : Persistance simple, données lisibles.
- **Avantages** : Fichier JSON lisible, portable.
- **Inconvénients** : Performance limitée pour gros volumes.

### 3. GORM Store (`type: "gorm"`)

- **Usage** : Production, gros volumes, requêtes complexes.
- **Avantages** : Base SQLite, migrations automatiques, performance.
- **Inconvénients** : Plus complexe pour des cas simples.

## Installation et configuration

### Prérequis

- Go 1.21+ installé
- PowerShell (Windows) ou terminal

### Installation

```powershell
# Se placer dans le dossier TP4
cd "C:\Users\Dev_Note\Desktop\Dev-Student\M2\Golang\Cours 1\exemple\TP4"

# Installer les dépendances
go mod tidy

# Optionnel : compiler l'exécutable
go build -o crm.exe .
```

### Configuration (config.yaml)

```yaml
storage:
  type: "json" # "memory" | "json" | "gorm"
  json_path: "data/contacts.json" # Chemin du fichier JSON
  gorm_path: "data/crm.db" # Chemin de la base SQLite
```

## Commandes disponibles

### Aide générale

```powershell
# Afficher l'aide principale
go run main.go --help
# ou avec l'exécutable
.\crm.exe --help
```

### Ajouter un contact (`add`)

```powershell
# Syntaxe
go run main.go add --nom="NOM" --email="email@domain.com"

# Exemples concrets
go run main.go add --nom="Tanjiro Kamado" --email="tanjiro@kimetsu.jp"
go run main.go add --nom="Nezuko Kamado" --email="nezuko@kimetsu.jp"
go run main.go add -n "Zenitsu Agatsuma" -e "zenitsu@kimetsu.jp"

# Équivalent avec exécutable
.\crm.exe add --nom="Inosuke Hashibira" --email="inosuke@kimetsu.jp"
```

**Flags disponibles :**

- `--nom, -n` : Nom du contact (obligatoire)
- `--email, -e` : Email du contact (obligatoire)

### Lister tous les contacts (`get_all`)

```powershell
# Lister tous les contacts
go run main.go get_all

# Avec exécutable
.\crm.exe get_all
```

**Sortie exemple :**

```
ID: 1, Nom: Tanjiro Kamado, Email: tanjiro@kimetsu.jp, Créé le: 2024-01-15 14:30:25, Mis à jour le: 2024-01-15 14:30:25
ID: 2, Nom: Nezuko Kamado, Email: nezuko@kimetsu.jp, Créé le: 2024-01-15 14:31:10, Mis à jour le: 2024-01-15 14:31:10
```

### Mettre à jour un contact (`update`)

```powershell
# Syntaxe
go run main.go update [ID] --nom="NOUVEAU_NOM" --email="nouvel@email.com"

# Exemples concrets
go run main.go update 1 --nom="Tanjiro (Pilier de l'Eau)"
go run main.go update 2 --email="nezuko.demon@kimetsu.jp"
go run main.go update 1 --nom="Tanjiro Kamado" --email="tanjiro.pilier@kimetsu.jp"

# Avec exécutable
.\crm.exe update 1 -n "Nouveau Nom" -e "nouveau@email.com"
```

**Flags disponibles :**

- `--nom, -n` : Nouveau nom (optionnel)
- `--email, -e` : Nouvel email (optionnel)
- Au moins un des deux flags est requis

### Supprimer un contact (`delete`)

```powershell
# Syntaxe
go run main.go delete [ID]

# Exemples concrets
go run main.go delete 1
go run main.go delete 2

# Avec exécutable
.\crm.exe delete 1
```

### Aide par commande

```powershell
# Aide spécifique pour chaque commande
go run main.go add --help
go run main.go get_all --help
go run main.go update --help
go run main.go delete --help
```

## Exemples d'utilisation complète

### Scénario 1 : Utilisation avec JSON Store (par défaut)

```powershell
# Vérifier la configuration (doit être type: "json")
Get-Content config.yaml

# Ajouter quelques contacts
go run main.go add --nom="Giyu Tomioka" --email="giyu@pilier.water"
go run main.go add --nom="Shinobu Kocho" --email="shinobu@pilier.insect"

# Lister pour vérifier
go run main.go get_all

# Modifier un contact
go run main.go update 1 --nom="Giyu Tomioka (Pilier de l'Eau)"

# Vérifier les modifications (noter le timestamp mis à jour)
go run main.go get_all

# Supprimer un contact
go run main.go delete 2

# Vérifier le fichier JSON créé
Get-Content data\contacts.json
```

### Scénario 2 : Basculer vers GORM/SQLite

```powershell
# Modifier la configuration
# Dans config.yaml, changer: type: "gorm"

# Les mêmes commandes fonctionnent
go run main.go add --nom="Kyojuro Rengoku" --email="kyojuro@pilier.flame"
go run main.go get_all

# Vérifier que la base SQLite est créée
Test-Path data\crm.db
```

### Scénario 3 : Mode Memory (pour tests)

```powershell
# Dans config.yaml, changer: type: "memory"

# Ajouter des contacts (perdus à la fermeture)
go run main.go add --nom="Contact Temporaire" --email="temp@test.com"
go run main.go get_all
```

## Concepts Go illustrés

### Architecture et design patterns

- **Interface Storer** : Contrat commun pour tous les types de stockage.
- **Injection de dépendances** : Factory pattern avec Viper pour créer le store.
- **Separation of Concerns** : CMD (interface) / Storage (persistance) / Models (données).

### Concurrence et sécurité

- **Thread-safety** : Mutex (`sync.RWMutex`) dans MemoryStore et JSONStore.
- **Copies défensives** : Éviter les mutations externes des données internes.

### Outils et libraries

- **Cobra** : CLI professionnelle avec sous-commandes et auto-completion.
- **Viper** : Configuration externe flexible (YAML, ENV, flags).
- **GORM** : ORM moderne avec migrations automatiques.

### Gestion d'erreurs

- **Propagation d'erreurs** : `RunE` dans Cobra pour gérer les erreurs proprement.
- **Validation** : Flags requis, vérification d'existence des contacts.

## Développement et tests

### Formatage et vérification

```powershell
# Formater le code
gofmt -w .

# Vérifier les erreurs potentielles
go vet ./...

# Détecter les race conditions
go run -race main.go get_all
```

### Auto-reload pendant le développement

```powershell
# Installer CompileDaemon (équivalent nodemon pour Go)
go install github.com/githubnemo/CompileDaemon@latest

# Lancer avec auto-reload
CompileDaemon -command "go run main.go"
```

### Structure des erreurs

- **Contact non trouvé** : Message explicite pour les IDs inexistants.
- **Validation** : Flags requis vérifiés par Cobra.
- **Stockage** : Erreurs de fichier/DB propagées proprement.

## Améliorations possibles

### Fonctionnalités

- **Recherche** : Commande `search` par nom ou email.
- **Import/Export** : CSV, Excel, JSON bulk operations.
- **Validation avancée** : Format email, numéros de téléphone.
- **Pagination** : Pour de gros volumes de contacts.

### Technique

- **Tests unitaires** : Suite complète pour chaque store.
- **CI/CD** : GitHub Actions pour tests automatiques.
- **API REST** : Exposer les fonctionnalités via HTTP.
- **Configuration avancée** : Profils (dev/prod), logs configurables.

## Dépannage

### Erreurs courantes

- **"type de stockage inconnu"** : Vérifier config.yaml et que Viper charge le fichier.
- **"Contact non trouvé"** : L'ID spécifié n'existe pas, utiliser `get_all` pour lister.
- **Permissions fichier** : S'assurer que le dossier `data/` est accessible en écriture.

### Debug

```powershell
# Vérifier la configuration chargée
go run main.go get_all
# Observe les messages DEBUG: storage.type

# Forcer un fichier de config spécifique
go run main.go --config="./config.yaml" get_all
```

## Critères de réussite validés ✅

- ✅ Programme compile et s'exécute sans erreur
- ✅ Toutes les sous-commandes (add, get_all, update, delete) fonctionnelles
- ✅ Base SQLite utilisée avec `type: "gorm"` (fichier .db créé et mis à jour)
- ✅ Basculement storage via config.yaml sans recompilation
- ✅ Code propre, formaté, et commenté
- ✅ Documentation complète et claire

## Auteur

TP4 réalisé dans le cadre du cours d'initiation Go. Evolution du Mini-CRM depuis un script simple (TP1) vers une application professionnelle avec architecture modulaire, CLI Cobra, configuration Viper et persistance multi-backends.
Réalisé par Fairytale-Dev(Farid-Efrei) (étudiant — Alternant à l'Efrei).

🦋 Paix à tous 🦋

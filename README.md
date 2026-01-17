# Ghost File System (GFS)

Ghost File System est un outil **CLI écrit en Go** qui implémente un système de **suppression logique de fichiers**.
Au lieu de supprimer définitivement un fichier, celui-ci est déplacé dans un espace sécurisé (`.ghost/`) afin de pouvoir être restauré ultérieurement.

Ce projet est orienté **système**, **robustesse** et **compréhension approfondie du module `os` de Go**.

---

## Objectif du projet

Dans les systèmes classiques, une suppression est irréversible.
Ghost File System introduit une couche de sécurité supplémentaire en permettant :

* la récupération de fichiers supprimés par erreur
* la traçabilité des suppressions
* un nettoyage contrôlé et sécurisé

L’objectif est de fournir un outil fiable, simple et proche du fonctionnement interne d’un système d’exploitation.

---

## Fonctionnalités principales

* Suppression logique (déplacement dans un espace fantôme)
* Restauration complète des fichiers supprimés
* Liste des fichiers fantômes avec métadonnées
* Suppression définitive contrôlée
* Nettoyage automatique des fichiers trop anciens
* Gestion stricte des erreurs système

---

## Fonctionnement général

Lorsqu’un fichier est supprimé via Ghost File System :

1. Le fichier n’est pas détruit
2. Il est déplacé vers un dossier caché `.ghost/`
3. Des métadonnées sont enregistrées (chemin original, date, permissions, taille)
4. Le fichier peut être restauré ou supprimé définitivement plus tard

---

## Structure du projet

```text
ghostfs/
├─ cmd/
│  └─ main.go
├─ internal/
│  ├─ ghost/
│  │  ├─ delete.go
│  │  ├─ restore.go
│  │  ├─ list.go
│  │  └─ clean.go
│  ├─ metadata/
│  │  └─ metadata.go
│  └─ utils/
│     └─ paths.go
├─ .ghost/
│  └─ metadata.json
├─ go.mod
└─ README.md
```

---

## Commandes disponibles

```bash
ghost delete <fichier>
ghost list
ghost restore <fichier>
ghost purge <fichier>
ghost clean --days <nombre>
```

### Détails des commandes

* `delete` : supprime logiquement un fichier
* `list` : affiche les fichiers supprimés
* `restore` : restaure un fichier à son emplacement d’origine
* `purge` : supprime définitivement un fichier fantôme
* `clean` : nettoie automatiquement les fichiers anciens

---

## Exemple de métadonnées

```json
{
  "example.txt": {
    "original_path": "/home/user/docs/example.txt",
    "deleted_at": "2026-01-17T14:32:00",
    "size": 2048,
    "permissions": "0644"
  }
}
```

---

## Technologies utilisées

* Go (standard library uniquement)
* Module `os`
* Module `encoding/json`

---

## Concepts abordés

* Manipulation avancée du système de fichiers
* Gestion des permissions et erreurs OS
* Conception d’un outil CLI robuste
* Architecture propre et maintenable
* Sécurité des opérations critiques

---

## Objectif pédagogique

Ce projet permet de maîtriser :

* le module `os` de Go en profondeur
* les opérations système réelles
* la gestion fiable des fichiers
* les bonnes pratiques d’outils système

Il est conçu comme un **projet de niveau ingénieur junior à intermédiaire**.

---

## Améliorations possibles

* Chiffrement des fichiers fantômes
* Versioning des fichiers supprimés
* Support multi-plateforme avancé
* Intégration avec Nexupload
* Interface graphique (Wails)

---

## Licence

Projet éducatif et open-source.

---


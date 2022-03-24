# HANGMAN-WEB

  - [Présentation du projet](#présentation-du-projet)
  - [Installation et utilisation](#installation-et-utilisation)
  - [Règle du jeu](#règle-du-jeu)
  - [Fonctionnalités](#fonctionnalités)

## Présentation du projet

Hangman-web est une application web permettant de jouer au pendu depuis un navigateur. Ce projet a plusieurs fonctionnalités supplémentaires comme faire une sauvegarde de sa partie ou une gestion de tableaux des scores.

## Installation et utilisation

L'installation est simple, il suffit de cloner le répertoire git en faisant la commande :
```bash
git clone https://git.ytrack.learn.ynov.com/LGRAND-MORCEL/hangman-web.git
```
Et d'ensuite aller dans le fichier `HANGMAN-WEB` et faire :
```bash
go run web/main.go
```
Une fois ces commandes faites vous pourrez immédiatement ouvrir votre navigateur et accéder au site en tapant `http://localhost:8080`. Vous allez arriver sur une page où vous devrez rentrer votre pseudo pour votre ScoreBoard. Ensuite, vous serez sur la page de sélection de niveaux. Vous aurez le choix entre trois niveaux de difficulté et si vous avez sauvegardé des parties, de pouvoir reprendre là où vous en étiez. Enfin, après avoir sélectionné ce que vous vouliez faire, vous arriverez sur la page de jeu. Si vous voulez accéder au ScoreBoard, ce sera à la fin de la partie.

## Règle du jeu

La règle du jeu est simple, il faut deviner le mot caché en soumettant des lettres tout en faisant le moins d'erreur sinon vous perdez. Pour soumettre une lettre ou un mot rien de plus simple entrez ce que vous voulez dans la barre et appuyez sur le bouton "send". Si la lettre est présente elle s'affiche et sinon vous avez une tentative en moins. Si vous avez soumis un mot et qu'il est bon, vous avez gagné la partie sinon vous avez deux tentatives en moins. La partie est perdue si vous avez fait de 10 erreurs. Si vous voulez voir l'avancement de votre pendu, vous pouvez cliquer sur le bouton en haut à droite en signe de corde. À tout moment, vous pouvez arrêter la partie en cliquant sur le bouton stop et ensuite, vous choisirez le nom de votre sauvegarde.

## Fonctionnalités

- Sélection du niveau au début de la partie
- Système de ScoreBoard avec classement des 5 premiers
- Système de sauvegarde des parties avec gestion des propriétaires des sauvegardes
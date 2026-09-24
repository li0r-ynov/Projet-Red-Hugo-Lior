# Tower of Duality

Projet Go reconstruit à partir des fichiers partagés dans la conversation. Les marqueurs de conflits Git ont été retirés.

## Lancer

Dans ce dossier : `go run .`

## Règles ajoutées

- Inventaire initial de 10 places. Chez le marchand, chaque amélioration coûte 30 pièces d'or et ajoute 10 places, au maximum trois fois.
- Victoires des étages 1 à 5 : 10, 20, 40, 80, puis 160 pièces d'or.
- Étage 3 : deux Plumes divines ; étage 4 : une Essence divine. Le Casque d'Arès demande une Essence divine et un Fer.
- Potion de Pâques : 10 pièces d'or, un seul achat par personnage. Après avoir atteint le palier Héros ou Démon, elle peut être offerte dans une seule Maison des Dieux, en échange de son arme divine.
- Les armes et les armures peuvent être équipées, remplacées ou retirées depuis l'inventaire.
- Première bataille à la moitié des PV maximum ; PV restaurés à la fin de chaque bataille. Une résurrection à la moitié des PV maximum par bataille.

## À savoir

- Cette version fonctionne en mémoire : quitter le programme ne sauvegarde pas encore la partie.
- Si l'inventaire est plein à la victoire des étages 3 ou 4, tout ou partie du butin ne peut pas être ramassé. Il faut libérer de la place avant ces combats.
- Le poison est conservé comme objet, mais son effet de combat reste à définir.

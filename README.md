# systeme-de-casier

<img src="https://media.tenor.com/images/44a4d937f5be76d516c7d909f25cd137/tenor.gif">

Le programme de casier interactif simple utilise la ligne de commande pour stocker les cartes d'identité des clients lors de la visite du bureau.

### comment utiliser

1. ouvrir l'invite de commande ou le terminal
2. ouvrez le programme à l'aide de la commande `./main`

### liste de commande

| Commande                                    | La description                                                                        |
| ------------------------------------------- | ------------------------------------------------------------------------------------- |
| init [nombre de casiers                     | faire le compte des casiers                                                           |
| status                                      | affiche l'état de chaque casier                                                       |
| input [type d'identité] [numéro d'identité] | insérer la carte d'identité                                                           |
| leave [numéro de casier]                    | vider le casier                                                                       |
| find [numéro d'identité]                    | affiche le numéro du casier en fonction du numéro d'identification                    |
| search [type d'identité]                    | affiche une liste de numéros d'identité en fonction de l'identité que vous recherchez |
| exit                                        | terminer le programme                                                                 |
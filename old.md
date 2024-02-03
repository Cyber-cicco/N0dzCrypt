## Comprendre l'architecture du backend

Le code java se trouve dans src/main/java/

### Les controlleurs

On peut trouver des classes suffixées de `Controller` dans  le package `pages`. C'est ici que
l'on va trouver les points d'entrée de  l'API.   

Chaque point d'entrée est identifié par une méthode annotée d'un verbe HTTP. Cette méthode va 
renvoyer une chaine de caractère représentant l'URL d'un template thymeleaf. Grâce à cette URL,
les méthodes vont construire du HTML à partir du template et le renvoyer à l'utilisateur quand
il fera des appels à l'API.

Chaque méthode renvoyant l'adresse d'un template prend en paramètre un objet de type `Model`.
Cet objet va permettre d'irriguer le template avec des variables issues de java.

### Les irrigateurs

L'irrigation ne se fait ni dans le controller, ni dans un service, mais dans un type de classe
dédié : un irrigateur. Ces classes sont toutes suffixées de `Irrigator`.

Une classe irrigator n'a qu'un seul but : faire des appels à des services et repositories pour
irriguer le model. Il doit contenir le moins de logique possible.

En effet, il est à prévoir pour un LMS que l'on ait besoin de créer un appli mobile, et donc 
des points d'entrée JSON. Il vaut mieux donc découpler au maximum le retrait, la modification
et la validation des données de l'API hypermédia, de façon à pouvoir réutiliser toute la logique
lorsque l'on aura besoin de créer une API JSON.

### Les services et repository

Il s'agit simplement des classes chargées de la manipulation et le retrait des données. Les 
repositories sont là pour donner des instructions directement à la base de données, tandis 
que les services sont là pour manipuler ces données dans le backend.

La philosophie des repositories est la suivante : plutôt que d'utiliser JPA, on utilise au maximum
des requêtes SQL natives. La raison est la suivante : il est possible que dans le futur, on 
ait besoin de migrer une partie du serveur en java vers un autre langage. 

Si par exemple on se retrouvait à n'avoir que des alternants faisant du python, il est possible
que l'on souhaite déplacer certaines fonctionalités nécessitant beaucoup de maintenance vers 
le framework django. Et pour cela, il peut être intéressant d'avoir déjà de disponible toute la
logique de retrait des données, et n'avoir qu'à intégrer le SQL à l'ORM

On essai donc également de **mettre la totalité de la logique de retrait dans la requête SQL**.
**On ne doit jamais récupérer plus que ce qui est nécessaire coté java pour ensuite utiliser
un filter ou un map**. On peut s'occuper de la récupération en fonction de certains fields et 
l'ordonancement directement dans la requête SQL. En plus de rendre l'app plus facile à migrer
vers un autre langage, cela est bénéfique pour les performances.

### Les entités

Classique, rien de particuliers à dire

### La sécurité et l'authentification

Le package security contient la configuration de sécurité. L'authentification se fait via un JWT.
Un service est à injecter dans les controlleurs pour sécuriser les différents points d'entrée de 
l'application : AuthenticationService. 

La méthode `getAuthInfos()` du AuthenticationService permet de récupérer un objet contenant les 
informations de l'utilisaeur connecté.

La méthode `mustBeOfRole(List<String> currentRoles, TypeRole expectedRole, HttpServletResponse response)`
permet de donner la liste de rôle de l'utilisateur connecté en premier argument, le rôle attendu
en deuxième argument, et un objet implémentant l'interface HttpServletResponse pour gérer
la réponse. La méthode va lever une exception si l'utilisateur n'est pas du rôle

La méthode `rolesMustMatchOne(List<String> currentRoles, List<TypeRole> acceptedRoles, HttpServletResponse response)`
se base sur le même principe sauf que l'on accepte plusieurs rôles possibles pour l'utilisateur.

## Comprendre l'architecture du frontend

On nomme frontend les éléments se trouvant dans `src/main/resources/templates` et
src/main/resources/static. 

### Les layout 

Les **layouts** sont des templates thymeleaf contenant une architecture qui sera commune à plusieurs 
**pages**. Il sont définis dans `src/main/resources/templates/layout`. Il contiennent un élément
HTML d'id `insert` dans lequel on insérera le contenu des pages. Un layout possède toutes les
métadonnées de la page ainsi que les liens vers les différentes feuilles de style et les 
différents scripts globaux à l'appli (notamment HTMX). L'idée serait de rendre l'appli le moins
possible dépendante de scripts globaux pour ne garder que HTMX et éventuellement les librairies,
mais ce n'est pas toujours facile. Le problème n'est pas encore réglé. Chaque layout se trouve
à la racine du dossier `layout`

### Les pages

Les **pages** sont des template thymeleaf dont l'URL sera poussée dans l'historique de navigation.
Cela peut correspondre à des pages complètes avec un tag body, un à des fragments qui seront 
insérés dans des layouts via HTMX. Par exemple, la page `login` est document HTML complet, 
alors que la page `home` est un document HTML devant être inséré dans un layout content 
la barrre de navigation pour être 
complète. On fait donc la différence entre **page complète** et **page fragmentée**.  

Une page fragmentée nécessite de créer deux points d'entrée dans le controlleur chargé de 
l'afficher : un renvoyant la page complète (donc le layout + la page fragmentée), et un
ne renvoyant que la page fragmentée.  

Prenons le cas d'une page fragmentée, `home`.

Mon layout contient la barre de navigation, que je n'ai pas besoin
de recharger quand je change de page. Quand je clique sur le lien de la barre de navigation
censé m'emmener vers la page d'accueil, plutôt que de faire un rechargement complet, je vais simplement
demander à HTMX de changer le HTML de la div d'id `insert` par le résultat de l'appel vers le
point d'entrée du back contenant la page fragmentée.  

Toutefois, si l'utilisateur recharge sa page, il faut qu'il puisse avoir la page complète et 
non la page fragmentée. Pour cela, je vais créer un point d'entrée `home/api` contenant ma
page fragmentée, et un point d'entrée `home` renvoyant ma page complète.

Dans le back, cela se traduit par ce code :

```java
@GetMapping({"api", "home/api"})
public String getHomePageApi( Model model){
        AuthenticationInfos userInfos = authenticationService.getAuthInfos();
        homePageIrrigator.irrigateModel(model, userInfos);
        return Routes.ADR_HOME;
        }
@GetMapping({"/", "", "home"})
public String getHomePage( Model model){
        AuthenticationInfos userInfos = authenticationService.getAuthInfos();
        layoutIrrigator.irrigateBaseLayout(model, userInfos, Routes.ADR_HOME);
        homePageIrrigator.irrigateModel(model, userInfos);
        return Routes.ADR_BASE_LAYOUT;
        }
```

D'un point de vue de l'organisation du code, on a pour chaque page un dossier nommé de la même
façon que la page, contenant un fichier html principal. Le dossier contient égalemment un dossier
nommé `fragments` contenant les **fragments** de la page

### Les fragments

Les fragments sont des templates thymeleaf présent dans une et une seule page, dont l'URL n'est pas
poussée dans l'historique lorsqu'ils sont affichés, mais qui peuvent être rechargés sans recharger
tout le contenu de la page

Par exemple, la page agenda contient un fragment `agenda.calendar`, correspondant au calendrier.
Le calendrier n'est jamais affiché seul (il est toujours accompagné d'une liste de cours à drag
& drop), et n'est pas affiché dans plusieurs pages. 

Quand je change de semaine, cela entraine un rechargement du calendrier, mais pas de toute ma page,
raison pour laquelle j'ai besoin d'en faire un template séparé du template principal.

Ainsi, le calendrier correspond à la définition d'un fragment.

D'un point de vue convention de nommage, on préfixera le fichier du nom du fichier auquel il
est rattaché suivi d'un point et du nom du fichier. Cela vaut également pour les fragments
dans des fragments. Par exemple, le calendrier de la page agenda est appelé `agenda.calendar`.

### Les composants

Un composant est un template thymeleaf ayant vocation à être utilisé dans plusieurs pages.
Par exemple, une modale peut être utilisée dans n'importe quelle page, de même pour 
un bouton, etc.

Thymeleaf rend assez facile le fait de réutiliser des composants. En effet il n'y
a pas vraiment de soucis à se faire en ce qui concerne la communication des informations
entre ceux-ci, vu qu'ils finissent de toute façon par être rendu en une seule page HTML.

D'une manière générale, il est conseillé de ne pas mettre trop de logique dans un composant 
pour éviter de créer un trop fort couplage entre des éléments qui n'ont pas lieu d'être couplé.

L'un des gros avantages d'une application basée sur HATEOAS est l'absence de couplage entre
les différentes UI de l'application. Il vaut parfois mieux faire des copiés collés un peu
sales que de surcharger un composant en logique pour éviter la duplication de code.

### Le style

On utilise tailwind pour gérer le style. D'une manière générale, ce framework est facile d'accès, 
intègre par défaut les bonnes pratiques du CSS, est extrêmement
flexible (bien plus que bootstrap vu qu'il n'y a que des classes utilitaires), et s'intègre bien
avec Thymeleaf. En effet, il est assez facile de créer des composants réutilisables, et l'utilisation
thymeleaf implique nécessairement de travailler sur des fichiers HTML très laids, il n'y a donc pas 
de soucis à la rendre encore plus moche avec des attributs `class` de quinze kilomètres et demi.

### L'orchestrateur de l'état de l'application : HTMX

HTMX ne remplis pas le rôle de "framework javascript" de l'application, dans le sens où il ne gère
(presque) pas d'état coté client et ne crée pas de HTML à partir d'un script JS. Il se contente de faire des
requêtes AJAX qui recupèrent du HTML, et échange le contenu d'un élément du DOM avec le HTML reçu.

Il fait cela grâce à des attributs HTML pouvant être ajoutés à n'importe quel élément du DOM, configurant
le déclencheur de la requête, sa cible, l'élément à échanger avec la réponse, etc.

Il y a également dans HTMX une gestion de l'historique de navigation et de l'URL, avec la possibilité de
mettre en cache les pages visitées et de déterminer à partir de certaines actions

### Petite introduction à HTMX

Prenons cet exemple de lien vers la page de cours de digi-learning :

```html
<h2 class="text-lg text-black">
    <a
        
        class="hover:cursor-pointer hover:text-primary"
        href="/cours"
        hx-get="/cours/api"
        hx-target="#insert"
        hx-swap="innerHTML"
        hx-push-url="cours"
    >
    Mes ressources e-learning
    </a>
<article class="text-sm text-black">

```
et mettons en parallèle avec le layout de base de l'application :

```html
<body class="flex w-full h-full overflow-hidden">
    <th:block th:replace="~{components/navbar}"></th:block>
    <main hx-history="false" id="router-outlet" class="w-full bg-secondary">
        <div id="insert" th:insert="${insert}" class="w-full h-screen"></div>
    </main>
    <th:block th:replace="pages/fragments/cours/dialogs/modal.cours"></th:block>
    <script src="/js/animation.post-forum.js"></script>
    <script src="/js/form-checking.js"></script>
    <script src="/js/animation.cours.js" type="module"></script>
    <script>
    </script>
</body>
</html>
```

Ici, nous pouvons voir ce que va faire HTMX lorsque l'on va cliquer sur le lien :
* Avec hx-get, il va faire une requête vers l'URL chemin `/cours/api` , qui va renvoyer
  du html.
* Avec hx-target, il vise l'élément du DOM possédant l'id `insert`
* Avec hx-swap, il précise que les éléments enfant de l'élément du DOM visé par hx-target devront être remplacés par le résultat de la requête AJAX. Cela veut dire que l'on va recharger la partie principale de l'application, mais que la navbar ne va elle pas être changée.
* Avec hx-push-url, on précise qu'il va falloir pousser l'URL `/cours` plutôt que `/cours/api` dans l'historique de navigation lorsque l'on aura cliqué sur le lien. Pourquoi ? Parce que si le chemin  `/cours/api` renvoie le HTML nécessaire pour mettre dans insert, il ne renvoie pas une page entière. Donc, si l'on réactualisait la page sur l'URL `/cours/api`, on se retrouverait avec une page partiel. L'URL `/cours` renvoie quant à elle une page entière, ce qui fait que l'on peut appuyer sur f5 et obtenir une page complète après avoir cliqué sur ce lien.

Et juste avec ces quatre attributs HTMX (plus hx-patch, put, post...), c'est 90% de ce que fait digilearning.

Voici d'ailleurs une liste exhaustive de tous les attributs HTMX utilisés dans le projet:

* hx-history
* hx-get
* hx-push-url
* hx-target
* hx-swap
* hx-select
* hx-select-oob
* hx-on::after-request
* hx-on::before-request
* hx-patch
* hx-delete
* hx-post
* hx-preserve

Coté back, il y a également quelques headers qu'il est possible de mettre sur les réponses pour changer le comportement de HTMX en fonction des événements coté back (par exemple rediriger la réponse vers un élément avec la classe erreur lorsque qu'un champ d'un formulaire est mal remplis)

Voilà donc le principal intérêt de HTMX : vous n'avez pas besoin d'écrire ce genre de chose :
```typescript
  put(candidatureDto: Partial<Candidature>) {
    return this.http.put<Candidature>(this.URL_CANDIDATURE, this.candidatureToDto(candidatureDto), {withCredentials: true})
      .pipe(
        map(candidature => this.dtoToCandidature(candidature) as Candidature)
      );
  }

```
pour faire un put vers votre API et récupérer un JSON que vous allez ensuite devoir manipuler avec typescript pour afficher ce que vous voulez dans le navigateur. Vous écrivez simplement :
```html
hx-put="/url"
```
et vous récupérez une réponse HTML générée par le back, et c'est suffisant.

Et vous esquivez également la séparation entre votre font et votre back, ce qui vous permet de ne travailler que sur une
seule application. Ce qui vous évite donc deux builds différents pour le front et le back, les soucis de compatibilité
de version entre front et back, le fait d'avoir à faire tourner deux serveurs en même temps pour tester votre
application, etc. La liste est encore longue, mais le simple gain de vélocité par rapport à framework javascript
classique serait en soit suffisant pour justifier l'utilisation de HTMX.

Mais également, cela permet de garder le serveur comme seule source de l'état des données. Ainsi, si votre page affiche
des données d'un utilisateur, ces données ont nécessairement été construites par le serveur, et donc reflètent
nécessairement l'état des données dans la base.

L'idée est de toujours récupérer et envoyer les données dépendantes du serveur via HTMX, et de ne jamais les altérer
avec autre chose que les appels AJAX effectués via HTMX.

## Interopérabilité entre HTMX et Thymeleaf

Thymeleaf est un moteur de templating java qui permet d'insérer des données issues du java dans n'importe quel élément
html, à n'importe quel endroit. Ce qui veut dire que par exemple, il est possible de définir l'URL vers laquelle un
hx-get va pointer en fonction de variables que l'on aura définies coté back. Par exemple, ce morceau de code :

```html
<a
    class="text-primary font-poppins hover:underline hover:cursor-pointer"
    th:text="${c.getTitre()}"
    th:hx-get="'/cours/sommaire/api?id=' + ${c.getId()}"
    th:hx-push-url="'/cours/sommaire?id=' + ${c.getId()}"
    hx-target="#liste-cours"
    hx-swap="outerHTML"
></a>
```

permet de définir la route vers laquelle HTMX va faire un appel en fonction de l'id de la variable c (ici, diminutif
de cours).

Ici, voici un morceau de html représentant une icône favori pour un cours donné d'une liste de cours :

```html
<div
    class="rounded-full hover:cursor-pointer hover:bg-lightAccent p-1"
    th:hx-patch="'/cours/bookmark?id=' + ${c.getId()}"
    th:hx-target="'#bookmarked-' + ${c.getId()}"
    hx-swap="outerHTML"
>
    <img
        class="min-w-[30px] min-h-[30px]"
        th:id="'bookmarked-' + ${c.getId()}"
        width="30"
        src="/img/icons/bookmark.svg"
        th:style="${c.getBoomarked()} ? 'filter: invert(73%) sepia(83%) saturate(1960%) hue-rotate(3deg) brightness(99%) contrast(108%);' : _"
    >
</div>
```
Ainsi, en utilisant des variables thymeleaf pour générer les attributs HTMX, leur faire faire des appels au back en
incluant des données permettant d'identifier les données sur lequels on souhaite interagir via l'appel HTTP.

Dans les exemples que nous avons montré, c représente en fait un élément d'une liste de cours. Thymeleaf permet en effet
de faire de l'itération sur les listes définies dans le back, comme ceci :


```html
<div class="bg-white p-5 shadow-dark shadow-md rounded-md" th:each="c, iterStat : ${cours}">
```

`cours` étant défini dans cette méthode d'un controller coté back :

```java
    @GetMapping("/liste")
    public String getListeCours(@CookieValue("AUTH-TOKEN") String token, @RequestParam("id") Long idSModule, @RequestParam("module") Long idModule, Model model){
        AuthenticationInfos userInfos = authenticationService.getAuthInfos(token);
        irrigateListeCours(userInfos, idSModule, idModule, model);
        model.addAttribute("insert", "pages/liste-cours");
        model.addAttribute("links", navBarService.getLinks(userInfos));
        return "base";
    }

    private void irrigateListeCours(AuthenticationInfos userInfos, Long idSModule, Long idModule, Model model) {
        model.addAttribute("cours", coursService.getCours(userInfos, idSModule));
        SousModule sousModule = sousModuleRepository.findById(idSModule).orElseThrow(EntityNotFoundException::new);
        model.addAttribute("smodule", sousModule);
        model.addAttribute("idModuleOrigine", idModule);
        model.addAttribute("bookmarked", coursRepository.getBookMarked(userInfos.getId()));
    }
```

L'idée derrière l'interaction entre Thymeleaf et HTMX est la suivante :

* On définit coté back des éléments de la base auquel l'utilisateur a accès
* On injecte des identifiants de ces éléments dans des attributs HTMX pour créer des liens vers d'autres pages
* Le back récupère ces éléments dans la requête AJAX lancée par HTMX pour charger des éléments de la base en fonction des informations passées dans la requête (en vérifiant bien que l'utilisateur a le droit d'avoir accès à ces éléments), puis les renvoie à l'utilisateur
* Et ainsi de suite...

## Le scripting coté client

Si HTMX a un rôle très important dans l'application, cela ne veut pas dire que les interactions
permises par le HTML et HTMX sont suffisantes pour faire toutes nos UI. En effet, si l'on veut faire des tableaux
kanban, si l'on veut créer des animations, si l'on veut un éditeur de texte riche, il sera nécessaire 
d'utiliser au moins un peu de javascript.

Mais là se trouve aussi l'intérêt de HTMX : maintenant qu'il n'y a plus d'état partagé dans toute l'application
coté client, on peut utiliser n'importe quel petit framework js pour créer des petits ilots d'interactivité.

On peut même n'utiliser aucun framework, et c'est à l'heure où j'écris le choix sur lequel je me suis porté
pour chacune des UI, même l'agenda et l'éditeur de texte. L'intérêt de n'utiliser aucun framework tiens au 
fait que l'on est sur que javascript ne sera jamais déprécié, contrairement à quelque chose comme Angular JS
ou les vieilles versions de React. En ne choisissant aucun framework, cela donne également d'office accès à
toutes les librairies javascript du monde, sans avoir à se demander s'il existe un adapteur de cette librairie
pour son propre framework.

Actuellement, l'idée selon laquelle les pages sont des petits ilots d'interactivité détachés du reste de l'application
n'est pas totalement vraie. Il y a certaines pages pour lesquelles j'ai du inclure les scripts dans le layout pour les faire
fonctionner. L'idée reste avant tout de faire fonctionner l'application. Mais actuellement, si on change le layout
d'une page fragmentée alors qu'elle dépend d'un script dans le layout, elle ne fonctionnera plus. J'aimerai n'avoir
que des librairies dans le layout.

L'une des solutions sur lesquelles je me penche actuellement est de créer un dossier `script` dans le
dossier des templates, pour définir des fichiers html contenant une simple balise script avec du javascript
dedans. En l'inculant dans un template thymeleaf, cela permettrait chaque fois que l'on recharge la page
d'avoir du code javascript déjà inclue dans la page, ce qui évite les problèmes de délai de chargement du script, 
de redéfinition de variables déjà définies si l'on revient sur une page utilisant un script déjà utilisé auparavant,
etc.

## Conclusion

Il s'agit là de lignes directrices pour la contribution au projet, je ne pense pas avoir trouvé "la bonne"
solution pour faire des applis avec HTMX, c'est juste l'architecture m'étant apparue comme cohérente au fur
et à mesure de la création de l'appli. C'est donc susceptible de changer.

Et d'une manière générale, étant donné que le stack est totalement expérimental, on doit un peu tout tester, 
et juste faire en sorte qu'une expérience ratée ne casse pas toute l'application.



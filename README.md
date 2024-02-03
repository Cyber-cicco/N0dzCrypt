# N0dzCrypt

**The Javascript / Java Fullstack Scripting framework.**

## What is N0dzCrypt ?

N0dzCrypt CLI tool to enhance Fullstack developpement of Java based webapps.

It makes use of particular set of old technologies and enhances them to allow you to create new WebApps 
based on the *REST* and *HATEOAS* principles while allowing you to create an UX similar to that of an SPA.

It also gives you ways to improve old Java webapps by offering a way to transform them into SPA-like app.

## What technologies  does it use ?

 - **Any SGBDR** as the database.
 - **Spring Boot** as the backend.
 - **Maven** as the java dependency manager.
 - **Thymeleaf** as the templating language.
 - **Tailwind** as the CSS framework.
 - **Vanilla Javascript** as the driver of client interactivity.
 - **HTMX** as the engine of application state.

 ## What can it be used for ?

  1) **Enhance**
    There is a lot of old MPA webapps still using Thymeleaf that could get some help getting into the modern
    age. Most of these apps look and feel old, and are no match for modern webapps. Most of them are considered
    legacy apps, and are replaced by a frontend using a classical Javascript framework. But if instead of rewriting the
    whole app, you could just enhance the old code to make it similar to every other SPAs ? That's the promise of the
    N0dzCrypt tool.

  2) **Create**
    BootStrapping a webapp usually involves creating two separate apps, one for the front, and another for the back.
    That means, even if the app ain't that large, you'll have to get a frontend and a backend team with quite a lot 
    of experience in ops, because the deployment and testing of two separate apps with multiple build steps can get quite
    complex (especially if your frontend uses things like typescript and JSX).  
    Also, if you want some of your frontend devs to be able to add or improve some backend code, you'll get locked into 
    a Javascript backend, which can be quite cumbersome, as it usually involves the slowness of node, mess of dependencies 
    of npm, and the weird quirckiness of Javascript.
    With N0dzCrypt, you can get rid of all of that, while maintaining the best of what the Javascript ecosystem can offer you.

## What is (or will be) included ?

  - create-nodzcrypt-app : a command line tool allowing you to bootstrap every necessary component of a N0dzCrypt app project, including
  the backend security dependencies and configuration, the configuration of the database connection, the installation and configuration of tailwind,
  and the layout of the component model.

  - a fullstack page generator : similarely to Angular, N0dzCrypt has a strong opinion on the way an app should be organized (while it doesn't enforce you to respect it). That way, same way as Angular, 
  it can allow you to create all the boilerplate needed to introduce a new page, a new component or new fragment in your webapp. The fun fact is it also lets you create the boilerplate of the server side code.

  - a component package manager, `nsc`. Similarely to what ShadCn/ui does for react, there is a way to import components using Tailwind as the CSS framework directly into our code. Except this one is Open Source,
  works more like npm, where anyone can share their own component on a central repository, and every component defines interfaces and endpoints in the backend if any is needed.
  (this will be hard asf to implement but I have no life).

  - an entity creation helper for the backend : it reuses some functionalities of an old tool I made, Spring-go, to let you create 
  jpa entities and repositories through the commande line.

## Organisation of a N0dzCrypt app :

### Understanding backend architecture

The backend uses the Spring boot framework convention for writing code. You can find the backend java code in 
`/src/main/java/${your_domain_name}`

### REST endpoints

REST endpoints can be created through the `nsg create -ctrl ${my_name}` or `nsg create -p ${my_name}`
command in the terminal. The first will only create a controller, the second will create an html page with a controller endpoint,
alongside an irrigator and a service. We will talk later about irrigators and services; for now, let's focus on the controllers.

By default, the route of the endpoint will be named after the name of the class. You can change that with the `-e` option.

In this java class, you can define methods to handle http verbs sent to this adress through the use of `@GetMapping`  `@PostMapping`, etc.

Every one of those methods need to return a String that will be 






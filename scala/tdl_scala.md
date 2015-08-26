# Test Driven Learning - Scala

## Environnement

La première étape consiste à installer son environnement de travail. Pour cela nous allons utiliser **sbt** (Simple Build Tool).

Cet outil est comparable à Maven avec une approche différente sur certains points. Nous n'allons pas rentrer dans ces détails pour le moment. Il nous permettra de nous concentrer sur l'apprentissage du language Scala.

### Installation

Pour installer sbt, je vous propose de suivre la documentation officielle (qui couvre l'ensemble des OS) : [Documentation officielle](http://www.scala-sbt.org/0.13/tutorial/Setup.html)

>En passant par sbt, il n'est pas nécessaire d'installer explicitement un SDK Scala, celui-ci est fourni par sbt

Pour valider l'installation de sbt, taper dans la console les commandes suivantes : 

	sbt console
	[info] Set current project to scala (in build file:***********)
	[info] Starting scala interpreter...
	[info]
	Welcome to Scala version 2.10.4 (Java HotSpot(TM) 64-Bit Server VM, Java 1.8.0_05).
	Type in expressions to have them evaluated.
	Type :help for more information.

	scala>

A l'invit de commande **scala>** taper le code suivant : 

	scala> println("Hello World")
	Hello World

Si vous obtenez le résultat ci-dessus alors c'est que tout est en ordre et que l'on peut passer à la suite.

Nous avons ici utilisé le REPL disponible avec le language Scala (*sbt console*). Le REPL permet de pouvoir exécuter à la volée du code. 

Par exemple, à l'aide du REPL on peut tester la création d'une fonction **add** qui a pour objectif d'ajouter 1 à l'entier passé en paramètre. Pour celà, on exécute successivement les commandes suivantes : 

Lancement du REPL

	sbt console

Création d'une variable i de type Int initialisée avec la valeur 2
	

	scala> val i:Int = 2
	i: Int = 2
	
Création de la fonction add
	

	scala> def add:Int=>Int = a=>a+1
	add: Int => Int
	
Utilisation de la fonction add

	scala> add(i)
	res1: Int = 3


### Création du projet : tdl-scala

Sbt suit les mêmes convention que Maven en terme de structure de projet. Nous allons donc créer la structure minimale de répertoire pour nous permettre l'écriture d'un test.

Première étape création du répertoire racine du projet : **tdl-scala**
Ensuite, il faut créer la structure (minimale) suivante :

	tdl-scala
		src/
  			main/
			    scala/
    	   			<main Scala sources>
			test/
    			scala/
					<test Scala sources>
    

Maintenant, il faut fournir le descripteur de projet à sbt. Ce fichier se nomme build.sbt et se trouve à la racine du projet : 

	tdl-scala
		src/
  			main/
			    scala/
    	   			<main Scala sources>
			test/
    			scala/
					<test Scala sources>
		build.sbt
	

Voilà notre environnement est prêt pour commencer à travailler, enfin presque. Il faut maintenant compléter le fichier build.sbt avec quelques informations.

### build.sbt

Les informations à renseigner sont les suivantes : 

	name := "tdl-scala"

	version := "1.0"

	scalaVersion := "2.10.2"

	libraryDependencies += "org.scalatest" % "scalatest_2.10" % "2.0"
	
On retrouve comme informations : 

 - le nom du projet
 - la version du projet
 - la version de scala à utiliser
 - les dépendances du projet.
 
La dépendance ajoutée au projet est la librairie standard permettant d'écrire des tests en Scala. Elle nous sera utile pour l'écriture du premier test.

## Première application : Hello World

Traditionnellement, nous allons commencer par écrire notre première application qui aura pour objectif d'afficher le message :

	Hello world
	
Pour cela, on va créer la classe main. Créez le fichier : 

	 src/main/scala/Main.scala

et ajouter le code suivant au fichier : 

```scala
	object Main extends App{
		println("hello world")
	}
```

Maintenant, vous pouvez lancer la commande 
	
	>sbt run
	
Vous devriez obtenir l'affichage du message (avec des informations supplémentaire) :
	
	Hello world

Vous venez d'écrire votre premier programme en Scala.

Votre environnement est prêt. Prochaine étape écrire un test 

## Premier test

Aujourd'hui, savoir coder c'est savoir tester son code.  Nous allons voir comment écrire un test en Scala. Pour cela nous allons utiliser la librairie **ScalaTest**. Cette librairie fournit une API très complète permettant d'écrire des tests selon une approche **Behavior-Driven Development** (BDD)
Dans ce premier exemple, nous allons utiliser *FunSuite*. Une version très proche d'un test Junit. Nous verrons plus tard l'ensemble des possibilités de la librairie sur une approche BDD.

Créez le fichier suivant : 
	
	src/test/scala/MyFirstTest.scala	

et compléter le avec le code suivant :


```scala
	import org.scalatest.{ShouldMatchers, FunSuite}

	class MyFirstTest extends FunSuite with ShouldMatchers {
		test("1 should not equal to 2") {
			1 should not equal(2)
		}
	}
```

Une fois le fichier sauvegardé, lancre la commande : 

	>sbt test

Vous devriez obtenir le message suivant : 

	[info] MyFirstTest:
	[info] - 1 should not equal to 2
	[info] Run completed in 342 milliseconds.
	[info] Total number of tests run: 1
	[info] Suites: completed 1, aborted 0
	[info] Tests: succeeded 1, failed 0, canceled 0, ignored 0, pending 0
	[info] All tests passed.
	
## Conclusion

Suite aux différentes étapes ci-dessus, nous avons à disposition un environnement opérationnel. Il nous permet d'écrire du code Scala, de l'exécuter et enfin de le tester.
Notons que nous avons utiliser également le REPL qui nous permet de pouvoir rapidement explorer avec le language






	

	




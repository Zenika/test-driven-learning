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

À l'invite de commande **scala>** taper le code suivant : 

	scala> 1+1
	res0: Int = 2

Si vous obtenez le résultat ci-dessus alors c'est que tout est en ordre et que l'on peut passer à la suite.

### Création du projet : tdl-scala

Sbt suit les mêmes conventions que Maven en terme de structure de projet. Nous allons donc créer la structure minimale de répertoire pour nous permettre l'écriture d'un test.

Première étape création du répertoire racine du projet : **tdl-scala**
Ensuite, il faut créer la structure (minimale) suivante :

```
tdl-scala/
└── src/
    ├── main/
    │   └── scala/
    │       └── <main Scala sources>
    └── test/
        └── scala/
            └── <test Scala sources>
```
Maintenant, il faut fournir un descripteur de projet à sbt. Ce fichier se nomme build.sbt et se trouve à la racine du projet : 

```
tdl-scala/
├── src/
│   ├── main/
│   │   └── scala/
│   │       └── <main Scala sources>
│   └── test/
│       └── scala/
│           └── <test Scala sources>
└── build.sbt
```		
	

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

## Premier test

Aujourd'hui, savoir coder c'est savoir tester son code.  Nous allons voir comment écrire un test en Scala. Pour cela nous allons utiliser la librairie **ScalaTest**. Cette librairie fournit une API très complète permettant d'écrire des tests selon une approche **Behavior-Driven Development** (BDD)
Dans ce premier exemple, nous allons utiliser *FunSuite*. Une version très proche d'un test Junit. Nous verrons plus tard l'ensemble des possibilités de la librairie sur une approche BDD.

Créer le fichier suivant : 
	
	src/test/scala/MyFirstTest.scala	

et compléter le avec le code suivant :


```scala
	import org.scalatest.{ShouldMatchers, FunSuite}

	class MyFirstTest extends FunSuite with ShouldMatchers {
		test("1 should equal to 2") {
			1 should equal(2)
		}
	}
```

Une fois le fichier sauvegardé, lancer la commande : 

	>sbt test

Vous devriez obtenir le message suivant : 

```
info] MyFirstTest:
[info] - 1 should  equal to 2 *** FAILED ***
[info]   1 did not equal 2 (MyFirstTest.scala:5)
[info] Run completed in 486 milliseconds.
[info] Total number of tests run: 1
[info] Suites: completed 1, aborted 0
[info] Tests: succeeded 0, failed 1, canceled 0, ignored 0, pending 0
[info] *** 1 TEST FAILED ***	
```


## Passer en mode Continuous testing

Avec sbt, il est possible de lancer l'exécution des tests en continu. C'est à dire que les tests vont se lancer automatiquement lors de la modification du code. C'est très pratique, cela offre une boucle de retour instantanée.

Pour lancer sbt dans ce mode là, il faut exécuter les deux commandes suivantes : 

	>sbt


On se retrouve dans le mode interactif de l'outil, taper la commande suivante : 


	>~test
	
Après l'exécution de cette commande, à chaque modification du code, les tests seront exécutés.

Nous allons donc maintenant modifier le code du test précédent : 

```scala
...
	test("2 should  equal to 2") {
          2 should  equal(2)
    }
...
```

et on devrait voir dans la console le test se relancer automatiquement et se terminer avec succès.

## Conclusion

Suite aux différentes étapes ci-dessus, nous avons à disposition un environnement opérationnel. Il nous permet d'écrire du code Scala, des tests et de les exécuter. 
La prochaine étape consistera à travailler sur un kata ce qui nous permettra d'enrichir nos connaissances sur ce language








	

	




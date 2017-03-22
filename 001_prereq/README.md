# VARIABLE DECLARATION
  - [short declaration operator](https://play.golang.org/p/j31K35scXO)
  - you can only use short declaration inside block scope
  - [scope: block, package](https://www.golang-book.com/books/web/01-02)
  - [multi-variable short declaration](https://play.golang.org/p/UnOcuIpJpn)
  - [var initializes to zero value](https://play.golang.org/p/Ewm2bVrNyr)
  
  *** 
  
# DATA STRUCTURES
  - slice
    - [shorthand declaration operator](https://play.golang.org/p/sjpM9uy3F5) & [composite literal](https://golang.org/ref/spec#Composite_literals)
    - [make](https://play.golang.org/p/uGGKaY7l7f)
      - append item
      - [append slice](https://play.golang.org/p/Z6odm9CrtT)
    - [accessing by index](https://play.golang.org/p/2lLQa4mu6M)
    - [for range](https://play.golang.org/p/dnJ0DpFgtw)
    - [slicing](https://play.golang.org/p/TdjtVDXCUb)
  - map
    - [shorthand declaration operator & composite literal](https://play.golang.org/p/2pLxzsM3e7)
    - [make](https://play.golang.org/p/22SaWl9KAJ)
    - [for range](https://play.golang.org/p/zELqotxH0C)
  - struct
    - composite data structure
    - [create programs through composition](https://en.wikipedia.org/wiki/Composition_over_inheritance)
    - [create your own named type with an underlying type & use a composite literal](https://play.golang.org/p/cZLX83I2e2)
    - [create your own anonymous type with an underlying type & use a composite literal](https://play.golang.org/p/PRC-qyZldf)
    - [embed types](https://play.golang.org/p/gjDVdOQkfe)
    - [methods](https://play.golang.org/p/kLz9RN_K8u)
    - [interface](https://play.golang.org/p/4z_wL36wxe)
    
    *** 
    
# HANDS-ON
## HANDS ON 1
    - create a type square
    - create a type circle
    - attach a method to each that calculates area and returns it
    - create a type shape which defines an interface as anything which has the area method
    - create a func info which takes type shape and then prints the area
    - create a value of type square
    - create a value of type circle
    - use func info to print the area of square
    - use func info to print the area of circle
   
   [SOLUTION](https://play.golang.org/p/1enChb7Kg5) 
    
## HANDS ON 2
    - create a struct that holds person fields
    - create a struct that holds secret agent fields and embeds person type
    - attach a method to person: pSpeak
    - attach a method to secret agent: saSpeak
    - create a variable of type person
    - create a variable of type secret agent
    - print a field from person
    - run pSpeak attached to the variable of type person
    - print a field from secret agent
    - run saSpeak attached to the variable of type secret agent
    - run pSpeak attached to the variable of type secret agent
  
  [SOLUTION](https://play.golang.org/p/RxrkCJw9Cd) 
    
## HANDS ON 3
    - create an interface type that both person and secretAgent implement
    - declare a func with a parameter of the interface’s type
    - call that func in main and pass in a value of type person
    - call that func in main and pass in a value of type secretAgent
    
  [SOLUTION](https://play.golang.org/p/dAP73m_elq)
  
  Thank you to @BenCarter78 for that cool solution.
   
## ADDITIONAL INFO
    - solution & optional additional info not necessary to know: assertions
    
  [SOLUTION](https://play.golang.org/p/0TX4o-u-_B)
    
## ADDITIONAL EXERCISES
  [Complete these exercises](https://docs.google.com/document/d/1AqD-5yfAw8P1aUwH6-07UTHc0FSSAnW9b44sXJEVoag/edit?usp=sharing)
  
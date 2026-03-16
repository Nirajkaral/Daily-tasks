# PARENT CLASS
class Animal:
    def __init__(self, name):
        self.name = name          

    def speak(self):              
        print(f"{self.name} makes a sound")

    def eat(self):                
        print(f"{self.name} is eating")

    def sleep(self):              
        print(f"{self.name} is sleeping")


# CHILD CLASS — Dog inherits from Animal
class Dog(Animal):
    def speak(self):              
        return "Woof"

    def fetch(self):              
        print(f"{self.name} is fetching the ball!")


# CHILD CLASS — Cat inherits from Animal
class Cat(Animal):
    def speak(self):              
        return "Meow"

    def purr(self):               
        print(f"{self.name} is purring... ")



if __name__ == "__main__":
    dog = Dog("Buddy")
    cat = Cat("Whiskers")

    print(dog.speak())    
    print(cat.speak())     
    dog.eat()             
    cat.sleep()            
    dog.fetch()            
    cat.purr()             
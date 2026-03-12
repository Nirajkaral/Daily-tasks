class Dog(Animal):
    def speak(self):
        return "Woof! "

class Cat(Animal):
    def speak(self):
        return "Meow! "

class Cow(Animal):
    def speak(self):
        return "Moo! "

class Snake(Animal):
    pass                  


animals = [Dog("Buddy"), Cat("Whiskers"), Cow("Bessie"), Snake("Sly")]

for animal in animals:
    print(f"{animal.name} says: {animal.speak()}")


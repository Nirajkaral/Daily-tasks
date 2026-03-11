class Dog:
    def speak(self):
        return "Woof!"

class Cat:
    def speak(self):
        return "Meow!"

class Duck:
    def speak(self):
        return "Quack!"

# All different objects — same method name speak()
animals = [Dog(), Cat(), Duck()]

for animal in animals:
    print(animal.speak())
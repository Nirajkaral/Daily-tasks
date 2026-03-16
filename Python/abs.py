from abc import ABC, abstractmethod


class Shape(ABC):

    
    def area(self):               
        pass

    @abstractmethod
    def perimeter(self):         
        pass

    def describe(self):           
        print(f"Area: {self.area():.2f} | Perimeter: {self.perimeter():.2f}")


class Circle(Shape):
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return 3.14 * self.radius ** 2

    def perimeter(self):
        return 2 * 3.14 * self.radius


class Rectangle(Shape):
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def area(self):
        return self.width * self.height

    def perimeter(self):
        return 2 * (self.width + self.height)


class Triangle(Shape):
    def __init__(self, a, b, c, base, height):
        self.a = a
        self.b = b
        self.c = c
        self.base = base
        self.height = height

    def area(self):
        return 0.5 * self.base * self.height

    def perimeter(self):
        return self.a + self.b + self.c


# ── Using Shapes ──
# shape = Shape()   

c = Circle(5)
r = Rectangle(4, 6)
t = Triangle(3, 4, 5, 4, 3)

c.describe()    
r.describe()    
t.describe()    
class Product:
    def __init__(self, name, price, stock):
        self.name = name
        self.price = price
        self.__stock = stock

    def buy(self, qty):
        if qty > self.__stock:
            print(f" Only {self.__stock} left in stock!")
        else:
            self.__stock -= qty
            print(f" Bought {qty} x {self.name}. Remaining stock: {self.__stock}")

class DigitalProduct(Product):       
    def buy(self, qty):              
        print(f" Downloaded {self.name} — no stock limit!")

p1 = Product("Laptop", 80000, 5)
p2 = DigitalProduct("eBook", 299, 999)

p1.buy(2)    
p1.buy(10)   
p2.buy(100)  
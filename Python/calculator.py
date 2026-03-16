class Calculator:
    def __init__(self):
        self.history = []        

    def add(self, a, b):
        result = a + b
        self.history.append(f"{a} + {b} = {result}")
        return result

    def subtract(self, a, b):
        result = a - b
        self.history.append(f"{a} - {b} = {result}")
        return result

    def multiply(self, a, b):
        result = a * b
        self.history.append(f"{a} * {b} = {result}")
        return result

    def divide(self, a, b):
        if b == 0:
            print(" Cannot divide by zero!")
            return None
        result = a / b
        self.history.append(f"{a} / {b} = {result}")
        return result

    def show_history(self):
        print("\n Calculation History:")
        for h in self.history:
            print(f"   {h}")

if __name__ == "__main__":
    calc = Calculator()
    print(calc.add(10, 5))         
    print(calc.subtract(10, 3))    
    print(calc.multiply(4, 6))     
    print(calc.divide(20, 4))      
    print(calc.divide(10, 0))      
    calc.show_history()
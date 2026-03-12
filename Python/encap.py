class BankAccount:
    def __init__(self, owner, balance):
        self.owner = owner           
        self.__balance = balance     
        self.__transactions = []     

    
    def get_balance(self):
        return self.__balance

    
    def set_balance(self, amount):
        if amount >= 0:               
            self.__balance = amount
        else:
            print("Balance cannot be negative!")

    def deposit(self, amount):
        if amount > 0:
            self.__balance += amount
            self.__transactions.append(f"Deposited: +{amount}")
            print(f" Deposited {amount}. Balance: {self.__balance}")
        else:
            print(" Deposit amount must be positive!")

    def withdraw(self, amount):
        if amount > self.__balance:
            print("Insufficient funds!")
        elif amount <= 0:
            print(" Withdrawal amount must be positive!")
        else:
            self.__balance -= amount
            self.__transactions.append(f"Withdrew: -{amount}")
            print(f" Withdrew {amount}. Balance: {self.__balance}")

    def get_transactions(self):
        print("\n Transaction History:")
        for t in self.__transactions:
            print(f"   {t}")


acc = BankAccount("Ravi", 5000)
print(acc.owner)          
# print(acc.__balance)    
print(acc.get_balance())   
acc.deposit(2000)
acc.withdraw(1000)
acc.withdraw(9000)         
acc.get_transactions()
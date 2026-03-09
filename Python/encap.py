class Bank:
    def __init__(self, balance):
        self.__balance = balance   # private

    def deposit(self, amount):
        self.__balance += amount

    def show_balance(self):
        print(self.__balance)
        
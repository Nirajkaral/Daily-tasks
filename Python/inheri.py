class Employee:                        
    def __init__(self, name, salary):
        self.name = name
        self.salary = salary

    def work(self):
        print(f"{self.name} is working.")

    def display(self):
        print(f"Name: {self.name}, Salary: {self.salary}")

class Manager(Employee):               
    def __init__(self, name, salary, team_size):
        super().__init__(name, salary) 
        self.team_size = team_size

    def manage(self):
        print(f"{self.name} manages {self.team_size} people.")

class Intern(Employee):                
    def __init__(self, name, salary, duration):
        super().__init__(name, salary)
        self.duration = duration

    def learn(self):
        print(f"{self.name} is learning for {self.duration} months.")

m = Manager("Bob", 80000, 10)
i = Intern("Sara", 20000, 6)

m.work()      
m.manage()    
i.work()      
i.learn()     
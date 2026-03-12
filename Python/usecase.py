from abc import ABC, abstractmethod

class Employee(ABC):
    def __init__(self, name, emp_id, base_salary):
        self.name = name
        self.emp_id = emp_id
        self.__base_salary = base_salary    # Encapsulation

    def get_base_salary(self):
        return self.__base_salary

    @abstractmethod
    def calculate_salary(self):
        pass

    @abstractmethod
    def role(self):
        pass

    def display_payslip(self):
        print(f"""
💼 Payslip:
   ID           : {self.emp_id}
   Name         : {self.name}
   Role         : {self.role()}
   Base Salary  : ₹{self.__base_salary}
   Total Salary : ₹{self.calculate_salary()}
        """)


# INHERITANCE
class FullTimeEmployee(Employee):
    def __init__(self, name, emp_id, base_salary, bonus):
        super().__init__(name, emp_id, base_salary)
        self.__bonus = bonus

    def role(self):
        return "Full Time Employee"

    def calculate_salary(self):
        return self.get_base_salary() + self.__bonus


class PartTimeEmployee(Employee):
    def __init__(self, name, emp_id, hourly_rate, hours_worked):
        super().__init__(name, emp_id, hourly_rate)
        self.__hours_worked = hours_worked

    def role(self):
        return "Part Time Employee"

    def calculate_salary(self):
        return self.get_base_salary() * self.__hours_worked


class Contractor(Employee):
    def __init__(self, name, emp_id, daily_rate, days_worked):
        super().__init__(name, emp_id, daily_rate)
        self.__days_worked = days_worked

    def role(self):
        return "Contractor"

    def calculate_salary(self):
        return self.get_base_salary() * self.__days_worked


# POLYMORPHISM
employees = [
    FullTimeEmployee("Ravi", "E001", 50000, 10000),
    PartTimeEmployee("Sara", "E002", 500, 80),
    Contractor("John", "E003", 3000, 20),
]

for emp in employees:
    emp.display_payslip()    # Same method — different salary calculation
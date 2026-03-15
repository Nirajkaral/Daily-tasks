from abc import ABC, abstractmethod


class LibraryItem(ABC):
    def __init__(self, title, author, item_id):
        self.title = title
        self.author = author
        self.item_id = item_id
        self.__available = True        

    def is_available(self):
        return self.__available

    def borrow(self):
        if not self.__available:
            print(f" '{self.title}' is not available!")
            return False
        self.__available = False
        print(f" '{self.title}' borrowed successfully!")
        return True

    def return_item(self):
        self.__available = True
        print(f" '{self.title}' returned successfully!")

  
    def item_type(self):
        pass

    
    def display_info(self):
        pass


class Book(LibraryItem):
    def __init__(self, title, author, item_id, pages):
        super().__init__(title, author, item_id)
        self.pages = pages

    def item_type(self):
        return " Book"

    def display_info(self):
        status = "Available" if self.is_available() else " Borrowed"
        print(f"""
{self.item_type()}
   ID      : {self.item_id}
   Title   : {self.title}
   Author  : {self.author}
   Pages   : {self.pages}
   Status  : {status}
        """)


class Magazine(LibraryItem):
    def __init__(self, title, author, item_id, issue):
        super().__init__(title, author, item_id)
        self.issue = issue

    def item_type(self):
        return " Magazine"

    def display_info(self):
        status = " Available" if self.is_available() else " Borrowed"
        print(f"""
{self.item_type()}
   ID      : {self.item_id}
   Title   : {self.title}
   Author  : {self.author}
   Issue   : {self.issue}
   Status  : {status}
        """)


class DVD(LibraryItem):
    def __init__(self, title, author, item_id, duration):
        super().__init__(title, author, item_id)
        self.duration = duration

    def item_type(self):
        return " DVD"

    def display_info(self):
        status = " Available" if self.is_available() else " Borrowed"
        print(f"""
{self.item_type()}
   ID       : {self.item_id}
   Title    : {self.title}
   Director : {self.author}
   Duration : {self.duration} mins
   Status   : {status}
        """)



class Library:
    def __init__(self, name):
        self.name = name
        self.__items = []               

    def add_item(self, item):
        self.__items.append(item)
        print(f" Added '{item.title}' to library")

    def show_all(self):
        print(f"\n {self.name} — All Items:")
        print("=" * 40)
        for item in self.__items:
            item.display_info()         

    def search(self, title):
        print(f"\n Searching for: {title}")
        for item in self.__items:
            if title.lower() in item.title.lower():
                item.display_info()
                return item
        print(f" '{title}' not found!")
        return None

    def borrow_item(self, title):
        item = self.search(title)
        if item:
            item.borrow()

    def return_item(self, title):
        item = self.search(title)
        if item:
            item.return_item()

    def show_available(self):
        print(f"\n Available Items in {self.name}:")
        for item in self.__items:
            if item.is_available():
                print(f"   {item.item_type()} — {item.title}")


lib = Library("City Central Library")


lib.add_item(Book("Python Crash Course", "Eric Matthes", "B001", 560))
lib.add_item(Book("Clean Code", "Robert Martin", "B002", 431))
lib.add_item(Magazine("Tech Today", "John Doe", "M001", "March 2026"))
lib.add_item(DVD("The Social Network", "David Fincher", "D001", 120))


lib.show_all()

lib.borrow_item("Python Crash Course")
lib.borrow_item("Python Crash Course")   


lib.show_available()


lib.return_item("Python Crash Course")
lib.show_available()
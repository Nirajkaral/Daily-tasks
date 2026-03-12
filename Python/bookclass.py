class Book:
    def __init__(self, title, author, pages):
        self.title = title          
        self.author = author        
        self.pages = pages          

    def display_info(self):         
        print(f"""
    Book Details:
    Title   : {self.title}
    Author  : {self.author}
    Pages   : {self.pages}
        """)


b1 = Book("Python Crash Course", "Eric Matthes", 560)
b2 = Book("The Go Programming Language", "Alan Donovan", 380)
b3 = Book("Clean Code", "Robert Martin", 431)


b1.display_info()
b2.display_info()
b3.display_info()


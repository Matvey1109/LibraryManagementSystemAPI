# Косяков Матвей гр. 253505
### Проект "Library Management System API"

***Описание:***
Система менеджмента библиотеки - это программное обеспечение, предназначенное для автоматизации основных библиотечных процессов. Она позволяет организовать и упростить хранение, поиск, выдачу и возврат книг, а также управление информацией о читателях библиотеки.

***Цель проекта:***
повысить эффективность работы библиотеки, улучшить доступ к ее ресурсам и облегчить взаимодействие между библиотекарями и читателями.

## 1. Модели данных:
![DataModels](https://raw.github.com/Matvey1109/LibraryManagementSystemAPI/Lab1/screenshots/DataModels.png)
### Модель Member:
- ID (string): Уникальный идентификатор члена библиотеки.
- Name (string): Имя члена библиотеки.
- Address (string): Адрес проживания члена библиотеки.
- Email (string): Электронная почта члена библиотеки.
- CreatedAt (time): Время регистрации члена библиотеки в системе.

### Модель Book:
- ID (string): Уникальный идентификатор книги.
- Title (string): Название книги.
- Author (string): Автор книги.
- PublicationYear (int): Год публикации книги.
- Genre (string): Жанр книги.
- AvailableCopies (int): Количество доступных копий книги.
- TotalCopies (int): Общее количество копий книги в библиотеке.


### Модель Borrowing:
- ID (string): Уникальный идентификатор записи о заимствовании книги.
- BookID (string): Идентификатор книги, которая была заимствована.
- MemberID (string): Идентификатор члена библиотеки, который заимствовал книгу.
- BorrowYear (int): Год, в котором книга была заимствована.
- ReturnYear (int): Год, в котором книга была возвращена.

Все модели используют строковый тип данных для идентификаторов, что позволяет им быть гибкими, использовать ObjectID() для генерации идентификаторов и иметь структуру для потенциального хранения их в MongoDB.
Модель Member хранит базовую информацию о читателях.
Модель Book хранит основную информацию о книгах.
Модель Borrowing отслеживает заимствования книг.
Все типы связи в моделях имеют отношение One-to-Many.

## 2. Инструменты для хранения данных:
![DataStorage](https://raw.github.com/Matvey1109/LibraryManagementSystemAPI/Lab1/screenshots/DataStorage.png)
## 2.1. Интерфейсы:
- **DataStorage:** Общий интерфейс для всех хранилищ данных. Определяет базовые методы для работы с данными: получение всех элементов, получение по идентификатору, добавление, удаление и обновление. Интерфейс DataStorage действует как адаптер, позволяя клиентам работать с различными реализациями хранилища через единый интерфейс. 
    - **MemberStorage:** Интерфейс для управления данными о читателях библиотеки.
    - **BookStorage:** Интерфейс для управления данными о книгах.
    - **BorrowingStorage:** Интерфейс для управления данными о заимствовании книг.

## 2.2. Реализации:
- **InMemoryDataStorage:** Хранилище данных локально. Является простой реализацией для небольших объемов данных.
  - **InMemoryMemberStorage, InMemoryBookStorage, InMemoryBorrowingStorage:** Конкретные реализации интерфейсов для хранения членов библиотеки, книг и заимствования книг локально.
- **MongoDBDataStorage:** Хранилище данных в базе данных MongoDB. Является более надежным и масштабируемым решением.
    - **MongoDBMemberStorage, MongoDBBookStorage, MongoDBBorrowingStorage:** Конкретные реализации интерфейсов для хранения членов библиотеки, книг и заимствования книг в MongoDB.

## 2.3. Методы:
MemberStorage:
- getAllMembers(): Возвращает список всех членов библиотеки.
- getMember(Member.ID): Возвращает одного члена библиотеки по его идентификатору.
- addMember(Member.Name, Member.Address, Member.Email): Добавляет нового члена библиотеки, указав его имя, адрес и электронную почту.
- deleteMember(Member.ID): Удаляет члена библиотеки по его идентификатору.
- updateMember(Member.ID, Member.Name, Member.Address, Member.Email): Обновляет информацию о члене библиотеки, используя его идентификатор и новые данные (имя, адрес, email).

BookStorage:
- getAllBooks(): Возвращает список всех книг в библиотеке.
- getBook(Book.ID): Возвращает одну книгу по ее идентификатору.
- addBook(Book.Title, Book.Author, Book.PublicationYear, Book.Genre, Book.TotalCopies): Добавляет новую книгу в библиотеку, указав ее название, автора, год издания, жанр и общее количество копий.
- deleteBook(Book.ID): Удаляет книгу из библиотеки по ее идентификатору.
- updateBook(Book.ID, Book.Title, Book.Author, Book.PublicationYear, Book.Genre, Book.AvailableCopies, Book.TotalCopies): Обновляет информацию о книге, используя ее идентификатор и новые данные (название, автор, год издания, жанр, доступные копии, общее количество копий).

BorrowingStorage:
- getAllBorrowings(): Возвращает список всех заимствований книг библиотеки.
- getMemberBooks(Borrowing.MemberID): Возвращает список всех книг, которые заимствовал указанный член библиотеки, используя его идентификатор.
- borrowBook(Borrowing.BookID, Borrowing.MemberID, Borrowing.BorrowDate): Записывает факт заимствования книги членом библиотеки, указывая идентификатор книги, идентификатор члена и дату заимствования.
- returnBook(Borrowing.ID): Записывает факт возврата книги, используя идентификатор записи о заимствовании.

## 2.4. Конструкторы:
- NewInMemoryDataStorage(): функция для создания объекта InMemoryDataStorage.
- NewMongoDBStorage(mongoURI string, dbName string): функция для создания объекта MongoDBStorage, указав URI и название базы данных.

## 3. Сервис, предоставляющий API:
![APIService](https://raw.github.com/Matvey1109/LibraryManagementSystemAPI/Lab1/screenshots/APIService.png)
## 3.1. Реализация:
- **APIService** - это сервис, который предоставляет API для доступа к данным о членах библиотеки, книгах и заимствованиях. Он использует хранилище данных (DataStorage), которое может быть реализовано локально (InMemoryDataStorage) или в базе данных (MongoDBDataStorage). APIService реализует обработку API-запросов. Его методы используют в своей реализации методы интерфейса DataStorage для работы с данными.

## 3.2. Методы (обработчики API-запросов):
Member:
- getAllMembersHandler(): использует метод getAllMembers. *GET /members*
- getMemberHandler(): использует метод getMember. *GET /members/{memberId}*
- addMemberHandler(): использует метод addMember. *POST /members*
- deleteMemberHandler(): использует метод deleteMember. *DELETE /members/{memberId}*
- updateMemberHandler(): использует метод updateMember. *PUT /members/{memberId}*

Книги:
- getAllBooksHandler(): использует метод getAllBooks. *GET /books*
- getBookHandler(): использует метод getBook. *GET /books/{bookId}*
- addBookHandler(): использует метод addBook. *POST /books*
- deleteBookHandler(): использует метод deleteBook. *DELETE /books/{bookId}*
- updateBookHandler(): использует метод updateBook. *PUT /books/{bookId}*

Заимствования:
- getAllBorrowingsHandler(): использует метод getMemberBooks. *GET /borrowings*
- getMemberBooksHandler(): использует метод getMemberBooks. *GET /borrowings/{memberId}*
- borrowBookHandler(): использует метод borrowBook. *POST /borrowings*
- returnBookHandler(): использует метод returnBook. *PUT /borrowings/{borrowingId}*

## 3.3. Конструктор:
- NewAPIService(storage DataStorage): Создает новый экземпляр сервиса APIService, используя указанное хранилище данных (DataStorage). Реализация конструктора создает singleton, что означает, что в рамках приложения будет существовать только один экземпляр APIService.

## 3.4. Дополнительные функции:
- registerAPIEndpoints(apiService *APIService): Регистрирует обработчики API-запросов для сервиса.
- startServer(): Запускает сервер.

## 4. Общая структура проекта:
![Architecture](https://raw.github.com/Matvey1109/LibraryManagementSystemAPI/Lab1/screenshots/Architecture.png)

## 5. Клиентская часть кода:
1. Создание хранилища данных.
2. Создание сервиса APIService.
3. Регистрация API-эндпоинтов.
4. Запуск сервера.

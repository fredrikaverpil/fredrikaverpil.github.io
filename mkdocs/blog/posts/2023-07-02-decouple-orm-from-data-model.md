---
date: 2023-07-02
draft: false
tags:
  - python
---

# Decoupling the ORM class from the data model class

I'm working on a rather large project where we would want to replace a [WSGI](https://en.wikipedia.org/wiki/Web_Server_Gateway_Interface) ORM with an [ASGI](https://en.wikipedia.org/wiki/Asynchronous_Server_Gateway_Interface) ORM, but it's tangled up into everything and called from all over the business logic. If the ORM would've been decoupled from the objects tossed around in the business logic, it would've been much easier to replace the ORM.

This blog post outlines an example of how this can be done with [Pydantic](https://github.com/pydantic/pydantic). I'm also including a bonus section on decoupling the data store communication from the business logic with the help of the "Repository pattern".

<!-- more -->

## Prerequisites

This post was written with Python 3.11 in mind. You'll need to install SQLAlchemy and Pydantic:

```python
pip install sqlalchemy==2.* pydantic==2.*
```

## Defining the ORM models

Let's start with defining ORM models and related functions in `orm.py`:

!!! example "orm.py"

    ```python
    from sqlalchemy import Column, Engine, Integer, String, create_engine
    from sqlalchemy.orm import declarative_base

    Base = declarative_base()


    class UserOrm(Base):
        __tablename__ = "users"

        id = Column(Integer, primary_key=True, nullable=False)
        name = Column(String(50), nullable=False)
        password = Column(String(50), nullable=False)
        email = Column(String(50), nullable=False, unique=True)
    ```

## Defining the business logic entity models

Let's now implement the internal "entity" model we'll use when passing around a user object in our business logic, in `entities.py`. Even if we replace the ORM or the database, this object likely won't change as it should still carry the same attributes and should not require refactorings to our business logic.

What's nice about using Pydantic models for such entities is we'll get the awesome validation Pydantic is known for, for "free".

!!! example "entities.py"

    ```python
    from pydantic import BaseModel, ConfigDict, Field


    class User(BaseModel):
        model_config = ConfigDict(from_attributes=True)

        id: int
        name: str = Field(max_length=50)
        email: str = Field(max_length=50)
    ```

!!! note

    Please note the `from_attributes` configuration. This is where the magic happens, as this enables creating the entity from the ORM object:

    ```python
    user_orm = UserOrm(name="John", email="johndoe@gmail.com", password="hashed_password")
    user = User.model_validate(user_orm)
    ```

## Defining the repositories

With the "Repository pattern", you want to define a tight scope around which code is responsible for communicating with e.g. a persistent data store, such as a database. Your business logic is not supposed to have this ownership. Instead, your business logic can actually run, regardless of which kind of repository you throw at it.

For example, you might want to use SQLAlchemy with Postgres in prod, but for tests maybe you want to use SQLAlchemy with an in-memory SQLite database for faster execution and less setup. Or maybe you want your app to gradually move over onto a different database, database driver, ORM or similar.

Let's define a couple of classes in `repositories.py`. First off, we define the abstract class `UserRepositoryAbc` that explains which required methods all user repositories must include. In this case it's the `create_user` and `get_all_users` methods. Then we implement the `UserRepository` class, which implements logic on how to communicate with our SQLite database using SQLAlchemy.

I'm just going to call this repository `UserRepository` for now, but imagine we could've had `UserSqlAlchemyRepository`, `UserMongoDbRepository`, `UserRedisRepository` or `UserFakeRepository`, all inheriting from `UserRepositoryAbc`. That last one, `UserFakeRepository`, could be used in tests and not even communicate with a real database.

!!! example "repositories.py"

    ```python
    import abc
    from typing import Self

    from sqlalchemy.orm import Session

    from entities import User
    from orm import Base, UserOrm


    class UserRepositoryAbc(abc.ABC):
        @abc.abstractmethod
        def create_user(self, name: str, email: str, hashed_password: str) -> User:
            raise NotImplementedError

        @abc.abstractmethod
        def get_all_users(self) -> list[User]:
            raise NotImplementedError


    class UserRepository(UserRepositoryAbc):
        @property
        def engine(self: Self) -> Engine:
            return create_engine("sqlite:///mydatabase.db", echo=True)

        def create_tables(self: Self, base) -> None:
            base.metadata.create_all(self.engine)

        def create_user(self: Self, name: str, email: str, hashed_password: str) -> User:
            with Session(self.engine) as session:
                user_orm = UserOrm(name=name, email=email, password=hashed_password)
                session.add(user_orm)
                session.commit()
                user = User.model_validate(user_orm)
            return user

        def get_all_users(self: Self) -> list[User]:
            with Session(self.engine) as session:
                users_orm: list[UserOrm] = session.query(UserOrm).all()
                users: list[User] = [
                    User.model_validate(user) for user in users_orm
                ]

            return users
    ```

!!! note "A note on the table creation and inclusion of engine"

    As you can see, I also added a methods `create_tables` and `engine`. These doesn't really belong on a users repository, and you might want to implement this on some general SQLAlchemy repository class or abstract the choice of database away from the ORM. But to avoid making this blog post too long and complicated, I just slapped them on there.

    Note that all ORM models inherting from `Base` in `orm.py` will have all their respective tables created when executing `Base.metadata.create_all(engine)`.

!!! note "A note on "

    The repository methods can take business logic entities (such as `User`) as input, or it can take strings, integers, booleans etc - or no arguments at all. It is likely desirable that it returns entities but that is no strict rule about this. Just have them return what makes the most sense. Just don't return the ORM objects!

    The idea is to limit all occurrences of ORM queries to the repositories and not implementing them in the business logic.

## Let's run some commands!

### Create the db tables

Let's begin by creating the db tables:

```python
>>> from orm import Base
>>> from repositories import UserRepository
>>> UserRepository().create_tables(base=Base)
2023-07-02 15:53:47,602 INFO sqlalchemy.engine.Engine BEGIN (implicit)
2023-07-02 15:53:47,602 INFO sqlalchemy.engine.Engine PRAGMA main.table_info("users")
2023-07-02 15:53:47,602 INFO sqlalchemy.engine.Engine [raw sql] ()
2023-07-02 15:53:47,603 INFO sqlalchemy.engine.Engine PRAGMA temp.table_info("users")
2023-07-02 15:53:47,603 INFO sqlalchemy.engine.Engine [raw sql] ()
2023-07-02 15:53:47,603 INFO sqlalchemy.engine.Engine
CREATE TABLE users (
        id INTEGER NOT NULL,
        name VARCHAR(50) NOT NULL,
        password VARCHAR(50) NOT NULL,
        email VARCHAR(50) NOT NULL,
        PRIMARY KEY (id),
        UNIQUE (email)
)


2023-07-02 15:53:47,603 INFO sqlalchemy.engine.Engine [no key 0.00006s] ()
2023-07-02 15:53:47,604 INFO sqlalchemy.engine.Engine COMMIT
```

### Communicate with db but return entities

Finally, we can now communicate with our database using the ORM but always return our entity objects rather than returning ORM objects directly. This is what our business logic would do, rather than call the ORM objects directly.

```python
>>> from repositories import UserRepository
>>> user = UserRepository().create_user(name="John Doe", email="johndoe@gmail.com", hashed_password="hashed_password")
2023-07-02 15:59:16,700 INFO sqlalchemy.engine.Engine BEGIN (implicit)
2023-07-02 15:59:16,702 INFO sqlalchemy.engine.Engine INSERT INTO users (name, password, email) VALUES (?, ?, ?)
2023-07-02 15:59:16,702 INFO sqlalchemy.engine.Engine [generated in 0.00024s] ('John Doe', 'hashed_password', 'johndoe@gmail.com')
2023-07-02 15:59:16,703 INFO sqlalchemy.engine.Engine COMMIT
2023-07-02 15:59:16,706 INFO sqlalchemy.engine.Engine BEGIN (implicit)
2023-07-02 15:59:16,708 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.name AS users_name, users.password AS users_password, u
sers.email AS users_email
FROM users
WHERE users.id = ?
2023-07-02 15:59:16,708 INFO sqlalchemy.engine.Engine [generated in 0.00015s] (1,)
2023-07-02 15:59:16,709 INFO sqlalchemy.engine.Engine ROLLBACK
>>> print(user)
id=1 name='John Doe' email='johndoe@gmail.com'
```

```python
>>> from repositories import UserRepository
>>> users =  UserRepository().get_all_users()
users =  UserRepository().get_all_users()
2023-07-02 21:11:27,858 INFO sqlalchemy.engine.Engine BEGIN (implicit)
2023-07-02 21:11:27,859 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.name AS users_name, users.password AS users_password
, users.email AS users_email
FROM users
2023-07-02 21:11:27,859 INFO sqlalchemy.engine.Engine [generated in 0.00013s] ()
2023-07-02 21:11:27,860 INFO sqlalchemy.engine.Engine ROLLBACK
>>> print(users)
[User(id=1, name='John Doe', email='johndoe@gmail.com')]
```

## A final note on switching out the repositories

In the above commands, I called the repository directly, just to show what the output would be like, and how it returns the `User` object rather than the ORM object.

But a more desirable pattern is to allow injection of the desired repository into the business logic. Imagine having business logic like below:

```python
def create_user(
    name: str,
    email: str,
    hashed_password: str,
    repository: UserRepositoryAbc = UserRepository(),
) -> User:
    return repository.create_user(
        name=name,
        email=email,
        hashed_password=hashed_password,
    )


def get_all_users(repository: UserRepositoryAbc = UserRepository()) -> list[User]:
    return repository.get_all_users()

```

Here you can see how the functions default to using our `UserRepository`, but but they can technically accept any other repository that abides by the abstract class of `UserRepositoryAbc`.

The above code snippets exhibits "Dependency injection" by allowing the repository to be provided externally, which promotes loose coupling and flexibility. It also aligns with the "Dependency inversion principle", where high-level modules (business logic) should not depend on low-level modules (repositories) directly but should instead depend on abstractions.

!!! note "To instantiate or not instantiate"

    You might have noticed how I've instantiated the classes in the signatures. You might want to consider using static methods in your repository classes instead, so you don't have to do this.

    ```python
    class UserRepository(UserRepositoryAbc):

        @staticmethod
        def create_user(name: str, email: str, hashed_password: str):
            ...

        @staticmethod
        def get_all_users():
            ...


    user = UserRepository.create_user(...)
    users = UserRepository.get_all_users()
    ```

    Or you can use other mechanisms to determine which repository to be used and instantiate the repositories only once.


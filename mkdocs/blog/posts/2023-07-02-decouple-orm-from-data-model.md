---
date: 2023-07-02
draft: false
tags:
  - python
---

# Decoupling the ORM class from the data model class

I'm working on a project where we want to replace a [WSGI](https://en.wikipedia.org/wiki/Web_Server_Gateway_Interface) ORM with an [ASGI](https://en.wikipedia.org/wiki/Asynchronous_Server_Gateway_Interface) ORM, but it's tangled into everything and called from all over the business logic. If the ORM would've been decoupled from the objects tossed around in the business logic, it would've been much easier to replace the ORM.

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

    from pydantic import BaseModel, ConfigDict, Field
    from sqlalchemy import Column, Engine, Integer, String, create_engine
    from sqlalchemy.orm import Session, declarative_base

    Base = declarative_base()


    class UserOrm(Base):
        __tablename__ = "users"

        id = Column(Integer, primary_key=True, nullable=False)
        name = Column(String(50), nullable=False)
        password = Column(String(50), nullable=False)
        email = Column(String(50), nullable=False, unique=True)
    ```

## Defining the entity models

Let's now implement the internal entity we'll use when passing around a user object in our business logic, in `entities.py`. Even if we replace the ORM or the database, this object likely won't change as it should still carry the same attributes and should not require refactorings to our business logic.

!!! example "entities.py"

    ```python
    from pydantic import BaseModel, ConfigDict, Field


    class UserEntity(BaseModel):
        model_config = ConfigDict(from_attributes=True)

        id: int
        name: str = Field(max_length=50)
        email: str = Field(max_length=50)
    ```

!!! note

    Please note the `from_attributes` configuration. This is where the magic happens, as this enables creating the entity from the ORM object:

    ```python
    user_orm = UserOrm(name="John", email="johndoe@gmail.com", hashed_password="hashed_password")
    user = UserEntity(user_orm)
    ```

## Defining the repositories

In the "Repository pattern", you want to isolate all code related to communicating with e.g. a persistent data store such as a database. The goal with this is to define a tight scope around which code owns the responsibility of talking to the data store, and how.

Let's define a couple of classes in `repositories.py`. First off, we define the abstract class `UserRepositoryABC` that explains which required methods all repositories must include. Then we implement the `UserSqlAlchemyRepository` class, which implements logic on how to communicate with our SQLite database using SQLAlchemy.

Imagine that you could here add in a `UserMongoDbRepository`, `UserRedisRepository` or `UserFakeRepository` which could be used by your business logic and/or tests.

!!! example "repositories.py"

    ```python
    import abc
    from typing import Self

    from sqlalchemy.orm import Session

    from entities import UserEntity
    from orm import UserOrm


    class UserRepositoryABC(abc.ABC):
        @abc.abstractmethod
        def add_user(self, name: str, email: str, hashed_password: str) -> UserEntity:
            raise NotImplementedError

        @abc.abstractmethod
        def get_all_users(self) -> list[UserEntity]:
            raise NotImplementedError


    class UserSqlAlchemyRepository(UserRepositoryABC):
        @property
        def engine(self: Self) -> Engine:
            return create_engine("sqlite:///mydatabase.db", echo=True)

        def create_tables(self: Self, base) -> None:
            base.metadata.create_all(self.engine)

        def add_user(self: Self, name: str, email: str, hashed_password: str) -> UserEntity:
            with Session(self.engine) as session:
                user = UserOrm(name=name, email=email, password=hashed_password)
                session.add(user)
                session.commit()
                user_model = UserEntity.model_validate(user)
            return user_model

        def get_all_users(self: Self) -> list[UserEntity]:
            with Session(self.engine) as session:
                users_orm: list[UserOrm] = session.query(UserOrm).all()
                users: list[UserEntity] = [
                    UserEntity.model_validate(user) for user in users_orm
                ]

            return users
    ```

## Let's run some commands!

### Create the db tables

Let's begin by creating the db tables:

```python
>>> from orm import Base
>>> from repositories import UserSqlAlchemyRepository
>>> UserSqlAlchemyRepository().create_tables(base=Base)
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
>>> from repositories import UserSqlAlchemyRepository
>>> user = UserSqlAlchemyRepository().add_user(name="John Doe", email="johndoe@gmail.com", hashed_password="hashed_password")
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
>>> from repositories import UserSqlAlchemyRepository
>>> users = UserSqlAlchemyRepository().get_all_users()
2023-07-02 16:01:06,250 INFO sqlalchemy.engine.Engine BEGIN (implicit)
2023-07-02 16:01:06,251 INFO sqlalchemy.engine.Engine SELECT users.id AS users_id, users.name AS users_name, users.password AS users_password, u
sers.email AS users_email
FROM users
2023-07-02 16:01:06,251 INFO sqlalchemy.engine.Engine [generated in 0.00032s] ()
2023-07-02 16:01:06,253 INFO sqlalchemy.engine.Engine ROLLBACK
>>> print(users)
[UserEntity(id=1, name='John Doe', email='johndoe@gmail.com')]
```

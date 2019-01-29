#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: SimonLuo <simonluo@thstack.com>
import os

from sqlalchemy import Column, DateTime, String, Integer, ForeignKey, func
from sqlalchemy.schema import ForeignKeyConstraint, UniqueConstraint
from sqlalchemy.orm import relationship, backref
from sqlalchemy.ext.declarative import declarative_base


Base = declarative_base()


class User(Base):
    __tablename__ = 'user'
    id = Column(Integer, primary_key=True, autoincrement=True)
    username = Column(String(32))

class Group(Base):
    __tablename__ = 'group'
    id = Column(Integer, primary_key=True, autoincrement=True)
    name = Column(String(32))

class Department(Base):
    __tablename__ = 'department'

    user_id = Column(Integer, ForeignKey('user.id'), primary_key=True)
    group_id = Column(Integer, ForeignKey('group.id'))
    # company = Column(Integer, ForeignKey('company.id'))

    name = Column(String)
    employees = relationship(
        'Employee',
        secondary='department_employee_link'
    )

class Company(Base):
    __tablename__ = 'company'

    user_id = Column(Integer, ForeignKey('user.id'), primary_key=True)
    group_id = Column(Integer, ForeignKey('group.id'), primary_key=True)

    name = Column(String)
    departments = relationship(
        Department,
    )

# class DepartmentEmployeeLink(Base):
#     __tablename__ = 'department_employee_link'
#     department_id = Column(Integer, ForeignKey('department.id'), primary_key=True)
#     employee_id = Column(Integer, ForeignKey('employee.id'), primary_key=True)


db_name = 'db_model2.sqlite'
if os.path.exists(db_name):
    os.remove(db_name)

from sqlalchemy import create_engine
engine = create_engine('sqlite:///' + db_name)

from sqlalchemy.orm import sessionmaker
session = sessionmaker()
session.configure(bind=engine)

if __name__ == "__main__":
    Base.metadata.create_all(engine)

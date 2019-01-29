#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: SimonLuo <simonluo@thstack.com>


import os

from sqlalchemy import Column, DateTime, String, Integer, ForeignKey, func
from sqlalchemy.orm import relationship, backref
from sqlalchemy.ext.declarative import declarative_base


Base = declarative_base()


class Department(Base):
    __tablename__ = 'department'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    employees = relationship(
        'Employee',
        secondary='department_employee_link'
    )


class Employee(Base):
    __tablename__ = 'employee'
    id = Column(Integer, primary_key=True)
    name = Column(String)
    hired_on = Column(DateTime, default=func.now())
    departments = relationship(
        Department,
        secondary='department_employee_link'
    )


class DepartmentEmployeeLink(Base):
    __tablename__ = 'department_employee_link'
    department_id = Column(Integer, ForeignKey('department.id'), primary_key=True)
    employee_id = Column(Integer, ForeignKey('employee.id'), primary_key=True)


db_name = 'alembic_sample.sqlite'
if os.path.exists(db_name):
    os.remove(db_name)

from sqlalchemy import create_engine
engine = create_engine('sqlite:///' + db_name)

from sqlalchemy.orm import sessionmaker
session = sessionmaker()
session.configure(bind=engine)

if __name__ == "__main__":
    Base.metadata.create_all(engine)

import pytest
import sqlite3
import pandas as pd

connection = sqlite3.connect('course.db')

courses = pd.read_sql('SELECT * FROM Course', connection)
sections = pd.read_sql('SELECT * FROM Section', connection)
meetings = pd.read_sql('SELECT * FROM Meeting', connection)
instructors = pd.read_sql('SELECT * FROM Instructor', connection)
classes = pd.read_sql('SELECT * FROM Class', connection)
gpa_entries = pd.read_sql('SELECT * FROM GPA_Entry', connection)

# See database design markdown file for reference


def test_columns():
    assert len(courses.columns) == 23
    assert len(sections.columns) == 11
    assert len(meetings.columns) == 8
    assert len(instructors.columns) == 3
    assert len(classes.columns) == 2
    assert len(gpa_entries.columns) == 18

    for table in [courses, sections, meetings, instructors, classes]:
        cols = table.columns
        for idx in range(1, len(cols)):
            assert table[cols[0]].size == table[cols[0]].size


def test_no_id_duplicates():
    for table in [courses, sections, meetings, instructors]:
        assert len(table['id'].unique()) == table['id'].size


connection.close()

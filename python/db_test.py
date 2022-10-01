import pytest
import sqlite3
import pandas as pd

# ensure columns are all the same length
# correct number of columns
def test_courses_size():
    connection = sqlite3.connect("gpa_dataset.db")
    df = pd.read_sql('SELECT * FROM courses', connection)
    cols = list(df)
    assert len(cols) == 6
    assert len(df[cols[0]]) == 5431
    for i in range(1, len(cols)):
        assert df[cols[0]].size == df[cols[i]].size
    connection.close()

# ensure columns are all the same length
def test_gpa_size():
    connection = sqlite3.connect("gpa_dataset.db")
    df = pd.read_sql('SELECT * FROM gpa', connection)
    cols = list(df)
    assert len(cols) == 17
    assert len(df[cols[0]]) == 61557
    for i in range(1, len(cols)):
        assert df[cols[0]].size == df[cols[i]].size
    connection.close()

def test_no_course_duplicates():
    connection = sqlite3.connect("gpa_dataset.db")
    df = pd.read_sql('SELECT * FROM courses', connection)
    assert len(df['Course Title'].unique()) == df['Course Title'].size
    connection.close()


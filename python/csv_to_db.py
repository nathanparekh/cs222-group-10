import sqlite3
import pandas as pd

# create database/connection
connection = sqlite3.connect("gpa_dataset.db")
cursor = connection.cursor()


# remove duplicates, keeping first occurance
def make_courses_table(connection):
    # read csv
    df = pd.read_csv("uiuc-gpa-dataset.csv")
    df = df.drop_duplicates(subset=['Course Title'])
    df = df[['Year', 'Term', 'YearTerm', 'Subject', 'Number', 'Course Title']]

    # populate db
    df.to_sql('courses', connection, if_exists='replace', index=False)



# gpa table
def make_gpa_table(connection):
    # read csv
    df = pd.read_csv("uiuc-gpa-dataset.csv")
    # create column with subject + number
    df['ID'] = df['Subject'] + df['Number'].astype('str')
    col_id = df.pop('ID')
    df.insert(0, 'ID', col_id)
    df = df.drop(['Year', 'Term', 'Sched Type', 'Course Title', 'Subject', 'Number'], axis=1)
    df.sort_values(by=['ID', 'YearTerm'], axis=0)
    # populate db
    df.to_sql('gpa', connection, if_exists='replace', index=False)
    # print(pd.read_sql('SELECT * FROM gpa WHERE ID="SOC275" LIMIT 30', connection))

connection.close()
import sqlite3
import pandas as pd

# create database/connection
connection = sqlite3.connect("gpa_dataset.db")
cursor = connection.cursor()

# read csv
df = pd.read_csv("uiuc-gpa-dataset.csv")
# remove duplicates, keeping first occurance
df_courses = df.drop_duplicates(subset=['Course Title'])
df_courses = df_courses[['Year', 'Term', 'YearTerm', 'Subject', 'Number', 'Course Title']]

# populate db
df_courses.to_sql('courses', connection, if_exists='replace', index=False)
print(pd.read_sql('SELECT * FROM courses LIMIT 10', connection))


# gpa table
# create column with subject + number
df['ID'] = df['Subject'] + df['Number'].astype('str')
col_id = df.pop('ID')
df.insert(0, 'ID', col_id)
df = df.drop(['Year', 'Term', 'Sched Type', 'Course Title', 'Subject', 'Number'], axis=1)

df.to_sql('gpa', connection, if_exists='replace', index=False)
print(pd.read_sql('SELECT * FROM gpa LIMIT 30', connection))


connection.close()
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
    df = df.drop(['Year', 'Term', 'Sched Type', 'Course Title'], axis=1)
    df.sort_values(by=['ID', 'YearTerm'], axis=0)
    # populate db
    df.to_sql('gpa', connection, if_exists='replace', index=False)
    cursor.executescript("""
        ALTER TABLE gpa
            ADD
                "Student Count" GENERATED ALWAYS AS ("A+" + "A" + "A-" + "B+" + "B" + "B-" + "C+" + "C" + "C-" + "D+" + "D" + "D-" + "F");
        ALTER TABLE gpa
            ADD "Total Quality Points" GENERATED ALWAYS AS (("A+" * 4) + ("A" * 4) + ("A-" * 3.67) + ("B+" * 3.33) + ("B" * 3) +
                                                            ("B-" * 2.67) + ("C+" * 2.33) + ("C" * 2) + ("C-" * 1.67) +
                                                            ("D+" * 1.33) + ("D" * 1) + ("D-" * .67) + ("F" * 0));
        ALTER TABLE gpa
            ADD "Average GPA" GENERATED ALWAYS AS ("Total Quality Points" / "Student Count" );
        ;
    """
    )
    # print(pd.read_sql('SELECT * FROM gpa WHERE ID="SOC275" LIMIT 30', connection))

# make_gpa_table(connection)
connection.close()
import sqlite3
import pandas as pd
import uuid
from cis_api import fetch_semester_as_json

# create database/connection
connection = sqlite3.connect('course.db')
cursor = connection.cursor()


# Replace Course, Section, Meeting, Instructor, Class tables in course.db
# with courses of passed-in year & semester
def make_courses_table(connection, year, semester):
    courses = fetch_semester_as_json(year, semester)
    courses, sections, meetings, instructors, classes = courses_to_tables(
        courses)

    course_df = pd.DataFrame(courses)
    section_df = pd.DataFrame(sections)
    meeting_df = pd.DataFrame(meetings)
    instr_df = pd.DataFrame(instructors)
    class_df = pd.DataFrame(classes)

    course_df.to_sql('Course', connection, if_exists='replace', index=False)
    section_df.to_sql('Section', connection, if_exists='replace', index=False)
    meeting_df.to_sql('Meeting', connection, if_exists='replace', index=False)
    instr_df.to_sql('Instructor', connection, if_exists='replace', index=False)
    class_df.to_sql('Class', connection, if_exists='replace', index=False)


# Convert array of courses to arrays of courses, sections, meetings, instructors, classes that follow the DB schema
def courses_to_tables(courses):
    sections = []
    meetings = []
    instructors = []
    classes = []

    for course in courses:
        course_id = str(uuid.uuid4())
        course.update({'id': course_id})

        for section in course.get('sections'):
            section_id = str(uuid.uuid4())
            section.update({'id': section_id})
            section.update({'course_id': course_id})

            for meeting in section.get('meetings'):
                meeting_id = str(uuid.uuid4())
                meeting.update({'id': meeting_id})
                meeting.update({'section_id': section_id})

                for instr in meeting.get('instructors'):
                    instructor = {}
                    instructor_id = str(uuid.uuid4())
                    instructor.update({'id': instructor_id})
                    instructor.update({'first_name': instr.get('firstName')})
                    instructor.update({'last_name': instr.get('lastName')})
                    instructors.append(instructor)

                    class_obj = {}
                    class_obj.update({'meeting_id': meeting_id})
                    class_obj.update({'instructor_id': instructor_id})
                    classes.append(class_obj)

                del meeting['instructors']
                meetings.append(meeting)
            del section['meetings']
            sections.append(section)
        del course['sections']

    return courses, sections, meetings, instructors, classes


# Replace GPA_Entry table in course.db based on contents of the csv
def make_gpa_table(connection):
    # read csv
    df = pd.read_csv("uiuc-gpa-dataset.csv")

    # rename columns
    df.rename(columns={
        'Sched Type': 'sched_type',
        'A+': 'a_plus',
        'A': 'a',
        'A-': 'a_minus',
        'B+': 'b_plus',
        'B': 'b',
        'B-': 'b_minus',
        'C+': 'c_plus',
        'C': 'c',
        'C-': 'c_minus',
        'D+': 'd_plus',
        'D': 'd',
        'D-': 'd_minus',
        'F': 'f',
        'W': 'w'
    }, inplace=True)

    # drop unnecessary columns
    cols_to_drop = ['Year', 'Term', 'YearTerm', 'Subject',
                    'Number', 'Course Title', 'Primary Instructor']
    df.drop(cols_to_drop, axis='columns', inplace=True)

    # add UUID (Primary Key)
    uuids = []
    for idx in range(len(df.index)):
        uuids.append(str(uuid.uuid4()))
    df.insert(0, 'id', uuids)

    # TODO: add course_id, instructor_id

    # populate db
    df.to_sql('GPA_Entry', connection, if_exists='replace', index=False)

    # # SQL code below should be part of the CLI's logic, because we don't want the
    # # database to have cols that are dependent on (or calculated from) other cols
    #
    # cursor.executescript("""
    #     ALTER TABLE GPA_Entry
    #         ADD
    #             "Student Count" GENERATED ALWAYS AS ("A+" + "A" + "A-" + "B+" + "B" + "B-" + "C+" + "C" + "C-" + "D+" + "D" + "D-" + "F");
    #     ALTER TABLE GPA_Entry
    #         ADD "Total Quality Points" GENERATED ALWAYS AS (("A+" * 4) + ("A" * 4) + ("A-" * 3.67) + ("B+" * 3.33) + ("B" * 3) +
    #                                                         ("B-" * 2.67) + ("C+" * 2.33) + ("C" * 2) + ("C-" * 1.67) +
    #                                                         ("D+" * 1.33) + ("D" * 1) + ("D-" * .67) + ("F" * 0));
    #     ALTER TABLE GPA_Entry
    #         ADD "Average GPA" GENERATED ALWAYS AS ("Total Quality Points" / "Student Count" );
    #     ;
    # """
    # )


# make_courses_table(connection, '2022', 'Winter')
# make_gpa_table(connection)

# df = pd.read_sql('SELECT * FROM Course', connection)
# print(df.head())
# df = pd.read_sql('SELECT * FROM Section', connection)
# print(df.head())
# df = pd.read_sql('SELECT * FROM Meeting', connection)
# print(df.head())
# df = pd.read_sql('SELECT * FROM Instructor', connection)
# print(df.head())
# df = pd.read_sql('SELECT * FROM Class', connection)
# print(df.head())
# df = pd.read_sql('SELECT * FROM GPA_Entry', connection)
# print(df.head())

connection.close()

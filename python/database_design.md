# The Course Database

**Purpose:** The course database keeps a list of course information and GPAs, compiled from different sources, for the purpose of searching, sorting, analyzing, and generating reports of the data.

**Database Engine:** SQLite3

## Tables

Below we specify each column of each table by its name, description, type (SQLite3 storage class), and an example.

### Course

| Column Name | Description | Type | Example |
|-------------|-------------|------|---------|
| id | Version 4 UUID (Primary key) | TEXT | ffb094c2-e9b5-450a-88a3-51249195ec63 |
| year | Year of this course offering | INTEGER | 2022 |
| term | Term of this course offering | TEXT | Fall |
| subject | Subject code | TEXT | CS |
| number | Course number | INTEGER | 225 |
| name | Course name | TEXT | Data Structures |
| description | Course description | TEXT | Data abstractions: elementary data structures (lists, stacks, queues, and trees) and their ... |
| credit_hours | Credit hours | TEXT | 4 hours. |
| advanced_comp | Fulfills Advanced Composition | INTEGER | 0 |
| non_western | Fulfills Non-Western Cultures | INTEGER | 0 |
| us_minority | Fulfills U.S. Minority Cultures | INTEGER | 0 |
| western | Fulfills Western/Comparative Cultures | INTEGER | 0 |
| hist_phil | Fulfills Historical & Philosophical Perspectives | INTEGER | 0 |
| lit_arts | Fulfills Literature & the Arts | INTEGER | 0 |
| life_sci | Fulfills Life Sciences | INTEGER | 0 |
| phys_sci | Fulfills Physical Sciences | INTEGER | 0 |
| quant_reas_1 | Fulfills Quantitative Reasoning I | INTEGER | 0 |
| quant_reas_2 | Fulfills Quantitative Reasoning II | INTEGER | 1 |
| behav_sci | Fulfills Behavioral Science | INTEGER | 0 |
| social_sci | Fulfills Advanced Composition | INTEGER | 0 |
| section_info | Other information about the course | TEXT | Credit is not given for CS 277 if credit for CS 225 has been earned. Prerequisite: ... |
| degree_attribs | Information about fulfilled degree requirements | TEXT | Quantitative Reasoning II course. |
| schedule_info | Information about registration | TEXT | Students must register for one lecture-discussion and one lecture section. |

### Section

| Column Name | Description | Type | Example |
|-------------|-------------|------|---------|
| id | Version 4 UUID (Primary key) | TEXT | ea5f9915-1d10-4d73-a133-8f337437fc6a |
| course_id | Primary key of the course this section is listed under | TEXT | ffb094c2-e9b5-450a-88a3-51249195ec63 |
| crn | Section CRN | INTEGER | 65054 |
| number | Section number | TEXT | ABA |
| description | Section text | TEXT | LAPTOP LAB SECTION -- Students are required to ... |
| status_code | Status code | TEXT | A |
| sect_status_code | Section status code | TEXT | A |
| part_of_term | Part of term | TEXT | 1 or A |
| enrollment_status | Enrollment status | TEXT | Open |
| start_date | Start date | TEXT | 2022-08-22Z |
| end_date | End date | TEXT | 2022-12-07Z |

### Meeting

| Column Name | Description | Type | Example |
|-------------|-------------|------|---------|
| id | Version 4 UUID (Primary key) | TEXT | 052a8875-9dc4-47b4-a676-522743873ca3 |
| section_id | Primary key of the section of this meeting | TEXT | ea5f9915-1d10-4d73-a133-8f337437fc6a |
| type | Meeting type | TEXT | Laboratory-Discussion |
| start_time | Start time | TEXT | 09:00 AM |
| end_time | End time | TEXT | 10:50 AM |
| days_of_week | Which days this meeting convenes | TEXT | R |
| room_num | Room number of meeting location | TEXT | 4029 or AUD |
| building_name | Building name of meeting location | TEXT | Campus Instructional Facility |

### Instructor

| Column Name | Description | Type | Example |
|-------------|-------------|------|---------|
| id | Version 4 UUID (Primary key) | TEXT | bfe81e58-7b97-4675-b713-4000b9a6d978 |
| first_name | First name | TEXT | C |
| last_name | Last name | TEXT | Evans |

### Class

| Column Name | Description | Type | Example |
|-------------|-------------|------|---------|
| meeting_id | Primary key of the meeting | TEXT | 052a8875-9dc4-47b4-a676-522743873ca3 |
| instructor_id | Primary key of the instructor | TEXT | bfe81e58-7b97-4675-b713-4000b9a6d978 |

A meeting (i.e. lecture or lab) and an instructor make a class. This is a junction table that defines the many-to-many relationship between meetings and instructors.

### GPA_Entry

| Column Name | Description | Type | Example |
|-------------|-------------|------|---------|
| id | Primary key | INTEGER | 7599854 |
| course_id | Primary key of the corresponding course | INTEGER | 9813204 |
| instructor_id | Primary key of the main instructor | INTEGER | 2825122 |
| sched_type | Schedule type | TEXT | LEC |
| a_plus | Number of students who received an A+ | INTEGER | 8 |
| a | A | INTEGER | 241 |
| a_minus | A- | INTEGER | 14 |
| b_plus | B+ | INTEGER | 22 |
| b | B | INTEGER | 15 |
| b_minus | B- | INTEGER | 6 |
| c_plus | C+ | INTEGER | 10 |
| c | C | INTEGER | 4 |
| c_minus | C- | INTEGER | 6 |
| d_plus | D+ | INTEGER | 5 |
| d | D | INTEGER | 1 |
| d_minus | D- | INTEGER | 3 |
| f | F | INTEGER | 10 |
| w | W | INTEGER | 5 |

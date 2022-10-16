import requests
from xml.etree import ElementTree

# Course Information Suite (CIS) API docs - https://courses.illinois.edu/cisdocs/

# Returns info on every course accessible through the CIS API; takes several hours
def fetch_schedule_history_as_json():
    # request for a list of school years offered by the CIS API
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule.xml'
    res = requests.get(endpoint)

    # parse the XML response into an ElementTree
    root = ElementTree.fromstring(res.content)

    # build the entire schedule JSON obj
    schedule = {}
    for year_elem in root.iter('calendarYear'):
        year = year_elem.attrib.get('id')
        semesters = fetch_year_as_json(year)
        schedule.update({year: semesters})

    return schedule

# Returns courses within a given school year; takes over 1 hour
def fetch_year_as_json(year):
    # request for the semesters of this school year
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}.xml'.format(year)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build school year JSON obj
    semesters = {}
    for sem_elem in root.iter('term'):
        semester = sem_elem.text.split()[0]
        subjects = fetch_semester_as_json(year, semester)
        semesters.update({sem: subjects})

    return semesters

# Returns courses within a given semester; takes over 30 mins for spring/fall semesters
def fetch_semester_as_json(year, semester):
    # request for subjects taught in this semester
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}/{}.xml'.format(year, semester)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build semester JSON obj, comprised of each subject mapped to its course offerings
    subjects = {}
    for subj_elem in root.iter('subject'):
        subj_code = subj_elem.attrib.get('id')
        courses = fetch_subj_as_json(year, semester, subj_code)
        subjects.update({subj_code: courses})

    return subjects

# Returns courses within a givin subject and semester
def fetch_subj_as_json(year, semester, subj_code):
    # request for course offerings of this subject
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}/{}/{}.xml'.format(year, semester, subj_code)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build array of course JSON objs
    courses = []
    for course_elem in root.iter('course'):
        course = fetch_course_as_json(year, semester, subj_code, course_elem.attrib.get('id'))
        courses.append(course)

    return courses

# Returns course
def fetch_course_as_json(year, semester, subj_code, course_num):
    # request for information about the course offering
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}/{}/{}/{}.xml'.format(year, semester, subj_code.upper(), course_num)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build the course JSON obj
    course = {}
    tags = ['label', 'description', 'creditHours', 'courseSectionInformation', 'sectionDegreeAttributes', 'classScheduleInformation']
    for tag in tags:
        elem = root.find(tag)
        if (elem is not None):
            course.update({tag: elem.text})

    # iterate through the course sections
    sections = []
    for section_elem in root.iter('section'):
        section = fetch_section_as_json(year, semester, subj_code, course_num, section_elem.attrib.get('id'))
        sections.append(section)

    course.update({'sections': sections})
    return course

# Returns section within a course
def fetch_section_as_json(year, semester, subj_code, course_num, crn):
    # request for information about the course section
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}/{}/{}/{}/{}.xml'.format(year, semester, subj_code.upper(), course_num, crn)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build section JSON obj
    section = {}
    tags = ['sectionNumber', 'statusCode', 'sectionText', 'partOfTerm', 'sectionStatusCode', 'enrollmentStatus', 'startDate', 'endDate']
    for tag in tags:
        elem = root.find(tag)
        if (elem is not None):
            section.update({tag: elem.text})

    # iterate through the meetings of this course section
    meetings = []
    for meeting_elem in root.iter('meeting'):
        # build meeting JSON obj (will be added to section JSON obj)
        meeting = {}
        tags = ['type', 'start', 'end', 'daysOfTheWeek', 'roomNumber', 'buildingName']
        for tag in tags:
            elem = meeting_elem.find(tag)
            if (elem is not None):
                meeting.update({tag: elem.text})

        # iterate through the instructors of this meeting
        instrs = []
        for instr_elem in meeting_elem.iter('instructor'):
            # build instructor JSON obj (will be added to meeting JSON obj)
            instr = {}
            instr.update({'firstName': instr_elem.attrib.get('firstName')})
            instr.update({'lastName': instr_elem.attrib.get('lastName')})
            instrs.append(instr)

        meeting.update({'instructors': instrs})
        meetings.append(meeting)

    section.update({'meetings': meetings})
    return section

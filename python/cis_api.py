import requests
from xml.etree import ElementTree
#import json

# Course Information Suite (CIS) API docs - https://courses.illinois.edu/cisdocs/


# Returns info on every course accessible through the CIS API; takes several hours
def fetch_schedule_history_as_json():
    # request for a list of school years offered by the CIS API
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule.xml'
    res = requests.get(endpoint)

    # parse the XML response into an ElementTree
    root = ElementTree.fromstring(res.content)

    # build array of course JSON objs
    courses = []
    for year_elem in root.iter('calendarYear'):
        year = year_elem.attrib.get('id')
        courses = courses + fetch_year_as_json(year)

    return courses


# Returns courses within a given school year; takes over 1 hour
def fetch_year_as_json(year):
    # request for the semesters of this school year
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}.xml'.format(
        year)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build array of course JSON objs
    courses = []
    for sem_elem in root.iter('term'):
        semester = sem_elem.text.split()[0]
        courses = courses + fetch_semester_as_json(year, semester)

    return courses


# Returns courses within a given semester; takes over 30 mins for spring/fall semesters
def fetch_semester_as_json(year, semester):
    # request for subjects taught in this semester
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}/{}.xml'.format(
        year, semester)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build array of course JSON objs
    courses = []
    for subj_elem in root.iter('subject'):
        subj_code = subj_elem.attrib.get('id')
        courses = courses + fetch_subj_as_json(year, semester, subj_code)

    return courses


# Returns courses within a givin subject and semester
def fetch_subj_as_json(year, semester, subj_code):
    # request for course offerings of this subject
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}/{}/{}.xml'.format(
        year, semester, subj_code)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build array of course JSON objs
    courses = []
    for course_elem in root.iter('course'):
        course = fetch_course_as_json(
            year, semester, subj_code, course_elem.attrib.get('id'))
        courses.append(course)

    return courses


def fetch_course_as_json(year, semester, subj_code, course_num):  # Returns course
    # request for information about the course offering
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}/{}/{}/{}.xml'.format(year,
                                                                                              semester, subj_code.upper(), course_num)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build the course JSON obj
    course = {}
    course.update({'year': int(year.strip())})
    course.update({'term': semester})
    course.update({'subject': subj_code})
    course.update({'number': int(course_num.strip())})

    tags = ['label', 'description', 'creditHours', 'courseSectionInformation',
            'sectionDegreeAttributes', 'classScheduleInformation']
    tag_aliases = ['name', 'description', 'credit_hours',
                   'section_info', 'degree_attribs', 'schedule_info']
    for idx, tag in enumerate(tags):
        elem = root.find(tag)
        if (elem is not None):
            course.update({tag_aliases[idx]: elem.text.strip()})

    geneds = ['1CLL', '1NW', '1US', '1WCC', '1HP', '1LA',
              '1LS', '1PS', '1QR1', '1QR2', '1BSC', '1SS']
    gened_aliases = ['advanced_comp', 'non_western', 'us_minority', 'western', 'hist_phil', 'lit_arts', 'life_sci',
                     'phys_sci', 'quant_res_1', 'quant_res_2', 'behav_sci', 'social_sci']
    geneds_fulfilled = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

    for gened_elem in root.iter('genEdAttribute'):
        gened = gened_elem.attrib.get('code')
        try:
            idx = geneds.index(gened)
            geneds_fulfilled[idx] = 1
        except ValueError:
            pass  # if gened code is not recognized, do nothing

    for idx, gened in enumerate(gened_aliases):
        course.update({gened: geneds_fulfilled[idx]})

    # iterate through the course sections
    sections = []
    for section_elem in root.iter('section'):
        section = fetch_section_as_json(
            year, semester, subj_code, course_num, section_elem.attrib.get('id'))
        sections.append(section)

    course.update({'sections': sections})
    return course


# Returns section within a course
def fetch_section_as_json(year, semester, subj_code, course_num, crn):
    # request for information about the course section
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}/{}/{}/{}/{}.xml'.format(
        year, semester, subj_code.upper(), course_num, crn)
    res = requests.get(endpoint)
    root = ElementTree.fromstring(res.content)

    # build section JSON obj
    section = {}
    section.update({'crn': int(crn.strip())})

    tags = ['sectionNumber', 'statusCode', 'sectionText', 'partOfTerm',
            'sectionStatusCode', 'enrollmentStatus', 'startDate', 'endDate']
    tag_aliases = ['number', 'status_code', 'description', 'part_of_term',
                   'sect_status_code', 'enrollment_status', 'start_date', 'end_date']
    for idx, tag in enumerate(tags):
        elem = root.find(tag)
        if (elem is not None):
            section.update({tag_aliases[idx]: elem.text.strip()})
        else:
            section.update({tag_aliases[idx]: None})

    # iterate through the meetings of this course section
    meetings = []
    for meeting_elem in root.iter('meeting'):
        # build meeting JSON obj (will be added to section JSON obj)
        meeting = {}
        tags = ['type', 'start', 'end', 'daysOfTheWeek',
                'roomNumber', 'buildingName']
        tag_aliases = ['type', 'start_time', 'end_time',
                       'days_of_week', 'room_num', 'building_name']
        for idx, tag in enumerate(tags):
            elem = meeting_elem.find(tag)
            if (elem is not None):
                meeting.update({tag_aliases[idx]: elem.text.strip()})
            else:
                meeting.update({tag_aliases[idx]: None})

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

#info = fetch_semester_as_json('2022', 'Winter')
#print(json.dumps(info, indent=2))

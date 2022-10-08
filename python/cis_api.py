import requests
from xml.etree import ElementTree

# Course Information Suite (CIS) API docs - https://courses.illinois.edu/cisdocs/

def fetch_course_as_json(year, semester, subj_code, course_num):
    # request for information about the course offering
    endpoint = 'https://courses.illinois.edu/cisapp/explorer/schedule/{}/{}/{}/{}.xml'.format(year, semester, subj_code.upper(), course_num)
    res = requests.get(endpoint)

    # parse the XML response into an ElementTree
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
        # request for information about the course section
        sect_endpoint = section_elem.attrib.get('href')
        sect_res = requests.get(sect_endpoint)
        sect_root = ElementTree.fromstring(sect_res.content)

        # build section JSON obj (will be added to course JSON obj)
        section = {}
        tags = ['sectionNumber', 'statusCode', 'sectionText', 'partOfTerm', 'sectionStatusCode', 'enrollmentStatus', 'startDate', 'endDate']
        for tag in tags:
            elem = sect_root.find(tag)
            if (elem is not None):
                section.update({tag: elem.text})

        # iterate through the meetings of this course section
        meetings = []
        for meeting_elem in sect_root.iter('meeting'):
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
        sections.append(section)

    course.update({'sections': sections})
    return course

# course = fetch_course_as_json('2022', 'Fall', 'CS', '225')
# print(course)

# Introduction
gocourse is a command line interface that allows a user to access UIUC course explorer using the terminal. It also integrates GPA data from Professor Fagen's website. Users can access a course's information, its past offerings, its grade distribution, and prerequisites.
# Technical architecture 
- The CLI application and data layer is built in Go. We use the cobra library for command line input, gookit/color and disiquiera/gotree for command line output, and Gorm to access the database in the data layer.
- We use a relational database which is accessed using SQLite. The initial database setup was done in python using pandas and xml.etree. The data was pulled from Course Explorer and Prof. Fagen's website.
# Installation
1. Clone the repository. This includes all source code as well as the database file.
2. Move into the cli directory.
3. Install required packages using `go get`.
4. You can now use the project, either by running `go run main.go` or the `./gocourse` executable.
# Group members
- Jamie: Wrote code to parse course names and numbers that were passed into commands as arguments. This code was reused and modified by others for their own commands. Wrote the `subject` and `course` command.
- Amanda: Created the initial database with Nathan. Refactored the process of pulling data from the database using Gorm. Wrote the `course history` and `course gpa` commands.
- Nathan: Created the inital database with Amanda. Created the database sections for meeting times. Wrote the `sections` command.
- David: Designed the database. Created the CIS API client, including parsing and cleaning data. Parsed prerequisite data from descriptions of courses. Wrote the `prereqs` command.
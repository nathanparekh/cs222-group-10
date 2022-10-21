https://cobra.dev/#concepts  
Commands usually follow the structure:  
`APPNAME VERB NOUN --ADJECTIVE.` or `APPNAME COMMAND ARG --FLAG`  
Example: `git clone URL --bare`  

https://www.bradcypert.com/testing-a-cobra-cli-in-go/  
Example template to make testing easier (second picture)  

https://clig.dev/  
Base commands should do the minimum default. Flags give the option to display more information  

https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/  
https://travis.media/how-to-use-subcommands-in-cobra-go-cobra-tutorial/  
https://dev.to/divrhino/adding-flags-to-a-command-line-tool-built-with-go-and-cobra-34f1  
https://stackoverflow.com/questions/38105859/make-a-cobra-command-flag-required  
Resource about making subcommands/flags  

https://go.dev/doc/database/querying  
https://pkg.go.dev/database/sql#DB.QueryRow  
https://www.allhandsontech.com/programming/golang/how-to-use-sqlite-with-go/  
Resource for sql  

## Examples:   
- `our-cli course --name="CS225"` display basic info on a course (name, description, professors, average class size)  
- `our-cli course history --name="Data Structures" -l` would return all the times CS225 was offered, from latest to oldest  
- `our-cli course gpa --name="Data Structures"` would return the grade breakdown/average gpa of CS225 in the latest semester  
- `our-cli course gpa --name="CS225" -a` would return all gpa data we have on cs225  
- `our-cli search geneds -s -g` would rank all geneds based on -s, --size and -g, --gpa  
- `our-cli search geneds -a` would just display all geneds in no particular order
- `our-cli search geneds --req=NW -s -g` would rank geneds fulfilling Non-western requirement based on -s, --size and -g, --gpa  
- `our-cli search geneds --req="NW,HP"` would list geneds fulfilling NW and HP requirement in no particular order  

1. The --name flag would be a persistent flag belonging to the "course" command and accept either a course number or title  
2. "gpa" and "history" are subcommands of course and can't be called on their own. This structure is easier to read.  
3. We can require the --name flag in the base command "course" since all subcommand of "course" need a speciifc course to function.  

```
Usage:
  course-project-group-10.git [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  course      Base command for searching course information
  geneds      Subcommand for filtering geneds
  gpa         Get the gpa breakdown of course in the most recent semester (As, Bs, etc., average gpa per instructor, section size)
  help        Help about any command
  history     Lists when a course was previously offered, by default returns most recent offering
  requisite   Display pre or post requisites of a course
  search      Base command for searching course attributes
  subject     Lists courses belonging to a particular subject

Flags:
  -h, --help     help for course-project-group-10.git
  -t, --toggle   Help message for toggle

Course Flags:
  -a, --all      display all data we have
  -l, --latest   sort result by latest first
  -n, --num      number of semesters to display (applied after sort)
      --name     course name to specify (required)
  -o, --oldest   sort result by oldest first
  -r, --recent   only return most recent result (default)

    Gpa Flags:
      -i, --instructor  filter by passed instructor
      -s, --size        sort by class size

Search Flags:
  -g, --gpa      sort by gpa
  -r, --req      string containing gened requirements to search for (passed as comma separated string?)
  -s, --size     sort by class size
```

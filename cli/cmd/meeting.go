package cmd

//
//import (
//	"errors"
//	"github.com/CS222-UIUC/course-project-group-10.git/data"
//	"github.com/spf13/cobra"
//)
//
//func getMeeting(args []string) ([]data.Meeting, error) {
//	if len(args) == 0 {
//	} else if len(args) == 1 { // argument is probably a course name
//		meeting, err := data.GetMeetings(map[string]interface{}{"crn": args[0]}, "LIMIT 1")
//		return meeting, err
//	}
//
//	return []data.Meeting{}, errors.New("malformed arguments")
//}
//
//func printMeeting(cmd *cobra.Command , args []string) {
//	//meetings, err := getMeeting(args)
//	//
//	//if err != nil {
//	//	fmt.Println("Error getting meeting:", err)
//	//	fmt.Println("Usage:")
//	//	fmt.Println("meeting [crn] to get meetings by name")
//	//} else {
//	//	//for i, meeting := range meetings {
//	//	//	fmt.Println()
//	//	//}
//	//}
//}
//
//var meetingCmd = &cobra.Command{
//	Use:   "meeting",
//	Short: "Lists meeting(s)",
//	Long: `Prints a list of meetings associated with a provided section`,
//	Run: printMeeting,
//}
//
//func init() {
//	rootCmd.AddCommand(meetingCmd)
//}

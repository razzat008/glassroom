package classroomauths

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/classroom/v1"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func GetClient() *http.Client {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, classroom.ClassroomCourseworkMeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil { // read the code from the stdin
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// function to send announcements
func SendAnnouncement(srv *classroom.Service, courseId *string) (*classroom.ListAnnouncementsResponse, error) {
	r, err := srv.Courses.Announcements.List(*courseId).PageSize(1).Do()
	if err != nil {
		log.Fatalf("Unable to send announcements: %v", err)
	}
	if len(r.Announcements) > 0 {
		return r, err
	} else {
		fmt.Print("No announcements found.")
	}
	return nil, err
}

func ListCourses(srv *classroom.Service) (*classroom.ListCoursesResponse, error) {
	r, err := srv.Courses.List().PageSize(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve courses. %v", err)
	}
	if len(r.Courses) > 0 {
		// printout the course
		// fmt.Print("Courses:\n")
		// for _, c := range r.Courses {
		// 	fmt.Printf("%s (%s)\n", c.Name, c.Id)
		// }
		return r, err
	} else {
		fmt.Print("No courses found.")
	}
	return nil, err
}

// a generic function to  conne
func CreateServiceToClassroom(client *http.Client) (*classroom.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // cancel the request if it takes longer than 10 seconds
	defer cancel()

	srv, err := classroom.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
		return nil, err
	}
	return srv, err
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type QuizResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Quiz struct {
			IsTagged bool `json:"isTagged"`
			IsLoved  bool `json:"isLoved"`
			Stats    struct {
				Played         int `json:"played"`
				TotalPlayers   int `json:"totalPlayers"`
				TotalCorrect   int `json:"totalCorrect"`
				TotalQuestions int `json:"totalQuestions"`
			} `json:"stats"`
			Love             int         `json:"love"`
			Cloned           bool        `json:"cloned"`
			ParentDetail     interface{} `json:"parentDetail"`
			Deleted          bool        `json:"deleted"`
			DraftVersion     interface{} `json:"draftVersion"`
			PublishedVersion string      `json:"publishedVersion"`
			IsShared         bool        `json:"isShared"`
			Type             string      `json:"type"`
			ID               string      `json:"_id"`
			CreatedBy        struct {
				Local struct {
					Username      string `json:"username"`
					CasedUsername string `json:"casedUsername"`
				} `json:"local"`
				Google struct {
					DisplayName string `json:"displayName"`
					Email       string `json:"email"`
					FirstName   string `json:"firstName"`
					Image       string `json:"image"`
					LastName    string `json:"lastName"`
					ProfileID   string `json:"profileId"`
				} `json:"google"`
				Student     interface{} `json:"student"`
				Deactivated bool        `json:"deactivated"`
				Deleted     bool        `json:"deleted"`
				LastName    string      `json:"lastName"`
				Media       string      `json:"media"`
				FirstName   string      `json:"firstName"`
				Country     string      `json:"country"`
				Occupation  string      `json:"occupation"`
				Title       string      `json:"title"`
				ID          string      `json:"id"`
			} `json:"createdBy"`
			Updated   time.Time `json:"updated"`
			CreatedAt time.Time `json:"createdAt"`
			Info      struct {
				ID         string    `json:"_id"`
				Lang       string    `json:"lang"`
				Name       string    `json:"name"`
				CreatedAt  time.Time `json:"createdAt"`
				Updated    time.Time `json:"updated"`
				Visibility bool      `json:"visibility"`
				Questions  []struct {
					ID        string `json:"_id"`
					Time      int    `json:"time"`
					Type      string `json:"type"`
					Published bool   `json:"published"`
					Structure struct {
						Settings struct {
							HasCorrectAnswer bool   `json:"hasCorrectAnswer"`
							FibDataType      string `json:"fibDataType"`
						} `json:"settings"`
						Explain interface{} `json:"explain"`
						Kind    string      `json:"kind"`
						Options []struct {
							Math struct {
								Latex []interface{} `json:"latex"`
							} `json:"math"`
							Type    string        `json:"type"`
							HasMath bool          `json:"hasMath"`
							Media   []interface{} `json:"media"`
							Text    string        `json:"text"`
						} `json:"options"`
						Query struct {
							Math struct {
								Latex []interface{} `json:"latex"`
							} `json:"math"`
							Type    string        `json:"type"`
							HasMath bool          `json:"hasMath"`
							Media   []interface{} `json:"media"`
							Text    string        `json:"text"`
						} `json:"query"`
						Answer int `json:"answer"`
					} `json:"structure"`
					Standards     []interface{} `json:"standards"`
					Topics        []interface{} `json:"topics"`
					IsSuperParent bool          `json:"isSuperParent"`
					CreatedAt     time.Time     `json:"createdAt"`
					Updated       time.Time     `json:"updated"`
					Cached        bool          `json:"cached"`
				} `json:"questions"`
				Subjects   []string      `json:"subjects"`
				Topics     []string      `json:"topics"`
				Subtopics  []string      `json:"subtopics"`
				Image      string        `json:"image"`
				Grade      []string      `json:"grade"`
				GradeLevel interface{}   `json:"gradeLevel"`
				Deleted    bool          `json:"deleted"`
				Standards  []interface{} `json:"standards"`
				Pref       struct {
					Time interface{} `json:"time"`
				} `json:"pref"`
				Traits struct {
					IsQuizWithoutCorrectAnswer bool `json:"isQuizWithoutCorrectAnswer"`
					TotalSlides                int  `json:"totalSlides"`
				} `json:"traits"`
				Theme struct {
					FontSize struct {
					} `json:"fontSize"`
					FontColor struct {
					} `json:"fontColor"`
					Background struct {
					} `json:"background"`
				} `json:"theme"`
				Cached         bool          `json:"cached"`
				QuestionDrafts []interface{} `json:"questionDrafts"`
				Courses        []interface{} `json:"courses"`
				IsProfane      bool          `json:"isProfane"`
				Whitelisted    bool          `json:"whitelisted"`
			} `json:"info"`
			HasPublishedVersion bool        `json:"hasPublishedVersion"`
			HasDraftVersion     bool        `json:"hasDraftVersion"`
			Lock                interface{} `json:"lock"`
		} `json:"quiz"`
		Draft interface{} `json:"draft"`
	} `json:"data"`
	Meta struct {
		Service string `json:"service"`
		Version string `json:"version"`
	} `json:"meta"`
}

func getQuiz(id string) QuizResponse {
	url := "https://quizizz.com/quiz/" + id
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result QuizResponse

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func main() {
	quiz := getQuiz("59be7b4a1557731200ccfa41")
	fmt.Println(quiz)

	router := gin.Default()
	err := router.Run()
	if err != nil {
		panic("Running server router failed!" + string(err.Error()))
	}
}

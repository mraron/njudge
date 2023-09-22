package handlers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mraron/njudge/internal/web/domain/problem"
	"github.com/mraron/njudge/internal/web/handlers/problemset"
	"github.com/mraron/njudge/pkg/problems"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/web/domain/submission"
	"github.com/mraron/njudge/internal/web/helpers"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/internal/web/helpers/roles"
	"github.com/mraron/njudge/internal/web/models"
	"github.com/mraron/njudge/internal/web/services"
)

type SubmissionSummary struct {
	problemset.StatusRow
	Code string `json:"code"`
}

type SubmissionGroup struct {
	Name      string               `json:"name"`
	Completed bool                 `json:"completed"`
	Failed    bool                 `json:"failed"`
	Score     float64              `json:"score"`
	MaxScore  float64              `json:"maxScore"`
	Scoring   problems.ScoringType `json:"scoring"`
	Testcases []SubmissionTestcase `json:"testCases"`
}

type SubmissionTestcase struct {
	Index          int                 `json:"index"`
	VerdictType    problem.VerdictType `json:"verdictType"`
	VerdictName    string              `json:"verdictName"`
	TimeSpent      string              `json:"timeSpent"`
	MemoryUsed     string              `json:"memoryUsed"`
	Score          float64             `json:"score"`
	MaxScore       float64             `json:"maxScore"`
	Output         string              `json:"output"`
	ExpectedOutput string              `json:"expectedOutput"`
	CheckerOutput  string              `json:"checkerOutput"`
}

type SubmissionStatus struct {
	Compiled     bool                  `json:"compiled"`
	FeedbackType problems.FeedbackType `json:"feedbackType"`
	Groups       []SubmissionGroup     `json:"groups"`
}

func SubmissionStatusFromSubmission(tr i18n.Translator, sub *submission.Submission) (*SubmissionStatus, error) {
	status := &problems.Status{}
	if err := status.Scan(sub.Status); err != nil {
		return nil, err
	}

	getTestcases := func(tcs []problems.Testcase) []SubmissionTestcase {
		res := make([]SubmissionTestcase, len(tcs))
		for i := range tcs {
			res[i] = SubmissionTestcase{
				Index:       tcs[i].Index,
				VerdictType: problem.Verdict(tcs[i].VerdictName).VerdictType(),
				VerdictName: problem.Verdict(tcs[i].VerdictName).Translate(tr),
				TimeSpent:   fmt.Sprintf("%d ms", tcs[i].TimeSpent/time.Millisecond),
				MemoryUsed:  fmt.Sprintf("%d KiB", tcs[i].MemoryUsed),
				Score:       tcs[i].Score,
				MaxScore:    tcs[i].MaxScore,
			}

			if status.FeedbackType == problems.FeedbackCF {
				res[i].Output = tcs[i].Output
				res[i].ExpectedOutput = tcs[i].ExpectedOutput
				res[i].CheckerOutput = tcs[i].CheckerOutput
			}
		}

		return res
	}

	res := &SubmissionStatus{}
	res.Compiled = status.Compiled
	res.FeedbackType = status.FeedbackType
	if len(status.Feedback) > 0 && status.FeedbackType != problems.FeedbackACM {
		for i := range status.Feedback[0].Groups {
			grp := status.Feedback[0].Groups[i]

			submissionGroup := SubmissionGroup{
				Testcases: getTestcases(grp.Testcases),
			}

			if status.FeedbackType != problems.FeedbackCF {
				submissionGroup.Name = grp.Name
				allRan := true
				allDRPCorAC := true
				for _, tc := range grp.Testcases {
					allRan = allRan && (tc.VerdictName != problems.VerdictDR)
					allDRPCorAC = allDRPCorAC &&
						(tc.VerdictName == problems.VerdictDR ||
							tc.VerdictName == problems.VerdictPC ||
							tc.VerdictName == problems.VerdictAC)
				}

				if status.FeedbackType == problems.FeedbackIOI {
					submissionGroup.Completed = allRan
				} else if status.FeedbackType == problems.FeedbackLazyIOI {
					submissionGroup.Completed = allRan || !allDRPCorAC
				}

				submissionGroup.Failed = !allDRPCorAC
				submissionGroup.Score = grp.Score()
				submissionGroup.MaxScore = grp.MaxScore()
				submissionGroup.Scoring = grp.Scoring
			}

			res.Groups = append(res.Groups, submissionGroup)
		}
	}

	return res, nil
}

func GetSubmission(DB *sqlx.DB, problemStore problems.Store, r submission.Repository) echo.HandlerFunc {
	type request struct {
		ID int `param:"id"`
	}

	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := &request{}
		if err := c.Bind(data); err != nil {
			return err
		}

		sub, err := r.Get(c.Request().Context(), data.ID)
		if err != nil {
			return err
		}

		statusRow, err := problemset.StatusRowFromSubmission(c.Request().Context(), DB.DB, problemStore, tr, sub)
		if err != nil {
			return err
		}

		status, err := SubmissionStatusFromSubmission(tr, sub)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, struct {
			Summary SubmissionSummary `json:"summary"`
			Status  SubmissionStatus  `json:"status"`
		}{
			Summary: SubmissionSummary{
				StatusRow: *statusRow,
				Code:      string(sub.Source),
			},
			Status: *status,
		})
	}
}

func RejudgeSubmission(rs services.RejudgeService) echo.HandlerFunc {
	type request struct {
		ID int `param:"id"`
	}

	return func(c echo.Context) error {
		u := c.Get("user").(*models.User)
		if !roles.Can(roles.Role(u.Role), roles.ActionCreate, "submissions/rejudge") {
			return helpers.UnauthorizedError(c)
		}

		data := &request{}
		if err := c.Bind(data); err != nil {
			return err
		}

		if err := rs.Rejudge(c.Request().Context(), data.ID); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getSubmission", data.ID))
	}
}

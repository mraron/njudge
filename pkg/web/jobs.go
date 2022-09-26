package web

import (
	"log"
	"math"
	"time"

	"github.com/mraron/njudge/pkg/web/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *Server) StartBackgroundJobs() {
	go s.runUpdateProblems()
	go s.runStatisticsUpdate()
}

func (s *Server) runUpdateProblems() {
	for {
		if err := s.ProblemStore.Update(); err != nil {
			log.Print(err)
		}

		time.Sleep(20 * time.Second)
	}
}

func (s *Server) runStatisticsUpdate() {
	for {
		probs, err := models.ProblemRels().All(s.DB)
		if err != nil {
			log.Print(err)
			continue
		}

		userPoints := make(map[int]float64)
		for _, p := range probs {
			points := math.Sqrt(1.0 / float64(p.SolverCount))
			solvedBy, err := models.Submissions(Distinct("user_id"), Where("verdict = 0"), Where("problemset = ?", p.Problemset), Where("problem = ?", p.Problem)).All(s.DB)
			if err != nil {
				log.Print(err)
				continue
			}

			for _, uid := range solvedBy {
				userPoints[uid.UserID] += points
			}

			sub, err := models.Submissions(Where("id=?", p.ID)).One(s.DB)
			if err != nil {
				log.Print(err)
				continue
			}

			if _, err = s.DB.Exec("UPDATE problem_rels SET solver_count = (SELECT COUNT(distinct user_id) from submissions where problemset = problem_rels.problemset and problem = problem_rels.problem and verdict = 0) WHERE problemset = $1 and problem = $2", sub.Problemset, sub.Problem); err != nil {
				log.Print(err)
				continue
			}
		}

		for uid, pts := range userPoints {
			var user models.User
			user.ID = uid
			user.Points = null.Float32{Float32: float32(pts), Valid: true}
			if _, err := user.Update(s.DB, boil.Whitelist("points")); err != nil {
				log.Print(err)
				continue
			}
		}

		time.Sleep(5 * time.Minute)
	}
}

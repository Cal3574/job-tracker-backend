package goals

import (
	"fmt"
	models "job_tracker/internal/models/goals"
	"job_tracker/pkg/utils"
)

//Func to create new goal

func CreateGoal(userId string, goalName string, goalDescription string, goalAction string, goalDeadline string, goalTarget int8, goalType string) (int, error) {
	var id int

	err := utils.DB.QueryRow(
		`INSERT INTO goals(goal_name, goal_description, goal_action, goal_deadline, goal_target, user_id, goal_type) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		goalName, goalDescription, goalAction, goalDeadline, goalTarget, userId, goalType,
	).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil

}

//Func to get all goals

func GetAllGoals(userId string) ([]models.Goal, error) {
	var goals []models.Goal
	rows, err := utils.DB.Query(`SELECT id, goal_name, goal_description, goal_action, goal_deadline, goal_progress, complete, goal_target, goal_type FROM goals WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var goal models.Goal
		if err := rows.Scan(&goal.ID, &goal.GoalName, &goal.GoalDescription, &goal.GoalAction, &goal.GoalDeadline, &goal.GoalProgress, &goal.Completed, &goal.GoalTarget, &goal.GoalType); err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}

	fmt.Println(goals, "goals here")
	return goals, nil

}

// Func to delete goal
func DeleteGoal(goalId int) error {
	_, err := utils.DB.Exec("DELETE from goals WHERE id = $1", goalId)
	return err
}

// Func to progress job goal
func ProgressGoal(goalId int) error {
	_, err := utils.DB.Exec("UPDATE goals SET goal_progress = goal_progress + 1 WHERE id = $1", goalId)
	return err
}

// Func mark goal as completed
func CompleteGoal(goalId int) error {
	_, err := utils.DB.Exec("UPDATE goals SET complete = true WHERE id = $1", goalId)
	return err
}

// Func get goal by id
func GetGoalById(goalId int) (models.Goal, error) {
	var goal models.Goal
	err := utils.DB.QueryRow(`SELECT id, goal_name, goal_description, goal_action, goal_deadline, goal_progress, complete, goal_target, goal_type FROM goals WHERE id = $1`, goalId).Scan(&goal.ID, &goal.GoalName, &goal.GoalDescription, &goal.GoalAction, &goal.GoalDeadline, &goal.GoalProgress, &goal.Completed, &goal.GoalTarget, &goal.GoalType)
	if err != nil {
		return models.Goal{}, err
	}
	return goal, nil
}

// Func get goals by goal type
func GetGoalsByType(userId string, goalType string) ([]models.Goal, error) {
	var goals []models.Goal
	rows, err := utils.DB.Query(`SELECT id, goal_name, goal_description, goal_action, goal_deadline, goal_progress, complete, goal_target, goal_type FROM goals WHERE user_id = $1 AND goal_type = $2`, userId, goalType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var goal models.Goal
		if err := rows.Scan(&goal.ID, &goal.GoalName, &goal.GoalDescription, &goal.GoalAction, &goal.GoalDeadline, &goal.GoalProgress, &goal.Completed, &goal.GoalTarget, &goal.GoalType); err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}

	return goals, nil
}

package goals

import (
	"errors"
	"fmt"
	models "job_tracker/internal/models/goals"
	repositories "job_tracker/internal/repositories/goals"
	"job_tracker/internal/shared"
	"time"
)

// Function to create a new goal
func CreateGoal(userId string, goalName string, goalDescription string, goalAction string, goalDeadline string, goalTarget int8, goalType string) (int, error) {
	return repositories.CreateGoal(userId, goalName, goalDescription, goalAction, goalDeadline, goalTarget, goalType)
}

// Function to get all goals
func GetAllGoals(userId string) ([]models.Goal, error) {
	return repositories.GetAllGoals(userId)
}

// Function to delete a specific goal by goalId
func DeleteGoal(goalId int) error {
	return repositories.DeleteGoal(goalId)

}

// Function to progress goals
func ProgressGoal(goalId int) error {
	goal, err := repositories.GetGoalById(goalId)
	if err != nil {
		return err
	}

	err = repositories.ProgressGoal(goalId)
	if err != nil {
		return err
	}

	if int8(goal.GoalProgress)+1 >= int8(goal.GoalTarget) {
		err = repositories.CompleteGoal(goalId)
		if err != nil {
			return err
		}

		// Notify the SSE endpoint that a goal has been completed
		fmt.Println("Goal completed")
		fmt.Println(goalId, "goal id")
		shared.GoalCompleteChannel <- goalId
	}
	return nil
}

// Function to check if user has active goals and progress if necessary
func HandleUserGoals(userId string, goalType string) error {
	fmt.Println(userId, "user id here")
	// Retrieve all goals for the user
	goals, err := repositories.GetGoalsByType(userId, goalType)
	if err != nil {
		return err
	}

	fmt.Println(goals, "goals here in handle user goals")
	if len(goals) == 0 {
		return errors.New("user has no goals")
	}

	// Get the current time for checking active goals
	currentTime := time.Now()

	// Loop through the user's goals
	for _, goal := range goals {
		// Check if the goal is active (i.e., not completed and within the deadline)
		fmt.Println(goal.GoalDeadline, "Goal deadline")
		goalDeadline, err := time.Parse(time.RFC3339, goal.GoalDeadline) // Parse the full date-time format
		if err != nil {
			return errors.New("invalid deadline format")
		}

		//ensure we don't try and progress completed goals
		if !goal.Completed && goalDeadline.After(currentTime) {
			// If goal is active, progress the goal
			err := ProgressGoal(int(goal.ID))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

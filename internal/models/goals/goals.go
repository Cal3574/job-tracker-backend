package goals

type Goal struct {
	ID              int    `json:"id"`
	GoalName        string `json:"goal_name"`
	GoalDescription string `json:"goal_description"`
	GoalAction      string `json:"goal_action"`
	GoalDeadline    string `json:"goal_deadline"`
	GoalProgress    int8   `json:"goal_progress"`
	Completed       bool   `json:"completed"`
	GoalTarget      int8   `json:"goal_target"`
	GoalType        string `json:"goal_type"`
}

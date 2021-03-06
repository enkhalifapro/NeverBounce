package nbDto

type AccountInfo struct {
	Status                string `json:"status"`
	BillingType           string `json:"billing_type"`
	Credits               int `json:"credits"`
	FreeApiCredits        int `json:"free_api_credits"`
	MonthlyApiUsage       int `json:"monthly_api_usage"`
	MonthlyDashboardUsage int `json:"monthly_dashboard_usage"`
	JobsCompleted         int `json:"jobs_completed"`
	JobsUnderReview       int `json:"jobs_under_review"`
	JobsQueued            int `json:"jobs_queued"`
	JobsProcessing        int `json:"jobs_processing"`
	ExecutionTime         int `json:"execution_time"`
}

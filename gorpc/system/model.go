package system

// Define the Loader struct to hold data from the `loader` table
type Loader struct {
	ID        int    `json:"id"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

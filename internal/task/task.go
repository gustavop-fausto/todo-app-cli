package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Done        bool   `json:"done"`
	Description string `json:"description"`
}

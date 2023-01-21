package pteroclient

type Server struct {
	Object     string `json:"object"`
	Attributes struct {
		ServerOwner bool   `json:"server_owner"`
		Identifier  string `json:"identifier"`
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		Node        string `json:"node"`
		SftpDetails struct {
			IP   string `json:"ip"`
			Port int    `json:"port"`
		} `json:"sftp_details"`
		Description string `json:"description"`
		Limits      struct {
			Memory int `json:"memory"`
			Swap   int `json:"swap"`
			Disk   int `json:"disk"`
			Io     int `json:"io"`
			Cpu    int `json:"cpu"`
		} `json:"limits"`
		FeatureLimits struct {
			Databases   int `json:"databases"`
			Allocations int `json:"allocations"`
			Backups     int `json:"backups"`
		} `json:"feature_limits"`
		IsSuspended   bool `json:"is_suspended"`
		IsInstalling  bool `json:"is_installing"`
		Relationships struct {
			Allocations struct {
				Object string `json:"object"`
				Data   []struct {
					Object     string `json:"object"`
					Attributes struct {
						Id        int    `json:"id"`
						IP        string `json:"ip"`
						IPAlias   string `json:"ip_alias"`
						Port      int    `json:"port"`
						Notes     string `json:"notes"`
						IsDefault bool   `json:"is_default"`
					} `json:"attributes"`
				} `json:"data"`
			} `json:"allocations"`
		} `json:"relationships"`
	} `json:"attributes"`
}

type ServerList struct {
	Object string   `json:"object"`
	Data   []Server `json:"data"`
	Meta   struct {
		Pagination struct {
			Total       int      `json:"total"`
			Count       int      `json:"count"`
			PerPage     int      `json:"per_page"`
			CurrentPage int      `json:"current_page"`
			TotalPages  int      `json:"total_pages"`
			Links       struct{} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
}

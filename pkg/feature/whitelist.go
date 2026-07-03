package feature

type IPWhiteList struct {
	Whitelist map[string]bool `json:whitelist`
}

func PopulateIPWhiteList(w *IPWhiteList, ipList []string){
	if len(ipList) > 0 && ipList[0] == "ALL" {
		
	}
}
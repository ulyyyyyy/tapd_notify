package condition

import "testing"

func Test_parseCondition(t *testing.T) {
	parseCondition(`project = SST and issuetype = "故障`)
}

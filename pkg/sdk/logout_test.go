package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogout(t *testing.T) {
	req := httptest.NewRequest("GET", "/logout", nil)
	w := httptest.NewRecorder()

	mockJson := `{"_benchmark":0.127513,"data":{"ID":1,"Email":"stevejobs@apple.com","Name":"Steve Jobs","Avatar":"","Payload":{},"Stripe":"","Terms":true,"Roles":[],"Settings":{},"Providers":null,"Teams":null,"URL":"/user/1","Initials":"SJ","TrialDays":0,"Pivot":null,"OrgRoles":null,"CurrentToken":null,"ActivatedAt":null,"OrgUsers":[],"Orgs":[],"CurrentOrg":null,"OrgID":null,"Org":null,"IsServiceUser":false,"HasEmployeeRole":false,"HasEmployeeAdminRole":false,"HasOrgManagerRole":false,"HasTrial":false,"HasInitial":false,"HasVuln":false,"HasAgent":false,"HasSbom":false,"OnlyCommunity":true,"created_at":"2024-03-11T13:05:48.049Z","updated_at":"2024-03-11T13:05:48.049Z"}}`

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockJson)
	})

	handler.ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		t.Error("error unmarshalling response")
	}

	t.Run("user response is parsed", func(t *testing.T) {
		assert.Equal(t, 1, response.Data.ID)
	})
}

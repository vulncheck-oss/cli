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

func TestGetIndexBackup(t *testing.T) {
	req := httptest.NewRequest("GET", "/index/backup", nil)
	w := httptest.NewRecorder()

	mockJson := `{"_benchmark":0.431983,"_meta":{"timestamp":"2024-02-12T23:12:49.009938487Z","index":"initial-access"},"data":[{"filename":"initial-access-1707769254529150054.zip","sha256":"27285d350260ccaf8e3adebd1ede53305e14d2e1b31a37d6f22f7d81e420c309","date_added":"2024-02-12T20:20:54.529Z","url":"https://mjwq58t95817e.mrap.accesspoint.s3-global.amazonaws.com/initial-access-1707769254529150054.zip?X-Amz-Algorithm=AWS4-ECDSA-P256-SHA256\u0026X-Amz-Credential=ASIAZKOAMLEDYPLCCK5U%2F20240212%2Fs3%2Faws4_request\u0026X-Amz-Date=20240212T231249Z\u0026X-Amz-Expires=900\u0026X-Amz-Region-Set=%2A\u0026X-Amz-Security-Token=IQoJb3JpZ2luX2VjEJD%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJHMEUCIQDZrFi6vhQUSX3btJhstRogbNykr2lp4PjMpASyFB%2BWZwIgO%2FnISolIR3Q4Dkgv4SRbziQUyT5R7iqi%2FaW%2FBx8VfQAq9AIIaBACGgw2NDA4OTA0NjA0MjMiDG5GDtB7G6sZZiGdGCrRAmA0Til8bdRJIHBabryj%2Brjj8tIPyGjcQsqlBSTowMNV3DGctOcPSlOz0sNDA0DNNsAbRPgdRtmnNtygnSmcZdwQFJ4GCpGRcy%2F9FH8hOcdObss248mU9DNoYP83Rm0Q7zsIANofCmsASUsev%2BcqDPVzff%2BSg33AGzFTEqDbWSFED%2FkCbWjq4Ah8gfswQMuW7NC%2FNX9PEmLncrS242eojhQHPyhNDO9QTea7SuoneZtPN2kQdIZQo7AOpXQuPVG4Q6%2BEKBdxjomNvvaMyj89fqWr75r2OyyXZeRgBnHOO%2FI7NCV9QJjA425oVrHTbQ9UTI911oVTb14yBpuS3hqjP3WU5DEOpd1O4fa1yN5cxC%2FzufLN%2FtL4eLH4h2vknuUkoOR8%2FHrSQnZi45g4Uc1aESnx3btRDHRShKnn5T5DumOyf5gDZEeeRmGY8tD%2F8187M8Uw78uqrgY6ngGcAVqoSnQFhTfgp3aT%2B6CaDqWHWoiQbsZOd7itGhPsp7uZ%2F71PsRWpdqrv4wsuiU0uvD0Krcx5xvdv9nflGKBXe4nk%2BKEJpNrBq12WOcPLm5809ojTNuqmy9wAOIcMUPtMmTTC32qiWMe%2FnpUYZAjSqsT4UVwAOBlMDdAJgL%2Bwrc27DuZSPVDqRjWtaTnkurQbWiQ0aDLJ0NXA2sRiag%3D%3D\u0026X-Amz-SignedHeaders=host\u0026x-id=GetObject\u0026X-Amz-Signature=30450221008ddf6a2e6244cdf9321ffbc66ecf5c9112eba6ff7d26964c512b70cd51eb91a702205803b43a6fe20f4574965d43e5a75906b186ea0504663370a09bccc513c16dde","url_ttl_minutes":60,"url_expires":"2024-02-12T21:20:54.529Z"}]}`

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockJson)
	})

	handler.ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	var backupResp BackupResponse
	err := json.Unmarshal(body, &backupResp)
	if err != nil {
		t.Error("error unmarshalling response")
	}

	t.Run("backup response is parsed", func(t *testing.T) {
		assert.Equal(t, "initial-access-1707769254529150054.zip", backupResp.Data[0].Filename)
	})
}

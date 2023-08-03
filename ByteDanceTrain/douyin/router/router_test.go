package router

/* this is a typical case
func TestRouter(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/someJSON", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	message := &gin.H{}
	if err = json.Unmarshal(w.Body.Bytes(), message); err != nil {
		t.Fatal(err)
	}
	//for key, val := range *message {
	//	log.Printf("key: %v %T\tval: %v %T", key, key, val, val)
	//}
	assert.Equal(t, gin.H{"message": "Hey", "status": float64(http.StatusOK)}, *message)
}
*/
